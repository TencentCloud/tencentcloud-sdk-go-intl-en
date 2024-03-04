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

package v20190924

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AccessVpc struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Private network access status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Private network access IP
	AccessIp *string `json:"AccessIp,omitnil,omitempty" name:"AccessIp"`
}

// Predefined struct for user
type CheckInstanceNameRequestParams struct {
	// Name of the instance to be created
	RegistryName *string `json:"RegistryName,omitnil,omitempty" name:"RegistryName"`
}

type CheckInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// Name of the instance to be created
	RegistryName *string `json:"RegistryName,omitnil,omitempty" name:"RegistryName"`
}

func (r *CheckInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckInstanceNameResponseParams struct {
	// Verification result. Valid values: true: Valid; false: Invalid.
	IsValidated *bool `json:"IsValidated,omitnil,omitempty" name:"IsValidated"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *CheckInstanceNameResponseParams `json:"Response"`
}

func (r *CheckInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckInstanceRequestParams struct {
	// ID of the instance to be verified.
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

type CheckInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be verified.
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

func (r *CheckInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckInstanceResponseParams struct {
	// Verification result. true: valid, false: invalid
	IsValidated *bool `json:"IsValidated,omitnil,omitempty" name:"IsValidated"`

	// ID of the region where the instance is located.
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CheckInstanceResponseParams `json:"Response"`
}

func (r *CheckInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageAccelerationServiceRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// ID of the VPC where the CFS file system to be created resides
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the subnet where the CFS file system to be created resides
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Storage class of the CFS file system to be created. Valid values: SD: Standard; HP: High-Performance.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// AZ name, such as `ap-beijing-1`. For more information, see the list of regions and AZs in Overview.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Cloud tag description
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

type CreateImageAccelerationServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// ID of the VPC where the CFS file system to be created resides
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the subnet where the CFS file system to be created resides
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Storage class of the CFS file system to be created. Valid values: SD: Standard; HP: High-Performance.
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// Permission group ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// AZ name, such as `ap-beijing-1`. For more information, see the list of regions and AZs in Overview.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Cloud tag description
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

func (r *CreateImageAccelerationServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageAccelerationServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "StorageType")
	delete(f, "PGroupId")
	delete(f, "Zone")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageAccelerationServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageAccelerationServiceResponseParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateImageAccelerationServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageAccelerationServiceResponseParams `json:"Response"`
}

func (r *CreateImageAccelerationServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageAccelerationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImmutableTagRulesRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Rule
	Rule *ImmutableTagRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type CreateImmutableTagRulesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Rule
	Rule *ImmutableTagRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

func (r *CreateImmutableTagRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImmutableTagRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImmutableTagRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImmutableTagRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateImmutableTagRulesResponseParams `json:"Response"`
}

func (r *CreateImmutableTagRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceCustomizedDomainRequestParams struct {
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Custom domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type CreateInstanceCustomizedDomainRequest struct {
	*tchttp.BaseRequest
	
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Custom domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *CreateInstanceCustomizedDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceCustomizedDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "DomainName")
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceCustomizedDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceCustomizedDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceCustomizedDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceCustomizedDomainResponseParams `json:"Response"`
}

func (r *CreateInstanceCustomizedDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceCustomizedDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// Enterprise Edition instance name
	RegistryName *string `json:"RegistryName,omitnil,omitempty" name:"RegistryName"`

	// Enterprise Edition instance type. Valid values: basic: Basic; standard: Standard; premium: Premium.
	RegistryType *string `json:"RegistryType,omitnil,omitempty" name:"RegistryType"`

	// Cloud tag description
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// Instance billing mode. Valid values: 0: Pay-as-you-go billing; 1: Prepaid. Default value: 0.
	RegistryChargeType *int64 `json:"RegistryChargeType,omitnil,omitempty" name:"RegistryChargeType"`

	// Auto-renewal setting and purchase period
	RegistryChargePrepaid *RegistryChargePrepaid `json:"RegistryChargePrepaid,omitnil,omitempty" name:"RegistryChargePrepaid"`

	// Whether to sync TCR cloud tags to the COS bucket
	SyncTag *bool `json:"SyncTag,omitnil,omitempty" name:"SyncTag"`

	// Whether to enable the COS Multi-AZ feature
	EnableCosMAZ *bool `json:"EnableCosMAZ,omitnil,omitempty" name:"EnableCosMAZ"`

	// Whether to enable deletion protection
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Enterprise Edition instance name
	RegistryName *string `json:"RegistryName,omitnil,omitempty" name:"RegistryName"`

	// Enterprise Edition instance type. Valid values: basic: Basic; standard: Standard; premium: Premium.
	RegistryType *string `json:"RegistryType,omitnil,omitempty" name:"RegistryType"`

	// Cloud tag description
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// Instance billing mode. Valid values: 0: Pay-as-you-go billing; 1: Prepaid. Default value: 0.
	RegistryChargeType *int64 `json:"RegistryChargeType,omitnil,omitempty" name:"RegistryChargeType"`

	// Auto-renewal setting and purchase period
	RegistryChargePrepaid *RegistryChargePrepaid `json:"RegistryChargePrepaid,omitnil,omitempty" name:"RegistryChargePrepaid"`

	// Whether to sync TCR cloud tags to the COS bucket
	SyncTag *bool `json:"SyncTag,omitnil,omitempty" name:"SyncTag"`

	// Whether to enable the COS Multi-AZ feature
	EnableCosMAZ *bool `json:"EnableCosMAZ,omitnil,omitempty" name:"EnableCosMAZ"`

	// Whether to enable deletion protection
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryName")
	delete(f, "RegistryType")
	delete(f, "TagSpecification")
	delete(f, "RegistryChargeType")
	delete(f, "RegistryChargePrepaid")
	delete(f, "SyncTag")
	delete(f, "EnableCosMAZ")
	delete(f, "DeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// Enterprise Edition instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceResponseParams `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceTokenRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Access credential type. Values: `longterm` and `temp` (default, valid for one hour)
	TokenType *string `json:"TokenType,omitnil,omitempty" name:"TokenType"`

	// Description of the long-term access credential
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type CreateInstanceTokenRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Access credential type. Values: `longterm` and `temp` (default, valid for one hour)
	TokenType *string `json:"TokenType,omitnil,omitempty" name:"TokenType"`

	// Description of the long-term access credential
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

func (r *CreateInstanceTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "TokenType")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceTokenResponseParams struct {
	// Username
	// Note: this field may return `null`, indicating that no valid value can be found.
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Access credential
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// Expiration timestamp of access credential. It is a string of numbers without unit.
	ExpTime *int64 `json:"ExpTime,omitnil,omitempty" name:"ExpTime"`

	// Token ID of long-term access credential. It is not available to temporary access credential.
	// Note: this field may return `null`, indicating that no valid value can be found.
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceTokenResponseParams `json:"Response"`
}

func (r *CreateInstanceTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultipleSecurityPolicyRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Security group policy
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
}

type CreateMultipleSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Security group policy
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *CreateMultipleSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultipleSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMultipleSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMultipleSecurityPolicyResponseParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMultipleSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateMultipleSecurityPolicyResponseParams `json:"Response"`
}

func (r *CreateMultipleSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMultipleSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespaceRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name, which can contain 2–30 lowercase letters, digits, and separators (".", "_", and "-") but can neither start or end with a separator nor contain consecutive separators.
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Whether to make public. Valid values: true: Yes; false: No.
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// Cloud tag description
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name, which can contain 2–30 lowercase letters, digits, and separators (".", "_", and "-") but can neither start or end with a separator nor contain consecutive separators.
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Whether to make public. Valid values: true: Yes; false: No.
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// Cloud tag description
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

func (r *CreateNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "IsPublic")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNamespaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNamespaceResponseParams `json:"Response"`
}

func (r *CreateNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReplicationInstanceRequestParams struct {
	// Master instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Region ID of the replication instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`

	// Region name of the replication instance
	ReplicationRegionName *string `json:"ReplicationRegionName,omitnil,omitempty" name:"ReplicationRegionName"`

	// Whether to sync TCR cloud tags to the COS Bucket
	SyncTag *bool `json:"SyncTag,omitnil,omitempty" name:"SyncTag"`
}

type CreateReplicationInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Master instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Region ID of the replication instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`

	// Region name of the replication instance
	ReplicationRegionName *string `json:"ReplicationRegionName,omitnil,omitempty" name:"ReplicationRegionName"`

	// Whether to sync TCR cloud tags to the COS Bucket
	SyncTag *bool `json:"SyncTag,omitnil,omitempty" name:"SyncTag"`
}

func (r *CreateReplicationInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplicationInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "ReplicationRegionId")
	delete(f, "ReplicationRegionName")
	delete(f, "SyncTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReplicationInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReplicationInstanceResponseParams struct {
	// Enterprise Registry Instance ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitnil,omitempty" name:"ReplicationRegistryId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateReplicationInstanceResponseParams `json:"Response"`
}

func (r *CreateReplicationInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRepositoryRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Brief repository description
	BriefDescription *string `json:"BriefDescription,omitnil,omitempty" name:"BriefDescription"`

	// Detailed repository description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Brief repository description
	BriefDescription *string `json:"BriefDescription,omitnil,omitempty" name:"BriefDescription"`

	// Detailed repository description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "BriefDescription")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRepositoryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *CreateRepositoryResponseParams `json:"Response"`
}

func (r *CreateRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityPolicyRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 192.168.0.0/24
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 192.168.0.0/24
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "CidrBlock")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityPolicyResponseParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityPolicyResponseParams `json:"Response"`
}

func (r *CreateSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceAccountRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Service account name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Policy list
	Permissions []*Permission `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// Service account description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Validity in days starting from the current day. It takes a higher priority than `ExpiresAt`.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Expiry time (timestamp, in milliseconds)
	ExpiresAt *int64 `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// Whether to disable the service account
	Disable *bool `json:"Disable,omitnil,omitempty" name:"Disable"`
}

type CreateServiceAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Service account name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Policy list
	Permissions []*Permission `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// Service account description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Validity in days starting from the current day. It takes a higher priority than `ExpiresAt`.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Expiry time (timestamp, in milliseconds)
	ExpiresAt *int64 `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// Whether to disable the service account
	Disable *bool `json:"Disable,omitnil,omitempty" name:"Disable"`
}

func (r *CreateServiceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Name")
	delete(f, "Permissions")
	delete(f, "Description")
	delete(f, "Duration")
	delete(f, "ExpiresAt")
	delete(f, "Disable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceAccountResponseParams struct {
	// Service account name (the prefix `tcr$` is automatically added)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Service account password, which is displayed only once
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Expiry time of the service account (timestamp)
	ExpiresAt *int64 `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// Creation time of the service account
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateServiceAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceAccountResponseParams `json:"Response"`
}

func (r *CreateServiceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignaturePolicyRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Policy name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// KMS key
	KmsId *string `json:"KmsId,omitnil,omitempty" name:"KmsId"`

	// Region of the KMS key
	KmsRegion *string `json:"KmsRegion,omitnil,omitempty" name:"KmsRegion"`

	// Custom domain name. If this parameter is left empty, the default domain name of the TCR instance will be used to generate the signature.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Whether to disable the signing policy. Default value: false.
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`
}

type CreateSignaturePolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Policy name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// KMS key
	KmsId *string `json:"KmsId,omitnil,omitempty" name:"KmsId"`

	// Region of the KMS key
	KmsRegion *string `json:"KmsRegion,omitnil,omitempty" name:"KmsRegion"`

	// Custom domain name. If this parameter is left empty, the default domain name of the TCR instance will be used to generate the signature.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Whether to disable the signing policy. Default value: false.
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`
}

func (r *CreateSignaturePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignaturePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Name")
	delete(f, "NamespaceName")
	delete(f, "KmsId")
	delete(f, "KmsRegion")
	delete(f, "Domain")
	delete(f, "Disabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSignaturePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignaturePolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSignaturePolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateSignaturePolicyResponseParams `json:"Response"`
}

func (r *CreateSignaturePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignaturePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignatureRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Tag name
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`
}

type CreateSignatureRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Tag name
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`
}

func (r *CreateSignatureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignatureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "ImageVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSignatureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignatureResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSignatureResponse struct {
	*tchttp.BaseResponse
	Response *CreateSignatureResponseParams `json:"Response"`
}

func (r *CreateSignatureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignatureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRetentionExecutionRequestParams struct {
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Tag retention rule ID
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// Whether the execution is simulated. Default value: false (not simulated)
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type CreateTagRetentionExecutionRequest struct {
	*tchttp.BaseRequest
	
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Tag retention rule ID
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// Whether the execution is simulated. Default value: false (not simulated)
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *CreateTagRetentionExecutionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRetentionExecutionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RetentionId")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTagRetentionExecutionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRetentionExecutionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTagRetentionExecutionResponse struct {
	*tchttp.BaseResponse
	Response *CreateTagRetentionExecutionResponseParams `json:"Response"`
}

func (r *CreateTagRetentionExecutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRetentionExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRetentionRuleRequestParams struct {
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace ID
	NamespaceId *int64 `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Retention policy
	RetentionRule *RetentionRule `json:"RetentionRule,omitnil,omitempty" name:"RetentionRule"`

	// Execution cycle. Valid values: manual, daily, weekly, monthly.
	CronSetting *string `json:"CronSetting,omitnil,omitempty" name:"CronSetting"`

	// Whether to disable the rule. Default value: false.
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`
}

type CreateTagRetentionRuleRequest struct {
	*tchttp.BaseRequest
	
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace ID
	NamespaceId *int64 `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Retention policy
	RetentionRule *RetentionRule `json:"RetentionRule,omitnil,omitempty" name:"RetentionRule"`

	// Execution cycle. Valid values: manual, daily, weekly, monthly.
	CronSetting *string `json:"CronSetting,omitnil,omitempty" name:"CronSetting"`

	// Whether to disable the rule. Default value: false.
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`
}

func (r *CreateTagRetentionRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRetentionRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceId")
	delete(f, "RetentionRule")
	delete(f, "CronSetting")
	delete(f, "Disabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTagRetentionRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTagRetentionRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTagRetentionRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateTagRetentionRuleResponseParams `json:"Response"`
}

func (r *CreateTagRetentionRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTagRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebhookTriggerRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Trigger parameter
	Trigger *WebhookTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type CreateWebhookTriggerRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Trigger parameter
	Trigger *WebhookTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *CreateWebhookTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebhookTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Trigger")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWebhookTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebhookTriggerResponseParams struct {
	// Newly created trigger
	Trigger *WebhookTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWebhookTriggerResponse struct {
	*tchttp.BaseResponse
	Response *CreateWebhookTriggerResponseParams `json:"Response"`
}

func (r *CreateWebhookTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomizedDomainInfo struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Certificate ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Domain name creation status. Valid values: SUCCESS, FAILURE, CREATING, DELETING.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteImageAccelerateServiceRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

type DeleteImageAccelerateServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

func (r *DeleteImageAccelerateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageAccelerateServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageAccelerateServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageAccelerateServiceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteImageAccelerateServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageAccelerateServiceResponseParams `json:"Response"`
}

func (r *DeleteImageAccelerateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageAccelerateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Image repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Image tag
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`
}

type DeleteImageRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Image repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Image tag
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`
}

func (r *DeleteImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RepositoryName")
	delete(f, "ImageVersion")
	delete(f, "NamespaceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteImageResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageResponseParams `json:"Response"`
}

func (r *DeleteImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImmutableTagRulesRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Rule ID
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteImmutableTagRulesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Rule ID
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *DeleteImmutableTagRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImmutableTagRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImmutableTagRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImmutableTagRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImmutableTagRulesResponseParams `json:"Response"`
}

func (r *DeleteImmutableTagRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceCustomizedDomainRequestParams struct {
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Custom domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type DeleteInstanceCustomizedDomainRequest struct {
	*tchttp.BaseRequest
	
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Custom domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *DeleteInstanceCustomizedDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceCustomizedDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "DomainName")
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceCustomizedDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceCustomizedDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceCustomizedDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceCustomizedDomainResponseParams `json:"Response"`
}

func (r *DeleteInstanceCustomizedDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceCustomizedDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Whether to delete the bucket. Default value: false.
	DeleteBucket *bool `json:"DeleteBucket,omitnil,omitempty" name:"DeleteBucket"`

	// Whether to enable the `dryRun` mode. Default value: false.
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Whether to delete the bucket. Default value: false.
	DeleteBucket *bool `json:"DeleteBucket,omitnil,omitempty" name:"DeleteBucket"`

	// Whether to enable the `dryRun` mode. Default value: false.
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "DeleteBucket")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceResponseParams `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceTokenRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Access credential ID
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`
}

type DeleteInstanceTokenRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Access credential ID
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`
}

func (r *DeleteInstanceTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "TokenId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceTokenResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceTokenResponseParams `json:"Response"`
}

func (r *DeleteInstanceTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMultipleSecurityPolicyRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Security group policy
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
}

type DeleteMultipleSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Security group policy
	SecurityGroupPolicySet []*SecurityPolicy `json:"SecurityGroupPolicySet,omitnil,omitempty" name:"SecurityGroupPolicySet"`
}

func (r *DeleteMultipleSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultipleSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMultipleSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMultipleSecurityPolicyResponseParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMultipleSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMultipleSecurityPolicyResponseParams `json:"Response"`
}

func (r *DeleteMultipleSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMultipleSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNamespaceRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`
}

type DeleteNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`
}

func (r *DeleteNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNamespaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNamespaceResponseParams `json:"Response"`
}

func (r *DeleteNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReplicationInstanceRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Replica instance ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitnil,omitempty" name:"ReplicationRegistryId"`

	// Region ID of the replica instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`
}

type DeleteReplicationInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Replica instance ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitnil,omitempty" name:"ReplicationRegistryId"`

	// Region ID of the replica instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`
}

func (r *DeleteReplicationInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReplicationInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "ReplicationRegistryId")
	delete(f, "ReplicationRegionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReplicationInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReplicationInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReplicationInstanceResponseParams `json:"Response"`
}

func (r *DeleteReplicationInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Image repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`
}

type DeleteRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Image repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`
}

func (r *DeleteRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRepositoryResponseParams `json:"Response"`
}

func (r *DeleteRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryTagsRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// List of tags. Up to 20 tags can be returned for a request.
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DeleteRepositoryTagsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// List of tags. Up to 20 tags can be returned for a request.
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DeleteRepositoryTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRepositoryTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRepositoryTagsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRepositoryTagsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRepositoryTagsResponseParams `json:"Response"`
}

func (r *DeleteRepositoryTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRepositoryTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityPolicyRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Allowlist ID
	PolicyIndex *int64 `json:"PolicyIndex,omitnil,omitempty" name:"PolicyIndex"`

	// Allowlist version
	PolicyVersion *string `json:"PolicyVersion,omitnil,omitempty" name:"PolicyVersion"`

	// IP range or IP address (mutually exclusive).
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`
}

type DeleteSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Allowlist ID
	PolicyIndex *int64 `json:"PolicyIndex,omitnil,omitempty" name:"PolicyIndex"`

	// Allowlist version
	PolicyVersion *string `json:"PolicyVersion,omitnil,omitempty" name:"PolicyVersion"`

	// IP range or IP address (mutually exclusive).
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`
}

func (r *DeleteSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "PolicyIndex")
	delete(f, "PolicyVersion")
	delete(f, "CidrBlock")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityPolicyResponseParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityPolicyResponseParams `json:"Response"`
}

func (r *DeleteSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceAccountRequestParams struct {
	// Instance ID	
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Service account name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteServiceAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID	
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Service account name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DeleteServiceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteServiceAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceAccountResponseParams `json:"Response"`
}

func (r *DeleteServiceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSignaturePolicyRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`
}

type DeleteSignaturePolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`
}

func (r *DeleteSignaturePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSignaturePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSignaturePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSignaturePolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSignaturePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSignaturePolicyResponseParams `json:"Response"`
}

func (r *DeleteSignaturePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSignaturePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagRetentionRuleRequestParams struct {
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Tag retention rule ID
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`
}

type DeleteTagRetentionRuleRequest struct {
	*tchttp.BaseRequest
	
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Tag retention rule ID
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`
}

func (r *DeleteTagRetentionRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagRetentionRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RetentionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTagRetentionRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTagRetentionRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTagRetentionRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTagRetentionRuleResponseParams `json:"Response"`
}

func (r *DeleteTagRetentionRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTagRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWebhookTriggerRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Trigger ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteWebhookTriggerRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Trigger ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteWebhookTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebhookTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Namespace")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWebhookTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWebhookTriggerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWebhookTriggerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWebhookTriggerResponseParams `json:"Response"`
}

func (r *DeleteWebhookTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChartDownloadInfoRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Chart name
	ChartName *string `json:"ChartName,omitnil,omitempty" name:"ChartName"`

	// Chart version
	ChartVersion *string `json:"ChartVersion,omitnil,omitempty" name:"ChartVersion"`
}

type DescribeChartDownloadInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Chart name
	ChartName *string `json:"ChartName,omitnil,omitempty" name:"ChartName"`

	// Chart version
	ChartVersion *string `json:"ChartVersion,omitnil,omitempty" name:"ChartVersion"`
}

func (r *DescribeChartDownloadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChartDownloadInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "ChartName")
	delete(f, "ChartVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChartDownloadInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChartDownloadInfoResponseParams struct {
	// Presigned URL for download
	PreSignedDownloadURL *string `json:"PreSignedDownloadURL,omitnil,omitempty" name:"PreSignedDownloadURL"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChartDownloadInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChartDownloadInfoResponseParams `json:"Response"`
}

func (r *DescribeChartDownloadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChartDownloadInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExternalEndpointStatusRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

type DescribeExternalEndpointStatusRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

func (r *DescribeExternalEndpointStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExternalEndpointStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExternalEndpointStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExternalEndpointStatusResponseParams struct {
	// Public network access status. Valid values: Opening, Opened, Closed.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Reason
	// Note: This field may return null, indicating that no valid values can be obtained.
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExternalEndpointStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExternalEndpointStatusResponseParams `json:"Response"`
}

func (r *DescribeExternalEndpointStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExternalEndpointStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGCJobsRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

type DescribeGCJobsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

func (r *DescribeGCJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGCJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGCJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGCJobsResponseParams struct {
	// List of GC jobs
	Jobs []*GCJobInfo `json:"Jobs,omitnil,omitempty" name:"Jobs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGCJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGCJobsResponseParams `json:"Response"`
}

func (r *DescribeGCJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGCJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAccelerateServiceRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

type DescribeImageAccelerateServiceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

func (r *DescribeImageAccelerateServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAccelerateServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageAccelerateServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageAccelerateServiceResponseParams struct {
	// Image acceleration status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// CFS VIP
	CFSVIP *string `json:"CFSVIP,omitnil,omitempty" name:"CFSVIP"`

	// Whether to enable
	IsEnable *bool `json:"IsEnable,omitnil,omitempty" name:"IsEnable"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageAccelerateServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageAccelerateServiceResponseParams `json:"Response"`
}

func (r *DescribeImageAccelerateServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageAccelerateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageManifestsRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Image repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Image tag
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`
}

type DescribeImageManifestsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Image repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Image tag
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`
}

func (r *DescribeImageManifestsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageManifestsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "ImageVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageManifestsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageManifestsResponseParams struct {
	// Image manifest information
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// Image configuration information
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageManifestsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageManifestsResponseParams `json:"Response"`
}

func (r *DescribeImageManifestsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageManifestsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Image repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Image tag specified for fuzzy search
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// Number of entries per page, which is used for pagination. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 1.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Image digest specified for search
	Digest *string `json:"Digest,omitnil,omitempty" name:"Digest"`

	// Whether to use exact matching. Valid values: `true` (exact matching), `null` (fuzzy matching).
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Image repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Image tag specified for fuzzy search
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// Number of entries per page, which is used for pagination. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number. Default value: 1.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Image digest specified for search
	Digest *string `json:"Digest,omitnil,omitempty" name:"Digest"`

	// Whether to use exact matching. Valid values: `true` (exact matching), `null` (fuzzy matching).
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "ImageVersion")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Digest")
	delete(f, "ExactMatch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesResponseParams struct {
	// List of container images
	ImageInfoList []*TcrImageInfo `json:"ImageInfoList,omitnil,omitempty" name:"ImageInfoList"`

	// Total number of container images
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagesResponseParams `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImmutableTagRulesRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

type DescribeImmutableTagRulesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

func (r *DescribeImmutableTagRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImmutableTagRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImmutableTagRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImmutableTagRulesResponseParams struct {
	// Rule list
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	Rules []*ImmutableTagRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// Namespace with no rules created
	// Note: this field may return `null`, indicating that no valid value can be obtained.
	EmptyNs []*string `json:"EmptyNs,omitnil,omitempty" name:"EmptyNs"`

	// Total rules
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImmutableTagRulesResponseParams `json:"Response"`
}

func (r *DescribeImmutableTagRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAllNamespacesRequestParams struct {
	// Number of entries per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Start position offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeInstanceAllNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// Number of entries per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Start position offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeInstanceAllNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAllNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceAllNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAllNamespacesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceAllNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceAllNamespacesResponseParams `json:"Response"`
}

func (r *DescribeInstanceAllNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAllNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCustomizedDomainRequestParams struct {
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeInstanceCustomizedDomainRequest struct {
	*tchttp.BaseRequest
	
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeInstanceCustomizedDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCustomizedDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceCustomizedDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCustomizedDomainResponseParams struct {
	// List of domain names
	// Note: This field may return null, indicating that no valid values can be obtained.
	DomainInfoList []*CustomizedDomainInfo `json:"DomainInfoList,omitnil,omitempty" name:"DomainInfoList"`

	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceCustomizedDomainResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceCustomizedDomainResponseParams `json:"Response"`
}

func (r *DescribeInstanceCustomizedDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCustomizedDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStatusRequestParams struct {
	// Array of instance IDs
	RegistryIds []*string `json:"RegistryIds,omitnil,omitempty" name:"RegistryIds"`
}

type DescribeInstanceStatusRequest struct {
	*tchttp.BaseRequest
	
	// Array of instance IDs
	RegistryIds []*string `json:"RegistryIds,omitnil,omitempty" name:"RegistryIds"`
}

func (r *DescribeInstanceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStatusResponseParams struct {
	// List of instance statuses
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegistryStatusSet []*RegistryStatus `json:"RegistryStatusSet,omitnil,omitempty" name:"RegistryStatusSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceStatusResponseParams `json:"Response"`
}

func (r *DescribeInstanceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTokenRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeInstanceTokenRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeInstanceTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTokenResponseParams struct {
	// Total number of long-term access credentials
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of long-term access credentials
	Tokens []*TcrInstanceToken `json:"Tokens,omitnil,omitempty" name:"Tokens"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceTokenResponseParams `json:"Response"`
}

func (r *DescribeInstanceTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// List of instance IDs (if it is empty,
	// it indicates to get all instances under the current account)
	Registryids []*string `json:"Registryids,omitnil,omitempty" name:"Registryids"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Whether to get the instances in all regions. Default value: False.
	AllRegion *bool `json:"AllRegion,omitnil,omitempty" name:"AllRegion"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs (if it is empty,
	// it indicates to get all instances under the current account)
	Registryids []*string `json:"Registryids,omitnil,omitempty" name:"Registryids"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Filters
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Whether to get the instances in all regions. Default value: False.
	AllRegion *bool `json:"AllRegion,omitnil,omitempty" name:"AllRegion"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Registryids")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "AllRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Total number of instances
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of instances
	// Note: This field may return null, indicating that no valid values can be obtained.
	Registries []*Registry `json:"Registries,omitnil,omitempty" name:"Registries"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternalEndpointsRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

type DescribeInternalEndpointsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

func (r *DescribeInternalEndpointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternalEndpointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInternalEndpointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInternalEndpointsResponseParams struct {
	// List of private network access addresses
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessVpcSet []*AccessVpc `json:"AccessVpcSet,omitnil,omitempty" name:"AccessVpcSet"`

	// Total number of private network access addresses
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInternalEndpointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInternalEndpointsResponseParams `json:"Response"`
}

func (r *DescribeInternalEndpointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInternalEndpointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespacesRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Specified namespace. If this parameter is left empty, all namespaces will be queried.
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page offset (page number from which to return the results)
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Whether to list all namespaces
	All *bool `json:"All,omitnil,omitempty" name:"All"`

	// Filters
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Whether to query only namespaces for which the KMS image signature is enabled
	KmsSignPolicy *bool `json:"KmsSignPolicy,omitnil,omitempty" name:"KmsSignPolicy"`
}

type DescribeNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Specified namespace. If this parameter is left empty, all namespaces will be queried.
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page offset (page number from which to return the results)
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Whether to list all namespaces
	All *bool `json:"All,omitnil,omitempty" name:"All"`

	// Filters
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Whether to query only namespaces for which the KMS image signature is enabled
	KmsSignPolicy *bool `json:"KmsSignPolicy,omitnil,omitempty" name:"KmsSignPolicy"`
}

func (r *DescribeNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "All")
	delete(f, "Filters")
	delete(f, "KmsSignPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespacesResponseParams struct {
	// List of namespaces
	NamespaceList []*TcrNamespaceInfo `json:"NamespaceList,omitnil,omitempty" name:"NamespaceList"`

	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNamespacesResponseParams `json:"Response"`
}

func (r *DescribeNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {

}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// Total number of returned results
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of regions
	Regions []*Region `json:"Regions,omitnil,omitempty" name:"Regions"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstanceCreateTasksRequestParams struct {
	// Replication instance ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitnil,omitempty" name:"ReplicationRegistryId"`

	// Region ID of the replication instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`
}

type DescribeReplicationInstanceCreateTasksRequest struct {
	*tchttp.BaseRequest
	
	// Replication instance ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitnil,omitempty" name:"ReplicationRegistryId"`

	// Region ID of the replication instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`
}

func (r *DescribeReplicationInstanceCreateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstanceCreateTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReplicationRegistryId")
	delete(f, "ReplicationRegionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationInstanceCreateTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstanceCreateTasksResponseParams struct {
	// Task details
	TaskDetail []*TaskDetail `json:"TaskDetail,omitnil,omitempty" name:"TaskDetail"`

	// Overall task status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReplicationInstanceCreateTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplicationInstanceCreateTasksResponseParams `json:"Response"`
}

func (r *DescribeReplicationInstanceCreateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstanceCreateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstanceSyncStatusRequestParams struct {
	// Master instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Replication instance ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitnil,omitempty" name:"ReplicationRegistryId"`

	// Region ID of the replication instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`

	// Whether to show the synchronization log
	ShowReplicationLog *bool `json:"ShowReplicationLog,omitnil,omitempty" name:"ShowReplicationLog"`

	// Page offset for log display. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 5, maximum value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeReplicationInstanceSyncStatusRequest struct {
	*tchttp.BaseRequest
	
	// Master instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Replication instance ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitnil,omitempty" name:"ReplicationRegistryId"`

	// Region ID of the replication instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`

	// Whether to show the synchronization log
	ShowReplicationLog *bool `json:"ShowReplicationLog,omitnil,omitempty" name:"ShowReplicationLog"`

	// Page offset for log display. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 5, maximum value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeReplicationInstanceSyncStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstanceSyncStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "ReplicationRegistryId")
	delete(f, "ReplicationRegionId")
	delete(f, "ShowReplicationLog")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationInstanceSyncStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstanceSyncStatusResponseParams struct {
	// Synchronization status
	ReplicationStatus *string `json:"ReplicationStatus,omitnil,omitempty" name:"ReplicationStatus"`

	// Synchronization completion time
	ReplicationTime *string `json:"ReplicationTime,omitnil,omitempty" name:"ReplicationTime"`

	// Synchronization log
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ReplicationLog *ReplicationLog `json:"ReplicationLog,omitnil,omitempty" name:"ReplicationLog"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReplicationInstanceSyncStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplicationInstanceSyncStatusResponseParams `json:"Response"`
}

func (r *DescribeReplicationInstanceSyncStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstanceSyncStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstancesRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 20, maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeReplicationInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: 20, maximum value: 100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeReplicationInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationInstancesResponseParams struct {
	// Total number of instances
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Replication instance list
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ReplicationRegistries []*ReplicationRegistry `json:"ReplicationRegistries,omitnil,omitempty" name:"ReplicationRegistries"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReplicationInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplicationInstancesResponseParams `json:"Response"`
}

func (r *DescribeReplicationInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoriesRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Specified namespace. If this parameter is left empty, image repositories in all namespaces will be queried.
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Specified image repository. If this parameter is left empty, all image repositories in the specified namespace will be queried.
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Page number, which is used for pagination
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page, which is used for pagination
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sort field. Valid values: -creation_time, -name, -update_time.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`
}

type DescribeRepositoriesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Specified namespace. If this parameter is left empty, image repositories in all namespaces will be queried.
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Specified image repository. If this parameter is left empty, all image repositories in the specified namespace will be queried.
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Page number, which is used for pagination
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page, which is used for pagination
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Sort field. Valid values: -creation_time, -name, -update_time.
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`
}

func (r *DescribeRepositoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRepositoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRepositoriesResponseParams struct {
	// Repository information list
	RepositoryList []*TcrRepositoryInfo `json:"RepositoryList,omitnil,omitempty" name:"RepositoryList"`

	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRepositoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRepositoriesResponseParams `json:"Response"`
}

func (r *DescribeRepositoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRepositoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPoliciesRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

type DescribeSecurityPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`
}

func (r *DescribeSecurityPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPoliciesResponseParams struct {
	// Instance security policy group
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecurityPolicySet []*SecurityPolicy `json:"SecurityPolicySet,omitnil,omitempty" name:"SecurityPolicySet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPoliciesResponseParams `json:"Response"`
}

func (r *DescribeSecurityPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceAccountsRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// All service accounts
	All *bool `json:"All,omitnil,omitempty" name:"All"`

	// Whether to fill in permission information.
	EmbedPermission *bool `json:"EmbedPermission,omitnil,omitempty" name:"EmbedPermission"`

	// Filters
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: `0`
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: `20`. Maximum value: `100`. The maximum value is automatically applied when a value exceeding it is entered.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeServiceAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// All service accounts
	All *bool `json:"All,omitnil,omitempty" name:"All"`

	// Whether to fill in permission information.
	EmbedPermission *bool `json:"EmbedPermission,omitnil,omitempty" name:"EmbedPermission"`

	// Filters
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset. Default value: `0`
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of output entries. Default value: `20`. Maximum value: `100`. The maximum value is automatically applied when a value exceeding it is entered.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeServiceAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "All")
	delete(f, "EmbedPermission")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceAccountsResponseParams struct {
	// List of service accounts
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ServiceAccounts []*ServiceAccount `json:"ServiceAccounts,omitnil,omitempty" name:"ServiceAccounts"`

	// Number of service level accounts.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServiceAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceAccountsResponseParams `json:"Response"`
}

func (r *DescribeServiceAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionExecutionRequestParams struct {
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Rule ID
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// `PageSize` for pagination
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTagRetentionExecutionRequest struct {
	*tchttp.BaseRequest
	
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Rule ID
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// `PageSize` for pagination
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTagRetentionExecutionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionExecutionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RetentionId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagRetentionExecutionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionExecutionResponseParams struct {
	// List of tag retention execution records
	RetentionExecutionList []*RetentionExecution `json:"RetentionExecutionList,omitnil,omitempty" name:"RetentionExecutionList"`

	// Total number of tag retention execution records
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTagRetentionExecutionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagRetentionExecutionResponseParams `json:"Response"`
}

func (r *DescribeTagRetentionExecutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionExecutionTaskRequestParams struct {
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Rule ID
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// Rule execution ID
	ExecutionId *int64 `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`

	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// `PageSize` for pagination
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTagRetentionExecutionTaskRequest struct {
	*tchttp.BaseRequest
	
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Rule ID
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// Rule execution ID
	ExecutionId *int64 `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`

	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// `PageSize` for pagination
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTagRetentionExecutionTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionExecutionTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RetentionId")
	delete(f, "ExecutionId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagRetentionExecutionTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionExecutionTaskResponseParams struct {
	// List of tag retention execution tasks
	RetentionTaskList []*RetentionTask `json:"RetentionTaskList,omitnil,omitempty" name:"RetentionTaskList"`

	// Total number of tag retention execution tasks
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTagRetentionExecutionTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagRetentionExecutionTaskResponseParams `json:"Response"`
}

func (r *DescribeTagRetentionExecutionTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionExecutionTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionRulesRequestParams struct {
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// `PageSize` for pagination
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTagRetentionRulesRequest struct {
	*tchttp.BaseRequest
	
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// `PageSize` for pagination
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTagRetentionRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagRetentionRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagRetentionRulesResponseParams struct {
	// List of tag retention policies
	RetentionPolicyList []*RetentionPolicy `json:"RetentionPolicyList,omitnil,omitempty" name:"RetentionPolicyList"`

	// Total number of tag retention policies
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTagRetentionRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagRetentionRulesResponseParams `json:"Response"`
}

func (r *DescribeTagRetentionRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagRetentionRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebhookTriggerLogRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Trigger ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeWebhookTriggerLogRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Trigger ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeWebhookTriggerLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebhookTriggerLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Namespace")
	delete(f, "Id")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebhookTriggerLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebhookTriggerLogResponseParams struct {
	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of logs
	Logs []*WebhookTriggerLog `json:"Logs,omitnil,omitempty" name:"Logs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWebhookTriggerLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebhookTriggerLogResponseParams `json:"Response"`
}

func (r *DescribeWebhookTriggerLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebhookTriggerLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebhookTriggerRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DescribeWebhookTriggerRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DescribeWebhookTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebhookTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebhookTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebhookTriggerResponseParams struct {
	// Total number of triggers
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of triggers
	Triggers []*WebhookTrigger `json:"Triggers,omitnil,omitempty" name:"Triggers"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWebhookTriggerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebhookTriggerResponseParams `json:"Response"`
}

func (r *DescribeWebhookTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadHelmChartRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Helm chart name
	ChartName *string `json:"ChartName,omitnil,omitempty" name:"ChartName"`

	// Helm chart version
	ChartVersion *string `json:"ChartVersion,omitnil,omitempty" name:"ChartVersion"`
}

type DownloadHelmChartRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Helm chart name
	ChartName *string `json:"ChartName,omitnil,omitempty" name:"ChartName"`

	// Helm chart version
	ChartVersion *string `json:"ChartVersion,omitnil,omitempty" name:"ChartVersion"`
}

func (r *DownloadHelmChartRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadHelmChartRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "ChartName")
	delete(f, "ChartVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadHelmChartRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadHelmChartResponseParams struct {
	// Temporary token
	TmpToken *string `json:"TmpToken,omitnil,omitempty" name:"TmpToken"`

	// Temporary `secretId`
	TmpSecretId *string `json:"TmpSecretId,omitnil,omitempty" name:"TmpSecretId"`

	// Temporary `secretKey`
	TmpSecretKey *string `json:"TmpSecretKey,omitnil,omitempty" name:"TmpSecretKey"`

	// Bucket information
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Instance ID
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Chart information
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// Start timestamp
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Token expiration timestamp
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DownloadHelmChartResponse struct {
	*tchttp.BaseResponse
	Response *DownloadHelmChartResponseParams `json:"Response"`
}

func (r *DownloadHelmChartResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadHelmChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DuplicateImageRequestParams struct {
	// Instance ID.
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Source namespace name.
	SourceNamespace *string `json:"SourceNamespace,omitnil,omitempty" name:"SourceNamespace"`

	// Source image repository name.
	SourceRepo *string `json:"SourceRepo,omitnil,omitempty" name:"SourceRepo"`

	// Source image tag or digest value. Currently, only tag is supported.
	SourceReference *string `json:"SourceReference,omitnil,omitempty" name:"SourceReference"`

	// Destination image version.
	DestinationTag *string `json:"DestinationTag,omitnil,omitempty" name:"DestinationTag"`

	// Destination namespace. Defaults to the same as the source if not filled.
	DestinationNamespace *string `json:"DestinationNamespace,omitnil,omitempty" name:"DestinationNamespace"`

	// Destination image repository. Defaults to the same as the source if not filled.
	DestinationRepo *string `json:"DestinationRepo,omitnil,omitempty" name:"DestinationRepo"`

	// Whether to overwrite.
	Override *bool `json:"Override,omitnil,omitempty" name:"Override"`
}

type DuplicateImageRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Source namespace name.
	SourceNamespace *string `json:"SourceNamespace,omitnil,omitempty" name:"SourceNamespace"`

	// Source image repository name.
	SourceRepo *string `json:"SourceRepo,omitnil,omitempty" name:"SourceRepo"`

	// Source image tag or digest value. Currently, only tag is supported.
	SourceReference *string `json:"SourceReference,omitnil,omitempty" name:"SourceReference"`

	// Destination image version.
	DestinationTag *string `json:"DestinationTag,omitnil,omitempty" name:"DestinationTag"`

	// Destination namespace. Defaults to the same as the source if not filled.
	DestinationNamespace *string `json:"DestinationNamespace,omitnil,omitempty" name:"DestinationNamespace"`

	// Destination image repository. Defaults to the same as the source if not filled.
	DestinationRepo *string `json:"DestinationRepo,omitnil,omitempty" name:"DestinationRepo"`

	// Whether to overwrite.
	Override *bool `json:"Override,omitnil,omitempty" name:"Override"`
}

func (r *DuplicateImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DuplicateImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "SourceNamespace")
	delete(f, "SourceRepo")
	delete(f, "SourceReference")
	delete(f, "DestinationTag")
	delete(f, "DestinationNamespace")
	delete(f, "DestinationRepo")
	delete(f, "Override")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DuplicateImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DuplicateImageResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DuplicateImageResponse struct {
	*tchttp.BaseResponse
	Response *DuplicateImageResponseParams `json:"Response"`
}

func (r *DuplicateImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DuplicateImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// Attribute name. If more than one filter exists, the logical relationship between these filters is `AND`.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Attribute value. If multiple values exist in one filter, the logical relationship between these values is `OR`.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type GCJobInfo struct {
	// Job ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Job status
	JobStatus *string `json:"JobStatus,omitnil,omitempty" name:"JobStatus"`

	// Creation time
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Scheduling information
	Schedule *Schedule `json:"Schedule,omitnil,omitempty" name:"Schedule"`
}

type Header struct {
	// Header Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Header Values
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type ImmutableTagRule struct {
	// Repository matching rule
	RepositoryPattern *string `json:"RepositoryPattern,omitnil,omitempty" name:"RepositoryPattern"`

	// Tag matching rule
	TagPattern *string `json:"TagPattern,omitnil,omitempty" name:"TagPattern"`

	// repoMatches or repoExcludes
	RepositoryDecoration *string `json:"RepositoryDecoration,omitnil,omitempty" name:"RepositoryDecoration"`

	// matches or excludes
	TagDecoration *string `json:"TagDecoration,omitnil,omitempty" name:"TagDecoration"`

	// Disabling rule
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`

	// Rule ID
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Namespace
	NsName *string `json:"NsName,omitnil,omitempty" name:"NsName"`
}

type KeyValueString struct {
	// Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ManageExternalEndpointRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Operation. Valid values: Create, Delete.
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

type ManageExternalEndpointRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Operation. Valid values: Create, Delete.
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

func (r *ManageExternalEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageExternalEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Operation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageExternalEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageExternalEndpointResponseParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ManageExternalEndpointResponse struct {
	*tchttp.BaseResponse
	Response *ManageExternalEndpointResponseParams `json:"Response"`
}

func (r *ManageExternalEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageExternalEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageInternalEndpointRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Create/Delete
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// ID of the VPC to be connected to
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the subnet to be connected to
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// ID of the requested region, which is used as the region of the replica instance
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Name of the requested region, which is used as the region of the replica instance
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type ManageInternalEndpointRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Create/Delete
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// ID of the VPC to be connected to
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// ID of the subnet to be connected to
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// ID of the requested region, which is used as the region of the replica instance
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Name of the requested region, which is used as the region of the replica instance
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

func (r *ManageInternalEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageInternalEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Operation")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "RegionId")
	delete(f, "RegionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageInternalEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageInternalEndpointResponseParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ManageInternalEndpointResponse struct {
	*tchttp.BaseResponse
	Response *ManageInternalEndpointResponseParams `json:"Response"`
}

func (r *ManageInternalEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageInternalEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageReplicationRequestParams struct {
	// Source instance ID
	SourceRegistryId *string `json:"SourceRegistryId,omitnil,omitempty" name:"SourceRegistryId"`

	// Destination instance ID
	DestinationRegistryId *string `json:"DestinationRegistryId,omitnil,omitempty" name:"DestinationRegistryId"`

	// Synchronization rule
	Rule *ReplicationRule `json:"Rule,omitnil,omitempty" name:"Rule"`

	// Rule description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Region ID of the destination instance. For example, `1` represents Guangzhou
	DestinationRegionId *uint64 `json:"DestinationRegionId,omitnil,omitempty" name:"DestinationRegionId"`

	// Configuration of the synchronization rule
	PeerReplicationOption *PeerReplicationOption `json:"PeerReplicationOption,omitnil,omitempty" name:"PeerReplicationOption"`
}

type ManageReplicationRequest struct {
	*tchttp.BaseRequest
	
	// Source instance ID
	SourceRegistryId *string `json:"SourceRegistryId,omitnil,omitempty" name:"SourceRegistryId"`

	// Destination instance ID
	DestinationRegistryId *string `json:"DestinationRegistryId,omitnil,omitempty" name:"DestinationRegistryId"`

	// Synchronization rule
	Rule *ReplicationRule `json:"Rule,omitnil,omitempty" name:"Rule"`

	// Rule description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Region ID of the destination instance. For example, `1` represents Guangzhou
	DestinationRegionId *uint64 `json:"DestinationRegionId,omitnil,omitempty" name:"DestinationRegionId"`

	// Configuration of the synchronization rule
	PeerReplicationOption *PeerReplicationOption `json:"PeerReplicationOption,omitnil,omitempty" name:"PeerReplicationOption"`
}

func (r *ManageReplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageReplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceRegistryId")
	delete(f, "DestinationRegistryId")
	delete(f, "Rule")
	delete(f, "Description")
	delete(f, "DestinationRegionId")
	delete(f, "PeerReplicationOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageReplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageReplicationResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ManageReplicationResponse struct {
	*tchttp.BaseResponse
	Response *ManageReplicationResponseParams `json:"Response"`
}

func (r *ManageReplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageReplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImmutableTagRulesRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Rule ID
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Rule
	Rule *ImmutableTagRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type ModifyImmutableTagRulesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Rule ID
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Rule
	Rule *ImmutableTagRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

func (r *ModifyImmutableTagRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImmutableTagRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RuleId")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyImmutableTagRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyImmutableTagRulesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyImmutableTagRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyImmutableTagRulesResponseParams `json:"Response"`
}

func (r *ModifyImmutableTagRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyImmutableTagRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Instance edition
	// Basic: `basic`
	// Standard: `standard`
	// Premium: `premium`
	RegistryType *string `json:"RegistryType,omitnil,omitempty" name:"RegistryType"`

	// Whether to enable deletion protection. It defaults to `false`. 
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Instance edition
	// Basic: `basic`
	// Standard: `standard`
	// Premium: `premium`
	RegistryType *string `json:"RegistryType,omitnil,omitempty" name:"RegistryType"`

	// Whether to enable deletion protection. It defaults to `false`. 
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RegistryType")
	delete(f, "DeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceTokenRequestParams struct {
	// ID of the long-term access credential of the instance
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Whether to enable the long-term access credential of the instance
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Access credential description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Valid values: 1: Modify the description; 2: Enable/Disable. Default value: 2.
	ModifyFlag *int64 `json:"ModifyFlag,omitnil,omitempty" name:"ModifyFlag"`
}

type ModifyInstanceTokenRequest struct {
	*tchttp.BaseRequest
	
	// ID of the long-term access credential of the instance
	TokenId *string `json:"TokenId,omitnil,omitempty" name:"TokenId"`

	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Whether to enable the long-term access credential of the instance
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Access credential description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Valid values: 1: Modify the description; 2: Enable/Disable. Default value: 2.
	ModifyFlag *int64 `json:"ModifyFlag,omitnil,omitempty" name:"ModifyFlag"`
}

func (r *ModifyInstanceTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TokenId")
	delete(f, "RegistryId")
	delete(f, "Enable")
	delete(f, "Desc")
	delete(f, "ModifyFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceTokenResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceTokenResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceTokenResponseParams `json:"Response"`
}

func (r *ModifyInstanceTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNamespaceRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Access level. Valid values: True: Public; False: Private.
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`
}

type ModifyNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Access level. Valid values: True: Public; False: Private.
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`
}

func (r *ModifyNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "IsPublic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNamespaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNamespaceResponseParams `json:"Response"`
}

func (r *ModifyNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRepositoryRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Image repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Brief repository description
	BriefDescription *string `json:"BriefDescription,omitnil,omitempty" name:"BriefDescription"`

	// Detailed repository description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyRepositoryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Image repository name
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// Brief repository description
	BriefDescription *string `json:"BriefDescription,omitnil,omitempty" name:"BriefDescription"`

	// Detailed repository description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyRepositoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRepositoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceName")
	delete(f, "RepositoryName")
	delete(f, "BriefDescription")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRepositoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRepositoryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRepositoryResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRepositoryResponseParams `json:"Response"`
}

func (r *ModifyRepositoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// PolicyId
	PolicyIndex *int64 `json:"PolicyIndex,omitnil,omitempty" name:"PolicyIndex"`

	// Allowed IP, such as `192.168.0.0/24`
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifySecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// PolicyId
	PolicyIndex *int64 `json:"PolicyIndex,omitnil,omitempty" name:"PolicyIndex"`

	// Allowed IP, such as `192.168.0.0/24`
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifySecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "PolicyIndex")
	delete(f, "CidrBlock")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyResponseParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityPolicyResponseParams `json:"Response"`
}

func (r *ModifySecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceAccountPasswordRequestParams struct {
	// Instance ID.
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Service level account name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Whether to randomly generate a password.
	Random *bool `json:"Random,omitnil,omitempty" name:"Random"`

	// Service level account password, 8 to 20 characters, contains at least one uppercase letter, one lowercase letter, and one number.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type ModifyServiceAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Service level account name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Whether to randomly generate a password.
	Random *bool `json:"Random,omitnil,omitempty" name:"Random"`

	// Service level account password, 8 to 20 characters, contains at least one uppercase letter, one lowercase letter, and one number.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *ModifyServiceAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Name")
	delete(f, "Random")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceAccountPasswordResponseParams struct {
	// Custom user password, displayed only once. Please make sure to retain it securely.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyServiceAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceAccountPasswordResponseParams `json:"Response"`
}

func (r *ModifyServiceAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceAccountRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Service account name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Service account description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Validity in days starting from the current day, It takes a higher priority than `ExpiresAt`.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Expiry time (timestamp, in milliseconds)
	ExpiresAt *int64 `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// Whether to disable the service account
	Disable *bool `json:"Disable,omitnil,omitempty" name:"Disable"`

	// Policy list
	Permissions []*Permission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

type ModifyServiceAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Service account name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Service account description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Validity in days starting from the current day, It takes a higher priority than `ExpiresAt`.
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Expiry time (timestamp, in milliseconds)
	ExpiresAt *int64 `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// Whether to disable the service account
	Disable *bool `json:"Disable,omitnil,omitempty" name:"Disable"`

	// Policy list
	Permissions []*Permission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

func (r *ModifyServiceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Duration")
	delete(f, "ExpiresAt")
	delete(f, "Disable")
	delete(f, "Permissions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyServiceAccountResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceAccountResponseParams `json:"Response"`
}

func (r *ModifyServiceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTagRetentionRuleRequestParams struct {
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// ID of the original namespace
	NamespaceId *int64 `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Retention policy
	RetentionRule *RetentionRule `json:"RetentionRule,omitnil,omitempty" name:"RetentionRule"`

	// Original execution cycle
	CronSetting *string `json:"CronSetting,omitnil,omitempty" name:"CronSetting"`

	// Rule ID
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// Whether to disable the rule
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`
}

type ModifyTagRetentionRuleRequest struct {
	*tchttp.BaseRequest
	
	// Primary instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// ID of the original namespace
	NamespaceId *int64 `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Retention policy
	RetentionRule *RetentionRule `json:"RetentionRule,omitnil,omitempty" name:"RetentionRule"`

	// Original execution cycle
	CronSetting *string `json:"CronSetting,omitnil,omitempty" name:"CronSetting"`

	// Rule ID
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// Whether to disable the rule
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`
}

func (r *ModifyTagRetentionRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTagRetentionRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "NamespaceId")
	delete(f, "RetentionRule")
	delete(f, "CronSetting")
	delete(f, "RetentionId")
	delete(f, "Disabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTagRetentionRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTagRetentionRuleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTagRetentionRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTagRetentionRuleResponseParams `json:"Response"`
}

func (r *ModifyTagRetentionRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTagRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebhookTriggerRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Trigger parameter
	Trigger *WebhookTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type ModifyWebhookTriggerRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Trigger parameter
	Trigger *WebhookTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// Namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *ModifyWebhookTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebhookTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "Trigger")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWebhookTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebhookTriggerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyWebhookTriggerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWebhookTriggerResponseParams `json:"Response"`
}

func (r *ModifyWebhookTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PeerReplicationOption struct {
	// UIN of the destination instance
	PeerRegistryUin *string `json:"PeerRegistryUin,omitnil,omitempty" name:"PeerRegistryUin"`

	// Permanent access Token for the destination instance
	PeerRegistryToken *string `json:"PeerRegistryToken,omitnil,omitempty" name:"PeerRegistryToken"`

	// Whether to enable cross-account synchronization
	EnablePeerReplication *bool `json:"EnablePeerReplication,omitnil,omitempty" name:"EnablePeerReplication"`
}

type Permission struct {
	// Resource path. Valid value: `Namespace`
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Actions. Currently only support: tcr:PushRepository, tcr:PullRepository, tcr:CreateRepository, tcr:CreateHelmChart, tcr:DescribeHelmCharts.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Actions []*string `json:"Actions,omitnil,omitempty" name:"Actions"`
}

type Region struct {
	// gz
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 1
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// ap-guangzhou
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// alluser
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// remark
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Creation time
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Update time
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type Registry struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Instance name
	RegistryName *string `json:"RegistryName,omitnil,omitempty" name:"RegistryName"`

	// Instance specification
	RegistryType *string `json:"RegistryType,omitnil,omitempty" name:"RegistryType"`

	// Instance status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Public access URL of the instance
	PublicDomain *string `json:"PublicDomain,omitnil,omitempty" name:"PublicDomain"`

	// Instance creation time
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Region name
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Region ID
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Whether to enable anonymity
	EnableAnonymous *bool `json:"EnableAnonymous,omitnil,omitempty" name:"EnableAnonymous"`

	// Token validity period
	TokenValidTime *uint64 `json:"TokenValidTime,omitnil,omitempty" name:"TokenValidTime"`

	// Internal access address of the instance
	InternalEndpoint *string `json:"InternalEndpoint,omitnil,omitempty" name:"InternalEndpoint"`

	// Cloud tag of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// Instance expiration time (for prepayment)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpiredAt *string `json:"ExpiredAt,omitnil,omitempty" name:"ExpiredAt"`

	// Instance billing mode. Valid values: 0: Postpayment; 1: Prepayment.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMod *int64 `json:"PayMod,omitnil,omitempty" name:"PayMod"`

	// Prepayment renewal flag. Valid values: 0: Manual renewal; 1: Auto-renewal; 2: No renewal and no notification.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Whether to enable instance deletion protection. false indicates it is not enabled.
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

type RegistryChargePrepaid struct {
	// Instance purchase duration in months
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Auto-renewal flag. Valid values: 0: Manual renewal; 1: Auto-renewal; 2: No renewal and no notification.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type RegistryCondition struct {
	// Instance creation process type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Instance creation process status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Reasons for transiting to the process
	// Note: This field may return null, indicating that no valid values can be obtained.
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type RegistryStatus struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Instance status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Additional status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Conditions []*RegistryCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

// Predefined struct for user
type RenewInstanceRequestParams struct {
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Auto-renewal flag and purchase duration in months for prepayment. Valid values: 0: Manual renewal; 1: Auto-renewal; 2: No renewal and no notification.
	RegistryChargePrepaid *RegistryChargePrepaid `json:"RegistryChargePrepaid,omitnil,omitempty" name:"RegistryChargePrepaid"`

	// Valid values: 0: renewal; 1: change from pay-as-you-go to monthly subscription billing
	Flag *int64 `json:"Flag,omitnil,omitempty" name:"Flag"`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Auto-renewal flag and purchase duration in months for prepayment. Valid values: 0: Manual renewal; 1: Auto-renewal; 2: No renewal and no notification.
	RegistryChargePrepaid *RegistryChargePrepaid `json:"RegistryChargePrepaid,omitnil,omitempty" name:"RegistryChargePrepaid"`

	// Valid values: 0: renewal; 1: change from pay-as-you-go to monthly subscription billing
	Flag *int64 `json:"Flag,omitnil,omitempty" name:"Flag"`
}

func (r *RenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistryId")
	delete(f, "RegistryChargePrepaid")
	delete(f, "Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceResponseParams struct {
	// Enterprise Edition instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewInstanceResponseParams `json:"Response"`
}

func (r *RenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicationFilter struct {
	// Type (`name`, `tag` and `resource`)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// It is left blank by default
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ReplicationLog struct {
	// Resource type
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Path of the source resource
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Path of the destination resource
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Destination *string `json:"Destination,omitnil,omitempty" name:"Destination"`

	// Synchronization status
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Start time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ReplicationRegistry struct {
	// Master instance ID
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Replication instance ID
	ReplicationRegistryId *string `json:"ReplicationRegistryId,omitnil,omitempty" name:"ReplicationRegistryId"`

	// Region ID of the replication instance
	ReplicationRegionId *uint64 `json:"ReplicationRegionId,omitnil,omitempty" name:"ReplicationRegionId"`

	// Region name of the replication instance
	ReplicationRegionName *string `json:"ReplicationRegionName,omitnil,omitempty" name:"ReplicationRegionName"`

	// Status of the replication instance
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Creation time
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`
}

type ReplicationRule struct {
	// Name of synchronization rule
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Destination namespace
	DestNamespace *string `json:"DestNamespace,omitnil,omitempty" name:"DestNamespace"`

	// Whether to override
	Override *bool `json:"Override,omitnil,omitempty" name:"Override"`

	// Synchronization filters
	Filters []*ReplicationFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type RetentionExecution struct {
	// Execution ID
	ExecutionId *int64 `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`

	// Rule ID
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// Execution start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Execution end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Execution status. Valid values: Failed, Succeed, Stopped, InProgress.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type RetentionPolicy struct {
	// Tag retention policy ID
	RetentionId *int64 `json:"RetentionId,omitnil,omitempty" name:"RetentionId"`

	// Namespace name
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// List of rules
	RetentionRuleList []*RetentionRule `json:"RetentionRuleList,omitnil,omitempty" name:"RetentionRuleList"`

	// Regular execution mode
	CronSetting *string `json:"CronSetting,omitnil,omitempty" name:"CronSetting"`

	// Whether to enable the rule
	Disabled *bool `json:"Disabled,omitnil,omitempty" name:"Disabled"`

	// The execution time of the next task based on the current time and `cronSetting`, which is for reference only
	NextExecutionTime *string `json:"NextExecutionTime,omitnil,omitempty" name:"NextExecutionTime"`
}

type RetentionRule struct {
	// Supported policy. Valid values: latestPushedK: Retain the latest specified number of pushed tags; nDaysSinceLastPush: Retain the tags pushed in the past specified number of days.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Rule value
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type RetentionTask struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Rule execution ID
	ExecutionId *int64 `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`

	// Task start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Task execution status. Valid values: Failed, Succeed, Stopped, InProgress.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Total number of tags
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Number of retained tags
	Retained *int64 `json:"Retained,omitnil,omitempty" name:"Retained"`

	// Application repository
	Repository *string `json:"Repository,omitnil,omitempty" name:"Repository"`
}

type Schedule struct {
	// Type. Valid values: Hourly, Daily, Weekly, Custom, Manual, Dryrun, None.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type SecurityPolicy struct {
	// Policy index
	PolicyIndex *int64 `json:"PolicyIndex,omitnil,omitempty" name:"PolicyIndex"`

	// Remarks
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// The public network IP address of the access source
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// The version of the security policy
	PolicyVersion *string `json:"PolicyVersion,omitnil,omitempty" name:"PolicyVersion"`
}

type ServiceAccount struct {
	// Service account name
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Whether to disable
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Disable *bool `json:"Disable,omitnil,omitempty" name:"Disable"`

	// Expiry time
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ExpiresAt *int64 `json:"ExpiresAt,omitnil,omitempty" name:"ExpiresAt"`

	// Creation time
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Update time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Policy
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Permissions []*Permission `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

type Tag struct {
	// Cloud tag key.Note: This field may return null, indicating that no valid value can be obtained.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Cloud tag value.Note: This field may return null, indicating that no valid value can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TagSpecification struct {
	// Default value: instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Cloud tag array
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type TaskDetail struct {
	// Task
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Task UUID
	TaskUUID *string `json:"TaskUUID,omitnil,omitempty" name:"TaskUUID"`

	// Task status
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Task details
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TaskMessage *string `json:"TaskMessage,omitnil,omitempty" name:"TaskMessage"`

	// Start time of the task
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// End time of the task
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FinishedTime *string `json:"FinishedTime,omitnil,omitempty" name:"FinishedTime"`
}

type TcrImageInfo struct {
	// Hash value
	Digest *string `json:"Digest,omitnil,omitempty" name:"Digest"`

	// Image size in bytes
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Tag name
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Artifact type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// KMS signature information
	// Note: This field may return null, indicating that no valid values can be obtained.
	KmsSignature *string `json:"KmsSignature,omitnil,omitempty" name:"KmsSignature"`
}

type TcrInstanceToken struct {
	// Token ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Token description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// ID of the instance of the token
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// Token status
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// Token creation time
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Token expiration timestamp
	ExpiredAt *int64 `json:"ExpiredAt,omitnil,omitempty" name:"ExpiredAt"`
}

type TcrNamespaceInfo struct {
	// Namespace name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Creation time
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// Access level
	Public *bool `json:"Public,omitnil,omitempty" name:"Public"`

	// Namespace ID
	NamespaceId *int64 `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// Cloud tag of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// Namespace metadata
	// Note: This field may return null, indicating that no valid values can be obtained.
	Metadata []*KeyValueString `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

type TcrRepositoryInfo struct {
	// Repository name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Namespace name
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Creation time, such as `2006-01-02 15:04:05.999999999 -0700 MST`
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// Whether to make public
	Public *bool `json:"Public,omitnil,omitempty" name:"Public"`

	// Detailed repository description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Brief description
	// Note: This field may return null, indicating that no valid values can be obtained.
	BriefDescription *string `json:"BriefDescription,omitnil,omitempty" name:"BriefDescription"`

	// Update time, such as `2006-01-02 15:04:05.999999999 -0700 MST`
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type WebhookTarget struct {
	// Target address
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Custom headers
	Headers []*Header `json:"Headers,omitnil,omitempty" name:"Headers"`
}

type WebhookTrigger struct {
	// Trigger name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Trigger target
	Targets []*WebhookTarget `json:"Targets,omitnil,omitempty" name:"Targets"`

	// Action to be triggered
	EventTypes []*string `json:"EventTypes,omitnil,omitempty" name:"EventTypes"`

	// Trigger rule
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Whether to enable the trigger
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// Trigger ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Trigger description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// ID of the namespace of the trigger
	NamespaceId *int64 `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`
}

type WebhookTriggerLog struct {
	// Log ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Trigger ID
	TriggerId *int64 `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// Event type
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// Notification type
	NotifyType *string `json:"NotifyType,omitnil,omitempty" name:"NotifyType"`

	// Details
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Creation time
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// Update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}