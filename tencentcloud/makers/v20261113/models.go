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

package v20261113

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type CosCredentials struct {
	// <p>Temporary SecretId</p>
	TmpSecretId *string `json:"TmpSecretId,omitnil,omitempty" name:"TmpSecretId"`

	// <p>Temporary SecretKey</p>
	TmpSecretKey *string `json:"TmpSecretKey,omitnil,omitempty" name:"TmpSecretKey"`

	// <p>Temporary Token</p>
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`
}

// Predefined struct for user
type CreateDeploymentRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Trigger source. Github, ReDeploy, Api, Webhook, Upload, Coding, Gitee, Gitlab</p>
	ViaMeta *string `json:"ViaMeta,omitnil,omitempty" name:"ViaMeta"`

	// <p>Environment slug. Production, Preview, or custom environment</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>Code source. For upload artifact, pass Upload.</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>Artifact type. Zip or Folder</p>
	DistType *string `json:"DistType,omitnil,omitempty" name:"DistType"`

	// <p>Object path of the uploaded artifact in the temporary bucket</p>
	TempBucketPath *string `json:"TempBucketPath,omitnil,omitempty" name:"TempBucketPath"`

	// <p>Build branch</p>
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>Submit hash</p>
	RepoCommitHash *string `json:"RepoCommitHash,omitnil,omitempty" name:"RepoCommitHash"`

	// <p>Commit description</p>
	RepoCommitMsg *string `json:"RepoCommitMsg,omitnil,omitempty" name:"RepoCommitMsg"`

	// <p>Build source</p>
	BuildFrom *string `json:"BuildFrom,omitnil,omitempty" name:"BuildFrom"`

	// <p>Template ID</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Build specification. The current server still reads from edgeone.json. This field is reserved for now.</p>
	CiSpec *string `json:"CiSpec,omitnil,omitempty" name:"CiSpec"`
}

type CreateDeploymentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Trigger source. Github, ReDeploy, Api, Webhook, Upload, Coding, Gitee, Gitlab</p>
	ViaMeta *string `json:"ViaMeta,omitnil,omitempty" name:"ViaMeta"`

	// <p>Environment slug. Production, Preview, or custom environment</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>Code source. For upload artifact, pass Upload.</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>Artifact type. Zip or Folder</p>
	DistType *string `json:"DistType,omitnil,omitempty" name:"DistType"`

	// <p>Object path of the uploaded artifact in the temporary bucket</p>
	TempBucketPath *string `json:"TempBucketPath,omitnil,omitempty" name:"TempBucketPath"`

	// <p>Build branch</p>
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>Submit hash</p>
	RepoCommitHash *string `json:"RepoCommitHash,omitnil,omitempty" name:"RepoCommitHash"`

	// <p>Commit description</p>
	RepoCommitMsg *string `json:"RepoCommitMsg,omitnil,omitempty" name:"RepoCommitMsg"`

	// <p>Build source</p>
	BuildFrom *string `json:"BuildFrom,omitnil,omitempty" name:"BuildFrom"`

	// <p>Template ID</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Build specification. The current server still reads from edgeone.json. This field is reserved for now.</p>
	CiSpec *string `json:"CiSpec,omitnil,omitempty" name:"CiSpec"`
}

func (r *CreateDeploymentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeploymentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "ViaMeta")
	delete(f, "Env")
	delete(f, "Provider")
	delete(f, "DistType")
	delete(f, "TempBucketPath")
	delete(f, "RepoBranch")
	delete(f, "RepoCommitHash")
	delete(f, "RepoCommitMsg")
	delete(f, "BuildFrom")
	delete(f, "TemplateId")
	delete(f, "CiSpec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeploymentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeploymentResponseParams struct {
	// <p>Deployment ID</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDeploymentResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeploymentResponseParams `json:"Response"`
}

func (r *CreateDeploymentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Create Source. Template, Custom, CLI, Guest</p>
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// <p>Code source. Upload, Github, Gitee, Gitlab, Coding</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>Accessing party identifier.</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Acceleration region. mainland, overseas, global</p>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// <p>Node.js version. 14.21.3, 16.20.2, 18.20.4, 20.18.0, 22.11.0, 22.17.1, 22.21.1, 24.5.0, 24.11.0, 24.18.0</p>
	NodeJsVersion *string `json:"NodeJsVersion,omitnil,omitempty" name:"NodeJsVersion"`

	// <p>Project root directory</p>
	RootDir *string `json:"RootDir,omitnil,omitempty" name:"RootDir"`

	// <p>Build artifact directory</p>
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// <p>Build command</p>
	BuildCmd *string `json:"BuildCmd,omitnil,omitempty" name:"BuildCmd"`

	// <p>Installation command</p>
	InstallCmd *string `json:"InstallCmd,omitnil,omitempty" name:"InstallCmd"`

	// <p>Framework</p>
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// <p>Template ID</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Repository address</p>
	RepoUrl *string `json:"RepoUrl,omitnil,omitempty" name:"RepoUrl"`

	// <p>Repository owner</p>
	RepoOwner *string `json:"RepoOwner,omitnil,omitempty" name:"RepoOwner"`

	// <p>Repository name</p>
	RepoName *string `json:"RepoName,omitnil,omitempty" name:"RepoName"`

	// <p>Production branch</p>
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>Initial environment variable.</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`

	// <p>Coding namespace. Used when the Provider is Coding.</p>
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// <p>Coding project ID. Used when the Provider is Coding</p>
	CodingProjectId *int64 `json:"CodingProjectId,omitnil,omitempty" name:"CodingProjectId"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Create Source. Template, Custom, CLI, Guest</p>
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// <p>Code source. Upload, Github, Gitee, Gitlab, Coding</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>Accessing party identifier.</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Acceleration region. mainland, overseas, global</p>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// <p>Node.js version. 14.21.3, 16.20.2, 18.20.4, 20.18.0, 22.11.0, 22.17.1, 22.21.1, 24.5.0, 24.11.0, 24.18.0</p>
	NodeJsVersion *string `json:"NodeJsVersion,omitnil,omitempty" name:"NodeJsVersion"`

	// <p>Project root directory</p>
	RootDir *string `json:"RootDir,omitnil,omitempty" name:"RootDir"`

	// <p>Build artifact directory</p>
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// <p>Build command</p>
	BuildCmd *string `json:"BuildCmd,omitnil,omitempty" name:"BuildCmd"`

	// <p>Installation command</p>
	InstallCmd *string `json:"InstallCmd,omitnil,omitempty" name:"InstallCmd"`

	// <p>Framework</p>
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// <p>Template ID</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Repository address</p>
	RepoUrl *string `json:"RepoUrl,omitnil,omitempty" name:"RepoUrl"`

	// <p>Repository owner</p>
	RepoOwner *string `json:"RepoOwner,omitnil,omitempty" name:"RepoOwner"`

	// <p>Repository name</p>
	RepoName *string `json:"RepoName,omitnil,omitempty" name:"RepoName"`

	// <p>Production branch</p>
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>Initial environment variable.</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`

	// <p>Coding namespace. Used when the Provider is Coding.</p>
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// <p>Coding project ID. Used when the Provider is Coding</p>
	CodingProjectId *int64 `json:"CodingProjectId,omitnil,omitempty" name:"CodingProjectId"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Name")
	delete(f, "Channel")
	delete(f, "Provider")
	delete(f, "Source")
	delete(f, "Area")
	delete(f, "NodeJsVersion")
	delete(f, "RootDir")
	delete(f, "OutputDir")
	delete(f, "BuildCmd")
	delete(f, "InstallCmd")
	delete(f, "Framework")
	delete(f, "TemplateId")
	delete(f, "RepoUrl")
	delete(f, "RepoOwner")
	delete(f, "RepoName")
	delete(f, "RepoBranch")
	delete(f, "EnvVars")
	delete(f, "Namespace")
	delete(f, "CodingProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateProjectResponseParams `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneCustomDomainRequestParams struct {
	// <p>Site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Custom domain name</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Environment slug. If not passed, it is determined by the backend.</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type CreateZoneCustomDomainRequest struct {
	*tchttp.BaseRequest
	
	// <p>Site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Custom domain name</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Environment slug. If not passed, it is determined by the backend.</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *CreateZoneCustomDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateZoneCustomDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "Domain")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateZoneCustomDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneCustomDomainResponseParams struct {
	// <p>Domain name.</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Domain configuration ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateZoneCustomDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateZoneCustomDomainResponseParams `json:"Response"`
}

func (r *CreateZoneCustomDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateZoneCustomDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomDomainInput struct {
	// <p>Domain name.</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Certificate ID</p>
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// <p>Verification status. Verify, Pass, Fail</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Whether not to apply for a certificate.</p>
	NoCertificate *bool `json:"NoCertificate,omitnil,omitempty" name:"NoCertificate"`
}

type CustomDomainItem struct {
	// <p>Domain name.</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

// Predefined struct for user
type DeleteProjectCustomDomainsRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Domain to unbind</p>
	CustomDomains []*string `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`
}

type DeleteProjectCustomDomainsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Domain to unbind</p>
	CustomDomains []*string `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`
}

func (r *DeleteProjectCustomDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectCustomDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "CustomDomains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectCustomDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectCustomDomainsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectCustomDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectCustomDomainsResponseParams `json:"Response"`
}

func (r *DeleteProjectCustomDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectCustomDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectEnvsRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Environment variables to delete. Pass Key (or Id). Value is required but not used for matching.</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

type DeleteProjectEnvsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Environment variables to delete. Pass Key (or Id). Value is required but not used for matching.</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

func (r *DeleteProjectEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "EnvVars")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectEnvsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectEnvsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectEnvsResponseParams `json:"Response"`
}

func (r *DeleteProjectEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectResponseParams `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Deployment struct {
	// <p>Deployment ID</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Success,Failed,Process,Invalid,Pending,Timeout,Cancelled</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Original status before deployment removal</p>
	OriginalStatus *string `json:"OriginalStatus,omitnil,omitempty" name:"OriginalStatus"`

	// <p>Build exit code. 0 indicates success</p>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>Build duration, in milliseconds.</p>
	BuildCost *string `json:"BuildCost,omitnil,omitempty" name:"BuildCost"`

	// <p>Trigger source. Github, ReDeploy, Api, Coding, Gitee, Upload, Gitlab</p>
	ViaMeta *string `json:"ViaMeta,omitnil,omitempty" name:"ViaMeta"`

	// <p>Environment slug</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>Branch</p>
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>Submit hash</p>
	RepoCommitHash *string `json:"RepoCommitHash,omitnil,omitempty" name:"RepoCommitHash"`

	// <p>Commit description</p>
	RepoCommitMsg *string `json:"RepoCommitMsg,omitnil,omitempty" name:"RepoCommitMsg"`

	// <p>Access address of this deployment</p>
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`

	// <p>Default domain address of the project</p>
	ProjectUrl *string `json:"ProjectUrl,omitnil,omitempty" name:"ProjectUrl"`

	// <p>Whether it is the current production version</p>
	UsedInProd *bool `json:"UsedInProd,omitnil,omitempty" name:"UsedInProd"`

	// <p>Address of a cover</p>
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// <p>Routing configuration, JSON string</p>
	MetaData *string `json:"MetaData,omitnil,omitempty" name:"MetaData"`

	// <p>Repository address used for this deployment</p>
	DplRepoUrl *string `json:"DplRepoUrl,omitnil,omitempty" name:"DplRepoUrl"`

	// <p>Creation time.</p>
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// <p>Update time.</p>
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`
}

type DeploymentFilter struct {
	// <p>Filter field. ProjectId, DeploymentId, Name, Status, StartTime, EndTime, RepoBranch</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Filter value</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type DescribeCosTempTokenRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project name. The temporary path is concatenated based on it.</p>
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// <p>Project ID. Used to reverse-check the project name when ProjectName is not passed in.</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeCosTempTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project name. The temporary path is concatenated based on it.</p>
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// <p>Project ID. Used to reverse-check the project name when ProjectName is not passed in.</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeCosTempTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosTempTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectName")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCosTempTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCosTempTokenResponseParams struct {
	// <p>COS bucket</p>
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// <p>COS region</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Upload path, format: appId/projectName/timestamp</p>
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// <p>Temporary key expiration time in Unix seconds</p>
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// <p>Time when the temporary key expires in ISO 8601</p>
	Expiration *string `json:"Expiration,omitnil,omitempty" name:"Expiration"`

	// <p>Temporary key</p>
	Credentials *CosCredentials `json:"Credentials,omitnil,omitempty" name:"Credentials"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCosTempTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCosTempTokenResponseParams `json:"Response"`
}

func (r *DescribeCosTempTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCosTempTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeploymentLogRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Deployment ID</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`
}

type DescribeDeploymentLogRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Deployment ID</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`
}

func (r *DescribeDeploymentLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeploymentLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "DeploymentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeploymentLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeploymentLogResponseParams struct {
	// <p>Build log address</p>
	LogUrl *string `json:"LogUrl,omitnil,omitempty" name:"LogUrl"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeploymentLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeploymentLogResponseParams `json:"Response"`
}

func (r *DescribeDeploymentLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeploymentLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeploymentsRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Query by deployment ID</p>
	DeploymentIds []*string `json:"DeploymentIds,omitnil,omitempty" name:"DeploymentIds"`

	// <p>Filter conditions.</p>
	Filters []*DeploymentFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Offset</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of items per page. Maximum value: 100.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>CreatedOn or ModifiedOn</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Asc or Desc</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>Start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>End time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeDeploymentsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Query by deployment ID</p>
	DeploymentIds []*string `json:"DeploymentIds,omitnil,omitempty" name:"DeploymentIds"`

	// <p>Filter conditions.</p>
	Filters []*DeploymentFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Offset</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of items per page. Maximum value: 100.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>CreatedOn or ModifiedOn</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Asc or Desc</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>Start time.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>End time.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeDeploymentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeploymentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "DeploymentIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeploymentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeploymentsResponseParams struct {
	// <p>Total.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Deployment list</p>
	Deployments []*Deployment `json:"Deployments,omitnil,omitempty" name:"Deployments"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeploymentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeploymentsResponseParams `json:"Response"`
}

func (r *DescribeDeploymentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeploymentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsResponseParams struct {
	// <p>Total.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Environment list</p>
	Environments []*Environment `json:"Environments,omitnil,omitempty" name:"Environments"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentsResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectEnvsRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Only return the variables of the environment slug</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeProjectEnvsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Only return the variables of the environment slug</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *DescribeProjectEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectEnvsResponseParams struct {
	// <p>Environment variable.</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectEnvsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectEnvsResponseParams `json:"Response"`
}

func (r *DescribeProjectEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Page number offset, starting from 0. The backend skips Offset * Limit entries</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of items per page. Maximum value: 100.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Filter criteria. Name parameter values: ProjectId, Name, Status, Provider</p>
	Filters []*ProjectFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Query by project ID</p>
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// <p>Sorting field. CreatedOn, ModifiedOn. Default: CreatedOn</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Sorting order. Asc, Desc, ASC, DESC. Default: DESC</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>Source. Only return the Upload project on drop</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Fuzzy search by project name</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Page number offset, starting from 0. The backend skips Offset * Limit entries</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Number of items per page. Maximum value: 100.</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Filter criteria. Name parameter values: ProjectId, Name, Status, Provider</p>
	Filters []*ProjectFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>Query by project ID</p>
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// <p>Sorting field. CreatedOn, ModifiedOn. Default: CreatedOn</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>Sorting order. Asc, Desc, ASC, DESC. Default: DESC</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>Source. Only return the Upload project on drop</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Fuzzy search by project name</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "ProjectIds")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "Source")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectsResponseParams struct {
	// <p>Total.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Project list</p>
	Projects []*Project `json:"Projects,omitnil,omitempty" name:"Projects"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectsResponseParams `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneCustomDomainsRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Only query these domain names. If not passed, return the domain name under the project.</p>
	CustomDomains []*CustomDomainInput `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`
}

type DescribeZoneCustomDomainsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Only query these domain names. If not passed, return the domain name under the project.</p>
	CustomDomains []*CustomDomainInput `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`
}

func (r *DescribeZoneCustomDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneCustomDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "CustomDomains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneCustomDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneCustomDomainsResponseParams struct {
	// <p>Total.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Domain name list.</p>
	PagesDomains []*ZoneDomain `json:"PagesDomains,omitnil,omitempty" name:"PagesDomains"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeZoneCustomDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneCustomDomainsResponseParams `json:"Response"`
}

func (r *DescribeZoneCustomDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneCustomDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DnsVerification struct {
	// <p>Record type, TXT</p>
	RecordType *string `json:"RecordType,omitnil,omitempty" name:"RecordType"`

	// <p>Host record</p>
	Subdomain *string `json:"Subdomain,omitnil,omitempty" name:"Subdomain"`

	// <p>Record value</p>
	RecordValue *string `json:"RecordValue,omitnil,omitempty" name:"RecordValue"`
}

type EnvVar struct {
	// <p>Variable name, up to 255 characters</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Variable value, up to 1000</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// <p>Variable ID. It is not passed in upon creation.</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>Remark. Up to 255 characters</p>
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// <p>Effective environment slug, for example, Production, Preview. Multi-line environments with the same Key must not overlap</p>
	Env []*string `json:"Env,omitnil,omitempty" name:"Env"`
}

type Environment struct {
	// <p>Environment identifier. It cannot be changed after creation.</p>
	Slug *string `json:"Slug,omitnil,omitempty" name:"Slug"`

	// <p>Environment name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Production,Preview,Custom</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Associate a branch. Wildcards are supported.</p>
	Branches []*string `json:"Branches,omitnil,omitempty" name:"Branches"`

	// <p>Whether to deploy automatically. Production is always true</p>
	AutoDeploy *bool `json:"AutoDeploy,omitnil,omitempty" name:"AutoDeploy"`

	// <p>Bound custom domain name</p>
	CustomDomains []*string `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`

	// <p>Environment variable.</p>
	EnvVars []*EnvironmentEnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`

	// <p>Creation time.</p>
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// <p>Update time.</p>
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`
}

type EnvironmentEnvVar struct {
	// <p>Variable ID</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>Variable name</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Variable value</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// <p>Remarks.</p>
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// <p>Effective environment slug</p>
	Env []*string `json:"Env,omitnil,omitempty" name:"Env"`
}

type FileVerification struct {
	// <p>Verification file path</p>
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// <p>Verify file content</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ForceRedirectConfig struct {
	// <p>on or off</p>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// <p>301,302,307,308</p>
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitnil,omitempty" name:"RedirectStatusCode"`
}

type HstsConfig struct {
	// <p>on or off</p>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// <p>Valid period in seconds.</p>
	MaxAge *int64 `json:"MaxAge,omitnil,omitempty" name:"MaxAge"`

	// <p>on or off</p>
	IncludeSubDomains *string `json:"IncludeSubDomains,omitnil,omitempty" name:"IncludeSubDomains"`

	// <p>on or off</p>
	Preload *string `json:"Preload,omitnil,omitempty" name:"Preload"`
}

// Predefined struct for user
type ModifyProjectEnvsRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Environment variable to write to</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

type ModifyProjectEnvsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Environment variable to write to</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

func (r *ModifyProjectEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "EnvVars")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectEnvsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProjectEnvsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectEnvsResponseParams `json:"Response"`
}

func (r *ModifyProjectEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectRequestParams struct {
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Project name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Project root directory</p>
	RootDir *string `json:"RootDir,omitnil,omitempty" name:"RootDir"`

	// <p>Build artifact directory</p>
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// <p>Build command</p>
	BuildCmd *string `json:"BuildCmd,omitnil,omitempty" name:"BuildCmd"`

	// <p>Installation command</p>
	InstallCmd *string `json:"InstallCmd,omitnil,omitempty" name:"InstallCmd"`

	// <p>Framework</p>
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// <p>Address of a cover</p>
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// <p>Node.js version</p>
	NodeJsVersion *string `json:"NodeJsVersion,omitnil,omitempty" name:"NodeJsVersion"`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest
	
	// <p>Makers site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Project name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Project root directory</p>
	RootDir *string `json:"RootDir,omitnil,omitempty" name:"RootDir"`

	// <p>Build artifact directory</p>
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// <p>Build command</p>
	BuildCmd *string `json:"BuildCmd,omitnil,omitempty" name:"BuildCmd"`

	// <p>Installation command</p>
	InstallCmd *string `json:"InstallCmd,omitnil,omitempty" name:"InstallCmd"`

	// <p>Framework</p>
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// <p>Address of a cover</p>
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// <p>Node.js version</p>
	NodeJsVersion *string `json:"NodeJsVersion,omitnil,omitempty" name:"NodeJsVersion"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "RootDir")
	delete(f, "OutputDir")
	delete(f, "BuildCmd")
	delete(f, "InstallCmd")
	delete(f, "Framework")
	delete(f, "CoverUrl")
	delete(f, "NodeJsVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectResponseParams `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneCustomDomainEnvRequestParams struct {
	// <p>Site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Domain name.</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Target environment slug</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type ModifyZoneCustomDomainEnvRequest struct {
	*tchttp.BaseRequest
	
	// <p>Site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Domain name.</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Target environment slug</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

func (r *ModifyZoneCustomDomainEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneCustomDomainEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProjectId")
	delete(f, "Domain")
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneCustomDomainEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneCustomDomainEnvResponseParams struct {
	// <p>Domain name.</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Environment slug after toggling</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyZoneCustomDomainEnvResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneCustomDomainEnvResponseParams `json:"Response"`
}

func (r *ModifyZoneCustomDomainEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneCustomDomainEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NsVerification struct {
	// <p>NS server</p>
	NameServers []*string `json:"NameServers,omitnil,omitempty" name:"NameServers"`
}

type OwnershipVerification struct {
	// <p>DNS verification</p>
	DnsVerification *DnsVerification `json:"DnsVerification,omitnil,omitempty" name:"DnsVerification"`

	// <p>File verification</p>
	FileVerification *FileVerification `json:"FileVerification,omitnil,omitempty" name:"FileVerification"`

	// <p>NS verification</p>
	NsVerification *NsVerification `json:"NsVerification,omitnil,omitempty" name:"NsVerification"`
}

type Project struct {
	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Project name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Project status. Normal, Ban, Unlink, Process, Pending, Offline</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Acceleration region. mainland, overseas, global</p>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// <p>Code source. Github, Gitee, Coding, Upload, Gitlab, Bitbucket, Cnb</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>Create Source. Template, Custom, CLI, Guest</p>
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`

	// <p>Preset domain name.</p>
	PresetDomain *string `json:"PresetDomain,omitnil,omitempty" name:"PresetDomain"`

	// <p>Preset domain name status</p>
	PresetDomainStatus *string `json:"PresetDomainStatus,omitnil,omitempty" name:"PresetDomainStatus"`

	// <p>Project root directory</p>
	RootDir *string `json:"RootDir,omitnil,omitempty" name:"RootDir"`

	// <p>Build artifact directory</p>
	OutputDir *string `json:"OutputDir,omitnil,omitempty" name:"OutputDir"`

	// <p>Build command</p>
	BuildCmd *string `json:"BuildCmd,omitnil,omitempty" name:"BuildCmd"`

	// <p>Installation command</p>
	InstallCmd *string `json:"InstallCmd,omitnil,omitempty" name:"InstallCmd"`

	// <p>Framework</p>
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// <p>Node.js version</p>
	NodejsVersion *string `json:"NodejsVersion,omitnil,omitempty" name:"NodejsVersion"`

	// <p>Repository address</p>
	RepoUrl *string `json:"RepoUrl,omitnil,omitempty" name:"RepoUrl"`

	// <p>Repository owner</p>
	RepoOwner *string `json:"RepoOwner,omitnil,omitempty" name:"RepoOwner"`

	// <p>Repository name</p>
	RepoName *string `json:"RepoName,omitnil,omitempty" name:"RepoName"`

	// <p>Production branch</p>
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>Preview Branch</p>
	RepoPreBranches []*string `json:"RepoPreBranches,omitnil,omitempty" name:"RepoPreBranches"`

	// <p>Deployment region</p>
	Regions *Regions `json:"Regions,omitnil,omitempty" name:"Regions"`

	// <p>Environment variable.</p>
	EnvVars []*ProjectEnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`

	// <p>Custom domain name</p>
	CustomDomains []*CustomDomainItem `json:"CustomDomains,omitnil,omitempty" name:"CustomDomains"`

	// <p>Last deployment</p>
	Deployment *ProjectDeployment `json:"Deployment,omitnil,omitempty" name:"Deployment"`

	// <p>Whether it is an Agent project. 1: yes, 0: no</p>
	IsAgent *int64 `json:"IsAgent,omitnil,omitempty" name:"IsAgent"`

	// <p>Agent framework. Non-Agent project is empty</p>
	AgentFramework *string `json:"AgentFramework,omitnil,omitempty" name:"AgentFramework"`

	// <p>Whether it hits a user-level wildcard domain name. 1: yes, 0: no</p>
	IsTld *int64 `json:"IsTld,omitnil,omitempty" name:"IsTld"`

	// <p>Creation time.</p>
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// <p>Update time.</p>
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`
}

type ProjectDeployment struct {
	// <p>Deployment ID</p>
	DeploymentId *string `json:"DeploymentId,omitnil,omitempty" name:"DeploymentId"`

	// <p>Project ID</p>
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Deployment state</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Build duration, in milliseconds.</p>
	BuildCost *string `json:"BuildCost,omitnil,omitempty" name:"BuildCost"`

	// <p>Trigger source.</p>
	ViaMeta *string `json:"ViaMeta,omitnil,omitempty" name:"ViaMeta"`

	// <p>Environment slug</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>Branch</p>
	RepoBranch *string `json:"RepoBranch,omitnil,omitempty" name:"RepoBranch"`

	// <p>Submit hash</p>
	RepoCommitHash *string `json:"RepoCommitHash,omitnil,omitempty" name:"RepoCommitHash"`

	// <p>Commit description</p>
	RepoCommitMsg *string `json:"RepoCommitMsg,omitnil,omitempty" name:"RepoCommitMsg"`

	// <p>Preview address</p>
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`

	// <p>Address of a cover</p>
	CoverUrl *string `json:"CoverUrl,omitnil,omitempty" name:"CoverUrl"`

	// <p>Creation time.</p>
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// <p>Update time.</p>
	ModifiedOn *string `json:"ModifiedOn,omitnil,omitempty" name:"ModifiedOn"`

	// <p>Repository address used for this deployment</p>
	DplRepoUrl *string `json:"DplRepoUrl,omitnil,omitempty" name:"DplRepoUrl"`
}

type ProjectEnvVar struct {
	// <p>Variable name</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Variable value</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ProjectFilter struct {
	// <p>Filter field. ProjectId, Name, Status, Provider</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Filter value</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Regions struct {
	// <p>Chinese mainland</p>
	Mainland []*string `json:"Mainland,omitnil,omitempty" name:"Mainland"`

	// <p>Overseas region</p>
	Overseas []*string `json:"Overseas,omitnil,omitempty" name:"Overseas"`
}

type SwitchConfig struct {
	// <p>on or off</p>
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ZoneDomain struct {
	// <p>Custom or Preset</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Domain name.</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Domain name status</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Bound environment slug. Default value: Production</p>
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// <p>Site ID</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>Acceleration region</p>
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// <p>CNAME</p>
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// <p>Primary domain name.</p>
	MainDomain *string `json:"MainDomain,omitnil,omitempty" name:"MainDomain"`

	// <p>IPv6</p>
	Ipv6 *SwitchConfig `json:"Ipv6,omitnil,omitempty" name:"Ipv6"`

	// <p>HTTP/2</p>
	Http2 *SwitchConfig `json:"Http2,omitnil,omitempty" name:"Http2"`

	// <p>OCSP</p>
	OcspStapling *SwitchConfig `json:"OcspStapling,omitnil,omitempty" name:"OcspStapling"`

	// <p>Force HTTPS</p>
	ForceRedirect *ForceRedirectConfig `json:"ForceRedirect,omitnil,omitempty" name:"ForceRedirect"`

	// <p>HSTS</p>
	Hsts *HstsConfig `json:"Hsts,omitnil,omitempty" name:"Hsts"`

	// <p>Ownership verification</p>
	OwnershipVerification *OwnershipVerification `json:"OwnershipVerification,omitnil,omitempty" name:"OwnershipVerification"`
}