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
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2026-11-13"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateDeploymentRequest() (request *CreateDeploymentRequest) {
    request = &CreateDeploymentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "CreateDeployment")
    
    
    return
}

func NewCreateDeploymentResponse() (response *CreateDeploymentResponse) {
    response = &CreateDeploymentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeployment
// This API is used to create a deployment. When uploading an artifact, pass Upload for ViaMeta and Provider, and pass the path of the uploaded object for TempBucketPath.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GITAUTHEXPIRED = "FailedOperation.GitAuthExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDeployment(request *CreateDeploymentRequest) (response *CreateDeploymentResponse, err error) {
    return c.CreateDeploymentWithContext(context.Background(), request)
}

// CreateDeployment
// This API is used to create a deployment. When uploading an artifact, pass Upload for ViaMeta and Provider, and pass the path of the uploaded object for TempBucketPath.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GITAUTHEXPIRED = "FailedOperation.GitAuthExpired"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateDeploymentWithContext(ctx context.Context, request *CreateDeploymentRequest) (response *CreateDeploymentResponse, err error) {
    if request == nil {
        request = NewCreateDeploymentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "CreateDeployment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeployment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeploymentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "CreateProject")
    
    
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProject
// This API is used to create a project. When uploading an artifact, pass Upload for Provider and Custom for Channel.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    return c.CreateProjectWithContext(context.Background(), request)
}

// CreateProject
// This API is used to create a project. When uploading an artifact, pass Upload for Provider and Custom for Channel.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "CreateProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateZoneCustomDomainRequest() (request *CreateZoneCustomDomainRequest) {
    request = &CreateZoneCustomDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "CreateZoneCustomDomain")
    
    
    return
}

func NewCreateZoneCustomDomainResponse() (response *CreateZoneCustomDomainResponse) {
    response = &CreateZoneCustomDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateZoneCustomDomain
// Bind a custom domain name to a specified environment of a project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateZoneCustomDomain(request *CreateZoneCustomDomainRequest) (response *CreateZoneCustomDomainResponse, err error) {
    return c.CreateZoneCustomDomainWithContext(context.Background(), request)
}

// CreateZoneCustomDomain
// Bind a custom domain name to a specified environment of a project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateZoneCustomDomainWithContext(ctx context.Context, request *CreateZoneCustomDomainRequest) (response *CreateZoneCustomDomainResponse, err error) {
    if request == nil {
        request = NewCreateZoneCustomDomainRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "CreateZoneCustomDomain")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateZoneCustomDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateZoneCustomDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
    request = &DeleteProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DeleteProject")
    
    
    return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
    response = &DeleteProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProject
// Delete a project. In-progress builds prevent deletion.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    return c.DeleteProjectWithContext(context.Background(), request)
}

// DeleteProject
// Delete a project. In-progress builds prevent deletion.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DeleteProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectCustomDomainsRequest() (request *DeleteProjectCustomDomainsRequest) {
    request = &DeleteProjectCustomDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DeleteProjectCustomDomains")
    
    
    return
}

func NewDeleteProjectCustomDomainsResponse() (response *DeleteProjectCustomDomainsResponse) {
    response = &DeleteProjectCustomDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProjectCustomDomains
// Unbind a custom domain name from a project by domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteProjectCustomDomains(request *DeleteProjectCustomDomainsRequest) (response *DeleteProjectCustomDomainsResponse, err error) {
    return c.DeleteProjectCustomDomainsWithContext(context.Background(), request)
}

// DeleteProjectCustomDomains
// Unbind a custom domain name from a project by domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteProjectCustomDomainsWithContext(ctx context.Context, request *DeleteProjectCustomDomainsRequest) (response *DeleteProjectCustomDomainsResponse, err error) {
    if request == nil {
        request = NewDeleteProjectCustomDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DeleteProjectCustomDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProjectCustomDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectCustomDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectEnvsRequest() (request *DeleteProjectEnvsRequest) {
    request = &DeleteProjectEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DeleteProjectEnvs")
    
    
    return
}

func NewDeleteProjectEnvsResponse() (response *DeleteProjectEnvsResponse) {
    response = &DeleteProjectEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProjectEnvs
// Delete environment variables by EnvVars. Match by Id or Key. Value is required but not used for matching. You can pass any value.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteProjectEnvs(request *DeleteProjectEnvsRequest) (response *DeleteProjectEnvsResponse, err error) {
    return c.DeleteProjectEnvsWithContext(context.Background(), request)
}

// DeleteProjectEnvs
// Delete environment variables by EnvVars. Match by Id or Key. Value is required but not used for matching. You can pass any value.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteProjectEnvsWithContext(ctx context.Context, request *DeleteProjectEnvsRequest) (response *DeleteProjectEnvsResponse, err error) {
    if request == nil {
        request = NewDeleteProjectEnvsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DeleteProjectEnvs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProjectEnvs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCosTempTokenRequest() (request *DescribeCosTempTokenRequest) {
    request = &DescribeCosTempTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeCosTempToken")
    
    
    return
}

func NewDescribeCosTempTokenResponse() (response *DescribeCosTempTokenResponse) {
    response = &DescribeCosTempTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCosTempToken
// Get the COS bucket, region, target path, and temporary key for uploading artifacts.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCosTempToken(request *DescribeCosTempTokenRequest) (response *DescribeCosTempTokenResponse, err error) {
    return c.DescribeCosTempTokenWithContext(context.Background(), request)
}

// DescribeCosTempToken
// Get the COS bucket, region, target path, and temporary key for uploading artifacts.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCosTempTokenWithContext(ctx context.Context, request *DescribeCosTempTokenRequest) (response *DescribeCosTempTokenResponse, err error) {
    if request == nil {
        request = NewDescribeCosTempTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeCosTempToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCosTempToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCosTempTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeploymentLogRequest() (request *DescribeDeploymentLogRequest) {
    request = &DescribeDeploymentLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeDeploymentLog")
    
    
    return
}

func NewDescribeDeploymentLogResponse() (response *DescribeDeploymentLogResponse) {
    response = &DescribeDeploymentLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeploymentLog
// Return the access address of the build log.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDeploymentLog(request *DescribeDeploymentLogRequest) (response *DescribeDeploymentLogResponse, err error) {
    return c.DescribeDeploymentLogWithContext(context.Background(), request)
}

// DescribeDeploymentLog
// Return the access address of the build log.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDeploymentLogWithContext(ctx context.Context, request *DescribeDeploymentLogRequest) (response *DescribeDeploymentLogResponse, err error) {
    if request == nil {
        request = NewDescribeDeploymentLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeDeploymentLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeploymentLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeploymentLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeploymentsRequest() (request *DescribeDeploymentsRequest) {
    request = &DescribeDeploymentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeDeployments")
    
    
    return
}

func NewDescribeDeploymentsResponse() (response *DescribeDeploymentsResponse) {
    response = &DescribeDeploymentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeployments
// Query deployment records of a project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDeployments(request *DescribeDeploymentsRequest) (response *DescribeDeploymentsResponse, err error) {
    return c.DescribeDeploymentsWithContext(context.Background(), request)
}

// DescribeDeployments
// Query deployment records of a project.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDeploymentsWithContext(ctx context.Context, request *DescribeDeploymentsRequest) (response *DescribeDeploymentsResponse, err error) {
    if request == nil {
        request = NewDescribeDeploymentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeDeployments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeployments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeploymentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentsRequest() (request *DescribeEnvironmentsRequest) {
    request = &DescribeEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeEnvironments")
    
    
    return
}

func NewDescribeEnvironmentsResponse() (response *DescribeEnvironmentsResponse) {
    response = &DescribeEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvironments
// Query environments in a project, including built-in Production, Preview, and custom environments.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    return c.DescribeEnvironmentsWithContext(context.Background(), request)
}

// DescribeEnvironments
// Query environments in a project, including built-in Production, Preview, and custom environments.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeEnvironmentsWithContext(ctx context.Context, request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeEnvironments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectEnvsRequest() (request *DescribeProjectEnvsRequest) {
    request = &DescribeProjectEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeProjectEnvs")
    
    
    return
}

func NewDescribeProjectEnvsResponse() (response *DescribeProjectEnvsResponse) {
    response = &DescribeProjectEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjectEnvs
// Query project environment variables. If Env is not specified, return all application scopes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeProjectEnvs(request *DescribeProjectEnvsRequest) (response *DescribeProjectEnvsResponse, err error) {
    return c.DescribeProjectEnvsWithContext(context.Background(), request)
}

// DescribeProjectEnvs
// Query project environment variables. If Env is not specified, return all application scopes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeProjectEnvsWithContext(ctx context.Context, request *DescribeProjectEnvsRequest) (response *DescribeProjectEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectEnvsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeProjectEnvs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectEnvs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeProjects")
    
    
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjects
// This API is used to query projects under a site by page. Offset is the page number offset, and the backend skips Offset * Limit items.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    return c.DescribeProjectsWithContext(context.Background(), request)
}

// DescribeProjects
// This API is used to query projects under a site by page. Offset is the page number offset, and the backend skips Offset * Limit items.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeProjectsWithContext(ctx context.Context, request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeProjects")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneCustomDomainsRequest() (request *DescribeZoneCustomDomainsRequest) {
    request = &DescribeZoneCustomDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "DescribeZoneCustomDomains")
    
    
    return
}

func NewDescribeZoneCustomDomainsResponse() (response *DescribeZoneCustomDomainsResponse) {
    response = &DescribeZoneCustomDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZoneCustomDomains
// Queries the custom and preset domain names of a project under a site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeZoneCustomDomains(request *DescribeZoneCustomDomainsRequest) (response *DescribeZoneCustomDomainsResponse, err error) {
    return c.DescribeZoneCustomDomainsWithContext(context.Background(), request)
}

// DescribeZoneCustomDomains
// Queries the custom and preset domain names of a project under a site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeZoneCustomDomainsWithContext(ctx context.Context, request *DescribeZoneCustomDomainsRequest) (response *DescribeZoneCustomDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeZoneCustomDomainsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "DescribeZoneCustomDomains")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneCustomDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneCustomDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectRequest() (request *ModifyProjectRequest) {
    request = &ModifyProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "ModifyProject")
    
    
    return
}

func NewModifyProjectResponse() (response *ModifyProjectResponse) {
    response = &ModifyProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProject
// Modify the project name and build configuration. Fields not passed retain their original values.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    return c.ModifyProjectWithContext(context.Background(), request)
}

// ModifyProject
// Modify the project name and build configuration. Fields not passed retain their original values.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyProjectWithContext(ctx context.Context, request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "ModifyProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectEnvsRequest() (request *ModifyProjectEnvsRequest) {
    request = &ModifyProjectEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "ModifyProjectEnvs")
    
    
    return
}

func NewModifyProjectEnvsResponse() (response *ModifyProjectEnvsResponse) {
    response = &ModifyProjectEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProjectEnvs
// Write environment variables based on the passed-in EnvVars.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyProjectEnvs(request *ModifyProjectEnvsRequest) (response *ModifyProjectEnvsResponse, err error) {
    return c.ModifyProjectEnvsWithContext(context.Background(), request)
}

// ModifyProjectEnvs
// Write environment variables based on the passed-in EnvVars.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyProjectEnvsWithContext(ctx context.Context, request *ModifyProjectEnvsRequest) (response *ModifyProjectEnvsResponse, err error) {
    if request == nil {
        request = NewModifyProjectEnvsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "ModifyProjectEnvs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProjectEnvs require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProjectEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneCustomDomainEnvRequest() (request *ModifyZoneCustomDomainEnvRequest) {
    request = &ModifyZoneCustomDomainEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("makers", APIVersion, "ModifyZoneCustomDomainEnv")
    
    
    return
}

func NewModifyZoneCustomDomainEnvResponse() (response *ModifyZoneCustomDomainEnvResponse) {
    response = &ModifyZoneCustomDomainEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyZoneCustomDomainEnv
// Switch a bound custom domain name to another environment.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyZoneCustomDomainEnv(request *ModifyZoneCustomDomainEnvRequest) (response *ModifyZoneCustomDomainEnvResponse, err error) {
    return c.ModifyZoneCustomDomainEnvWithContext(context.Background(), request)
}

// ModifyZoneCustomDomainEnv
// Switch a bound custom domain name to another environment.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyZoneCustomDomainEnvWithContext(ctx context.Context, request *ModifyZoneCustomDomainEnvRequest) (response *ModifyZoneCustomDomainEnvResponse, err error) {
    if request == nil {
        request = NewModifyZoneCustomDomainEnvRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "makers", APIVersion, "ModifyZoneCustomDomainEnv")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneCustomDomainEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneCustomDomainEnvResponse()
    err = c.Send(request, response)
    return
}
