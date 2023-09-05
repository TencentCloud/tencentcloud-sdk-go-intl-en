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

package v20210519

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AddLabelRequestParams struct {
	// The label ID.
	LabelId *uint64 `json:"LabelId,omitnil" name:"LabelId"`


	Did *string `json:"Did,omitnil" name:"Did"`
}

type AddLabelRequest struct {
	*tchttp.BaseRequest
	
	// The label ID.
	LabelId *uint64 `json:"LabelId,omitnil" name:"LabelId"`

	Did *string `json:"Did,omitnil" name:"Did"`
}

func (r *AddLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LabelId")
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddLabelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddLabelResponse struct {
	*tchttp.BaseResponse
	Response *AddLabelResponseParams `json:"Response"`
}

func (r *AddLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Authority struct {
	// The authority ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// The DID.
	DidId *uint64 `json:"DidId,omitnil" name:"DidId"`

	// The details of the DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// The authority name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Whether the authority is certified. `1`: Yes; `2`: No.
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// The DID service ID.
	DidServiceId *uint64 `json:"DidServiceId,omitnil" name:"DidServiceId"`

	// The application ID.
	ContractAppId *uint64 `json:"ContractAppId,omitnil" name:"ContractAppId"`

	// Remarks
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// The registration time.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RegisterTime *string `json:"RegisterTime,omitnil" name:"RegisterTime"`

	// The recognition time.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RecognizeTime *string `json:"RecognizeTime,omitnil" name:"RecognizeTime"`

	// The creation time.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// The last updated time.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The application name.
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// The on-chain label.
	LabelName *string `json:"LabelName,omitnil" name:"LabelName"`
}

type BcosClusterItem struct {
	// The network ID.
	ChainId *uint64 `json:"ChainId,omitnil" name:"ChainId"`

	// The network name.
	ChainName *string `json:"ChainName,omitnil" name:"ChainName"`

	// The number of organizations.
	AgencyCount *uint64 `json:"AgencyCount,omitnil" name:"AgencyCount"`

	// The consortium ID.
	ConsortiumId *uint64 `json:"ConsortiumId,omitnil" name:"ConsortiumId"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// The expiration time.
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// The network status.
	ChainStatus *uint64 `json:"ChainStatus,omitnil" name:"ChainStatus"`

	// The resource ID.
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// The cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The consortium name.
	ConsortiumName *string `json:"ConsortiumName,omitnil" name:"ConsortiumName"`

	// The organization ID.
	AgencyId *uint64 `json:"AgencyId,omitnil" name:"AgencyId"`

	// The renewal status.
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// The network mode.
	TotalNetworkNode *uint64 `json:"TotalNetworkNode,omitnil" name:"TotalNetworkNode"`

	// The number of nodes created.
	TotalCreateNode *uint64 `json:"TotalCreateNode,omitnil" name:"TotalCreateNode"`

	// The total number of groups.
	TotalGroups *uint64 `json:"TotalGroups,omitnil" name:"TotalGroups"`
}

// Predefined struct for user
type CancelAuthorityIssuerRequestParams struct {
	// The details of the DID.
	Did *string `json:"Did,omitnil" name:"Did"`
}

type CancelAuthorityIssuerRequest struct {
	*tchttp.BaseRequest
	
	// The details of the DID.
	Did *string `json:"Did,omitnil" name:"Did"`
}

func (r *CancelAuthorityIssuerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAuthorityIssuerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelAuthorityIssuerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelAuthorityIssuerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelAuthorityIssuerResponse struct {
	*tchttp.BaseResponse
	Response *CancelAuthorityIssuerResponseParams `json:"Response"`
}

func (r *CancelAuthorityIssuerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAuthorityIssuerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckChainRequestParams struct {
	// The group ID.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The name of the DID organization.
	AgencyName *string `json:"AgencyName,omitnil" name:"AgencyName"`
}

type CheckChainRequest struct {
	*tchttp.BaseRequest
	
	// The group ID.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The name of the DID organization.
	AgencyName *string `json:"AgencyName,omitnil" name:"AgencyName"`
}

func (r *CheckChainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckChainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ClusterId")
	delete(f, "AgencyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckChainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckChainResponseParams struct {
	// Whether you are the owner of the consortium. `1`: Yes; `0`: No.
	RoleType *int64 `json:"RoleType,omitnil" name:"RoleType"`

	// The chain ID.
	ChainId *string `json:"ChainId,omitnil" name:"ChainId"`

	// The application name.
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckChainResponse struct {
	*tchttp.BaseResponse
	Response *CheckChainResponseParams `json:"Response"`
}

func (r *CheckChainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckChainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDidDeployRequestParams struct {
	// The task ID.
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`
}

type CheckDidDeployRequest struct {
	*tchttp.BaseRequest
	
	// The task ID.
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *CheckDidDeployRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDidDeployRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckDidDeployRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDidDeployResponseParams struct {
	// The service information.
	Task *Task `json:"Task,omitnil" name:"Task"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckDidDeployResponse struct {
	*tchttp.BaseResponse
	Response *CheckDidDeployResponseParams `json:"Response"`
}

func (r *CheckDidDeployResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDidDeployResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsortiumItem struct {
	// The consortium ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// The consortium name.
	Name *string `json:"Name,omitnil" name:"Name"`
}

type Contract struct {
	// The application name.
	ApplyName *string `json:"ApplyName,omitnil" name:"ApplyName"`

	// The contract status. `true`: Enabled; `false`: Not enabled.
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// The CNS address of the contract.
	Hash *string `json:"Hash,omitnil" name:"Hash"`

	// The desensitized CNS address.
	HashShow *string `json:"HashShow,omitnil" name:"HashShow"`

	// The DID of the organization that deployed the contract.
	WeId *string `json:"WeId,omitnil" name:"WeId"`

	// The name of the organization that deployed the contract.
	DeployName *string `json:"DeployName,omitnil" name:"DeployName"`

	// The group.
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// The deployment time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type CptIssueRank struct {
	// The (claim protocol type) CPT name.
	CptName *string `json:"CptName,omitnil" name:"CptName"`

	// The ranking.
	Rank *int64 `json:"Rank,omitnil" name:"Rank"`

	// The number of credentials issued.
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// The application name.
	ApplyName *string `json:"ApplyName,omitnil" name:"ApplyName"`

	// The application ID.
	ApplyId *uint64 `json:"ApplyId,omitnil" name:"ApplyId"`
}

type CptListData struct {
	// The CPT ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// The name of the claim protocol type (CPT).
	Name *string `json:"Name,omitnil" name:"Name"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The service ID.
	ServiceId *uint64 `json:"ServiceId,omitnil" name:"ServiceId"`

	// The application ID of the contract.
	ContractAppId *uint64 `json:"ContractAppId,omitnil" name:"ContractAppId"`

	// The claim protocol type (CPT) ID.
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`

	// The type. `1`: System CPTs; `2`: Authorization CPTs; `3`: General CPTs
	CptType *uint64 `json:"CptType,omitnil" name:"CptType"`

	// The description of the claim protocol type (CPT).
	Description *string `json:"Description,omitnil" name:"Description"`

	// The JSON file of the claim protocol type (CPT).
	CptJson *string `json:"CptJson,omitnil" name:"CptJson"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// The last updated time.
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// The DID of the creator.
	CreatorDid *string `json:"CreatorDid,omitnil" name:"CreatorDid"`

	// The application name.
	AppName *string `json:"AppName,omitnil" name:"AppName"`
}

// Predefined struct for user
type CreateCredentialRequestParams struct {
	// A parameter set. For details, see the example.
	FunctionArg *FunctionArg `json:"FunctionArg,omitnil" name:"FunctionArg"`

	// A parameter set. For details, see the example.
	TransactionArg *TransactionArg `json:"TransactionArg,omitnil" name:"TransactionArg"`

	// The version.
	VersionCredential *string `json:"VersionCredential,omitnil" name:"VersionCredential"`

	// Whether the credential is unsigned.
	UnSigned *bool `json:"UnSigned,omitnil" name:"UnSigned"`
}

type CreateCredentialRequest struct {
	*tchttp.BaseRequest
	
	// A parameter set. For details, see the example.
	FunctionArg *FunctionArg `json:"FunctionArg,omitnil" name:"FunctionArg"`

	// A parameter set. For details, see the example.
	TransactionArg *TransactionArg `json:"TransactionArg,omitnil" name:"TransactionArg"`

	// The version.
	VersionCredential *string `json:"VersionCredential,omitnil" name:"VersionCredential"`

	// Whether the credential is unsigned.
	UnSigned *bool `json:"UnSigned,omitnil" name:"UnSigned"`
}

func (r *CreateCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionArg")
	delete(f, "TransactionArg")
	delete(f, "VersionCredential")
	delete(f, "UnSigned")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCredentialResponseParams struct {
	// The information of the credential.
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateCredentialResponseParams `json:"Response"`
}

func (r *CreateCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDidServiceRequestParams struct {
	// The consortium name.
	ConsortiumName *string `json:"ConsortiumName,omitnil" name:"ConsortiumName"`

	// The consortium ID.
	ConsortiumId *int64 `json:"ConsortiumId,omitnil" name:"ConsortiumId"`

	// The group ID.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// The organization name.
	AgencyName *string `json:"AgencyName,omitnil" name:"AgencyName"`

	// The application name.
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group name.
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`
}

type CreateDidServiceRequest struct {
	*tchttp.BaseRequest
	
	// The consortium name.
	ConsortiumName *string `json:"ConsortiumName,omitnil" name:"ConsortiumName"`

	// The consortium ID.
	ConsortiumId *int64 `json:"ConsortiumId,omitnil" name:"ConsortiumId"`

	// The group ID.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// The organization name.
	AgencyName *string `json:"AgencyName,omitnil" name:"AgencyName"`

	// The application name.
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group name.
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`
}

func (r *CreateDidServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDidServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConsortiumName")
	delete(f, "ConsortiumId")
	delete(f, "GroupId")
	delete(f, "AgencyName")
	delete(f, "AppName")
	delete(f, "ClusterId")
	delete(f, "GroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDidServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDidServiceResponseParams struct {
	// The service information.
	Task *Task `json:"Task,omitnil" name:"Task"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDidServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDidServiceResponseParams `json:"Response"`
}

func (r *CreateDidServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDidServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLabelRequestParams struct {
	// The label name.
	LabelName *string `json:"LabelName,omitnil" name:"LabelName"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`
}

type CreateLabelRequest struct {
	*tchttp.BaseRequest
	
	// The label name.
	LabelName *string `json:"LabelName,omitnil" name:"LabelName"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`
}

func (r *CreateLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LabelName")
	delete(f, "ClusterId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLabelResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateLabelResponse struct {
	*tchttp.BaseResponse
	Response *CreateLabelResponseParams `json:"Response"`
}

func (r *CreateLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSelectiveCredentialRequestParams struct {
	// A parameter set.
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitnil" name:"FunctionArg"`

	// The disclosure policy ID.
	PolicyId *uint64 `json:"PolicyId,omitnil" name:"PolicyId"`
}

type CreateSelectiveCredentialRequest struct {
	*tchttp.BaseRequest
	
	// A parameter set.
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitnil" name:"FunctionArg"`

	// The disclosure policy ID.
	PolicyId *uint64 `json:"PolicyId,omitnil" name:"PolicyId"`
}

func (r *CreateSelectiveCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSelectiveCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionArg")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSelectiveCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSelectiveCredentialResponseParams struct {
	// The credential string.
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSelectiveCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateSelectiveCredentialResponseParams `json:"Response"`
}

func (r *CreateSelectiveCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSelectiveCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPrivateKeyRequestParams struct {
	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The private key.
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`
}

type CreateTDidByPrivateKeyRequest struct {
	*tchttp.BaseRequest
	
	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The private key.
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`
}

func (r *CreateTDidByPrivateKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPrivateKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "PrivateKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByPrivateKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPrivateKeyResponseParams struct {
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTDidByPrivateKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidByPrivateKeyResponseParams `json:"Response"`
}

func (r *CreateTDidByPrivateKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPrivateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPublicKeyRequestParams struct {
	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The public key.
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// The encrypted public key.
	EncryptPubKey *string `json:"EncryptPubKey,omitnil" name:"EncryptPubKey"`
}

type CreateTDidByPublicKeyRequest struct {
	*tchttp.BaseRequest
	
	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The public key.
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// The encrypted public key.
	EncryptPubKey *string `json:"EncryptPubKey,omitnil" name:"EncryptPubKey"`
}

func (r *CreateTDidByPublicKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPublicKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "PublicKey")
	delete(f, "EncryptPubKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByPublicKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPublicKeyResponseParams struct {
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTDidByPublicKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidByPublicKeyResponseParams `json:"Response"`
}

func (r *CreateTDidByPublicKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidRequestParams struct {
	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	//  
	Relegation *uint64 `json:"Relegation,omitnil" name:"Relegation"`
}

type CreateTDidRequest struct {
	*tchttp.BaseRequest
	
	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	//  
	Relegation *uint64 `json:"Relegation,omitnil" name:"Relegation"`
}

func (r *CreateTDidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ClusterId")
	delete(f, "Relegation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidResponseParams struct {
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTDidResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidResponseParams `json:"Response"`
}

func (r *CreateTDidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CredentialStatus struct {
	// The credential ID.
	CredentialId *string `json:"CredentialId,omitnil" name:"CredentialId"`

	// The credential status. `0`: Invalid; `1`: Valid.
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// The DID of the issuer.
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// A summary of the credential.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Digest *string `json:"Digest,omitnil" name:"Digest"`

	// The credential signature.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Signature *string `json:"Signature,omitnil" name:"Signature"`

	// The last updated timestamp.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TimeStamp *uint64 `json:"TimeStamp,omitnil" name:"TimeStamp"`
}

// Predefined struct for user
type DeployByNameRequestParams struct {
	// The application name.
	ApplicationName *string `json:"ApplicationName,omitnil" name:"ApplicationName"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`
}

type DeployByNameRequest struct {
	*tchttp.BaseRequest
	
	// The application name.
	ApplicationName *string `json:"ApplicationName,omitnil" name:"ApplicationName"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`
}

func (r *DeployByNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployByNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationName")
	delete(f, "ClusterId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployByNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployByNameResponseParams struct {
	// The hash value.
	Hash *string `json:"Hash,omitnil" name:"Hash"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeployByNameResponse struct {
	*tchttp.BaseResponse
	Response *DeployByNameResponseParams `json:"Response"`
}

func (r *DeployByNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployByNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DidCluster struct {
	// The chain ID.
	ChainId *uint64 `json:"ChainId,omitnil" name:"ChainId"`

	// The chain name.
	ChainName *string `json:"ChainName,omitnil" name:"ChainName"`

	// The number of organizations.
	AgencyCount *uint64 `json:"AgencyCount,omitnil" name:"AgencyCount"`

	// The consortium ID.
	ConsortiumId *uint64 `json:"ConsortiumId,omitnil" name:"ConsortiumId"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// The expiration time.
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// The network status.
	ChainStatus *uint64 `json:"ChainStatus,omitnil" name:"ChainStatus"`

	// The resource ID.
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The consortium name.
	ConsortiumName *string `json:"ConsortiumName,omitnil" name:"ConsortiumName"`

	// The organization ID.
	AgencyId *uint64 `json:"AgencyId,omitnil" name:"AgencyId"`

	// Whether auto-renewal is enabled.
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// The total number of network nodes.
	TotalNetworkNode *uint64 `json:"TotalNetworkNode,omitnil" name:"TotalNetworkNode"`

	// The number of nodes of the current organization.
	TotalCreateNode *uint64 `json:"TotalCreateNode,omitnil" name:"TotalCreateNode"`

	// The total number of groups.
	TotalGroups *uint64 `json:"TotalGroups,omitnil" name:"TotalGroups"`

	// The total number of DIDs.
	DidCount *uint64 `json:"DidCount,omitnil" name:"DidCount"`
}

type DidData struct {
	// The service ID.
	ServiceId *uint64 `json:"ServiceId,omitnil" name:"ServiceId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The application name.
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// Remarks
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// The status of the authority. `1`: Not registered; `2`: Not certified; `3`: Certified.
	AuthorityState *int64 `json:"AuthorityState,omitnil" name:"AuthorityState"`

	// The label of the DID.
	LabelName *string `json:"LabelName,omitnil" name:"LabelName"`

	// The DID creation time.
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The consortium name.
	AllianceName *string `json:"AllianceName,omitnil" name:"AllianceName"`

	// The label ID.
	LabelId *uint64 `json:"LabelId,omitnil" name:"LabelId"`
}

type DidServiceInfo struct {
	// The DID service ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// The application ID.
	Appid *uint64 `json:"Appid,omitnil" name:"Appid"`

	// The account ID.
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// The consortium ID.
	ConsortiumId *int64 `json:"ConsortiumId,omitnil" name:"ConsortiumId"`

	// The consortium name.
	ConsortiumName *string `json:"ConsortiumName,omitnil" name:"ConsortiumName"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// The chain ID.
	ChainId *string `json:"ChainId,omitnil" name:"ChainId"`

	// Whether you are the owner of the consortium. `1`: Yes; `0`: No.
	RoleType *uint64 `json:"RoleType,omitnil" name:"RoleType"`

	// The organization DID.
	AgencyDid *string `json:"AgencyDid,omitnil" name:"AgencyDid"`

	// The organization name.
	CreateOrg *string `json:"CreateOrg,omitnil" name:"CreateOrg"`

	// The endpoint.
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// The last updated time.
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// The group name.
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`
}

// Predefined struct for user
type DownCptRequestParams struct {
	// The claim protocol type (CPT) index.
	CptIndex *uint64 `json:"CptIndex,omitnil" name:"CptIndex"`
}

type DownCptRequest struct {
	*tchttp.BaseRequest
	
	// The claim protocol type (CPT) index.
	CptIndex *uint64 `json:"CptIndex,omitnil" name:"CptIndex"`
}

func (r *DownCptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownCptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CptIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownCptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownCptResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DownCptResponse struct {
	*tchttp.BaseResponse
	Response *DownCptResponseParams `json:"Response"`
}

func (r *DownCptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownCptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableHashRequestParams struct {
	// The CNS address of the contract.
	Hash *string `json:"Hash,omitnil" name:"Hash"`
}

type EnableHashRequest struct {
	*tchttp.BaseRequest
	
	// The CNS address of the contract.
	Hash *string `json:"Hash,omitnil" name:"Hash"`
}

func (r *EnableHashRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableHashRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Hash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableHashRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableHashResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableHashResponse struct {
	*tchttp.BaseResponse
	Response *EnableHashResponseParams `json:"Response"`
}

func (r *EnableHashResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableHashResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FunctionArg struct {
	// The claim protocol type (CPT) ID.
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`

	// The DID of the issuer.
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// The expiration time.
	ExpirationDate *string `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// The claim.
	ClaimJson *string `json:"ClaimJson,omitnil" name:"ClaimJson"`
}

// Predefined struct for user
type GetAgencyTDidRequestParams struct {
	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type GetAgencyTDidRequest struct {
	*tchttp.BaseRequest
	
	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *GetAgencyTDidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAgencyTDidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAgencyTDidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAgencyTDidResponseParams struct {
	// The prefix (fixed).
	Prefix *string `json:"Prefix,omitnil" name:"Prefix"`

	// The details of the DID.
	Identity []*Identity `json:"Identity,omitnil" name:"Identity"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetAgencyTDidResponse struct {
	*tchttp.BaseResponse
	Response *GetAgencyTDidResponseParams `json:"Response"`
}

func (r *GetAgencyTDidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAgencyTDidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAuthoritiesListRequestParams struct {
	// The page number, beginning from 1.
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The number of records per page.
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// Whether to query certified or uncertified authorities. `1`: Certified; `2`: Uncertified.
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type GetAuthoritiesListRequest struct {
	*tchttp.BaseRequest
	
	// The page number, beginning from 1.
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The number of records per page.
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// Whether to query certified or uncertified authorities. `1`: Certified; `2`: Uncertified.
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

func (r *GetAuthoritiesListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthoritiesListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Did")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAuthoritiesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAuthoritiesListResponseParams struct {
	// A data set.
	ResultList []*Authority `json:"ResultList,omitnil" name:"ResultList"`

	// The total number of records.
	AllCount *int64 `json:"AllCount,omitnil" name:"AllCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetAuthoritiesListResponse struct {
	*tchttp.BaseResponse
	Response *GetAuthoritiesListResponseParams `json:"Response"`
}

func (r *GetAuthoritiesListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthoritiesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAuthorityIssuerRequestParams struct {
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`
}

type GetAuthorityIssuerRequest struct {
	*tchttp.BaseRequest
	
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`
}

func (r *GetAuthorityIssuerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthorityIssuerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAuthorityIssuerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAuthorityIssuerResponseParams struct {
	// The authority name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// The blockchain network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The blockchain group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The DID of the authority.
	Did *string `json:"Did,omitnil" name:"Did"`

	// Remarks.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// The registration time.
	RegisterTime *string `json:"RegisterTime,omitnil" name:"RegisterTime"`

	// The recognition time.
	RecognizeTime *string `json:"RecognizeTime,omitnil" name:"RecognizeTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetAuthorityIssuerResponse struct {
	*tchttp.BaseResponse
	Response *GetAuthorityIssuerResponseParams `json:"Response"`
}

func (r *GetAuthorityIssuerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAuthorityIssuerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetConsortiumClusterListRequestParams struct {
	// The consortium ID.
	ConsortiumId *uint64 `json:"ConsortiumId,omitnil" name:"ConsortiumId"`
}

type GetConsortiumClusterListRequest struct {
	*tchttp.BaseRequest
	
	// The consortium ID.
	ConsortiumId *uint64 `json:"ConsortiumId,omitnil" name:"ConsortiumId"`
}

func (r *GetConsortiumClusterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetConsortiumClusterListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConsortiumId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetConsortiumClusterListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetConsortiumClusterListResponseParams struct {
	// A list of networks.
	ClusterList []*BcosClusterItem `json:"ClusterList,omitnil" name:"ClusterList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetConsortiumClusterListResponse struct {
	*tchttp.BaseResponse
	Response *GetConsortiumClusterListResponseParams `json:"Response"`
}

func (r *GetConsortiumClusterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetConsortiumClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetConsortiumListRequestParams struct {

}

type GetConsortiumListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetConsortiumListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetConsortiumListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetConsortiumListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetConsortiumListResponseParams struct {
	// A list of the consortiums.
	ConsortiumList []*ConsortiumItem `json:"ConsortiumList,omitnil" name:"ConsortiumList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetConsortiumListResponse struct {
	*tchttp.BaseResponse
	Response *GetConsortiumListResponseParams `json:"Response"`
}

func (r *GetConsortiumListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetConsortiumListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCptInfoRequestParams struct {
	// The claim protocol type (CPT) index.
	CptIndex *uint64 `json:"CptIndex,omitnil" name:"CptIndex"`
}

type GetCptInfoRequest struct {
	*tchttp.BaseRequest
	
	// The claim protocol type (CPT) index.
	CptIndex *uint64 `json:"CptIndex,omitnil" name:"CptIndex"`
}

func (r *GetCptInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCptInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CptIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCptInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCptInfoResponseParams struct {
	// The JSON data of the claim protocol type (CPT).
	CptJsonData *string `json:"CptJsonData,omitnil" name:"CptJsonData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetCptInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetCptInfoResponseParams `json:"Response"`
}

func (r *GetCptInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCptInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCptListRequestParams struct {
	// The start.
	DisplayStart *uint64 `json:"DisplayStart,omitnil" name:"DisplayStart"`

	// The maximum number of records to return.
	DisplayLength *uint64 `json:"DisplayLength,omitnil" name:"DisplayLength"`

	// The type. `0`: All CPTs; `1`: System CPTs; `2`: Authorization CPTs; `3`: General CPTs
	CptType *uint64 `json:"CptType,omitnil" name:"CptType"`
}

type GetCptListRequest struct {
	*tchttp.BaseRequest
	
	// The start.
	DisplayStart *uint64 `json:"DisplayStart,omitnil" name:"DisplayStart"`

	// The maximum number of records to return.
	DisplayLength *uint64 `json:"DisplayLength,omitnil" name:"DisplayLength"`

	// The type. `0`: All CPTs; `1`: System CPTs; `2`: Authorization CPTs; `3`: General CPTs
	CptType *uint64 `json:"CptType,omitnil" name:"CptType"`
}

func (r *GetCptListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCptListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayStart")
	delete(f, "DisplayLength")
	delete(f, "CptType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCptListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCptListResponseParams struct {
	// The information of claim protocol types (CPT).
	CptDataList []*CptListData `json:"CptDataList,omitnil" name:"CptDataList"`

	// The total number of claim protocol types (CPT).
	AllCount *uint64 `json:"AllCount,omitnil" name:"AllCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetCptListResponse struct {
	*tchttp.BaseResponse
	Response *GetCptListResponseParams `json:"Response"`
}

func (r *GetCptListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCptListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialCptRankRequestParams struct {
	// The start date (as early as 2021-4-23).
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date (as early as 2021-4-23).
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type GetCredentialCptRankRequest struct {
	*tchttp.BaseRequest
	
	// The start date (as early as 2021-4-23).
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date (as early as 2021-4-23).
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *GetCredentialCptRankRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialCptRankRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCredentialCptRankRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialCptRankResponseParams struct {
	// The rankings.
	RankIssueResult []*CptIssueRank `json:"RankIssueResult,omitnil" name:"RankIssueResult"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetCredentialCptRankResponse struct {
	*tchttp.BaseResponse
	Response *GetCredentialCptRankResponseParams `json:"Response"`
}

func (r *GetCredentialCptRankResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialCptRankResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialIssueRankRequestParams struct {
	// The start date (as early as 2021-4-23).
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date (as early as 2021-4-23).
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type GetCredentialIssueRankRequest struct {
	*tchttp.BaseRequest
	
	// The start date (as early as 2021-4-23).
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date (as early as 2021-4-23).
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *GetCredentialIssueRankRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialIssueRankRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCredentialIssueRankRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialIssueRankResponseParams struct {
	// The rankings.
	RankIssueResult []*CptIssueRank `json:"RankIssueResult,omitnil" name:"RankIssueResult"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetCredentialIssueRankResponse struct {
	*tchttp.BaseResponse
	Response *GetCredentialIssueRankResponseParams `json:"Response"`
}

func (r *GetCredentialIssueRankResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialIssueRankResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialIssueTrendRequestParams struct {
	// The start date (as early as 2021-4-23).
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date (as early as 2021-4-23).
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type GetCredentialIssueTrendRequest struct {
	*tchttp.BaseRequest
	
	// The start date (as early as 2021-4-23).
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date (as early as 2021-4-23).
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *GetCredentialIssueTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialIssueTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCredentialIssueTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialIssueTrendResponseParams struct {
	// The trend information.
	Trend []*Trend `json:"Trend,omitnil" name:"Trend"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetCredentialIssueTrendResponse struct {
	*tchttp.BaseResponse
	Response *GetCredentialIssueTrendResponseParams `json:"Response"`
}

func (r *GetCredentialIssueTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialIssueTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialStatusRequestParams struct {
	// The credential ID.
	CredentialId *string `json:"CredentialId,omitnil" name:"CredentialId"`
}

type GetCredentialStatusRequest struct {
	*tchttp.BaseRequest
	
	// The credential ID.
	CredentialId *string `json:"CredentialId,omitnil" name:"CredentialId"`
}

func (r *GetCredentialStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CredentialId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCredentialStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialStatusResponseParams struct {
	// The credential status.
	CredentialStatus *CredentialStatus `json:"CredentialStatus,omitnil" name:"CredentialStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetCredentialStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetCredentialStatusResponseParams `json:"Response"`
}

func (r *GetCredentialStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataPanelRequestParams struct {
	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type GetDataPanelRequest struct {
	*tchttp.BaseRequest
	
	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *GetDataPanelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataPanelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDataPanelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDataPanelResponseParams struct {
	// The number of blockchain networks.
	BlockNetworkCount *int64 `json:"BlockNetworkCount,omitnil" name:"BlockNetworkCount"`

	// The blockchain network name.
	BlockNetworkName *string `json:"BlockNetworkName,omitnil" name:"BlockNetworkName"`

	// The current block height.
	BlockHeight *int64 `json:"BlockHeight,omitnil" name:"BlockHeight"`

	// The blockchain network type.
	BlockNetworkType *int64 `json:"BlockNetworkType,omitnil" name:"BlockNetworkType"`

	// The number of DIDs.
	DidCount *int64 `json:"DidCount,omitnil" name:"DidCount"`

	// The number of claim protocol types (CPT).
	CptCount *int64 `json:"CptCount,omitnil" name:"CptCount"`

	// The number of certified authorities.
	CertificatedAuthCount *int64 `json:"CertificatedAuthCount,omitnil" name:"CertificatedAuthCount"`

	// The number of credentials issued.
	IssueCptCount *int64 `json:"IssueCptCount,omitnil" name:"IssueCptCount"`

	// The number of new DIDs in the current week.
	NewDidCount *int64 `json:"NewDidCount,omitnil" name:"NewDidCount"`

	// The number of BCOS networks.
	BcosCount *int64 `json:"BcosCount,omitnil" name:"BcosCount"`

	// The number of Fabric networks.
	FabricCount *int64 `json:"FabricCount,omitnil" name:"FabricCount"`

	// The number of ChainMaker networks.
	ChainMakerCount *int64 `json:"ChainMakerCount,omitnil" name:"ChainMakerCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDataPanelResponse struct {
	*tchttp.BaseResponse
	Response *GetDataPanelResponseParams `json:"Response"`
}

func (r *GetDataPanelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDataPanelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeployInfoRequestParams struct {
	// The CNS address of the contract.
	Hash *string `json:"Hash,omitnil" name:"Hash"`
}

type GetDeployInfoRequest struct {
	*tchttp.BaseRequest
	
	// The CNS address of the contract.
	Hash *string `json:"Hash,omitnil" name:"Hash"`
}

func (r *GetDeployInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeployInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Hash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeployInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeployInfoResponseParams struct {
	// The CNS address of the contract.
	Hash *string `json:"Hash,omitnil" name:"Hash"`

	// The main group ID of the contract.
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// The DID of the organization that deployed the contract.
	DeployDid *string `json:"DeployDid,omitnil" name:"DeployDid"`

	// The TDID SDK version.
	SdkVersion *string `json:"SdkVersion,omitnil" name:"SdkVersion"`

	// The TDID contract version.
	ContractVersion *string `json:"ContractVersion,omitnil" name:"ContractVersion"`

	// The blockchain node version.
	BlockVersion *string `json:"BlockVersion,omitnil" name:"BlockVersion"`

	// The IP address of the blockchain node.
	BlockIp *string `json:"BlockIp,omitnil" name:"BlockIp"`

	// The address of the DID contract.
	DidAddress *string `json:"DidAddress,omitnil" name:"DidAddress"`

	// The address of the claim protocol type (CPT) contract.
	CptAddress *string `json:"CptAddress,omitnil" name:"CptAddress"`

	// The address of the authority.
	AuthorityAddress *string `json:"AuthorityAddress,omitnil" name:"AuthorityAddress"`

	// The evidence contract address.
	EvidenceAddress *string `json:"EvidenceAddress,omitnil" name:"EvidenceAddress"`

	// The contract address of the specific issuer.
	SpecificAddress *string `json:"SpecificAddress,omitnil" name:"SpecificAddress"`

	// The chain ID.
	ChainId *string `json:"ChainId,omitnil" name:"ChainId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDeployInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetDeployInfoResponseParams `json:"Response"`
}

func (r *GetDeployInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeployInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeployListRequestParams struct {
	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The start.
	DisplayStart *uint64 `json:"DisplayStart,omitnil" name:"DisplayStart"`

	// The maximum number of records to return.
	DisplayLength *uint64 `json:"DisplayLength,omitnil" name:"DisplayLength"`
}

type GetDeployListRequest struct {
	*tchttp.BaseRequest
	
	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The start.
	DisplayStart *uint64 `json:"DisplayStart,omitnil" name:"DisplayStart"`

	// The maximum number of records to return.
	DisplayLength *uint64 `json:"DisplayLength,omitnil" name:"DisplayLength"`
}

func (r *GetDeployListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeployListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "GroupId")
	delete(f, "DisplayStart")
	delete(f, "DisplayLength")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDeployListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDeployListResponseParams struct {
	// The total number of contracts.
	AllCount *uint64 `json:"AllCount,omitnil" name:"AllCount"`

	// A list of deployed contracts.
	Result []*Contract `json:"Result,omitnil" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDeployListResponse struct {
	*tchttp.BaseResponse
	Response *GetDeployListResponseParams `json:"Response"`
}

func (r *GetDeployListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDeployListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidClusterDetailRequestParams struct {
	// The DID network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type GetDidClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// The DID network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *GetDidClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidClusterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidClusterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidClusterDetailResponseParams struct {
	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The consortium name.
	ConsortiumName *string `json:"ConsortiumName,omitnil" name:"ConsortiumName"`

	// The name of the blockchain organization.
	ChainAgency *string `json:"ChainAgency,omitnil" name:"ChainAgency"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDidClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *GetDidClusterDetailResponseParams `json:"Response"`
}

func (r *GetDidClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidClusterListRequestParams struct {

}

type GetDidClusterListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetDidClusterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidClusterListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidClusterListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidClusterListResponseParams struct {
	// A list of the DID networks.
	DidClusterList []*DidCluster `json:"DidClusterList,omitnil" name:"DidClusterList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDidClusterListResponse struct {
	*tchttp.BaseResponse
	Response *GetDidClusterListResponseParams `json:"Response"`
}

func (r *GetDidClusterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidDetailRequestParams struct {
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`
}

type GetDidDetailRequest struct {
	*tchttp.BaseRequest
	
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`
}

func (r *GetDidDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidDetailResponseParams struct {
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// Remarks
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// The public key.
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// Whether the DID is a certified authority.
	AuthorityState *int64 `json:"AuthorityState,omitnil" name:"AuthorityState"`

	// The consortium ID.
	ConsortiumId *int64 `json:"ConsortiumId,omitnil" name:"ConsortiumId"`

	// The consortium name.
	ConsortiumName *string `json:"ConsortiumName,omitnil" name:"ConsortiumName"`

	// The group ID.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The BCOS resource ID.
	ResChainId *string `json:"ResChainId,omitnil" name:"ResChainId"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDidDetailResponse struct {
	*tchttp.BaseResponse
	Response *GetDidDetailResponseParams `json:"Response"`
}

func (r *GetDidDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidDocumentRequestParams struct {
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`
}

type GetDidDocumentRequest struct {
	*tchttp.BaseRequest
	
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`
}

func (r *GetDidDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidDocumentResponseParams struct {
	// The name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// The DID document.
	Document *string `json:"Document,omitnil" name:"Document"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDidDocumentResponse struct {
	*tchttp.BaseResponse
	Response *GetDidDocumentResponseParams `json:"Response"`
}

func (r *GetDidDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidListRequestParams struct {
	// The number of records per page.
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// The page number, beginning from 1.
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`
}

type GetDidListRequest struct {
	*tchttp.BaseRequest
	
	// The number of records per page.
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// The page number, beginning from 1.
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`
}

func (r *GetDidListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Did")
	delete(f, "ClusterId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidListResponseParams struct {
	// A list of DIDs.
	DataList []*DidData `json:"DataList,omitnil" name:"DataList"`

	// The total number of records.
	AllCount *int64 `json:"AllCount,omitnil" name:"AllCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDidListResponse struct {
	*tchttp.BaseResponse
	Response *GetDidListResponseParams `json:"Response"`
}

func (r *GetDidListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidRegisterTrendRequestParams struct {
	// The start date (as early as 2021-4-23).
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date (as early as 2021-4-23).
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type GetDidRegisterTrendRequest struct {
	*tchttp.BaseRequest
	
	// The start date (as early as 2021-4-23).
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date (as early as 2021-4-23).
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *GetDidRegisterTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidRegisterTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidRegisterTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidRegisterTrendResponseParams struct {
	// The trend information.
	Trend []*Trend `json:"Trend,omitnil" name:"Trend"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDidRegisterTrendResponse struct {
	*tchttp.BaseResponse
	Response *GetDidRegisterTrendResponseParams `json:"Response"`
}

func (r *GetDidRegisterTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidRegisterTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidServiceDetailRequestParams struct {
	// The DID service ID.
	ServiceId *uint64 `json:"ServiceId,omitnil" name:"ServiceId"`
}

type GetDidServiceDetailRequest struct {
	*tchttp.BaseRequest
	
	// The DID service ID.
	ServiceId *uint64 `json:"ServiceId,omitnil" name:"ServiceId"`
}

func (r *GetDidServiceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidServiceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidServiceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidServiceDetailResponseParams struct {
	// The DID service information.
	DidService *DidServiceInfo `json:"DidService,omitnil" name:"DidService"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDidServiceDetailResponse struct {
	*tchttp.BaseResponse
	Response *GetDidServiceDetailResponseParams `json:"Response"`
}

func (r *GetDidServiceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidServiceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidServiceListRequestParams struct {
	// `1`: Return results at the network level; `0`: Return results at the service level.
	Type *uint64 `json:"Type,omitnil" name:"Type"`
}

type GetDidServiceListRequest struct {
	*tchttp.BaseRequest
	
	// `1`: Return results at the network level; `0`: Return results at the service level.
	Type *uint64 `json:"Type,omitnil" name:"Type"`
}

func (r *GetDidServiceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidServiceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDidServiceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDidServiceListResponseParams struct {
	// A list of DID services.
	DidServiceList []*DidServiceInfo `json:"DidServiceList,omitnil" name:"DidServiceList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDidServiceListResponse struct {
	*tchttp.BaseResponse
	Response *GetDidServiceListResponseParams `json:"Response"`
}

func (r *GetDidServiceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDidServiceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupListRequestParams struct {
	// `0`: Groups with no DID services; `1`: Groups with DID services.
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type GetGroupListRequest struct {
	*tchttp.BaseRequest
	
	// `0`: Groups with no DID services; `1`: Groups with DID services.
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *GetGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupListResponseParams struct {
	// A list of groups.
	Result []*Group `json:"Result,omitnil" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *GetGroupListResponseParams `json:"Response"`
}

func (r *GetGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLabelListRequestParams struct {
	// The number of records per page.
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// The page number, beginning from 1.
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`
}

type GetLabelListRequest struct {
	*tchttp.BaseRequest
	
	// The number of records per page.
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// The page number, beginning from 1.
	PageNumber *int64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`
}

func (r *GetLabelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLabelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "ClusterId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLabelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLabelListResponseParams struct {
	// A data set.
	Result []*Label `json:"Result,omitnil" name:"Result"`

	// The total number of records.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetLabelListResponse struct {
	*tchttp.BaseResponse
	Response *GetLabelListResponseParams `json:"Response"`
}

func (r *GetLabelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLabelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPolicyListRequestParams struct {
	// The start.
	DisplayStart *int64 `json:"DisplayStart,omitnil" name:"DisplayStart"`

	// The maximum number of records to return.
	DisplayLength *int64 `json:"DisplayLength,omitnil" name:"DisplayLength"`
}

type GetPolicyListRequest struct {
	*tchttp.BaseRequest
	
	// The start.
	DisplayStart *int64 `json:"DisplayStart,omitnil" name:"DisplayStart"`

	// The maximum number of records to return.
	DisplayLength *int64 `json:"DisplayLength,omitnil" name:"DisplayLength"`
}

func (r *GetPolicyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPolicyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayStart")
	delete(f, "DisplayLength")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPolicyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPolicyListResponseParams struct {
	// A list of disclosure policies.
	PolicyDataList []*Policy `json:"PolicyDataList,omitnil" name:"PolicyDataList"`

	// The total number of records.
	AllCount *int64 `json:"AllCount,omitnil" name:"AllCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetPolicyListResponse struct {
	*tchttp.BaseResponse
	Response *GetPolicyListResponseParams `json:"Response"`
}

func (r *GetPolicyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPublicKeyRequestParams struct {
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`
}

type GetPublicKeyRequest struct {
	*tchttp.BaseRequest
	
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`
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
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPublicKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPublicKeyResponseParams struct {
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// The public key.
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type Group struct {
	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The number of nodes.
	NodeCount *uint64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// The number of nodes of the organization.
	NodeCountOfAgency *uint64 `json:"NodeCountOfAgency,omitnil" name:"NodeCountOfAgency"`

	// The description of the group.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Whether you are the owner of the consortium or not.
	RoleType *uint64 `json:"RoleType,omitnil" name:"RoleType"`

	// The chain ID.
	ChainId *string `json:"ChainId,omitnil" name:"ChainId"`
}

type Identity struct {
	// The account identifier.
	AccountIdentifier *string `json:"AccountIdentifier,omitnil" name:"AccountIdentifier"`

	// The chain ID.
	ChainID *string `json:"ChainID,omitnil" name:"ChainID"`

	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The group name.
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`
}

type Label struct {
	// The label ID.
	LabelId *uint64 `json:"LabelId,omitnil" name:"LabelId"`

	// The label name.
	LabelName *string `json:"LabelName,omitnil" name:"LabelName"`

	// The number of DIDs.
	DidCount *int64 `json:"DidCount,omitnil" name:"DidCount"`

	// The DID of the creator.
	Did *string `json:"Did,omitnil" name:"Did"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// The group ID.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`
}

type Policy struct {
	// The policy index.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// The policy name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The service ID.
	ServiceId *uint64 `json:"ServiceId,omitnil" name:"ServiceId"`

	// The application ID of the contract.
	ContractAppId *uint64 `json:"ContractAppId,omitnil" name:"ContractAppId"`

	// The policy ID.
	PolicyId *uint64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// The claim protocol type (CPT) ID.
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`

	// The JSON data.
	PolicyJson *string `json:"PolicyJson,omitnil" name:"PolicyJson"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// The last updated time.
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// The DID of the creator.
	CreatorDid *string `json:"CreatorDid,omitnil" name:"CreatorDid"`

	// The application name.
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// The claim protocol type (CPT) index.
	CptIndex *uint64 `json:"CptIndex,omitnil" name:"CptIndex"`
}

type Proof struct {
	// The creation time.
	Created *int64 `json:"Created,omitnil" name:"Created"`

	// The DID of the creator.
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// The salt value.
	SaltJson *string `json:"SaltJson,omitnil" name:"SaltJson"`

	// The signature.
	SignatureValue *string `json:"SignatureValue,omitnil" name:"SignatureValue"`

	// The type.
	Type *string `json:"Type,omitnil" name:"Type"`
}

// Predefined struct for user
type QueryPolicyRequestParams struct {
	// The policy index.
	PolicyIndex *int64 `json:"PolicyIndex,omitnil" name:"PolicyIndex"`
}

type QueryPolicyRequest struct {
	*tchttp.BaseRequest
	
	// The policy index.
	PolicyIndex *int64 `json:"PolicyIndex,omitnil" name:"PolicyIndex"`
}

func (r *QueryPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryPolicyResponseParams struct {
	// The policy index.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// The policy ID.
	PolicyId *uint64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// The claim protocol type (CPT) ID.
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`

	// The content of the policy.
	PolicyData *string `json:"PolicyData,omitnil" name:"PolicyData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryPolicyResponse struct {
	*tchttp.BaseResponse
	Response *QueryPolicyResponseParams `json:"Response"`
}

func (r *QueryPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeAuthorityIssuerRequestParams struct {
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`
}

type RecognizeAuthorityIssuerRequest struct {
	*tchttp.BaseRequest
	
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`
}

func (r *RecognizeAuthorityIssuerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeAuthorityIssuerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeAuthorityIssuerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeAuthorityIssuerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecognizeAuthorityIssuerResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeAuthorityIssuerResponseParams `json:"Response"`
}

func (r *RecognizeAuthorityIssuerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeAuthorityIssuerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterClaimPolicyRequestParams struct {
	// The claim protocol type (CPT) index.
	CptIndex *uint64 `json:"CptIndex,omitnil" name:"CptIndex"`

	// The disclosure policy.
	Policy *string `json:"Policy,omitnil" name:"Policy"`
}

type RegisterClaimPolicyRequest struct {
	*tchttp.BaseRequest
	
	// The claim protocol type (CPT) index.
	CptIndex *uint64 `json:"CptIndex,omitnil" name:"CptIndex"`

	// The disclosure policy.
	Policy *string `json:"Policy,omitnil" name:"Policy"`
}

func (r *RegisterClaimPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterClaimPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CptIndex")
	delete(f, "Policy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterClaimPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterClaimPolicyResponseParams struct {
	// The policy index.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// The policy ID.
	PolicyId *uint64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RegisterClaimPolicyResponse struct {
	*tchttp.BaseResponse
	Response *RegisterClaimPolicyResponseParams `json:"Response"`
}

func (r *RegisterClaimPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterClaimPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterCptRequestParams struct {
	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The JSON data of the credential claim type (CPT).
	CptJson *string `json:"CptJson,omitnil" name:"CptJson"`

	// If you do not specify this, the ID will auto increment.
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`
}

type RegisterCptRequest struct {
	*tchttp.BaseRequest
	
	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The JSON data of the credential claim type (CPT).
	CptJson *string `json:"CptJson,omitnil" name:"CptJson"`

	// If you do not specify this, the ID will auto increment.
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`
}

func (r *RegisterCptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterCptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "ClusterId")
	delete(f, "CptJson")
	delete(f, "CptId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterCptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterCptResponseParams struct {
	// The claim protocol type (CPT) index.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// The claim protocol type (CPT) ID.
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RegisterCptResponse struct {
	*tchttp.BaseResponse
	Response *RegisterCptResponseParams `json:"Response"`
}

func (r *RegisterCptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterCptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterIssuerRequestParams struct {
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// The issuing authority name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Remarks
	Description *string `json:"Description,omitnil" name:"Description"`
}

type RegisterIssuerRequest struct {
	*tchttp.BaseRequest
	
	// The DID.
	Did *string `json:"Did,omitnil" name:"Did"`

	// The issuing authority name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Remarks
	Description *string `json:"Description,omitnil" name:"Description"`
}

func (r *RegisterIssuerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterIssuerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterIssuerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterIssuerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RegisterIssuerResponse struct {
	*tchttp.BaseResponse
	Response *RegisterIssuerResponseParams `json:"Response"`
}

func (r *RegisterIssuerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterIssuerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveHashRequestParams struct {
	// The CNS address of the contract.
	Hash *string `json:"Hash,omitnil" name:"Hash"`
}

type RemoveHashRequest struct {
	*tchttp.BaseRequest
	
	// The CNS address of the contract.
	Hash *string `json:"Hash,omitnil" name:"Hash"`
}

func (r *RemoveHashRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveHashRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Hash")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveHashRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveHashResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RemoveHashResponse struct {
	*tchttp.BaseResponse
	Response *RemoveHashResponseParams `json:"Response"`
}

func (r *RemoveHashResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveHashResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCredentialStatusRequestParams struct {
	// The credential status.
	CredentialStatus *CredentialStatus `json:"CredentialStatus,omitnil" name:"CredentialStatus"`
}

type SetCredentialStatusRequest struct {
	*tchttp.BaseRequest
	
	// The credential status.
	CredentialStatus *CredentialStatus `json:"CredentialStatus,omitnil" name:"CredentialStatus"`
}

func (r *SetCredentialStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCredentialStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CredentialStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetCredentialStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCredentialStatusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SetCredentialStatusResponse struct {
	*tchttp.BaseResponse
	Response *SetCredentialStatusResponseParams `json:"Response"`
}

func (r *SetCredentialStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCredentialStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Task struct {
	// The task ID.
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// The application ID.
	AppId *uint64 `json:"AppId,omitnil" name:"AppId"`

	// The network ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitnil" name:"GroupId"`

	// The service ID.
	ServiceId *uint64 `json:"ServiceId,omitnil" name:"ServiceId"`

	// `0`: Under deployment; `1`: Deployed successfully; other values: Deployment failed.
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// The error code.
	ErrorCode *string `json:"ErrorCode,omitnil" name:"ErrorCode"`

	// The error message.
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// The last updated time.
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type TransactionArg struct {
	// The credential ID.
	InvokerTDid *string `json:"InvokerTDid,omitnil" name:"InvokerTDid"`
}

type Trend struct {
	// The time point.
	Time *string `json:"Time,omitnil" name:"Time"`

	// The count.
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

// Predefined struct for user
type VerifyCredentialRequestParams struct {
	// A parameter set.
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitnil" name:"FunctionArg"`
}

type VerifyCredentialRequest struct {
	*tchttp.BaseRequest
	
	// A parameter set.
	FunctionArg *VerifyFunctionArg `json:"FunctionArg,omitnil" name:"FunctionArg"`
}

func (r *VerifyCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionArg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyCredentialResponseParams struct {
	// Whether the verification is successful.
	Result *bool `json:"Result,omitnil" name:"Result"`

	// The verification code.
	VerifyCode *uint64 `json:"VerifyCode,omitnil" name:"VerifyCode"`

	// The verification message.
	VerifyMessage *string `json:"VerifyMessage,omitnil" name:"VerifyMessage"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VerifyCredentialResponse struct {
	*tchttp.BaseResponse
	Response *VerifyCredentialResponseParams `json:"Response"`
}

func (r *VerifyCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyFunctionArg struct {
	// The claim protocol type (CPT) ID.
	CptId *uint64 `json:"CptId,omitnil" name:"CptId"`

	// The issuer DID.
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// The expiration time.
	ExpirationDate *int64 `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// The claim.
	ClaimJson *string `json:"ClaimJson,omitnil" name:"ClaimJson"`

	// The time when the credential was issued.
	IssuanceDate *int64 `json:"IssuanceDate,omitnil" name:"IssuanceDate"`

	// The context.
	Context *string `json:"Context,omitnil" name:"Context"`

	// The ID.
	Id *string `json:"Id,omitnil" name:"Id"`

	// The signature.
	Proof *Proof `json:"Proof,omitnil" name:"Proof"`

	// The type.
	Type []*string `json:"Type,omitnil" name:"Type"`
}