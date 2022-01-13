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

package v20201112

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-11-12"

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


func NewAssociateAccessGroupsRequest() (request *AssociateAccessGroupsRequest) {
    request = &AssociateAccessGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "AssociateAccessGroups")
    
    
    return
}

func NewAssociateAccessGroupsResponse() (response *AssociateAccessGroupsResponse) {
    response = &AssociateAccessGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateAccessGroups
// This API is used to bind multiple permission groups to a mount point.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssociateAccessGroups(request *AssociateAccessGroupsRequest) (response *AssociateAccessGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateAccessGroupsRequest()
    }
    
    response = NewAssociateAccessGroupsResponse()
    err = c.Send(request, response)
    return
}

// AssociateAccessGroups
// This API is used to bind multiple permission groups to a mount point.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AssociateAccessGroupsWithContext(ctx context.Context, request *AssociateAccessGroupsRequest) (response *AssociateAccessGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateAccessGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssociateAccessGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccessGroupRequest() (request *CreateAccessGroupRequest) {
    request = &CreateAccessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateAccessGroup")
    
    
    return
}

func NewCreateAccessGroupResponse() (response *CreateAccessGroupResponse) {
    response = &CreateAccessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccessGroup
// This API is used to create a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPNAME = "InvalidParameterValue.InvalidAccessGroupName"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDVPCID = "InvalidParameterValue.InvalidVpcId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXISTS = "ResourceNotFound.VpcNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccessGroup(request *CreateAccessGroupRequest) (response *CreateAccessGroupResponse, err error) {
    if request == nil {
        request = NewCreateAccessGroupRequest()
    }
    
    response = NewCreateAccessGroupResponse()
    err = c.Send(request, response)
    return
}

// CreateAccessGroup
// This API is used to create a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPNAME = "InvalidParameterValue.InvalidAccessGroupName"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDVPCID = "InvalidParameterValue.InvalidVpcId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXISTS = "ResourceNotFound.VpcNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccessGroupWithContext(ctx context.Context, request *CreateAccessGroupRequest) (response *CreateAccessGroupResponse, err error) {
    if request == nil {
        request = NewCreateAccessGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateAccessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccessRulesRequest() (request *CreateAccessRulesRequest) {
    request = &CreateAccessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateAccessRules")
    
    
    return
}

func NewCreateAccessRulesResponse() (response *CreateAccessRulesResponse) {
    response = &CreateAccessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccessRules
// This API is used to batch create permission rules. You don't need to enter the permission rule IDs and creation time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDACCESSRULEADDRESS = "InvalidParameterValue.InvalidAccessRuleAddress"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccessRules(request *CreateAccessRulesRequest) (response *CreateAccessRulesResponse, err error) {
    if request == nil {
        request = NewCreateAccessRulesRequest()
    }
    
    response = NewCreateAccessRulesResponse()
    err = c.Send(request, response)
    return
}

// CreateAccessRules
// This API is used to batch create permission rules. You don't need to enter the permission rule IDs and creation time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDACCESSRULEADDRESS = "InvalidParameterValue.InvalidAccessRuleAddress"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccessRulesWithContext(ctx context.Context, request *CreateAccessRulesRequest) (response *CreateAccessRulesResponse, err error) {
    if request == nil {
        request = NewCreateAccessRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateAccessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileSystemRequest() (request *CreateFileSystemRequest) {
    request = &CreateFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateFileSystem")
    
    
    return
}

func NewCreateFileSystemResponse() (response *CreateFileSystemResponse) {
    response = &CreateFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFileSystem
// This API is used to create a file system (asynchronously).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCAPACITYQUOTA = "InvalidParameterValue.InvalidCapacityQuota"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMNAME = "InvalidParameterValue.InvalidFileSystemName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateFileSystem(request *CreateFileSystemRequest) (response *CreateFileSystemResponse, err error) {
    if request == nil {
        request = NewCreateFileSystemRequest()
    }
    
    response = NewCreateFileSystemResponse()
    err = c.Send(request, response)
    return
}

// CreateFileSystem
// This API is used to create a file system (asynchronously).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCAPACITYQUOTA = "InvalidParameterValue.InvalidCapacityQuota"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMNAME = "InvalidParameterValue.InvalidFileSystemName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateFileSystemWithContext(ctx context.Context, request *CreateFileSystemRequest) (response *CreateFileSystemResponse, err error) {
    if request == nil {
        request = NewCreateFileSystemRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLifeCycleRulesRequest() (request *CreateLifeCycleRulesRequest) {
    request = &CreateLifeCycleRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateLifeCycleRules")
    
    
    return
}

func NewCreateLifeCycleRulesResponse() (response *CreateLifeCycleRulesResponse) {
    response = &CreateLifeCycleRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLifeCycleRules
// This API is used to batch create lifecycle rules. You don't need to enter the lifecycle rule IDs and creation time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateLifeCycleRules(request *CreateLifeCycleRulesRequest) (response *CreateLifeCycleRulesResponse, err error) {
    if request == nil {
        request = NewCreateLifeCycleRulesRequest()
    }
    
    response = NewCreateLifeCycleRulesResponse()
    err = c.Send(request, response)
    return
}

// CreateLifeCycleRules
// This API is used to batch create lifecycle rules. You don't need to enter the lifecycle rule IDs and creation time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateLifeCycleRulesWithContext(ctx context.Context, request *CreateLifeCycleRulesRequest) (response *CreateLifeCycleRulesResponse, err error) {
    if request == nil {
        request = NewCreateLifeCycleRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateLifeCycleRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMountPointRequest() (request *CreateMountPointRequest) {
    request = &CreateMountPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateMountPoint")
    
    
    return
}

func NewCreateMountPointResponse() (response *CreateMountPointResponse) {
    response = &CreateMountPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMountPoint
// This API is used to create a mount point for a successfully created file system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTNAME = "InvalidParameterValue.InvalidMountPointName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateMountPoint(request *CreateMountPointRequest) (response *CreateMountPointResponse, err error) {
    if request == nil {
        request = NewCreateMountPointRequest()
    }
    
    response = NewCreateMountPointResponse()
    err = c.Send(request, response)
    return
}

// CreateMountPoint
// This API is used to create a mount point for a successfully created file system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTNAME = "InvalidParameterValue.InvalidMountPointName"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateMountPointWithContext(ctx context.Context, request *CreateMountPointRequest) (response *CreateMountPointResponse, err error) {
    if request == nil {
        request = NewCreateMountPointRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateMountPointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRestoreTasksRequest() (request *CreateRestoreTasksRequest) {
    request = &CreateRestoreTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateRestoreTasks")
    
    
    return
}

func NewCreateRestoreTasksResponse() (response *CreateRestoreTasksResponse) {
    response = &CreateRestoreTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRestoreTasks
// This API is used to batch create restoration tasks. You don't need to enter the restoration task IDs, status, and creation time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateRestoreTasks(request *CreateRestoreTasksRequest) (response *CreateRestoreTasksResponse, err error) {
    if request == nil {
        request = NewCreateRestoreTasksRequest()
    }
    
    response = NewCreateRestoreTasksResponse()
    err = c.Send(request, response)
    return
}

// CreateRestoreTasks
// This API is used to batch create restoration tasks. You don't need to enter the restoration task IDs, status, and creation time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateRestoreTasksWithContext(ctx context.Context, request *CreateRestoreTasksRequest) (response *CreateRestoreTasksResponse, err error) {
    if request == nil {
        request = NewCreateRestoreTasksRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateRestoreTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccessGroupRequest() (request *DeleteAccessGroupRequest) {
    request = &DeleteAccessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteAccessGroup")
    
    
    return
}

func NewDeleteAccessGroupResponse() (response *DeleteAccessGroupResponse) {
    response = &DeleteAccessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccessGroup
// This API is used to delete a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSGROUPBOUND = "FailedOperation.AccessGroupBound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAccessGroup(request *DeleteAccessGroupRequest) (response *DeleteAccessGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAccessGroupRequest()
    }
    
    response = NewDeleteAccessGroupResponse()
    err = c.Send(request, response)
    return
}

// DeleteAccessGroup
// This API is used to delete a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSGROUPBOUND = "FailedOperation.AccessGroupBound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAccessGroupWithContext(ctx context.Context, request *DeleteAccessGroupRequest) (response *DeleteAccessGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAccessGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteAccessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccessRulesRequest() (request *DeleteAccessRulesRequest) {
    request = &DeleteAccessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteAccessRules")
    
    
    return
}

func NewDeleteAccessRulesResponse() (response *DeleteAccessRulesResponse) {
    response = &DeleteAccessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccessRules
// This API is used to batch delete permission rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSRULENOTEXISTS = "ResourceNotFound.AccessRuleNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAccessRules(request *DeleteAccessRulesRequest) (response *DeleteAccessRulesResponse, err error) {
    if request == nil {
        request = NewDeleteAccessRulesRequest()
    }
    
    response = NewDeleteAccessRulesResponse()
    err = c.Send(request, response)
    return
}

// DeleteAccessRules
// This API is used to batch delete permission rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSRULENOTEXISTS = "ResourceNotFound.AccessRuleNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteAccessRulesWithContext(ctx context.Context, request *DeleteAccessRulesRequest) (response *DeleteAccessRulesResponse, err error) {
    if request == nil {
        request = NewDeleteAccessRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteAccessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileSystemRequest() (request *DeleteFileSystemRequest) {
    request = &DeleteFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteFileSystem")
    
    
    return
}

func NewDeleteFileSystemResponse() (response *DeleteFileSystemResponse) {
    response = &DeleteFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteFileSystem
// This API is used to delete a file system. Non-empty file systems cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FILESYSTEMNOTEMPTY = "FailedOperation.FileSystemNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteFileSystem(request *DeleteFileSystemRequest) (response *DeleteFileSystemResponse, err error) {
    if request == nil {
        request = NewDeleteFileSystemRequest()
    }
    
    response = NewDeleteFileSystemResponse()
    err = c.Send(request, response)
    return
}

// DeleteFileSystem
// This API is used to delete a file system. Non-empty file systems cannot be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FILESYSTEMNOTEMPTY = "FailedOperation.FileSystemNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteFileSystemWithContext(ctx context.Context, request *DeleteFileSystemRequest) (response *DeleteFileSystemResponse, err error) {
    if request == nil {
        request = NewDeleteFileSystemRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLifeCycleRulesRequest() (request *DeleteLifeCycleRulesRequest) {
    request = &DeleteLifeCycleRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteLifeCycleRules")
    
    
    return
}

func NewDeleteLifeCycleRulesResponse() (response *DeleteLifeCycleRulesResponse) {
    response = &DeleteLifeCycleRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLifeCycleRules
// This API is used to batch delete lifecycle rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLifeCycleRules(request *DeleteLifeCycleRulesRequest) (response *DeleteLifeCycleRulesResponse, err error) {
    if request == nil {
        request = NewDeleteLifeCycleRulesRequest()
    }
    
    response = NewDeleteLifeCycleRulesResponse()
    err = c.Send(request, response)
    return
}

// DeleteLifeCycleRules
// This API is used to batch delete lifecycle rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteLifeCycleRulesWithContext(ctx context.Context, request *DeleteLifeCycleRulesRequest) (response *DeleteLifeCycleRulesResponse, err error) {
    if request == nil {
        request = NewDeleteLifeCycleRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteLifeCycleRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMountPointRequest() (request *DeleteMountPointRequest) {
    request = &DeleteMountPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteMountPoint")
    
    
    return
}

func NewDeleteMountPointResponse() (response *DeleteMountPointResponse) {
    response = &DeleteMountPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMountPoint
// This API is used to delete a mount point.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteMountPoint(request *DeleteMountPointRequest) (response *DeleteMountPointResponse, err error) {
    if request == nil {
        request = NewDeleteMountPointRequest()
    }
    
    response = NewDeleteMountPointResponse()
    err = c.Send(request, response)
    return
}

// DeleteMountPoint
// This API is used to delete a mount point.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteMountPointWithContext(ctx context.Context, request *DeleteMountPointRequest) (response *DeleteMountPointResponse, err error) {
    if request == nil {
        request = NewDeleteMountPointRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteMountPointResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessGroupRequest() (request *DescribeAccessGroupRequest) {
    request = &DescribeAccessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeAccessGroup")
    
    
    return
}

func NewDescribeAccessGroupResponse() (response *DescribeAccessGroupResponse) {
    response = &DescribeAccessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessGroup
// This API is used to view the details of a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAccessGroup(request *DescribeAccessGroupRequest) (response *DescribeAccessGroupResponse, err error) {
    if request == nil {
        request = NewDescribeAccessGroupRequest()
    }
    
    response = NewDescribeAccessGroupResponse()
    err = c.Send(request, response)
    return
}

// DescribeAccessGroup
// This API is used to view the details of a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAccessGroupWithContext(ctx context.Context, request *DescribeAccessGroupRequest) (response *DescribeAccessGroupResponse, err error) {
    if request == nil {
        request = NewDescribeAccessGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAccessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessGroupsRequest() (request *DescribeAccessGroupsRequest) {
    request = &DescribeAccessGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeAccessGroups")
    
    
    return
}

func NewDescribeAccessGroupsResponse() (response *DescribeAccessGroupsResponse) {
    response = &DescribeAccessGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessGroups
// This API is used to view the list of permission groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDVPCID = "InvalidParameterValue.InvalidVpcId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAccessGroups(request *DescribeAccessGroupsRequest) (response *DescribeAccessGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAccessGroupsRequest()
    }
    
    response = NewDescribeAccessGroupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeAccessGroups
// This API is used to view the list of permission groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDVPCID = "InvalidParameterValue.InvalidVpcId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAccessGroupsWithContext(ctx context.Context, request *DescribeAccessGroupsRequest) (response *DescribeAccessGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAccessGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAccessGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessRulesRequest() (request *DescribeAccessRulesRequest) {
    request = &DescribeAccessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeAccessRules")
    
    
    return
}

func NewDescribeAccessRulesResponse() (response *DescribeAccessRulesResponse) {
    response = &DescribeAccessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessRules
// This API is used to view the list of permission rules by permission group ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAccessRules(request *DescribeAccessRulesRequest) (response *DescribeAccessRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAccessRulesRequest()
    }
    
    response = NewDescribeAccessRulesResponse()
    err = c.Send(request, response)
    return
}

// DescribeAccessRules
// This API is used to view the list of permission rules by permission group ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAccessRulesWithContext(ctx context.Context, request *DescribeAccessRulesRequest) (response *DescribeAccessRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAccessRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAccessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileSystemRequest() (request *DescribeFileSystemRequest) {
    request = &DescribeFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeFileSystem")
    
    
    return
}

func NewDescribeFileSystemResponse() (response *DescribeFileSystemResponse) {
    response = &DescribeFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileSystem
// This API is used to view the details of a file system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeFileSystem(request *DescribeFileSystemRequest) (response *DescribeFileSystemResponse, err error) {
    if request == nil {
        request = NewDescribeFileSystemRequest()
    }
    
    response = NewDescribeFileSystemResponse()
    err = c.Send(request, response)
    return
}

// DescribeFileSystem
// This API is used to view the details of a file system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeFileSystemWithContext(ctx context.Context, request *DescribeFileSystemRequest) (response *DescribeFileSystemResponse, err error) {
    if request == nil {
        request = NewDescribeFileSystemRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileSystemsRequest() (request *DescribeFileSystemsRequest) {
    request = &DescribeFileSystemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeFileSystems")
    
    
    return
}

func NewDescribeFileSystemsResponse() (response *DescribeFileSystemsResponse) {
    response = &DescribeFileSystemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileSystems
// This API is used to view the list of file systems.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeFileSystems(request *DescribeFileSystemsRequest) (response *DescribeFileSystemsResponse, err error) {
    if request == nil {
        request = NewDescribeFileSystemsRequest()
    }
    
    response = NewDescribeFileSystemsResponse()
    err = c.Send(request, response)
    return
}

// DescribeFileSystems
// This API is used to view the list of file systems.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeFileSystemsWithContext(ctx context.Context, request *DescribeFileSystemsRequest) (response *DescribeFileSystemsResponse, err error) {
    if request == nil {
        request = NewDescribeFileSystemsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFileSystemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLifeCycleRulesRequest() (request *DescribeLifeCycleRulesRequest) {
    request = &DescribeLifeCycleRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeLifeCycleRules")
    
    
    return
}

func NewDescribeLifeCycleRulesResponse() (response *DescribeLifeCycleRulesResponse) {
    response = &DescribeLifeCycleRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLifeCycleRules
// This API is used to view the list of lifecycle rules by file system ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLifeCycleRules(request *DescribeLifeCycleRulesRequest) (response *DescribeLifeCycleRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLifeCycleRulesRequest()
    }
    
    response = NewDescribeLifeCycleRulesResponse()
    err = c.Send(request, response)
    return
}

// DescribeLifeCycleRules
// This API is used to view the list of lifecycle rules by file system ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeLifeCycleRulesWithContext(ctx context.Context, request *DescribeLifeCycleRulesRequest) (response *DescribeLifeCycleRulesResponse, err error) {
    if request == nil {
        request = NewDescribeLifeCycleRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeLifeCycleRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMountPointRequest() (request *DescribeMountPointRequest) {
    request = &DescribeMountPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeMountPoint")
    
    
    return
}

func NewDescribeMountPointResponse() (response *DescribeMountPointResponse) {
    response = &DescribeMountPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMountPoint
// This API is used to view the details of a mount point.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMountPoint(request *DescribeMountPointRequest) (response *DescribeMountPointResponse, err error) {
    if request == nil {
        request = NewDescribeMountPointRequest()
    }
    
    response = NewDescribeMountPointResponse()
    err = c.Send(request, response)
    return
}

// DescribeMountPoint
// This API is used to view the details of a mount point.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMountPointWithContext(ctx context.Context, request *DescribeMountPointRequest) (response *DescribeMountPointResponse, err error) {
    if request == nil {
        request = NewDescribeMountPointRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMountPointResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMountPointsRequest() (request *DescribeMountPointsRequest) {
    request = &DescribeMountPointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeMountPoints")
    
    
    return
}

func NewDescribeMountPointsResponse() (response *DescribeMountPointsResponse) {
    response = &DescribeMountPointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMountPoints
// This API is used to view the list of mount points.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMountPoints(request *DescribeMountPointsRequest) (response *DescribeMountPointsResponse, err error) {
    if request == nil {
        request = NewDescribeMountPointsRequest()
    }
    
    response = NewDescribeMountPointsResponse()
    err = c.Send(request, response)
    return
}

// DescribeMountPoints
// This API is used to view the list of mount points.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMountPointsWithContext(ctx context.Context, request *DescribeMountPointsRequest) (response *DescribeMountPointsResponse, err error) {
    if request == nil {
        request = NewDescribeMountPointsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeMountPointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceTagsRequest() (request *DescribeResourceTagsRequest) {
    request = &DescribeResourceTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeResourceTags")
    
    
    return
}

func NewDescribeResourceTagsResponse() (response *DescribeResourceTagsResponse) {
    response = &DescribeResourceTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceTags
// This API is used to view the list of resource tags by file system ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeResourceTags(request *DescribeResourceTagsRequest) (response *DescribeResourceTagsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTagsRequest()
    }
    
    response = NewDescribeResourceTagsResponse()
    err = c.Send(request, response)
    return
}

// DescribeResourceTags
// This API is used to view the list of resource tags by file system ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeResourceTagsWithContext(ctx context.Context, request *DescribeResourceTagsRequest) (response *DescribeResourceTagsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTagsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeResourceTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRestoreTasksRequest() (request *DescribeRestoreTasksRequest) {
    request = &DescribeRestoreTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeRestoreTasks")
    
    
    return
}

func NewDescribeRestoreTasksResponse() (response *DescribeRestoreTasksResponse) {
    response = &DescribeRestoreTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRestoreTasks
// This API is used to view the list of restoration tasks by file system ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRestoreTasks(request *DescribeRestoreTasksRequest) (response *DescribeRestoreTasksResponse, err error) {
    if request == nil {
        request = NewDescribeRestoreTasksRequest()
    }
    
    response = NewDescribeRestoreTasksResponse()
    err = c.Send(request, response)
    return
}

// DescribeRestoreTasks
// This API is used to view the list of restoration tasks by file system ID.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRestoreTasksWithContext(ctx context.Context, request *DescribeRestoreTasksRequest) (response *DescribeRestoreTasksResponse, err error) {
    if request == nil {
        request = NewDescribeRestoreTasksRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRestoreTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateAccessGroupsRequest() (request *DisassociateAccessGroupsRequest) {
    request = &DisassociateAccessGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DisassociateAccessGroups")
    
    
    return
}

func NewDisassociateAccessGroupsResponse() (response *DisassociateAccessGroupsResponse) {
    response = &DisassociateAccessGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateAccessGroups
// This API is used to unbind multiple permission groups from a mount point.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisassociateAccessGroups(request *DisassociateAccessGroupsRequest) (response *DisassociateAccessGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateAccessGroupsRequest()
    }
    
    response = NewDisassociateAccessGroupsResponse()
    err = c.Send(request, response)
    return
}

// DisassociateAccessGroups
// This API is used to unbind multiple permission groups from a mount point.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisassociateAccessGroupsWithContext(ctx context.Context, request *DisassociateAccessGroupsRequest) (response *DisassociateAccessGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateAccessGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisassociateAccessGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessGroupRequest() (request *ModifyAccessGroupRequest) {
    request = &ModifyAccessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyAccessGroup")
    
    
    return
}

func NewModifyAccessGroupResponse() (response *ModifyAccessGroupResponse) {
    response = &ModifyAccessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccessGroup
// This API is used to modify the attributes of a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPNAME = "InvalidParameterValue.InvalidAccessGroupName"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAccessGroup(request *ModifyAccessGroupRequest) (response *ModifyAccessGroupResponse, err error) {
    if request == nil {
        request = NewModifyAccessGroupRequest()
    }
    
    response = NewModifyAccessGroupResponse()
    err = c.Send(request, response)
    return
}

// ModifyAccessGroup
// This API is used to modify the attributes of a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPID = "InvalidParameterValue.InvalidAccessGroupId"
//  INVALIDPARAMETERVALUE_INVALIDACCESSGROUPNAME = "InvalidParameterValue.InvalidAccessGroupName"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSGROUPNOTEXISTS = "ResourceNotFound.AccessGroupNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAccessGroupWithContext(ctx context.Context, request *ModifyAccessGroupRequest) (response *ModifyAccessGroupResponse, err error) {
    if request == nil {
        request = NewModifyAccessGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAccessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessRulesRequest() (request *ModifyAccessRulesRequest) {
    request = &ModifyAccessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyAccessRules")
    
    
    return
}

func NewModifyAccessRulesResponse() (response *ModifyAccessRulesResponse) {
    response = &ModifyAccessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccessRules
// This API is used to batch modify the attributes of permission rules, such as address, access mode, and priority. You must specify the permission rule IDs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSRULEADDRESS = "InvalidParameterValue.InvalidAccessRuleAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSRULENOTEXISTS = "ResourceNotFound.AccessRuleNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAccessRules(request *ModifyAccessRulesRequest) (response *ModifyAccessRulesResponse, err error) {
    if request == nil {
        request = NewModifyAccessRulesRequest()
    }
    
    response = NewModifyAccessRulesResponse()
    err = c.Send(request, response)
    return
}

// ModifyAccessRules
// This API is used to batch modify the attributes of permission rules, such as address, access mode, and priority. You must specify the permission rule IDs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDACCESSRULEADDRESS = "InvalidParameterValue.InvalidAccessRuleAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ACCESSRULENOTEXISTS = "ResourceNotFound.AccessRuleNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyAccessRulesWithContext(ctx context.Context, request *ModifyAccessRulesRequest) (response *ModifyAccessRulesResponse, err error) {
    if request == nil {
        request = NewModifyAccessRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAccessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFileSystemRequest() (request *ModifyFileSystemRequest) {
    request = &ModifyFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyFileSystem")
    
    
    return
}

func NewModifyFileSystemResponse() (response *ModifyFileSystemResponse) {
    response = &ModifyFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFileSystem
// This API is used to modify the attributes of a successfully created file system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUOTALESSTHANCURRENTUSED = "FailedOperation.QuotaLessThanCurrentUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCAPACITYQUOTA = "InvalidParameterValue.InvalidCapacityQuota"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMNAME = "InvalidParameterValue.InvalidFileSystemName"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyFileSystem(request *ModifyFileSystemRequest) (response *ModifyFileSystemResponse, err error) {
    if request == nil {
        request = NewModifyFileSystemRequest()
    }
    
    response = NewModifyFileSystemResponse()
    err = c.Send(request, response)
    return
}

// ModifyFileSystem
// This API is used to modify the attributes of a successfully created file system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUOTALESSTHANCURRENTUSED = "FailedOperation.QuotaLessThanCurrentUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDCAPACITYQUOTA = "InvalidParameterValue.InvalidCapacityQuota"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMNAME = "InvalidParameterValue.InvalidFileSystemName"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyFileSystemWithContext(ctx context.Context, request *ModifyFileSystemRequest) (response *ModifyFileSystemResponse, err error) {
    if request == nil {
        request = NewModifyFileSystemRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLifeCycleRulesRequest() (request *ModifyLifeCycleRulesRequest) {
    request = &ModifyLifeCycleRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyLifeCycleRules")
    
    
    return
}

func NewModifyLifeCycleRulesResponse() (response *ModifyLifeCycleRulesResponse) {
    response = &ModifyLifeCycleRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLifeCycleRules
// This API is used to batch modify the attributes of lifecycle rules, such as name, path, transition list, and status. You must specify the lifecycle rule IDs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLifeCycleRules(request *ModifyLifeCycleRulesRequest) (response *ModifyLifeCycleRulesResponse, err error) {
    if request == nil {
        request = NewModifyLifeCycleRulesRequest()
    }
    
    response = NewModifyLifeCycleRulesResponse()
    err = c.Send(request, response)
    return
}

// ModifyLifeCycleRules
// This API is used to batch modify the attributes of lifecycle rules, such as name, path, transition list, and status. You must specify the lifecycle rule IDs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyLifeCycleRulesWithContext(ctx context.Context, request *ModifyLifeCycleRulesRequest) (response *ModifyLifeCycleRulesResponse, err error) {
    if request == nil {
        request = NewModifyLifeCycleRulesRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyLifeCycleRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMountPointRequest() (request *ModifyMountPointRequest) {
    request = &ModifyMountPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyMountPoint")
    
    
    return
}

func NewModifyMountPointResponse() (response *ModifyMountPointResponse) {
    response = &ModifyMountPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMountPoint
// This API is used to modify the attributes of a mount point.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTNAME = "InvalidParameterValue.InvalidMountPointName"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyMountPoint(request *ModifyMountPointRequest) (response *ModifyMountPointResponse, err error) {
    if request == nil {
        request = NewModifyMountPointRequest()
    }
    
    response = NewModifyMountPointResponse()
    err = c.Send(request, response)
    return
}

// ModifyMountPoint
// This API is used to modify the attributes of a mount point.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTID = "InvalidParameterValue.InvalidMountPointId"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTPOINTNAME = "InvalidParameterValue.InvalidMountPointName"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MOUNTPOINTNOTEXISTS = "ResourceNotFound.MountPointNotExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyMountPointWithContext(ctx context.Context, request *ModifyMountPointRequest) (response *ModifyMountPointResponse, err error) {
    if request == nil {
        request = NewModifyMountPointRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyMountPointResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceTagsRequest() (request *ModifyResourceTagsRequest) {
    request = &ModifyResourceTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyResourceTags")
    
    
    return
}

func NewModifyResourceTagsResponse() (response *ModifyResourceTagsResponse) {
    response = &ModifyResourceTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyResourceTags
// This API is used to modify the list of resource tags by overwriting them all.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyResourceTags(request *ModifyResourceTagsRequest) (response *ModifyResourceTagsResponse, err error) {
    if request == nil {
        request = NewModifyResourceTagsRequest()
    }
    
    response = NewModifyResourceTagsResponse()
    err = c.Send(request, response)
    return
}

// ModifyResourceTags
// This API is used to modify the list of resource tags by overwriting them all.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTEXISTS = "ResourceNotFound.FileSystemNotExists"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyResourceTagsWithContext(ctx context.Context, request *ModifyResourceTagsRequest) (response *ModifyResourceTagsResponse, err error) {
    if request == nil {
        request = NewModifyResourceTagsRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyResourceTagsResponse()
    err = c.Send(request, response)
    return
}
