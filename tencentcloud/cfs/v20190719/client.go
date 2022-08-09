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

package v20190719

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-19"

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


func NewCreateCfsFileSystemRequest() (request *CreateCfsFileSystemRequest) {
    request = &CreateCfsFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "CreateCfsFileSystem")
    
    
    return
}

func NewCreateCfsFileSystemResponse() (response *CreateCfsFileSystemResponse) {
    response = &CreateCfsFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCfsFileSystem
// This API is used to create a file system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDRESOURCEPKGFAILED = "FailedOperation.BindResourcePkgFailed"
//  FAILEDOPERATION_CLIENTTOKENINUSE = "FailedOperation.ClientTokenInUse"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFSFAILED = "InternalError.CreateFsFailed"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLIENTTOKENLIMITEXCEEDED = "InvalidParameterValue.ClientTokenLimitExceeded"
//  INVALIDPARAMETERVALUE_FSNAMELIMITEXCEEDED = "InvalidParameterValue.FsNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDCLIENTTOKEN = "InvalidParameterValue.InvalidClientToken"
//  INVALIDPARAMETERVALUE_INVALIDENCRYPTED = "InvalidParameterValue.InvalidEncrypted"
//  INVALIDPARAMETERVALUE_INVALIDFSNAME = "InvalidParameterValue.InvalidFsName"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTTARGETIP = "InvalidParameterValue.InvalidMountTargetIp"
//  INVALIDPARAMETERVALUE_INVALIDNETINTERFACE = "InvalidParameterValue.InvalidNetInterface"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPID = "InvalidParameterValue.InvalidPgroupId"
//  INVALIDPARAMETERVALUE_INVALIDPROTOCOL = "InvalidParameterValue.InvalidProtocol"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_INVALIDRESOURCETAGS = "InvalidParameterValue.InvalidResourceTags"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  INVALIDPARAMETERVALUE_INVALIDSTORAGERESOURCEPKG = "InvalidParameterValue.InvalidStorageResourcePkg"
//  INVALIDPARAMETERVALUE_INVALIDSTORAGETYPE = "InvalidParameterValue.InvalidStorageType"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_INVALIDTAGKEY = "InvalidParameterValue.InvalidTagKey"
//  INVALIDPARAMETERVALUE_INVALIDVIP = "InvalidParameterValue.InvalidVip"
//  INVALIDPARAMETERVALUE_INVALIDVPCID = "InvalidParameterValue.InvalidVpcId"
//  INVALIDPARAMETERVALUE_INVALIDVPCPARAMETER = "InvalidParameterValue.InvalidVpcParameter"
//  INVALIDPARAMETERVALUE_INVALIDZONEID = "InvalidParameterValue.InvalidZoneId"
//  INVALIDPARAMETERVALUE_INVALIDZONEORZONEID = "InvalidParameterValue.InvalidZoneOrZoneId"
//  INVALIDPARAMETERVALUE_MISSINGKMSKEYID = "InvalidParameterValue.MissingKmsKeyId"
//  INVALIDPARAMETERVALUE_MISSINGSTORAGERESOURCEPKG = "InvalidParameterValue.MissingStorageResourcePkg"
//  INVALIDPARAMETERVALUE_MISSINGSUBNETIDORUNSUBNETID = "InvalidParameterValue.MissingSubnetidOrUnsubnetid"
//  INVALIDPARAMETERVALUE_MISSINGVPCPARAMETER = "InvalidParameterValue.MissingVpcParameter"
//  INVALIDPARAMETERVALUE_MISSINGVPCIDORUNVPCID = "InvalidParameterValue.MissingVpcidOrUnvpcid"
//  INVALIDPARAMETERVALUE_MISSINGZONEID = "InvalidParameterValue.MissingZoneId"
//  INVALIDPARAMETERVALUE_MISSINGZONEORZONEID = "InvalidParameterValue.MissingZoneOrZoneId"
//  INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUELIMITEXCEEDED = "InvalidParameterValue.TagValueLimitExceeded"
//  INVALIDPARAMETERVALUE_UNAVAILABLEREGION = "InvalidParameterValue.UnavailableRegion"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDPARAMETERVALUE_ZONEIDREGIONNOTMATCH = "InvalidParameterValue.ZoneIdRegionNotMatch"
//  RESOURCEINSUFFICIENT_FILESYSTEMLIMITEXCEEDED = "ResourceInsufficient.FileSystemLimitExceeded"
//  RESOURCEINSUFFICIENT_REGIONSOLDOUT = "ResourceInsufficient.RegionSoldOut"
//  RESOURCEINSUFFICIENT_SUBNETIPALLOCCUPIED = "ResourceInsufficient.SubnetIpAllOccupied"
//  RESOURCEINSUFFICIENT_TAGLIMITEXCEEDED = "ResourceInsufficient.TagLimitExceeded"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BASICNETINTERFACENOTSUPPORTED = "UnsupportedOperation.BasicNetInterfaceNotSupported"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsFileSystem(request *CreateCfsFileSystemRequest) (response *CreateCfsFileSystemResponse, err error) {
    return c.CreateCfsFileSystemWithContext(context.Background(), request)
}

// CreateCfsFileSystem
// This API is used to create a file system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDRESOURCEPKGFAILED = "FailedOperation.BindResourcePkgFailed"
//  FAILEDOPERATION_CLIENTTOKENINUSE = "FailedOperation.ClientTokenInUse"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFSFAILED = "InternalError.CreateFsFailed"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLIENTTOKENLIMITEXCEEDED = "InvalidParameterValue.ClientTokenLimitExceeded"
//  INVALIDPARAMETERVALUE_FSNAMELIMITEXCEEDED = "InvalidParameterValue.FsNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDCLIENTTOKEN = "InvalidParameterValue.InvalidClientToken"
//  INVALIDPARAMETERVALUE_INVALIDENCRYPTED = "InvalidParameterValue.InvalidEncrypted"
//  INVALIDPARAMETERVALUE_INVALIDFSNAME = "InvalidParameterValue.InvalidFsName"
//  INVALIDPARAMETERVALUE_INVALIDMOUNTTARGETIP = "InvalidParameterValue.InvalidMountTargetIp"
//  INVALIDPARAMETERVALUE_INVALIDNETINTERFACE = "InvalidParameterValue.InvalidNetInterface"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPID = "InvalidParameterValue.InvalidPgroupId"
//  INVALIDPARAMETERVALUE_INVALIDPROTOCOL = "InvalidParameterValue.InvalidProtocol"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_INVALIDRESOURCETAGS = "InvalidParameterValue.InvalidResourceTags"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  INVALIDPARAMETERVALUE_INVALIDSTORAGERESOURCEPKG = "InvalidParameterValue.InvalidStorageResourcePkg"
//  INVALIDPARAMETERVALUE_INVALIDSTORAGETYPE = "InvalidParameterValue.InvalidStorageType"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_INVALIDTAGKEY = "InvalidParameterValue.InvalidTagKey"
//  INVALIDPARAMETERVALUE_INVALIDVIP = "InvalidParameterValue.InvalidVip"
//  INVALIDPARAMETERVALUE_INVALIDVPCID = "InvalidParameterValue.InvalidVpcId"
//  INVALIDPARAMETERVALUE_INVALIDVPCPARAMETER = "InvalidParameterValue.InvalidVpcParameter"
//  INVALIDPARAMETERVALUE_INVALIDZONEID = "InvalidParameterValue.InvalidZoneId"
//  INVALIDPARAMETERVALUE_INVALIDZONEORZONEID = "InvalidParameterValue.InvalidZoneOrZoneId"
//  INVALIDPARAMETERVALUE_MISSINGKMSKEYID = "InvalidParameterValue.MissingKmsKeyId"
//  INVALIDPARAMETERVALUE_MISSINGSTORAGERESOURCEPKG = "InvalidParameterValue.MissingStorageResourcePkg"
//  INVALIDPARAMETERVALUE_MISSINGSUBNETIDORUNSUBNETID = "InvalidParameterValue.MissingSubnetidOrUnsubnetid"
//  INVALIDPARAMETERVALUE_MISSINGVPCPARAMETER = "InvalidParameterValue.MissingVpcParameter"
//  INVALIDPARAMETERVALUE_MISSINGVPCIDORUNVPCID = "InvalidParameterValue.MissingVpcidOrUnvpcid"
//  INVALIDPARAMETERVALUE_MISSINGZONEID = "InvalidParameterValue.MissingZoneId"
//  INVALIDPARAMETERVALUE_MISSINGZONEORZONEID = "InvalidParameterValue.MissingZoneOrZoneId"
//  INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUELIMITEXCEEDED = "InvalidParameterValue.TagValueLimitExceeded"
//  INVALIDPARAMETERVALUE_UNAVAILABLEREGION = "InvalidParameterValue.UnavailableRegion"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDPARAMETERVALUE_ZONEIDREGIONNOTMATCH = "InvalidParameterValue.ZoneIdRegionNotMatch"
//  RESOURCEINSUFFICIENT_FILESYSTEMLIMITEXCEEDED = "ResourceInsufficient.FileSystemLimitExceeded"
//  RESOURCEINSUFFICIENT_REGIONSOLDOUT = "ResourceInsufficient.RegionSoldOut"
//  RESOURCEINSUFFICIENT_SUBNETIPALLOCCUPIED = "ResourceInsufficient.SubnetIpAllOccupied"
//  RESOURCEINSUFFICIENT_TAGLIMITEXCEEDED = "ResourceInsufficient.TagLimitExceeded"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BASICNETINTERFACENOTSUPPORTED = "UnsupportedOperation.BasicNetInterfaceNotSupported"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsFileSystemWithContext(ctx context.Context, request *CreateCfsFileSystemRequest) (response *CreateCfsFileSystemResponse, err error) {
    if request == nil {
        request = NewCreateCfsFileSystemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCfsFileSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCfsFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCfsPGroupRequest() (request *CreateCfsPGroupRequest) {
    request = &CreateCfsPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "CreateCfsPGroup")
    
    
    return
}

func NewCreateCfsPGroupResponse() (response *CreateCfsPGroupResponse) {
    response = &CreateCfsPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCfsPGroup
// This API is used to create a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_DUPLICATEDPGROUPNAME = "InvalidParameterValue.DuplicatedPgroupName"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPNAME = "InvalidParameterValue.InvalidPgroupName"
//  INVALIDPARAMETERVALUE_MISSINGPGROUPNAME = "InvalidParameterValue.MissingPgroupName"
//  INVALIDPARAMETERVALUE_PGROUPDESCINFOLIMITEXCEEDED = "InvalidParameterValue.PgroupDescinfoLimitExceeded"
//  INVALIDPARAMETERVALUE_PGROUPNAMELIMITEXCEEDED = "InvalidParameterValue.PgroupNameLimitExceeded"
//  RESOURCEINSUFFICIENT_PGROUPNUMBERLIMITEXCEEDED = "ResourceInsufficient.PgroupNumberLimitExceeded"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsPGroup(request *CreateCfsPGroupRequest) (response *CreateCfsPGroupResponse, err error) {
    return c.CreateCfsPGroupWithContext(context.Background(), request)
}

// CreateCfsPGroup
// This API is used to create a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_DUPLICATEDPGROUPNAME = "InvalidParameterValue.DuplicatedPgroupName"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPNAME = "InvalidParameterValue.InvalidPgroupName"
//  INVALIDPARAMETERVALUE_MISSINGPGROUPNAME = "InvalidParameterValue.MissingPgroupName"
//  INVALIDPARAMETERVALUE_PGROUPDESCINFOLIMITEXCEEDED = "InvalidParameterValue.PgroupDescinfoLimitExceeded"
//  INVALIDPARAMETERVALUE_PGROUPNAMELIMITEXCEEDED = "InvalidParameterValue.PgroupNameLimitExceeded"
//  RESOURCEINSUFFICIENT_PGROUPNUMBERLIMITEXCEEDED = "ResourceInsufficient.PgroupNumberLimitExceeded"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsPGroupWithContext(ctx context.Context, request *CreateCfsPGroupRequest) (response *CreateCfsPGroupResponse, err error) {
    if request == nil {
        request = NewCreateCfsPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCfsPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCfsPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCfsRuleRequest() (request *CreateCfsRuleRequest) {
    request = &CreateCfsRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "CreateCfsRule")
    
    
    return
}

func NewCreateCfsRuleResponse() (response *CreateCfsRuleResponse) {
    response = &CreateCfsRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCfsRule
// This API is used to create a permission group rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"
//  INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"
//  RESOURCEINSUFFICIENT_RULELIMITEXCEEDED = "ResourceInsufficient.RuleLimitExceeded"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsRule(request *CreateCfsRuleRequest) (response *CreateCfsRuleResponse, err error) {
    return c.CreateCfsRuleWithContext(context.Background(), request)
}

// CreateCfsRule
// This API is used to create a permission group rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"
//  INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"
//  RESOURCEINSUFFICIENT_RULELIMITEXCEEDED = "ResourceInsufficient.RuleLimitExceeded"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsRuleWithContext(ctx context.Context, request *CreateCfsRuleRequest) (response *CreateCfsRuleResponse, err error) {
    if request == nil {
        request = NewCreateCfsRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCfsRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCfsRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCfsFileSystemRequest() (request *DeleteCfsFileSystemRequest) {
    request = &DeleteCfsFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteCfsFileSystem")
    
    
    return
}

func NewDeleteCfsFileSystemResponse() (response *DeleteCfsFileSystemResponse) {
    response = &DeleteCfsFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCfsFileSystem
// This API is used to delete a file system.
//
// error code that may be returned:
//  FAILEDOPERATION_MOUNTTARGETEXISTS = "FailedOperation.MountTargetExists"
//  FAILEDOPERATION_UNTAGRESOURCEFAILED = "FailedOperation.UntagResourceFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsFileSystem(request *DeleteCfsFileSystemRequest) (response *DeleteCfsFileSystemResponse, err error) {
    return c.DeleteCfsFileSystemWithContext(context.Background(), request)
}

// DeleteCfsFileSystem
// This API is used to delete a file system.
//
// error code that may be returned:
//  FAILEDOPERATION_MOUNTTARGETEXISTS = "FailedOperation.MountTargetExists"
//  FAILEDOPERATION_UNTAGRESOURCEFAILED = "FailedOperation.UntagResourceFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsFileSystemWithContext(ctx context.Context, request *DeleteCfsFileSystemRequest) (response *DeleteCfsFileSystemResponse, err error) {
    if request == nil {
        request = NewDeleteCfsFileSystemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCfsFileSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCfsFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCfsPGroupRequest() (request *DeleteCfsPGroupRequest) {
    request = &DeleteCfsPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteCfsPGroup")
    
    
    return
}

func NewDeleteCfsPGroupResponse() (response *DeleteCfsPGroupResponse) {
    response = &DeleteCfsPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCfsPGroup
// This API is used to delete a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsPGroup(request *DeleteCfsPGroupRequest) (response *DeleteCfsPGroupResponse, err error) {
    return c.DeleteCfsPGroupWithContext(context.Background(), request)
}

// DeleteCfsPGroup
// This API is used to delete a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsPGroupWithContext(ctx context.Context, request *DeleteCfsPGroupRequest) (response *DeleteCfsPGroupResponse, err error) {
    if request == nil {
        request = NewDeleteCfsPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCfsPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCfsPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCfsRuleRequest() (request *DeleteCfsRuleRequest) {
    request = &DeleteCfsRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteCfsRule")
    
    
    return
}

func NewDeleteCfsRuleResponse() (response *DeleteCfsRuleResponse) {
    response = &DeleteCfsRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCfsRule
// This API is used to delete a permission group rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"
//  INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  RESOURCENOTFOUND_RULENOTFOUND = "ResourceNotFound.RuleNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsRule(request *DeleteCfsRuleRequest) (response *DeleteCfsRuleResponse, err error) {
    return c.DeleteCfsRuleWithContext(context.Background(), request)
}

// DeleteCfsRule
// This API is used to delete a permission group rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"
//  INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  RESOURCENOTFOUND_RULENOTFOUND = "ResourceNotFound.RuleNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsRuleWithContext(ctx context.Context, request *DeleteCfsRuleRequest) (response *DeleteCfsRuleResponse, err error) {
    if request == nil {
        request = NewDeleteCfsRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCfsRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCfsRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMountTargetRequest() (request *DeleteMountTargetRequest) {
    request = &DeleteMountTargetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteMountTarget")
    
    
    return
}

func NewDeleteMountTargetResponse() (response *DeleteMountTargetResponse) {
    response = &DeleteMountTargetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMountTarget
// This API is used to delete a mount target.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteMountTarget(request *DeleteMountTargetRequest) (response *DeleteMountTargetResponse, err error) {
    return c.DeleteMountTargetWithContext(context.Background(), request)
}

// DeleteMountTarget
// This API is used to delete a mount target.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteMountTargetWithContext(ctx context.Context, request *DeleteMountTargetRequest) (response *DeleteMountTargetResponse, err error) {
    if request == nil {
        request = NewDeleteMountTargetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMountTarget require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMountTargetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableZoneInfoRequest() (request *DescribeAvailableZoneInfoRequest) {
    request = &DescribeAvailableZoneInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeAvailableZoneInfo")
    
    
    return
}

func NewDescribeAvailableZoneInfoResponse() (response *DescribeAvailableZoneInfoResponse) {
    response = &DescribeAvailableZoneInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAvailableZoneInfo
// This API is used to query the availability of a region.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeAvailableZoneInfo(request *DescribeAvailableZoneInfoRequest) (response *DescribeAvailableZoneInfoResponse, err error) {
    return c.DescribeAvailableZoneInfoWithContext(context.Background(), request)
}

// DescribeAvailableZoneInfo
// This API is used to query the availability of a region.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeAvailableZoneInfoWithContext(ctx context.Context, request *DescribeAvailableZoneInfoRequest) (response *DescribeAvailableZoneInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableZoneInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableZoneInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableZoneInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsFileSystemClientsRequest() (request *DescribeCfsFileSystemClientsRequest) {
    request = &DescribeCfsFileSystemClientsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsFileSystemClients")
    
    
    return
}

func NewDescribeCfsFileSystemClientsResponse() (response *DescribeCfsFileSystemClientsResponse) {
    response = &DescribeCfsFileSystemClientsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfsFileSystemClients
// This API is used to query clients on which this file system is mounted. To do so, the client needs to have the CFS monitoring plugin installed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCfsFileSystemClients(request *DescribeCfsFileSystemClientsRequest) (response *DescribeCfsFileSystemClientsResponse, err error) {
    return c.DescribeCfsFileSystemClientsWithContext(context.Background(), request)
}

// DescribeCfsFileSystemClients
// This API is used to query clients on which this file system is mounted. To do so, the client needs to have the CFS monitoring plugin installed.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCfsFileSystemClientsWithContext(ctx context.Context, request *DescribeCfsFileSystemClientsRequest) (response *DescribeCfsFileSystemClientsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsFileSystemClientsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsFileSystemClients require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsFileSystemClientsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsFileSystemsRequest() (request *DescribeCfsFileSystemsRequest) {
    request = &DescribeCfsFileSystemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsFileSystems")
    
    
    return
}

func NewDescribeCfsFileSystemsResponse() (response *DescribeCfsFileSystemsResponse) {
    response = &DescribeCfsFileSystemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfsFileSystems
// This API is used to query file systems.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMIDORREGION = "InvalidParameterValue.MissingFileSystemIdOrRegion"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  INVALIDPARAMETERVALUE_TAGKEYFILTERLIMITEXCEEDED = "InvalidParameterValue.TagKeyFilterLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"
//  INVALIDPARAMETERVALUE_UNAVAILABLEREGION = "InvalidParameterValue.UnavailableRegion"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsFileSystems(request *DescribeCfsFileSystemsRequest) (response *DescribeCfsFileSystemsResponse, err error) {
    return c.DescribeCfsFileSystemsWithContext(context.Background(), request)
}

// DescribeCfsFileSystems
// This API is used to query file systems.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMIDORREGION = "InvalidParameterValue.MissingFileSystemIdOrRegion"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  INVALIDPARAMETERVALUE_TAGKEYFILTERLIMITEXCEEDED = "InvalidParameterValue.TagKeyFilterLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"
//  INVALIDPARAMETERVALUE_UNAVAILABLEREGION = "InvalidParameterValue.UnavailableRegion"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsFileSystemsWithContext(ctx context.Context, request *DescribeCfsFileSystemsRequest) (response *DescribeCfsFileSystemsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsFileSystemsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsFileSystems require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsFileSystemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsPGroupsRequest() (request *DescribeCfsPGroupsRequest) {
    request = &DescribeCfsPGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsPGroups")
    
    
    return
}

func NewDescribeCfsPGroupsResponse() (response *DescribeCfsPGroupsResponse) {
    response = &DescribeCfsPGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfsPGroups
// This API is used to query the list of permission groups.
//
// error code that may be returned:
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsPGroups(request *DescribeCfsPGroupsRequest) (response *DescribeCfsPGroupsResponse, err error) {
    return c.DescribeCfsPGroupsWithContext(context.Background(), request)
}

// DescribeCfsPGroups
// This API is used to query the list of permission groups.
//
// error code that may be returned:
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsPGroupsWithContext(ctx context.Context, request *DescribeCfsPGroupsRequest) (response *DescribeCfsPGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsPGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsPGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsPGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsRulesRequest() (request *DescribeCfsRulesRequest) {
    request = &DescribeCfsRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsRules")
    
    
    return
}

func NewDescribeCfsRulesResponse() (response *DescribeCfsRulesResponse) {
    response = &DescribeCfsRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfsRules
// This API is used to query the list of permission group rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsRules(request *DescribeCfsRulesRequest) (response *DescribeCfsRulesResponse, err error) {
    return c.DescribeCfsRulesWithContext(context.Background(), request)
}

// DescribeCfsRules
// This API is used to query the list of permission group rules.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsRulesWithContext(ctx context.Context, request *DescribeCfsRulesRequest) (response *DescribeCfsRulesResponse, err error) {
    if request == nil {
        request = NewDescribeCfsRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsServiceStatusRequest() (request *DescribeCfsServiceStatusRequest) {
    request = &DescribeCfsServiceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsServiceStatus")
    
    
    return
}

func NewDescribeCfsServiceStatusResponse() (response *DescribeCfsServiceStatusResponse) {
    response = &DescribeCfsServiceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCfsServiceStatus
// This API is used to query the status of the CFS service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsServiceStatus(request *DescribeCfsServiceStatusRequest) (response *DescribeCfsServiceStatusResponse, err error) {
    return c.DescribeCfsServiceStatusWithContext(context.Background(), request)
}

// DescribeCfsServiceStatus
// This API is used to query the status of the CFS service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsServiceStatusWithContext(ctx context.Context, request *DescribeCfsServiceStatusRequest) (response *DescribeCfsServiceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeCfsServiceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsServiceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsServiceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMountTargetsRequest() (request *DescribeMountTargetsRequest) {
    request = &DescribeMountTargetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeMountTargets")
    
    
    return
}

func NewDescribeMountTargetsResponse() (response *DescribeMountTargetsResponse) {
    response = &DescribeMountTargetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMountTargets
// This API is used to query the mount targets of a file system.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeMountTargets(request *DescribeMountTargetsRequest) (response *DescribeMountTargetsResponse, err error) {
    return c.DescribeMountTargetsWithContext(context.Background(), request)
}

// DescribeMountTargets
// This API is used to query the mount targets of a file system.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeMountTargetsWithContext(ctx context.Context, request *DescribeMountTargetsRequest) (response *DescribeMountTargetsResponse, err error) {
    if request == nil {
        request = NewDescribeMountTargetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMountTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMountTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewSignUpCfsServiceRequest() (request *SignUpCfsServiceRequest) {
    request = &SignUpCfsServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "SignUpCfsService")
    
    
    return
}

func NewSignUpCfsServiceResponse() (response *SignUpCfsServiceResponse) {
    response = &SignUpCfsServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SignUpCfsService
// This API is used to activate the CFS service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) SignUpCfsService(request *SignUpCfsServiceRequest) (response *SignUpCfsServiceResponse, err error) {
    return c.SignUpCfsServiceWithContext(context.Background(), request)
}

// SignUpCfsService
// This API is used to activate the CFS service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) SignUpCfsServiceWithContext(ctx context.Context, request *SignUpCfsServiceRequest) (response *SignUpCfsServiceResponse, err error) {
    if request == nil {
        request = NewSignUpCfsServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SignUpCfsService require credential")
    }

    request.SetContext(ctx)
    
    response = NewSignUpCfsServiceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCfsFileSystemNameRequest() (request *UpdateCfsFileSystemNameRequest) {
    request = &UpdateCfsFileSystemNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsFileSystemName")
    
    
    return
}

func NewUpdateCfsFileSystemNameResponse() (response *UpdateCfsFileSystemNameResponse) {
    response = &UpdateCfsFileSystemNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCfsFileSystemName
// This API is used to update a file system name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FSNAMELIMITEXCEEDED = "InvalidParameterValue.FsNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSNAME = "InvalidParameterValue.InvalidFsName"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsFileSystemName(request *UpdateCfsFileSystemNameRequest) (response *UpdateCfsFileSystemNameResponse, err error) {
    return c.UpdateCfsFileSystemNameWithContext(context.Background(), request)
}

// UpdateCfsFileSystemName
// This API is used to update a file system name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FSNAMELIMITEXCEEDED = "InvalidParameterValue.FsNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSNAME = "InvalidParameterValue.InvalidFsName"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsFileSystemNameWithContext(ctx context.Context, request *UpdateCfsFileSystemNameRequest) (response *UpdateCfsFileSystemNameResponse, err error) {
    if request == nil {
        request = NewUpdateCfsFileSystemNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsFileSystemName require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsFileSystemNameResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCfsFileSystemPGroupRequest() (request *UpdateCfsFileSystemPGroupRequest) {
    request = &UpdateCfsFileSystemPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsFileSystemPGroup")
    
    
    return
}

func NewUpdateCfsFileSystemPGroupResponse() (response *UpdateCfsFileSystemPGroupResponse) {
    response = &UpdateCfsFileSystemPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCfsFileSystemPGroup
// This API is used to update the permission group used by a file system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPID = "InvalidParameterValue.InvalidPgroupId"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsFileSystemPGroup(request *UpdateCfsFileSystemPGroupRequest) (response *UpdateCfsFileSystemPGroupResponse, err error) {
    return c.UpdateCfsFileSystemPGroupWithContext(context.Background(), request)
}

// UpdateCfsFileSystemPGroup
// This API is used to update the permission group used by a file system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPID = "InvalidParameterValue.InvalidPgroupId"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsFileSystemPGroupWithContext(ctx context.Context, request *UpdateCfsFileSystemPGroupRequest) (response *UpdateCfsFileSystemPGroupResponse, err error) {
    if request == nil {
        request = NewUpdateCfsFileSystemPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsFileSystemPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsFileSystemPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCfsFileSystemSizeLimitRequest() (request *UpdateCfsFileSystemSizeLimitRequest) {
    request = &UpdateCfsFileSystemSizeLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsFileSystemSizeLimit")
    
    
    return
}

func NewUpdateCfsFileSystemSizeLimitResponse() (response *UpdateCfsFileSystemSizeLimitResponse) {
    response = &UpdateCfsFileSystemSizeLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCfsFileSystemSizeLimit
// This API is used to update the capacity limit of a file system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FSSIZELIMITEXCEEDED = "InvalidParameterValue.FsSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSIZELIMIT = "InvalidParameterValue.InvalidFsSizeLimit"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsFileSystemSizeLimit(request *UpdateCfsFileSystemSizeLimitRequest) (response *UpdateCfsFileSystemSizeLimitResponse, err error) {
    return c.UpdateCfsFileSystemSizeLimitWithContext(context.Background(), request)
}

// UpdateCfsFileSystemSizeLimit
// This API is used to update the capacity limit of a file system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FSSIZELIMITEXCEEDED = "InvalidParameterValue.FsSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSIZELIMIT = "InvalidParameterValue.InvalidFsSizeLimit"
//  INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"
//  INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsFileSystemSizeLimitWithContext(ctx context.Context, request *UpdateCfsFileSystemSizeLimitRequest) (response *UpdateCfsFileSystemSizeLimitResponse, err error) {
    if request == nil {
        request = NewUpdateCfsFileSystemSizeLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsFileSystemSizeLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsFileSystemSizeLimitResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCfsPGroupRequest() (request *UpdateCfsPGroupRequest) {
    request = &UpdateCfsPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsPGroup")
    
    
    return
}

func NewUpdateCfsPGroupResponse() (response *UpdateCfsPGroupResponse) {
    response = &UpdateCfsPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCfsPGroup
// This API is used to update the information of a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MOUNTTARGETEXISTS = "FailedOperation.MountTargetExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEDPGROUPNAME = "InvalidParameterValue.DuplicatedPgroupName"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPNAME = "InvalidParameterValue.InvalidPgroupName"
//  INVALIDPARAMETERVALUE_MISSINGNAMEORDESCINFO = "InvalidParameterValue.MissingNameOrDescinfo"
//  INVALIDPARAMETERVALUE_MISSINGPGROUPNAME = "InvalidParameterValue.MissingPgroupName"
//  INVALIDPARAMETERVALUE_PGROUPDESCINFOLIMITEXCEEDED = "InvalidParameterValue.PgroupDescinfoLimitExceeded"
//  INVALIDPARAMETERVALUE_PGROUPNAMELIMITEXCEEDED = "InvalidParameterValue.PgroupNameLimitExceeded"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsPGroup(request *UpdateCfsPGroupRequest) (response *UpdateCfsPGroupResponse, err error) {
    return c.UpdateCfsPGroupWithContext(context.Background(), request)
}

// UpdateCfsPGroup
// This API is used to update the information of a permission group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MOUNTTARGETEXISTS = "FailedOperation.MountTargetExists"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEDPGROUPNAME = "InvalidParameterValue.DuplicatedPgroupName"
//  INVALIDPARAMETERVALUE_INVALIDPGROUPNAME = "InvalidParameterValue.InvalidPgroupName"
//  INVALIDPARAMETERVALUE_MISSINGNAMEORDESCINFO = "InvalidParameterValue.MissingNameOrDescinfo"
//  INVALIDPARAMETERVALUE_MISSINGPGROUPNAME = "InvalidParameterValue.MissingPgroupName"
//  INVALIDPARAMETERVALUE_PGROUPDESCINFOLIMITEXCEEDED = "InvalidParameterValue.PgroupDescinfoLimitExceeded"
//  INVALIDPARAMETERVALUE_PGROUPNAMELIMITEXCEEDED = "InvalidParameterValue.PgroupNameLimitExceeded"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsPGroupWithContext(ctx context.Context, request *UpdateCfsPGroupRequest) (response *UpdateCfsPGroupResponse, err error) {
    if request == nil {
        request = NewUpdateCfsPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCfsRuleRequest() (request *UpdateCfsRuleRequest) {
    request = &UpdateCfsRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsRule")
    
    
    return
}

func NewUpdateCfsRuleResponse() (response *UpdateCfsRuleResponse) {
    response = &UpdateCfsRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCfsRule
// This API is used to update a permission rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  FAILEDOPERATION_PGROUPLINKCFSV10 = "FailedOperation.PgroupLinkCfsv10"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"
//  INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"
//  INVALIDPARAMETERVALUE_RULENOTMATCHPGROUP = "InvalidParameterValue.RuleNotMatchPgroup"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsRule(request *UpdateCfsRuleRequest) (response *UpdateCfsRuleResponse, err error) {
    return c.UpdateCfsRuleWithContext(context.Background(), request)
}

// UpdateCfsRule
// This API is used to update a permission rule.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"
//  FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"
//  FAILEDOPERATION_PGROUPLINKCFSV10 = "FailedOperation.PgroupLinkCfsv10"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"
//  INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"
//  INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"
//  INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"
//  INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"
//  INVALIDPARAMETERVALUE_RULENOTMATCHPGROUP = "InvalidParameterValue.RuleNotMatchPgroup"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateCfsRuleWithContext(ctx context.Context, request *UpdateCfsRuleRequest) (response *UpdateCfsRuleResponse, err error) {
    if request == nil {
        request = NewUpdateCfsRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsRuleResponse()
    err = c.Send(request, response)
    return
}
