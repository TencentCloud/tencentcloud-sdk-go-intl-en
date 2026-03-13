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

package v20190719

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
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


func NewApplyPathLifecyclePolicyRequest() (request *ApplyPathLifecyclePolicyRequest) {
    request = &ApplyPathLifecyclePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "ApplyPathLifecyclePolicy")
    
    
    return
}

func NewApplyPathLifecyclePolicyResponse() (response *ApplyPathLifecyclePolicyResponse) {
    response = &ApplyPathLifecyclePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyPathLifecyclePolicy
// Configure the directory list associated with the lifecycle policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_BINDINFREQUENTACCESSFIRST = "InvalidParameterValue.BindInfrequentaccessFirst"
//  INVALIDPARAMETERVALUE_EXTERNALSTORAGESUPPORTDATAFLOWONLY = "InvalidParameterValue.ExternalStorageSupportDataflowOnly"
//  INVALIDPARAMETERVALUE_INVALIDPOLICYFSPATH = "InvalidParameterValue.InvalidPolicyFsPath"
//  INVALIDPARAMETERVALUE_PATHUSEDINDATAFLOW = "InvalidParameterValue.PathUsedInDataflow"
//  RESOURCEINSUFFICIENT_POLICYFSLIMITEXCEEDED = "ResourceInsufficient.PolicyFsLimitExceeded"
//  RESOURCEINSUFFICIENT_POLICYFSPATHLIMITEXCEEDED = "ResourceInsufficient.PolicyFsPathLimitExceeded"
//  RESOURCEINSUFFICIENT_POLICYLIMITEXCEEDED = "ResourceInsufficient.PolicyLimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ApplyPathLifecyclePolicy(request *ApplyPathLifecyclePolicyRequest) (response *ApplyPathLifecyclePolicyResponse, err error) {
    return c.ApplyPathLifecyclePolicyWithContext(context.Background(), request)
}

// ApplyPathLifecyclePolicy
// Configure the directory list associated with the lifecycle policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_BINDINFREQUENTACCESSFIRST = "InvalidParameterValue.BindInfrequentaccessFirst"
//  INVALIDPARAMETERVALUE_EXTERNALSTORAGESUPPORTDATAFLOWONLY = "InvalidParameterValue.ExternalStorageSupportDataflowOnly"
//  INVALIDPARAMETERVALUE_INVALIDPOLICYFSPATH = "InvalidParameterValue.InvalidPolicyFsPath"
//  INVALIDPARAMETERVALUE_PATHUSEDINDATAFLOW = "InvalidParameterValue.PathUsedInDataflow"
//  RESOURCEINSUFFICIENT_POLICYFSLIMITEXCEEDED = "ResourceInsufficient.PolicyFsLimitExceeded"
//  RESOURCEINSUFFICIENT_POLICYFSPATHLIMITEXCEEDED = "ResourceInsufficient.PolicyFsPathLimitExceeded"
//  RESOURCEINSUFFICIENT_POLICYLIMITEXCEEDED = "ResourceInsufficient.PolicyLimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ApplyPathLifecyclePolicyWithContext(ctx context.Context, request *ApplyPathLifecyclePolicyRequest) (response *ApplyPathLifecyclePolicyResponse, err error) {
    if request == nil {
        request = NewApplyPathLifecyclePolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "ApplyPathLifecyclePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyPathLifecyclePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyPathLifecyclePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewBindAutoSnapshotPolicyRequest() (request *BindAutoSnapshotPolicyRequest) {
    request = &BindAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "BindAutoSnapshotPolicy")
    
    
    return
}

func NewBindAutoSnapshotPolicyResponse() (response *BindAutoSnapshotPolicyResponse) {
    response = &BindAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindAutoSnapshotPolicy
// This API is used to bind one or multiple file systems to a snapshot policy. A file system can be bound to only one policy.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE_AUTOPOLICYNOTFOUND = "InvalidParameterValue.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDSNAPPOLICYSTATUS = "InvalidParameterValue.InvalidSnapPolicyStatus"
//  RESOURCEINSUFFICIENT_SNAPSHOTSIZELIMITEXCEEDED = "ResourceInsufficient.SnapshotSizeLimitExceeded"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) BindAutoSnapshotPolicy(request *BindAutoSnapshotPolicyRequest) (response *BindAutoSnapshotPolicyResponse, err error) {
    return c.BindAutoSnapshotPolicyWithContext(context.Background(), request)
}

// BindAutoSnapshotPolicy
// This API is used to bind one or multiple file systems to a snapshot policy. A file system can be bound to only one policy.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE_AUTOPOLICYNOTFOUND = "InvalidParameterValue.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDSNAPPOLICYSTATUS = "InvalidParameterValue.InvalidSnapPolicyStatus"
//  RESOURCEINSUFFICIENT_SNAPSHOTSIZELIMITEXCEEDED = "ResourceInsufficient.SnapshotSizeLimitExceeded"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) BindAutoSnapshotPolicyWithContext(ctx context.Context, request *BindAutoSnapshotPolicyRequest) (response *BindAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewBindAutoSnapshotPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "BindAutoSnapshotPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoSnapshotPolicyRequest() (request *CreateAutoSnapshotPolicyRequest) {
    request = &CreateAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "CreateAutoSnapshotPolicy")
    
    
    return
}

func NewCreateAutoSnapshotPolicyResponse() (response *CreateAutoSnapshotPolicyResponse) {
    response = &CreateAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAutoSnapshotPolicy
// This API is used to create a scheduled snapshot policy.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDALIVEDDAYS = "InvalidParameter.InvalidAlivedDays"
//  INVALIDPARAMETER_INVALIDPARAMDAYOFWEEK = "InvalidParameter.InvalidParamDayofWeek"
//  INVALIDPARAMETER_INVALIDPARAMHOUR = "InvalidParameter.InvalidParamHour"
//  INVALIDPARAMETER_INVALIDSNAPSHOTPOLICYNAME = "InvalidParameter.InvalidSnapshotPolicyName"
//  INVALIDPARAMETER_MISSINGPOLICYPARAM = "InvalidParameter.MissingPolicyParam"
//  INVALIDPARAMETER_SNAPSHOTPOLICYNAMELIMITEXCEEDED = "InvalidParameter.SnapshotPolicyNameLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUTOPOLICYNOTFOUND = "InvalidParameterValue.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE_INVALIDALIVEDAYS = "InvalidParameterValue.InvalidAliveDays"
//  INVALIDPARAMETERVALUE_INVALIDDESTINATIONREGIONS = "InvalidParameterValue.InvalidDestinationRegions"
//  INVALIDPARAMETERVALUE_INVALIDPARAMDAYOFMONTH = "InvalidParameterValue.InvalidParamDayOfMonth"
//  INVALIDPARAMETERVALUE_INVALIDPARAMDAYOFWEEK = "InvalidParameterValue.InvalidParamDayOfWeek"
//  INVALIDPARAMETERVALUE_INVALIDPARAMINTERVALDAYS = "InvalidParameterValue.InvalidParamIntervalDays"
//  INVALIDPARAMETERVALUE_INVALIDSNAPPOLICYSTATUS = "InvalidParameterValue.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTNAME = "InvalidParameterValue.InvalidSnapshotName"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTPOLICYNAME = "InvalidParameterValue.InvalidSnapshotPolicyName"
//  INVALIDPARAMETERVALUE_MISSINGSNAPNAMEORALIVEDAY = "InvalidParameterValue.MissingSnapNameOrAliveDay"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMELIMITEXCEEDED = "InvalidParameterValue.SnapshotNameLimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTPOLICYNAMELIMITEXCEEDED = "InvalidParameterValue.SnapshotPolicyNameLimitExceeded"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (response *CreateAutoSnapshotPolicyResponse, err error) {
    return c.CreateAutoSnapshotPolicyWithContext(context.Background(), request)
}

// CreateAutoSnapshotPolicy
// This API is used to create a scheduled snapshot policy.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDALIVEDDAYS = "InvalidParameter.InvalidAlivedDays"
//  INVALIDPARAMETER_INVALIDPARAMDAYOFWEEK = "InvalidParameter.InvalidParamDayofWeek"
//  INVALIDPARAMETER_INVALIDPARAMHOUR = "InvalidParameter.InvalidParamHour"
//  INVALIDPARAMETER_INVALIDSNAPSHOTPOLICYNAME = "InvalidParameter.InvalidSnapshotPolicyName"
//  INVALIDPARAMETER_MISSINGPOLICYPARAM = "InvalidParameter.MissingPolicyParam"
//  INVALIDPARAMETER_SNAPSHOTPOLICYNAMELIMITEXCEEDED = "InvalidParameter.SnapshotPolicyNameLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUTOPOLICYNOTFOUND = "InvalidParameterValue.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE_INVALIDALIVEDAYS = "InvalidParameterValue.InvalidAliveDays"
//  INVALIDPARAMETERVALUE_INVALIDDESTINATIONREGIONS = "InvalidParameterValue.InvalidDestinationRegions"
//  INVALIDPARAMETERVALUE_INVALIDPARAMDAYOFMONTH = "InvalidParameterValue.InvalidParamDayOfMonth"
//  INVALIDPARAMETERVALUE_INVALIDPARAMDAYOFWEEK = "InvalidParameterValue.InvalidParamDayOfWeek"
//  INVALIDPARAMETERVALUE_INVALIDPARAMINTERVALDAYS = "InvalidParameterValue.InvalidParamIntervalDays"
//  INVALIDPARAMETERVALUE_INVALIDSNAPPOLICYSTATUS = "InvalidParameterValue.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTNAME = "InvalidParameterValue.InvalidSnapshotName"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTPOLICYNAME = "InvalidParameterValue.InvalidSnapshotPolicyName"
//  INVALIDPARAMETERVALUE_MISSINGSNAPNAMEORALIVEDAY = "InvalidParameterValue.MissingSnapNameOrAliveDay"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMELIMITEXCEEDED = "InvalidParameterValue.SnapshotNameLimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTPOLICYNAMELIMITEXCEEDED = "InvalidParameterValue.SnapshotPolicyNameLimitExceeded"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateAutoSnapshotPolicyWithContext(ctx context.Context, request *CreateAutoSnapshotPolicyRequest) (response *CreateAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAutoSnapshotPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "CreateAutoSnapshotPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
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
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE_CLIENTTOKENLIMITEXCEEDED = "InvalidParameterValue.ClientTokenLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATEDTAGKEY = "InvalidParameterValue.DuplicatedTagKey"
//  INVALIDPARAMETERVALUE_FSNAMELIMITEXCEEDED = "InvalidParameterValue.FsNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDCLIENTTOKEN = "InvalidParameterValue.InvalidClientToken"
//  INVALIDPARAMETERVALUE_INVALIDENCRYPTED = "InvalidParameterValue.InvalidEncrypted"
//  INVALIDPARAMETERVALUE_INVALIDFSNAME = "InvalidParameterValue.InvalidFsName"
//  INVALIDPARAMETERVALUE_INVALIDMETATYPE = "InvalidParameterValue.InvalidMetaType"
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
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUE = "InvalidParameterValue.InvalidTagValue"
//  INVALIDPARAMETERVALUE_INVALIDTURBOCAPACITY = "InvalidParameterValue.InvalidTurboCapacity"
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
//  INVALIDPARAMETERVALUE_TAGKEYFILTERLIMITEXCEEDED = "InvalidParameterValue.TagKeyFilterLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUEFILTERLIMITEXCEEDED = "InvalidParameterValue.TagValueFilterLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUELIMITEXCEEDED = "InvalidParameterValue.TagValueLimitExceeded"
//  INVALIDPARAMETERVALUE_UNAVAILABLEREGION = "InvalidParameterValue.UnavailableRegion"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDPARAMETERVALUE_ZONEIDREGIONNOTMATCH = "InvalidParameterValue.ZoneIdRegionNotMatch"
//  RESOURCEINSUFFICIENT_FILESYSTEMLIMITEXCEEDED = "ResourceInsufficient.FileSystemLimitExceeded"
//  RESOURCEINSUFFICIENT_REGIONSOLDOUT = "ResourceInsufficient.RegionSoldOut"
//  RESOURCEINSUFFICIENT_SUBNETIPALLOCCUPIED = "ResourceInsufficient.SubnetIpAllOccupied"
//  RESOURCEINSUFFICIENT_TAGLIMITEXCEEDED = "ResourceInsufficient.TagLimitExceeded"
//  RESOURCEINSUFFICIENT_TAGQUOTASEXCEEDED = "ResourceInsufficient.TagQuotasExceeded"
//  RESOURCEINSUFFICIENT_TURBOSPECIALCAPACITYFILESYSTEMCOUNTLIMIT = "ResourceInsufficient.TurboSpecialCapacityFileSystemCountLimit"
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
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE_CLIENTTOKENLIMITEXCEEDED = "InvalidParameterValue.ClientTokenLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATEDTAGKEY = "InvalidParameterValue.DuplicatedTagKey"
//  INVALIDPARAMETERVALUE_FSNAMELIMITEXCEEDED = "InvalidParameterValue.FsNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDCLIENTTOKEN = "InvalidParameterValue.InvalidClientToken"
//  INVALIDPARAMETERVALUE_INVALIDENCRYPTED = "InvalidParameterValue.InvalidEncrypted"
//  INVALIDPARAMETERVALUE_INVALIDFSNAME = "InvalidParameterValue.InvalidFsName"
//  INVALIDPARAMETERVALUE_INVALIDMETATYPE = "InvalidParameterValue.InvalidMetaType"
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
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUE = "InvalidParameterValue.InvalidTagValue"
//  INVALIDPARAMETERVALUE_INVALIDTURBOCAPACITY = "InvalidParameterValue.InvalidTurboCapacity"
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
//  INVALIDPARAMETERVALUE_TAGKEYFILTERLIMITEXCEEDED = "InvalidParameterValue.TagKeyFilterLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUEFILTERLIMITEXCEEDED = "InvalidParameterValue.TagValueFilterLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUELIMITEXCEEDED = "InvalidParameterValue.TagValueLimitExceeded"
//  INVALIDPARAMETERVALUE_UNAVAILABLEREGION = "InvalidParameterValue.UnavailableRegion"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDPARAMETERVALUE_ZONEIDREGIONNOTMATCH = "InvalidParameterValue.ZoneIdRegionNotMatch"
//  RESOURCEINSUFFICIENT_FILESYSTEMLIMITEXCEEDED = "ResourceInsufficient.FileSystemLimitExceeded"
//  RESOURCEINSUFFICIENT_REGIONSOLDOUT = "ResourceInsufficient.RegionSoldOut"
//  RESOURCEINSUFFICIENT_SUBNETIPALLOCCUPIED = "ResourceInsufficient.SubnetIpAllOccupied"
//  RESOURCEINSUFFICIENT_TAGLIMITEXCEEDED = "ResourceInsufficient.TagLimitExceeded"
//  RESOURCEINSUFFICIENT_TAGQUOTASEXCEEDED = "ResourceInsufficient.TagQuotasExceeded"
//  RESOURCEINSUFFICIENT_TURBOSPECIALCAPACITYFILESYSTEMCOUNTLIMIT = "ResourceInsufficient.TurboSpecialCapacityFileSystemCountLimit"
//  RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BASICNETINTERFACENOTSUPPORTED = "UnsupportedOperation.BasicNetInterfaceNotSupported"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsFileSystemWithContext(ctx context.Context, request *CreateCfsFileSystemRequest) (response *CreateCfsFileSystemResponse, err error) {
    if request == nil {
        request = NewCreateCfsFileSystemRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "CreateCfsFileSystem")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "CreateCfsPGroup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "CreateCfsRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCfsRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCfsRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCfsSnapshotRequest() (request *CreateCfsSnapshotRequest) {
    request = &CreateCfsSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "CreateCfsSnapshot")
    
    
    return
}

func NewCreateCfsSnapshotResponse() (response *CreateCfsSnapshotResponse) {
    response = &CreateCfsSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCfsSnapshot
// This API is used to create a file system snapshot.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSNAPSHOTNAME = "InvalidParameter.InvalidSnapshotName"
//  INVALIDPARAMETER_SNAPSHOTNAMELIMITEXCEEDED = "InvalidParameter.SnapshotNameLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FSSIZELIMITEXCEEDED = "InvalidParameterValue.FsSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_INVALIDRESOURCETAGS = "InvalidParameterValue.InvalidResourceTags"
//  INVALIDPARAMETERVALUE_INVALIDTAGKEY = "InvalidParameterValue.InvalidTagKey"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUE = "InvalidParameterValue.InvalidTagValue"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  INVALIDPARAMETERVALUE_TAGKEYFILTERLIMITEXCEEDED = "InvalidParameterValue.TagKeyFilterLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUEFILTERLIMITEXCEEDED = "InvalidParameterValue.TagValueFilterLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUELIMITEXCEEDED = "InvalidParameterValue.TagValueLimitExceeded"
//  RESOURCEINSUFFICIENT_SNAPSHOTSIZELIMITEXCEEDED = "ResourceInsufficient.SnapshotSizeLimitExceeded"
//  RESOURCEINSUFFICIENT_TAGLIMITEXCEEDED = "ResourceInsufficient.TagLimitExceeded"
//  RESOURCEINSUFFICIENT_TAGQUOTASEXCEEDED = "ResourceInsufficient.TagQuotasExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsSnapshot(request *CreateCfsSnapshotRequest) (response *CreateCfsSnapshotResponse, err error) {
    return c.CreateCfsSnapshotWithContext(context.Background(), request)
}

// CreateCfsSnapshot
// This API is used to create a file system snapshot.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSNAPSHOTNAME = "InvalidParameter.InvalidSnapshotName"
//  INVALIDPARAMETER_SNAPSHOTNAMELIMITEXCEEDED = "InvalidParameter.SnapshotNameLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_FSSIZELIMITEXCEEDED = "InvalidParameterValue.FsSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_INVALIDRESOURCETAGS = "InvalidParameterValue.InvalidResourceTags"
//  INVALIDPARAMETERVALUE_INVALIDTAGKEY = "InvalidParameterValue.InvalidTagKey"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUE = "InvalidParameterValue.InvalidTagValue"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  INVALIDPARAMETERVALUE_TAGKEYFILTERLIMITEXCEEDED = "InvalidParameterValue.TagKeyFilterLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUEFILTERLIMITEXCEEDED = "InvalidParameterValue.TagValueFilterLimitExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUELIMITEXCEEDED = "InvalidParameterValue.TagValueLimitExceeded"
//  RESOURCEINSUFFICIENT_SNAPSHOTSIZELIMITEXCEEDED = "ResourceInsufficient.SnapshotSizeLimitExceeded"
//  RESOURCEINSUFFICIENT_TAGLIMITEXCEEDED = "ResourceInsufficient.TagLimitExceeded"
//  RESOURCEINSUFFICIENT_TAGQUOTASEXCEEDED = "ResourceInsufficient.TagQuotasExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) CreateCfsSnapshotWithContext(ctx context.Context, request *CreateCfsSnapshotRequest) (response *CreateCfsSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateCfsSnapshotRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "CreateCfsSnapshot")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCfsSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCfsSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataFlowRequest() (request *CreateDataFlowRequest) {
    request = &CreateDataFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "CreateDataFlow")
    
    
    return
}

func NewCreateDataFlowResponse() (response *CreateDataFlowResponse) {
    response = &CreateDataFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDataFlow
// This API is used to create a data flow.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWNAME = "InvalidParameterValue.InvalidDataFlowName"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWSOURCEINFO = "InvalidParameterValue.InvalidDataFlowSourceInfo"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWSOURCESTORAGETYPE = "InvalidParameterValue.InvalidDataFlowSourceStorageType"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWTARGETPATH = "InvalidParameterValue.InvalidDataFlowTargetPath"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLEDATATASKTYPE = "InvalidParameterValue.InvalidLifecycleDataTaskType"
//  RESOURCEINSUFFICIENT_DATAFLOWLIMITEXCEEDED = "ResourceInsufficient.DataFlowLimitExceeded"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
func (c *Client) CreateDataFlow(request *CreateDataFlowRequest) (response *CreateDataFlowResponse, err error) {
    return c.CreateDataFlowWithContext(context.Background(), request)
}

// CreateDataFlow
// This API is used to create a data flow.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWNAME = "InvalidParameterValue.InvalidDataFlowName"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWSOURCEINFO = "InvalidParameterValue.InvalidDataFlowSourceInfo"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWSOURCESTORAGETYPE = "InvalidParameterValue.InvalidDataFlowSourceStorageType"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWTARGETPATH = "InvalidParameterValue.InvalidDataFlowTargetPath"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLEDATATASKTYPE = "InvalidParameterValue.InvalidLifecycleDataTaskType"
//  RESOURCEINSUFFICIENT_DATAFLOWLIMITEXCEEDED = "ResourceInsufficient.DataFlowLimitExceeded"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
func (c *Client) CreateDataFlowWithContext(ctx context.Context, request *CreateDataFlowRequest) (response *CreateDataFlowResponse, err error) {
    if request == nil {
        request = NewCreateDataFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "CreateDataFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLifecycleDataTaskRequest() (request *CreateLifecycleDataTaskRequest) {
    request = &CreateLifecycleDataTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "CreateLifecycleDataTask")
    
    
    return
}

func NewCreateLifecycleDataTaskResponse() (response *CreateLifecycleDataTaskResponse) {
    response = &CreateLifecycleDataTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLifecycleDataTask
// Support proactive settlement/preheat API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWID = "InvalidParameterValue.InvalidDataFlowId"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWSOURCESTORAGETYPE = "InvalidParameterValue.InvalidDataFlowSourceStorageType"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWTARGETPATH = "InvalidParameterValue.InvalidDataFlowTargetPath"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLEDATATASKTYPE = "InvalidParameterValue.InvalidLifecycleDataTaskType"
//  INVALIDPARAMETERVALUE_WAITINGTASKLIMITEXCEEDED = "InvalidParameterValue.WaitingTaskLimitExceeded"
func (c *Client) CreateLifecycleDataTask(request *CreateLifecycleDataTaskRequest) (response *CreateLifecycleDataTaskResponse, err error) {
    return c.CreateLifecycleDataTaskWithContext(context.Background(), request)
}

// CreateLifecycleDataTask
// Support proactive settlement/preheat API.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWID = "InvalidParameterValue.InvalidDataFlowId"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWSOURCESTORAGETYPE = "InvalidParameterValue.InvalidDataFlowSourceStorageType"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWTARGETPATH = "InvalidParameterValue.InvalidDataFlowTargetPath"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLEDATATASKTYPE = "InvalidParameterValue.InvalidLifecycleDataTaskType"
//  INVALIDPARAMETERVALUE_WAITINGTASKLIMITEXCEEDED = "InvalidParameterValue.WaitingTaskLimitExceeded"
func (c *Client) CreateLifecycleDataTaskWithContext(ctx context.Context, request *CreateLifecycleDataTaskRequest) (response *CreateLifecycleDataTaskResponse, err error) {
    if request == nil {
        request = NewCreateLifecycleDataTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "CreateLifecycleDataTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLifecycleDataTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLifecycleDataTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLifecyclePolicyRequest() (request *CreateLifecyclePolicyRequest) {
    request = &CreateLifecyclePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "CreateLifecyclePolicy")
    
    
    return
}

func NewCreateLifecyclePolicyResponse() (response *CreateLifecyclePolicyResponse) {
    response = &CreateLifecyclePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLifecyclePolicy
// This API is used to create a file storage lifecycle policy.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINSUFFICIENT_POLICYLIMITEXCEEDED = "ResourceInsufficient.PolicyLimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateLifecyclePolicy(request *CreateLifecyclePolicyRequest) (response *CreateLifecyclePolicyResponse, err error) {
    return c.CreateLifecyclePolicyWithContext(context.Background(), request)
}

// CreateLifecyclePolicy
// This API is used to create a file storage lifecycle policy.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINSUFFICIENT_POLICYLIMITEXCEEDED = "ResourceInsufficient.PolicyLimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateLifecyclePolicyWithContext(ctx context.Context, request *CreateLifecyclePolicyRequest) (response *CreateLifecyclePolicyResponse, err error) {
    if request == nil {
        request = NewCreateLifecyclePolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "CreateLifecyclePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLifecyclePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLifecyclePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLifecyclePolicyDownloadTaskRequest() (request *CreateLifecyclePolicyDownloadTaskRequest) {
    request = &CreateLifecyclePolicyDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "CreateLifecyclePolicyDownloadTask")
    
    
    return
}

func NewCreateLifecyclePolicyDownloadTaskResponse() (response *CreateLifecyclePolicyDownloadTaskResponse) {
    response = &CreateLifecyclePolicyDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLifecyclePolicyDownloadTask
// Download the file list in a lifecycle task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINSUFFICIENT_POLICYLIMITEXCEEDED = "ResourceInsufficient.PolicyLimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateLifecyclePolicyDownloadTask(request *CreateLifecyclePolicyDownloadTaskRequest) (response *CreateLifecyclePolicyDownloadTaskResponse, err error) {
    return c.CreateLifecyclePolicyDownloadTaskWithContext(context.Background(), request)
}

// CreateLifecyclePolicyDownloadTask
// Download the file list in a lifecycle task.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINSUFFICIENT_POLICYLIMITEXCEEDED = "ResourceInsufficient.PolicyLimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateLifecyclePolicyDownloadTaskWithContext(ctx context.Context, request *CreateLifecyclePolicyDownloadTaskRequest) (response *CreateLifecyclePolicyDownloadTaskResponse, err error) {
    if request == nil {
        request = NewCreateLifecyclePolicyDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "CreateLifecyclePolicyDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLifecyclePolicyDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLifecyclePolicyDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMigrationTaskRequest() (request *CreateMigrationTaskRequest) {
    request = &CreateMigrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "CreateMigrationTask")
    
    
    return
}

func NewCreateMigrationTaskResponse() (response *CreateMigrationTaskResponse) {
    response = &CreateMigrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMigrationTask
// This API is used to create a migration task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateMigrationTask(request *CreateMigrationTaskRequest) (response *CreateMigrationTaskResponse, err error) {
    return c.CreateMigrationTaskWithContext(context.Background(), request)
}

// CreateMigrationTask
// This API is used to create a migration task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) CreateMigrationTaskWithContext(ctx context.Context, request *CreateMigrationTaskRequest) (response *CreateMigrationTaskResponse, err error) {
    if request == nil {
        request = NewCreateMigrationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "CreateMigrationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMigrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMigrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAutoSnapshotPolicyRequest() (request *DeleteAutoSnapshotPolicyRequest) {
    request = &DeleteAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteAutoSnapshotPolicy")
    
    
    return
}

func NewDeleteAutoSnapshotPolicyResponse() (response *DeleteAutoSnapshotPolicyResponse) {
    response = &DeleteAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAutoSnapshotPolicy
// This API is used to delete a scheduled snapshot policy.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUTOPOLICYNOTFOUND = "InvalidParameterValue.AutoPolicyNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteAutoSnapshotPolicy(request *DeleteAutoSnapshotPolicyRequest) (response *DeleteAutoSnapshotPolicyResponse, err error) {
    return c.DeleteAutoSnapshotPolicyWithContext(context.Background(), request)
}

// DeleteAutoSnapshotPolicy
// This API is used to delete a scheduled snapshot policy.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUTOPOLICYNOTFOUND = "InvalidParameterValue.AutoPolicyNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteAutoSnapshotPolicyWithContext(ctx context.Context, request *DeleteAutoSnapshotPolicyRequest) (response *DeleteAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteAutoSnapshotPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DeleteAutoSnapshotPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAutoSnapshotPolicyResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DeleteCfsFileSystem")
    
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
// This API is used to delete a permission group. Only permission groups not bound to a file system can be deleted by this API.
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
// This API is used to delete a permission group. Only permission groups not bound to a file system can be deleted by this API.
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DeleteCfsPGroup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DeleteCfsRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCfsRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCfsRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCfsSnapshotRequest() (request *DeleteCfsSnapshotRequest) {
    request = &DeleteCfsSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteCfsSnapshot")
    
    
    return
}

func NewDeleteCfsSnapshotResponse() (response *DeleteCfsSnapshotResponse) {
    response = &DeleteCfsSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCfsSnapshot
// This API is used to delete a file system snapshot.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsSnapshot(request *DeleteCfsSnapshotRequest) (response *DeleteCfsSnapshotResponse, err error) {
    return c.DeleteCfsSnapshotWithContext(context.Background(), request)
}

// DeleteCfsSnapshot
// This API is used to delete a file system snapshot.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DeleteCfsSnapshotWithContext(ctx context.Context, request *DeleteCfsSnapshotRequest) (response *DeleteCfsSnapshotResponse, err error) {
    if request == nil {
        request = NewDeleteCfsSnapshotRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DeleteCfsSnapshot")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCfsSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCfsSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDataFlowRequest() (request *DeleteDataFlowRequest) {
    request = &DeleteDataFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteDataFlow")
    
    
    return
}

func NewDeleteDataFlowResponse() (response *DeleteDataFlowResponse) {
    response = &DeleteDataFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDataFlow
// Delete a data flow.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWID = "InvalidParameterValue.InvalidDataFlowId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_INVALIDLIFECYCLEDATATASKSTATUS = "UnsupportedOperation.InvalidLifecycleDataTaskStatus"
func (c *Client) DeleteDataFlow(request *DeleteDataFlowRequest) (response *DeleteDataFlowResponse, err error) {
    return c.DeleteDataFlowWithContext(context.Background(), request)
}

// DeleteDataFlow
// Delete a data flow.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWID = "InvalidParameterValue.InvalidDataFlowId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_INVALIDLIFECYCLEDATATASKSTATUS = "UnsupportedOperation.InvalidLifecycleDataTaskStatus"
func (c *Client) DeleteDataFlowWithContext(ctx context.Context, request *DeleteDataFlowRequest) (response *DeleteDataFlowResponse, err error) {
    if request == nil {
        request = NewDeleteDataFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DeleteDataFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDataFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLifecyclePolicyRequest() (request *DeleteLifecyclePolicyRequest) {
    request = &DeleteLifecyclePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteLifecyclePolicy")
    
    
    return
}

func NewDeleteLifecyclePolicyResponse() (response *DeleteLifecyclePolicyResponse) {
    response = &DeleteLifecyclePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLifecyclePolicy
// Delete a lifecycle management policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWID = "InvalidParameterValue.InvalidDataFlowId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_INVALIDLIFECYCLEDATATASKSTATUS = "UnsupportedOperation.InvalidLifecycleDataTaskStatus"
func (c *Client) DeleteLifecyclePolicy(request *DeleteLifecyclePolicyRequest) (response *DeleteLifecyclePolicyResponse, err error) {
    return c.DeleteLifecyclePolicyWithContext(context.Background(), request)
}

// DeleteLifecyclePolicy
// Delete a lifecycle management policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWID = "InvalidParameterValue.InvalidDataFlowId"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_INVALIDLIFECYCLEDATATASKSTATUS = "UnsupportedOperation.InvalidLifecycleDataTaskStatus"
func (c *Client) DeleteLifecyclePolicyWithContext(ctx context.Context, request *DeleteLifecyclePolicyRequest) (response *DeleteLifecyclePolicyResponse, err error) {
    if request == nil {
        request = NewDeleteLifecyclePolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DeleteLifecyclePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLifecyclePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLifecyclePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMigrationTaskRequest() (request *DeleteMigrationTaskRequest) {
    request = &DeleteMigrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DeleteMigrationTask")
    
    
    return
}

func NewDeleteMigrationTaskResponse() (response *DeleteMigrationTaskResponse) {
    response = &DeleteMigrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMigrationTask
// This API is used to delete migration tasks. Deletion is not supported for tasks in the status of Waiting, creating, running, canceling, or terminating.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMigrationTask(request *DeleteMigrationTaskRequest) (response *DeleteMigrationTaskResponse, err error) {
    return c.DeleteMigrationTaskWithContext(context.Background(), request)
}

// DeleteMigrationTask
// This API is used to delete migration tasks. Deletion is not supported for tasks in the status of Waiting, creating, running, canceling, or terminating.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMigrationTaskWithContext(ctx context.Context, request *DeleteMigrationTaskRequest) (response *DeleteMigrationTaskResponse, err error) {
    if request == nil {
        request = NewDeleteMigrationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DeleteMigrationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMigrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMigrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoSnapshotPoliciesRequest() (request *DescribeAutoSnapshotPoliciesRequest) {
    request = &DescribeAutoSnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeAutoSnapshotPolicies")
    
    
    return
}

func NewDescribeAutoSnapshotPoliciesResponse() (response *DescribeAutoSnapshotPoliciesResponse) {
    response = &DescribeAutoSnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoSnapshotPolicies
// This API is used to query the list of scheduled snapshot policies of a file system.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeAutoSnapshotPolicies(request *DescribeAutoSnapshotPoliciesRequest) (response *DescribeAutoSnapshotPoliciesResponse, err error) {
    return c.DescribeAutoSnapshotPoliciesWithContext(context.Background(), request)
}

// DescribeAutoSnapshotPolicies
// This API is used to query the list of scheduled snapshot policies of a file system.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeAutoSnapshotPoliciesWithContext(ctx context.Context, request *DescribeAutoSnapshotPoliciesRequest) (response *DescribeAutoSnapshotPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoSnapshotPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeAutoSnapshotPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoSnapshotPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoSnapshotPoliciesResponse()
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
//  AUTHFAILURE_GETROLEFAILED = "AuthFailure.GetRoleFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeAvailableZoneInfo(request *DescribeAvailableZoneInfoRequest) (response *DescribeAvailableZoneInfoResponse, err error) {
    return c.DescribeAvailableZoneInfoWithContext(context.Background(), request)
}

// DescribeAvailableZoneInfo
// This API is used to query the availability of a region.
//
// error code that may be returned:
//  AUTHFAILURE_GETROLEFAILED = "AuthFailure.GetRoleFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeAvailableZoneInfoWithContext(ctx context.Context, request *DescribeAvailableZoneInfoRequest) (response *DescribeAvailableZoneInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableZoneInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeAvailableZoneInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableZoneInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableZoneInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBucketListRequest() (request *DescribeBucketListRequest) {
    request = &DescribeBucketListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeBucketList")
    
    
    return
}

func NewDescribeBucketListResponse() (response *DescribeBucketListResponse) {
    response = &DescribeBucketListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBucketList
// To obtain the list of data source buckets.
//
// error code that may be returned:
//  AUTHFAILURE_GETROLEFAILED = "AuthFailure.GetRoleFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeBucketList(request *DescribeBucketListRequest) (response *DescribeBucketListResponse, err error) {
    return c.DescribeBucketListWithContext(context.Background(), request)
}

// DescribeBucketList
// To obtain the list of data source buckets.
//
// error code that may be returned:
//  AUTHFAILURE_GETROLEFAILED = "AuthFailure.GetRoleFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeBucketListWithContext(ctx context.Context, request *DescribeBucketListRequest) (response *DescribeBucketListResponse, err error) {
    if request == nil {
        request = NewDescribeBucketListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeBucketList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBucketList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBucketListResponse()
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
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
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
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
func (c *Client) DescribeCfsFileSystemClientsWithContext(ctx context.Context, request *DescribeCfsFileSystemClientsRequest) (response *DescribeCfsFileSystemClientsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsFileSystemClientsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeCfsFileSystemClients")
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsFileSystemsWithContext(ctx context.Context, request *DescribeCfsFileSystemsRequest) (response *DescribeCfsFileSystemsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsFileSystemsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeCfsFileSystems")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeCfsPGroups")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeCfsRules")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeCfsServiceStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsServiceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsServiceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsSnapshotOverviewRequest() (request *DescribeCfsSnapshotOverviewRequest) {
    request = &DescribeCfsSnapshotOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsSnapshotOverview")
    
    
    return
}

func NewDescribeCfsSnapshotOverviewResponse() (response *DescribeCfsSnapshotOverviewResponse) {
    response = &DescribeCfsSnapshotOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCfsSnapshotOverview
// This API is used to get the snapshot overview of a file system.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsSnapshotOverview(request *DescribeCfsSnapshotOverviewRequest) (response *DescribeCfsSnapshotOverviewResponse, err error) {
    return c.DescribeCfsSnapshotOverviewWithContext(context.Background(), request)
}

// DescribeCfsSnapshotOverview
// This API is used to get the snapshot overview of a file system.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsSnapshotOverviewWithContext(ctx context.Context, request *DescribeCfsSnapshotOverviewRequest) (response *DescribeCfsSnapshotOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeCfsSnapshotOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeCfsSnapshotOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsSnapshotOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsSnapshotOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCfsSnapshotsRequest() (request *DescribeCfsSnapshotsRequest) {
    request = &DescribeCfsSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeCfsSnapshots")
    
    
    return
}

func NewDescribeCfsSnapshotsResponse() (response *DescribeCfsSnapshotsResponse) {
    response = &DescribeCfsSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCfsSnapshots
// This API is used to query the list of snapshots of a file system.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsSnapshots(request *DescribeCfsSnapshotsRequest) (response *DescribeCfsSnapshotsResponse, err error) {
    return c.DescribeCfsSnapshotsWithContext(context.Background(), request)
}

// DescribeCfsSnapshots
// This API is used to query the list of snapshots of a file system.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeCfsSnapshotsWithContext(ctx context.Context, request *DescribeCfsSnapshotsRequest) (response *DescribeCfsSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeCfsSnapshotsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeCfsSnapshots")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCfsSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCfsSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataFlowRequest() (request *DescribeDataFlowRequest) {
    request = &DescribeDataFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeDataFlow")
    
    
    return
}

func NewDescribeDataFlowResponse() (response *DescribeDataFlowResponse) {
    response = &DescribeDataFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataFlow
// This API is used to query data flow information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWID = "InvalidParameterValue.InvalidDataFlowId"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWNAME = "InvalidParameterValue.InvalidDataFlowName"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
func (c *Client) DescribeDataFlow(request *DescribeDataFlowRequest) (response *DescribeDataFlowResponse, err error) {
    return c.DescribeDataFlowWithContext(context.Background(), request)
}

// DescribeDataFlow
// This API is used to query data flow information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWID = "InvalidParameterValue.InvalidDataFlowId"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWNAME = "InvalidParameterValue.InvalidDataFlowName"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
func (c *Client) DescribeDataFlowWithContext(ctx context.Context, request *DescribeDataFlowRequest) (response *DescribeDataFlowResponse, err error) {
    if request == nil {
        request = NewDescribeDataFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeDataFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLifecycleDataTaskRequest() (request *DescribeLifecycleDataTaskRequest) {
    request = &DescribeLifecycleDataTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeLifecycleDataTask")
    
    
    return
}

func NewDescribeLifecycleDataTaskResponse() (response *DescribeLifecycleDataTaskResponse) {
    response = &DescribeLifecycleDataTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLifecycleDataTask
// This API is used to query the lifecycle task. It only supports querying task data within 3 months.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MISSINGTASKIDORREGION = "InvalidParameterValue.MissingTaskIdOrRegion"
func (c *Client) DescribeLifecycleDataTask(request *DescribeLifecycleDataTaskRequest) (response *DescribeLifecycleDataTaskResponse, err error) {
    return c.DescribeLifecycleDataTaskWithContext(context.Background(), request)
}

// DescribeLifecycleDataTask
// This API is used to query the lifecycle task. It only supports querying task data within 3 months.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MISSINGTASKIDORREGION = "InvalidParameterValue.MissingTaskIdOrRegion"
func (c *Client) DescribeLifecycleDataTaskWithContext(ctx context.Context, request *DescribeLifecycleDataTaskRequest) (response *DescribeLifecycleDataTaskResponse, err error) {
    if request == nil {
        request = NewDescribeLifecycleDataTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeLifecycleDataTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLifecycleDataTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLifecycleDataTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLifecyclePoliciesRequest() (request *DescribeLifecyclePoliciesRequest) {
    request = &DescribeLifecyclePoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeLifecyclePolicies")
    
    
    return
}

func NewDescribeLifecyclePoliciesResponse() (response *DescribeLifecyclePoliciesResponse) {
    response = &DescribeLifecyclePoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLifecyclePolicies
// Query lifecycle management policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MISSINGTASKIDORREGION = "InvalidParameterValue.MissingTaskIdOrRegion"
func (c *Client) DescribeLifecyclePolicies(request *DescribeLifecyclePoliciesRequest) (response *DescribeLifecyclePoliciesResponse, err error) {
    return c.DescribeLifecyclePoliciesWithContext(context.Background(), request)
}

// DescribeLifecyclePolicies
// Query lifecycle management policies.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MISSINGTASKIDORREGION = "InvalidParameterValue.MissingTaskIdOrRegion"
func (c *Client) DescribeLifecyclePoliciesWithContext(ctx context.Context, request *DescribeLifecyclePoliciesRequest) (response *DescribeLifecyclePoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeLifecyclePoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeLifecyclePolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLifecyclePolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLifecyclePoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationTasksRequest() (request *DescribeMigrationTasksRequest) {
    request = &DescribeMigrationTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeMigrationTasks")
    
    
    return
}

func NewDescribeMigrationTasksResponse() (response *DescribeMigrationTasksResponse) {
    response = &DescribeMigrationTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigrationTasks
// This API is used to get the list of migration tasks.
//
// To use this API, submit a ticket for us to add you to the allowlist.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MISSINGTASKIDORREGION = "InvalidParameterValue.MissingTaskIdOrRegion"
func (c *Client) DescribeMigrationTasks(request *DescribeMigrationTasksRequest) (response *DescribeMigrationTasksResponse, err error) {
    return c.DescribeMigrationTasksWithContext(context.Background(), request)
}

// DescribeMigrationTasks
// This API is used to get the list of migration tasks.
//
// To use this API, submit a ticket for us to add you to the allowlist.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_MISSINGTASKIDORREGION = "InvalidParameterValue.MissingTaskIdOrRegion"
func (c *Client) DescribeMigrationTasksWithContext(ctx context.Context, request *DescribeMigrationTasksRequest) (response *DescribeMigrationTasksResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeMigrationTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrationTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrationTasksResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeMountTargets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMountTargets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMountTargetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotOperationLogsRequest() (request *DescribeSnapshotOperationLogsRequest) {
    request = &DescribeSnapshotOperationLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DescribeSnapshotOperationLogs")
    
    
    return
}

func NewDescribeSnapshotOperationLogsResponse() (response *DescribeSnapshotOperationLogsResponse) {
    response = &DescribeSnapshotOperationLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSnapshotOperationLogs
// This API is used to query the operation logs of a snapshot.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeSnapshotOperationLogs(request *DescribeSnapshotOperationLogsRequest) (response *DescribeSnapshotOperationLogsResponse, err error) {
    return c.DescribeSnapshotOperationLogsWithContext(context.Background(), request)
}

// DescribeSnapshotOperationLogs
// This API is used to query the operation logs of a snapshot.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DescribeSnapshotOperationLogsWithContext(ctx context.Context, request *DescribeSnapshotOperationLogsRequest) (response *DescribeSnapshotOperationLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotOperationLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DescribeSnapshotOperationLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotOperationLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotOperationLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDoDirectoryOperationRequest() (request *DoDirectoryOperationRequest) {
    request = &DoDirectoryOperationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "DoDirectoryOperation")
    
    
    return
}

func NewDoDirectoryOperationResponse() (response *DoDirectoryOperationResponse) {
    response = &DoDirectoryOperationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DoDirectoryOperation
// File system directory operation API. Currently, only the Turbo series file system supports calling this API to perform directory operations. The Universal Series file system (including the enhanced version) does not support calling.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DoDirectoryOperation(request *DoDirectoryOperationRequest) (response *DoDirectoryOperationResponse, err error) {
    return c.DoDirectoryOperationWithContext(context.Background(), request)
}

// DoDirectoryOperation
// File system directory operation API. Currently, only the Turbo series file system supports calling this API to perform directory operations. The Universal Series file system (including the enhanced version) does not support calling.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) DoDirectoryOperationWithContext(ctx context.Context, request *DoDirectoryOperationRequest) (response *DoDirectoryOperationResponse, err error) {
    if request == nil {
        request = NewDoDirectoryOperationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "DoDirectoryOperation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DoDirectoryOperation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDoDirectoryOperationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDataFlowRequest() (request *ModifyDataFlowRequest) {
    request = &ModifyDataFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "ModifyDataFlow")
    
    
    return
}

func NewModifyDataFlowResponse() (response *ModifyDataFlowResponse) {
    response = &ModifyDataFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDataFlow
// Modify data flow parameters.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWID = "InvalidParameterValue.InvalidDataFlowId"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWNAME = "InvalidParameterValue.InvalidDataFlowName"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWSOURCEINFO = "InvalidParameterValue.InvalidDataFlowSourceInfo"
func (c *Client) ModifyDataFlow(request *ModifyDataFlowRequest) (response *ModifyDataFlowResponse, err error) {
    return c.ModifyDataFlowWithContext(context.Background(), request)
}

// ModifyDataFlow
// Modify data flow parameters.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWID = "InvalidParameterValue.InvalidDataFlowId"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWNAME = "InvalidParameterValue.InvalidDataFlowName"
//  INVALIDPARAMETERVALUE_INVALIDDATAFLOWSOURCEINFO = "InvalidParameterValue.InvalidDataFlowSourceInfo"
func (c *Client) ModifyDataFlowWithContext(ctx context.Context, request *ModifyDataFlowRequest) (response *ModifyDataFlowResponse, err error) {
    if request == nil {
        request = NewModifyDataFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "ModifyDataFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDataFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDataFlowResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFileSystemAutoScaleUpRuleRequest() (request *ModifyFileSystemAutoScaleUpRuleRequest) {
    request = &ModifyFileSystemAutoScaleUpRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "ModifyFileSystemAutoScaleUpRule")
    
    
    return
}

func NewModifyFileSystemAutoScaleUpRuleResponse() (response *ModifyFileSystemAutoScaleUpRuleResponse) {
    response = &ModifyFileSystemAutoScaleUpRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFileSystemAutoScaleUpRule
// This API is used to set the file system scaling policy. It supports only the turbo file system.
//
// error code that may be returned:
//  AUTHFAILURE_GETROLEFAILED = "AuthFailure.GetRoleFailed"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALEUPPARAMS = "InvalidParameterValue.InvalidAutoScaleUpParams"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  RESOURCENOTFOUND_FSNOTEXIST = "ResourceNotFound.FsNotExist"
//  RESOURCENOTFOUND_RESOURCEPACKAGENOTFOUND = "ResourceNotFound.ResourcePackageNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyFileSystemAutoScaleUpRule(request *ModifyFileSystemAutoScaleUpRuleRequest) (response *ModifyFileSystemAutoScaleUpRuleResponse, err error) {
    return c.ModifyFileSystemAutoScaleUpRuleWithContext(context.Background(), request)
}

// ModifyFileSystemAutoScaleUpRule
// This API is used to set the file system scaling policy. It supports only the turbo file system.
//
// error code that may be returned:
//  AUTHFAILURE_GETROLEFAILED = "AuthFailure.GetRoleFailed"
//  INVALIDPARAMETERVALUE_INVALIDAUTOSCALEUPPARAMS = "InvalidParameterValue.InvalidAutoScaleUpParams"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  RESOURCENOTFOUND_FSNOTEXIST = "ResourceNotFound.FsNotExist"
//  RESOURCENOTFOUND_RESOURCEPACKAGENOTFOUND = "ResourceNotFound.ResourcePackageNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyFileSystemAutoScaleUpRuleWithContext(ctx context.Context, request *ModifyFileSystemAutoScaleUpRuleRequest) (response *ModifyFileSystemAutoScaleUpRuleResponse, err error) {
    if request == nil {
        request = NewModifyFileSystemAutoScaleUpRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "ModifyFileSystemAutoScaleUpRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFileSystemAutoScaleUpRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFileSystemAutoScaleUpRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLifecyclePolicyRequest() (request *ModifyLifecyclePolicyRequest) {
    request = &ModifyLifecyclePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "ModifyLifecyclePolicy")
    
    
    return
}

func NewModifyLifecyclePolicyResponse() (response *ModifyLifecyclePolicyResponse) {
    response = &ModifyLifecyclePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyLifecyclePolicy
// Update a file storage lifecycle policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CFSPATH = "InvalidParameterValue.CfsPath"
//  INVALIDPARAMETERVALUE_POLICYRULESTORAGETYPEINVALID = "InvalidParameterValue.PolicyRuleStorageTypeInvalid"
//  INVALIDPARAMETERVALUE_STORAGETYPEINVALID = "InvalidParameterValue.StorageTypeInvalid"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyLifecyclePolicy(request *ModifyLifecyclePolicyRequest) (response *ModifyLifecyclePolicyResponse, err error) {
    return c.ModifyLifecyclePolicyWithContext(context.Background(), request)
}

// ModifyLifecyclePolicy
// Update a file storage lifecycle policy.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CFSPATH = "InvalidParameterValue.CfsPath"
//  INVALIDPARAMETERVALUE_POLICYRULESTORAGETYPEINVALID = "InvalidParameterValue.PolicyRuleStorageTypeInvalid"
//  INVALIDPARAMETERVALUE_STORAGETYPEINVALID = "InvalidParameterValue.StorageTypeInvalid"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyLifecyclePolicyWithContext(ctx context.Context, request *ModifyLifecyclePolicyRequest) (response *ModifyLifecyclePolicyResponse, err error) {
    if request == nil {
        request = NewModifyLifecyclePolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "ModifyLifecyclePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLifecyclePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLifecyclePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewScaleUpFileSystemRequest() (request *ScaleUpFileSystemRequest) {
    request = &ScaleUpFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "ScaleUpFileSystem")
    
    
    return
}

func NewScaleUpFileSystemResponse() (response *ScaleUpFileSystemResponse) {
    response = &ScaleUpFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleUpFileSystem
// This API is used to expand the turbo file system. It only supports expansion and does not support scale-down. The scaling increment for the turbo standard type is 10240 GIB, and for the turbo performance type, it is 5120 GIB.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_INVALIDSCALEUPTARGETCAPACITY = "InvalidParameterValue.InvalidScaleupTargetCapacity"
//  INVALIDPARAMETERVALUE_INVALIDTURBOCAPACITY = "InvalidParameterValue.InvalidTurboCapacity"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_FSNOTEXIST = "ResourceNotFound.FsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MISSINGKMSACCESSPERMISSION = "UnsupportedOperation.MissingKmsAccessPermission"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) ScaleUpFileSystem(request *ScaleUpFileSystemRequest) (response *ScaleUpFileSystemResponse, err error) {
    return c.ScaleUpFileSystemWithContext(context.Background(), request)
}

// ScaleUpFileSystem
// This API is used to expand the turbo file system. It only supports expansion and does not support scale-down. The scaling increment for the turbo standard type is 10240 GIB, and for the turbo performance type, it is 5120 GIB.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"
//  INVALIDPARAMETERVALUE_INVALIDSCALEUPTARGETCAPACITY = "InvalidParameterValue.InvalidScaleupTargetCapacity"
//  INVALIDPARAMETERVALUE_INVALIDTURBOCAPACITY = "InvalidParameterValue.InvalidTurboCapacity"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  RESOURCENOTFOUND_FSNOTEXIST = "ResourceNotFound.FsNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MISSINGKMSACCESSPERMISSION = "UnsupportedOperation.MissingKmsAccessPermission"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) ScaleUpFileSystemWithContext(ctx context.Context, request *ScaleUpFileSystemRequest) (response *ScaleUpFileSystemResponse, err error) {
    if request == nil {
        request = NewScaleUpFileSystemRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "ScaleUpFileSystem")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleUpFileSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleUpFileSystemResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "SignUpCfsService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SignUpCfsService require credential")
    }

    request.SetContext(ctx)
    
    response = NewSignUpCfsServiceResponse()
    err = c.Send(request, response)
    return
}

func NewStopLifecycleDataTaskRequest() (request *StopLifecycleDataTaskRequest) {
    request = &StopLifecycleDataTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "StopLifecycleDataTask")
    
    
    return
}

func NewStopLifecycleDataTaskResponse() (response *StopLifecycleDataTaskResponse) {
    response = &StopLifecycleDataTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopLifecycleDataTask
// This API is used to terminate a lifecycle task.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLETASKID = "InvalidParameterValue.InvalidLifecycleTaskId"
func (c *Client) StopLifecycleDataTask(request *StopLifecycleDataTaskRequest) (response *StopLifecycleDataTaskResponse, err error) {
    return c.StopLifecycleDataTaskWithContext(context.Background(), request)
}

// StopLifecycleDataTask
// This API is used to terminate a lifecycle task.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDLIFECYCLETASKID = "InvalidParameterValue.InvalidLifecycleTaskId"
func (c *Client) StopLifecycleDataTaskWithContext(ctx context.Context, request *StopLifecycleDataTaskRequest) (response *StopLifecycleDataTaskResponse, err error) {
    if request == nil {
        request = NewStopLifecycleDataTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "StopLifecycleDataTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopLifecycleDataTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopLifecycleDataTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStopMigrationTaskRequest() (request *StopMigrationTaskRequest) {
    request = &StopMigrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "StopMigrationTask")
    
    
    return
}

func NewStopMigrationTaskResponse() (response *StopMigrationTaskResponse) {
    response = &StopMigrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopMigrationTask
// This API is used to terminate a migration task, including tasks in the Waiting or Running status.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) StopMigrationTask(request *StopMigrationTaskRequest) (response *StopMigrationTaskResponse, err error) {
    return c.StopMigrationTaskWithContext(context.Background(), request)
}

// StopMigrationTask
// This API is used to terminate a migration task, including tasks in the Waiting or Running status.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) StopMigrationTaskWithContext(ctx context.Context, request *StopMigrationTaskRequest) (response *StopMigrationTaskResponse, err error) {
    if request == nil {
        request = NewStopMigrationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "StopMigrationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopMigrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopMigrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindAutoSnapshotPolicyRequest() (request *UnbindAutoSnapshotPolicyRequest) {
    request = &UnbindAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "UnbindAutoSnapshotPolicy")
    
    
    return
}

func NewUnbindAutoSnapshotPolicyResponse() (response *UnbindAutoSnapshotPolicyResponse) {
    response = &UnbindAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindAutoSnapshotPolicy
// This API is used to unbind a snapshot policy from a file system.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UnbindAutoSnapshotPolicy(request *UnbindAutoSnapshotPolicyRequest) (response *UnbindAutoSnapshotPolicyResponse, err error) {
    return c.UnbindAutoSnapshotPolicyWithContext(context.Background(), request)
}

// UnbindAutoSnapshotPolicy
// This API is used to unbind a snapshot policy from a file system.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UnbindAutoSnapshotPolicyWithContext(ctx context.Context, request *UnbindAutoSnapshotPolicyRequest) (response *UnbindAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewUnbindAutoSnapshotPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "UnbindAutoSnapshotPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAutoSnapshotPolicyRequest() (request *UpdateAutoSnapshotPolicyRequest) {
    request = &UpdateAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateAutoSnapshotPolicy")
    
    
    return
}

func NewUpdateAutoSnapshotPolicyResponse() (response *UpdateAutoSnapshotPolicyResponse) {
    response = &UpdateAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAutoSnapshotPolicy
// This API is used to update a scheduled snapshot policy.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDPARAMDAYOFWEEK = "InvalidParameter.InvalidParamDayofWeek"
//  INVALIDPARAMETER_INVALIDPARAMHOUR = "InvalidParameter.InvalidParamHour"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETER_INVALIDSNAPSHOTPOLICYNAME = "InvalidParameter.InvalidSnapshotPolicyName"
//  INVALIDPARAMETER_MISSINGPOLICYPARAM = "InvalidParameter.MissingPolicyParam"
//  INVALIDPARAMETER_SNAPSHOTPOLICYNAMELIMITEXCEEDED = "InvalidParameter.SnapshotPolicyNameLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUTOPOLICYNOTFOUND = "InvalidParameterValue.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE_INVALIDDESTINATIONREGIONS = "InvalidParameterValue.InvalidDestinationRegions"
//  INVALIDPARAMETERVALUE_MISSINGPOLICYPARAM = "InvalidParameterValue.MissingPolicyParam"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateAutoSnapshotPolicy(request *UpdateAutoSnapshotPolicyRequest) (response *UpdateAutoSnapshotPolicyResponse, err error) {
    return c.UpdateAutoSnapshotPolicyWithContext(context.Background(), request)
}

// UpdateAutoSnapshotPolicy
// This API is used to update a scheduled snapshot policy.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"
//  INVALIDPARAMETER_INVALIDPARAMDAYOFWEEK = "InvalidParameter.InvalidParamDayofWeek"
//  INVALIDPARAMETER_INVALIDPARAMHOUR = "InvalidParameter.InvalidParamHour"
//  INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"
//  INVALIDPARAMETER_INVALIDSNAPSHOTPOLICYNAME = "InvalidParameter.InvalidSnapshotPolicyName"
//  INVALIDPARAMETER_MISSINGPOLICYPARAM = "InvalidParameter.MissingPolicyParam"
//  INVALIDPARAMETER_SNAPSHOTPOLICYNAMELIMITEXCEEDED = "InvalidParameter.SnapshotPolicyNameLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUTOPOLICYNOTFOUND = "InvalidParameterValue.AutoPolicyNotFound"
//  INVALIDPARAMETERVALUE_INVALIDDESTINATIONREGIONS = "InvalidParameterValue.InvalidDestinationRegions"
//  INVALIDPARAMETERVALUE_MISSINGPOLICYPARAM = "InvalidParameterValue.MissingPolicyParam"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
//  UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"
//  UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
func (c *Client) UpdateAutoSnapshotPolicyWithContext(ctx context.Context, request *UpdateAutoSnapshotPolicyRequest) (response *UpdateAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewUpdateAutoSnapshotPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "UpdateAutoSnapshotPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAutoSnapshotPolicyResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "UpdateCfsFileSystemName")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "UpdateCfsFileSystemPGroup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "UpdateCfsFileSystemSizeLimit")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "UpdateCfsPGroup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "UpdateCfsRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCfsSnapshotAttributeRequest() (request *UpdateCfsSnapshotAttributeRequest) {
    request = &UpdateCfsSnapshotAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfs", APIVersion, "UpdateCfsSnapshotAttribute")
    
    
    return
}

func NewUpdateCfsSnapshotAttributeResponse() (response *UpdateCfsSnapshotAttributeResponse) {
    response = &UpdateCfsSnapshotAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCfsSnapshotAttribute
// This API is used to update the name and retention period of a file system snapshot.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSNAPSHOTNAME = "InvalidParameter.InvalidSnapshotName"
//  INVALIDPARAMETER_SNAPSHOTNAMELIMITEXCEEDED = "InvalidParameter.SnapshotNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTNAME = "InvalidParameterValue.InvalidSnapshotName"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
func (c *Client) UpdateCfsSnapshotAttribute(request *UpdateCfsSnapshotAttributeRequest) (response *UpdateCfsSnapshotAttributeResponse, err error) {
    return c.UpdateCfsSnapshotAttributeWithContext(context.Background(), request)
}

// UpdateCfsSnapshotAttribute
// This API is used to update the name and retention period of a file system snapshot.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDSNAPSHOTNAME = "InvalidParameter.InvalidSnapshotName"
//  INVALIDPARAMETER_SNAPSHOTNAMELIMITEXCEEDED = "InvalidParameter.SnapshotNameLimitExceeded"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTNAME = "InvalidParameterValue.InvalidSnapshotName"
//  INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"
func (c *Client) UpdateCfsSnapshotAttributeWithContext(ctx context.Context, request *UpdateCfsSnapshotAttributeRequest) (response *UpdateCfsSnapshotAttributeResponse, err error) {
    if request == nil {
        request = NewUpdateCfsSnapshotAttributeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cfs", APIVersion, "UpdateCfsSnapshotAttribute")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCfsSnapshotAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCfsSnapshotAttributeResponse()
    err = c.Send(request, response)
    return
}
