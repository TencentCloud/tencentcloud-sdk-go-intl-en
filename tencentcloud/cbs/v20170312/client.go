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

package v20170312

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2017-03-12"

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


func NewApplySnapshotRequest() (request *ApplySnapshotRequest) {
    request = &ApplySnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "ApplySnapshot")
    
    
    return
}

func NewApplySnapshotResponse() (response *ApplySnapshotResponse) {
    response = &ApplySnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplySnapshot
// This API (ApplySnapshot) is used to roll back a snapshot to the original cloud disk.
//
// 
//
// * The snapshot can only be rolled back to the original cloud disk. For data disk snapshots, if you need to copy the snapshot data to other cloud disks, use the API [CreateDisks](https://intl.cloud.tencent.com/document/product/362/16312?from_cn_redirect=1) to create an elastic cloud disk and then copy the snapshot data to the created cloud disk. 
//
// * The snapshot for rollback must be in NORMAL status. The snapshot status can be queried in the SnapshotState field in the output parameters through the API [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
//
// * For elastic cloud disks, the cloud disk must be in UNMOUNTED status. The cloud disk status can be queried in the Attached field returned by the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1). For non-elastic cloud disks purchased together with instances, the instance must be in SHUTDOWN status. The instance status can be queried through the API [DescribeInstancesStatus](https://intl.cloud.tencent.com/document/product/213/15738?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_SNAPSHOTCREATING = "InvalidDisk.SnapshotCreating"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER_DISKSIZENOTMATCH = "InvalidParameter.DiskSizeNotMatch"
//  INVALIDPARAMETER_SHOULDCONVERTSNAPSHOTTOIMAGE = "InvalidParameter.ShouldConvertSnapshotToImage"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
func (c *Client) ApplySnapshot(request *ApplySnapshotRequest) (response *ApplySnapshotResponse, err error) {
    return c.ApplySnapshotWithContext(context.Background(), request)
}

// ApplySnapshot
// This API (ApplySnapshot) is used to roll back a snapshot to the original cloud disk.
//
// 
//
// * The snapshot can only be rolled back to the original cloud disk. For data disk snapshots, if you need to copy the snapshot data to other cloud disks, use the API [CreateDisks](https://intl.cloud.tencent.com/document/product/362/16312?from_cn_redirect=1) to create an elastic cloud disk and then copy the snapshot data to the created cloud disk. 
//
// * The snapshot for rollback must be in NORMAL status. The snapshot status can be queried in the SnapshotState field in the output parameters through the API [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1).
//
// * For elastic cloud disks, the cloud disk must be in UNMOUNTED status. The cloud disk status can be queried in the Attached field returned by the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1). For non-elastic cloud disks purchased together with instances, the instance must be in SHUTDOWN status. The instance status can be queried through the API [DescribeInstancesStatus](https://intl.cloud.tencent.com/document/product/213/15738?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_SNAPSHOTCREATING = "InvalidDisk.SnapshotCreating"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER_DISKSIZENOTMATCH = "InvalidParameter.DiskSizeNotMatch"
//  INVALIDPARAMETER_SHOULDCONVERTSNAPSHOTTOIMAGE = "InvalidParameter.ShouldConvertSnapshotToImage"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
func (c *Client) ApplySnapshotWithContext(ctx context.Context, request *ApplySnapshotRequest) (response *ApplySnapshotResponse, err error) {
    if request == nil {
        request = NewApplySnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplySnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplySnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewAttachDisksRequest() (request *AttachDisksRequest) {
    request = &AttachDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "AttachDisks")
    
    
    return
}

func NewAttachDisksResponse() (response *AttachDisksResponse) {
    response = &AttachDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachDisks
// This API is used to mount one or more cloud disks.
//
//  
//
// * Batch operation is supported. You can mount multiple cloud disks to one CVM in a single request. If any of these cloud disks cannot be mounted, the operation fails and a specific error code returns.
//
// * This is an async API. A successful request indicates that the mounting is initiated. You can call the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API to query the status of cloud disks. If the status changes from `ATTACHING` to `ATTACHED`, the mounting is successful.
//
// error code that may be returned:
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_INSTANCEATTACHEDDISK = "LimitExceeded.InstanceAttachedDisk"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_ATTACHED = "ResourceUnavailable.Attached"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"
//  RESOURCEUNAVAILABLE_ZONENOTMATCH = "ResourceUnavailable.ZoneNotMatch"
func (c *Client) AttachDisks(request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
    return c.AttachDisksWithContext(context.Background(), request)
}

// AttachDisks
// This API is used to mount one or more cloud disks.
//
//  
//
// * Batch operation is supported. You can mount multiple cloud disks to one CVM in a single request. If any of these cloud disks cannot be mounted, the operation fails and a specific error code returns.
//
// * This is an async API. A successful request indicates that the mounting is initiated. You can call the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API to query the status of cloud disks. If the status changes from `ATTACHING` to `ATTACHED`, the mounting is successful.
//
// error code that may be returned:
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  LIMITEXCEEDED_INSTANCEATTACHEDDISK = "LimitExceeded.InstanceAttachedDisk"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_ATTACHED = "ResourceUnavailable.Attached"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"
//  RESOURCEUNAVAILABLE_ZONENOTMATCH = "ResourceUnavailable.ZoneNotMatch"
func (c *Client) AttachDisksWithContext(ctx context.Context, request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
    if request == nil {
        request = NewAttachDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachDisksResponse()
    err = c.Send(request, response)
    return
}

func NewBindAutoSnapshotPolicyRequest() (request *BindAutoSnapshotPolicyRequest) {
    request = &BindAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "BindAutoSnapshotPolicy")
    
    
    return
}

func NewBindAutoSnapshotPolicyResponse() (response *BindAutoSnapshotPolicyResponse) {
    response = &BindAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindAutoSnapshotPolicy
// This API (BindAutoSnapshotPolicy) is used to bind the cloud disk to the specified scheduled snapshot policy.
//
// 
//
// * For the scheduled snapshot policy limit of each region, see [Scheduled Snapshots](https://intl.cloud.tencent.com/document/product/362/8191?from_cn_redirect=1).
//
// * When a cloud disk that is bound to a scheduled snapshot policy is in the unused state (that is, an elastic cloud disk has not been mounted or the server of an inelastic disk is powered off) scheduled snapshots are not created.
//
// error code that may be returned:
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDDISK_ALREADYBOUND = "InvalidDisk.AlreadyBound"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BINDDISKLIMITEXCEEDED = "InvalidParameterValue.BindDiskLimitExceeded"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) BindAutoSnapshotPolicy(request *BindAutoSnapshotPolicyRequest) (response *BindAutoSnapshotPolicyResponse, err error) {
    return c.BindAutoSnapshotPolicyWithContext(context.Background(), request)
}

// BindAutoSnapshotPolicy
// This API (BindAutoSnapshotPolicy) is used to bind the cloud disk to the specified scheduled snapshot policy.
//
// 
//
// * For the scheduled snapshot policy limit of each region, see [Scheduled Snapshots](https://intl.cloud.tencent.com/document/product/362/8191?from_cn_redirect=1).
//
// * When a cloud disk that is bound to a scheduled snapshot policy is in the unused state (that is, an elastic cloud disk has not been mounted or the server of an inelastic disk is powered off) scheduled snapshots are not created.
//
// error code that may be returned:
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDDISK_ALREADYBOUND = "InvalidDisk.AlreadyBound"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BINDDISKLIMITEXCEEDED = "InvalidParameterValue.BindDiskLimitExceeded"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) BindAutoSnapshotPolicyWithContext(ctx context.Context, request *BindAutoSnapshotPolicyRequest) (response *BindAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewBindAutoSnapshotPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCopySnapshotCrossRegionsRequest() (request *CopySnapshotCrossRegionsRequest) {
    request = &CopySnapshotCrossRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "CopySnapshotCrossRegions")
    
    
    return
}

func NewCopySnapshotCrossRegionsResponse() (response *CopySnapshotCrossRegionsResponse) {
    response = &CopySnapshotCrossRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CopySnapshotCrossRegions
// This API is used to replicate a snapshot to another region.
//
// 
//
// * This is an async API. A new snapshot ID is issued when the cross-region replication task is generated. It does not mean that the snapshot has been replicated successfully. You can all the [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1) API in the destination region to check for this snapshot. If the snapshot status is `NORMAL`, the snapshot is replicated successfully.
//
// * The snapshot cross-region replication service will be commercialized in the Q3 of 2022. We will notify users about the commercialization in advance. Please check your messages in the Message Center.
//
// error code that may be returned:
//  INSUFFICIENTSNAPSHOTQUOTA = "InsufficientSnapshotQuota"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_COPYSNAPSHOTCONFLICT = "ResourceInUse.CopySnapshotConflict"
//  UNSUPPORTEDOPERATION_SNAPSHOTNOTSUPPORTCOPY = "UnsupportedOperation.SnapshotNotSupportCopy"
func (c *Client) CopySnapshotCrossRegions(request *CopySnapshotCrossRegionsRequest) (response *CopySnapshotCrossRegionsResponse, err error) {
    return c.CopySnapshotCrossRegionsWithContext(context.Background(), request)
}

// CopySnapshotCrossRegions
// This API is used to replicate a snapshot to another region.
//
// 
//
// * This is an async API. A new snapshot ID is issued when the cross-region replication task is generated. It does not mean that the snapshot has been replicated successfully. You can all the [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1) API in the destination region to check for this snapshot. If the snapshot status is `NORMAL`, the snapshot is replicated successfully.
//
// * The snapshot cross-region replication service will be commercialized in the Q3 of 2022. We will notify users about the commercialization in advance. Please check your messages in the Message Center.
//
// error code that may be returned:
//  INSUFFICIENTSNAPSHOTQUOTA = "InsufficientSnapshotQuota"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_COPYSNAPSHOTCONFLICT = "ResourceInUse.CopySnapshotConflict"
//  UNSUPPORTEDOPERATION_SNAPSHOTNOTSUPPORTCOPY = "UnsupportedOperation.SnapshotNotSupportCopy"
func (c *Client) CopySnapshotCrossRegionsWithContext(ctx context.Context, request *CopySnapshotCrossRegionsRequest) (response *CopySnapshotCrossRegionsResponse, err error) {
    if request == nil {
        request = NewCopySnapshotCrossRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopySnapshotCrossRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopySnapshotCrossRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoSnapshotPolicyRequest() (request *CreateAutoSnapshotPolicyRequest) {
    request = &CreateAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "CreateAutoSnapshotPolicy")
    
    
    return
}

func NewCreateAutoSnapshotPolicyResponse() (response *CreateAutoSnapshotPolicyResponse) {
    response = &CreateAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAutoSnapshotPolicy
// This API (CreateAutoSnapshotPolicy) is used to create a scheduled snapshot policy.
//
// 
//
// * For the limits on the number of scheduled snapshot policies that can be created in each region, see [Scheduled Snapshots](https://intl.cloud.tencent.com/document/product/362/8191?from_cn_redirect=1).
//
// * The quantity and capacity of the snapshots that can be created in each region are limited. For more information, see the **Snapshots** page on the Tencent Cloud Console. If the number of snapshots exceeds the quota, the creation of the scheduled snapshots will fail.
//
// error code that may be returned:
//  AUTOSNAPSHOTPOLICYOUTOFQUOTA = "AutoSnapshotPolicyOutOfQuota"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (response *CreateAutoSnapshotPolicyResponse, err error) {
    return c.CreateAutoSnapshotPolicyWithContext(context.Background(), request)
}

// CreateAutoSnapshotPolicy
// This API (CreateAutoSnapshotPolicy) is used to create a scheduled snapshot policy.
//
// 
//
// * For the limits on the number of scheduled snapshot policies that can be created in each region, see [Scheduled Snapshots](https://intl.cloud.tencent.com/document/product/362/8191?from_cn_redirect=1).
//
// * The quantity and capacity of the snapshots that can be created in each region are limited. For more information, see the **Snapshots** page on the Tencent Cloud Console. If the number of snapshots exceeds the quota, the creation of the scheduled snapshots will fail.
//
// error code that may be returned:
//  AUTOSNAPSHOTPOLICYOUTOFQUOTA = "AutoSnapshotPolicyOutOfQuota"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateAutoSnapshotPolicyWithContext(ctx context.Context, request *CreateAutoSnapshotPolicyRequest) (response *CreateAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAutoSnapshotPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDisksRequest() (request *CreateDisksRequest) {
    request = &CreateDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "CreateDisks")
    
    
    return
}

func NewCreateDisksResponse() (response *CreateDisksResponse) {
    response = &CreateDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDisks
// This API is used to create one or more cloud disks.
//
// 
//
// * This API supports creating a cloud disk with a data disk snapshot so that the snapshot data can be copied to the purchased cloud disk.
//
// * This API is an async API. A cloud disk ID list will be returned when a request is made successfully, but it does not mean that the creation has been completed. You can call the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API to query cloud disks by `DiskId`. If a new cloud disk can be found and its state is 'UNATTACHED' or 'ATTACHED', it means that the cloud disk has been created successfully.
//
// error code that may be returned:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETER_PROJECTIDNOTEXIST = "InvalidParameter.ProjectIdNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
//  UNAUTHORIZEDOPERATION_ROLENOTEXISTS = "UnauthorizedOperation.RoleNotExists"
func (c *Client) CreateDisks(request *CreateDisksRequest) (response *CreateDisksResponse, err error) {
    return c.CreateDisksWithContext(context.Background(), request)
}

// CreateDisks
// This API is used to create one or more cloud disks.
//
// 
//
// * This API supports creating a cloud disk with a data disk snapshot so that the snapshot data can be copied to the purchased cloud disk.
//
// * This API is an async API. A cloud disk ID list will be returned when a request is made successfully, but it does not mean that the creation has been completed. You can call the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API to query cloud disks by `DiskId`. If a new cloud disk can be found and its state is 'UNATTACHED' or 'ATTACHED', it means that the cloud disk has been created successfully.
//
// error code that may be returned:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETER_PROJECTIDNOTEXIST = "InvalidParameter.ProjectIdNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
//  UNAUTHORIZEDOPERATION_ROLENOTEXISTS = "UnauthorizedOperation.RoleNotExists"
func (c *Client) CreateDisksWithContext(ctx context.Context, request *CreateDisksRequest) (response *CreateDisksResponse, err error) {
    if request == nil {
        request = NewCreateDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDisksResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSnapshotRequest() (request *CreateSnapshotRequest) {
    request = &CreateSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "CreateSnapshot")
    
    
    return
}

func NewCreateSnapshotResponse() (response *CreateSnapshotResponse) {
    response = &CreateSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSnapshot
// This API (CreateSnapshot) is used to create a snapshot of a specified cloud disk.
//
// 
//
// * Snapshots can only be created for cloud disks with the snapshot capability. To check whether a cloud disk has the snapshot capability, see the SnapshotAbility field returned by the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1).
//
// * For the number of snapshots that can be created, please see [Product Usage Restriction](https://intl.cloud.tencent.com/doc/product/362/5145?from_cn_redirect=1).
//
// error code that may be returned:
//  INSUFFICIENTSNAPSHOTQUOTA = "InsufficientSnapshotQuota"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTSUPPORTSNAPSHOT = "InvalidDisk.NotSupportSnapshot"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_SNAPSHOTCREATING = "InvalidDisk.SnapshotCreating"
//  INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  LIMITEXCEEDED_INSTANCEATTACHEDDISK = "LimitExceeded.InstanceAttachedDisk"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"
//  RESOURCEINSUFFICIENT_OVERQUOTA = "ResourceInsufficient.OverQuota"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  RESOURCEUNAVAILABLE_TOOMANYCREATINGSNAPSHOT = "ResourceUnavailable.TooManyCreatingSnapshot"
//  UNSUPPORTEDOPERATION_DISKENCRYPT = "UnsupportedOperation.DiskEncrypt"
func (c *Client) CreateSnapshot(request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
    return c.CreateSnapshotWithContext(context.Background(), request)
}

// CreateSnapshot
// This API (CreateSnapshot) is used to create a snapshot of a specified cloud disk.
//
// 
//
// * Snapshots can only be created for cloud disks with the snapshot capability. To check whether a cloud disk has the snapshot capability, see the SnapshotAbility field returned by the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1).
//
// * For the number of snapshots that can be created, please see [Product Usage Restriction](https://intl.cloud.tencent.com/doc/product/362/5145?from_cn_redirect=1).
//
// error code that may be returned:
//  INSUFFICIENTSNAPSHOTQUOTA = "InsufficientSnapshotQuota"
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_NOTSUPPORTSNAPSHOT = "InvalidDisk.NotSupportSnapshot"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_SNAPSHOTCREATING = "InvalidDisk.SnapshotCreating"
//  INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  LIMITEXCEEDED_INSTANCEATTACHEDDISK = "LimitExceeded.InstanceAttachedDisk"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"
//  RESOURCEINSUFFICIENT_OVERQUOTA = "ResourceInsufficient.OverQuota"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  RESOURCEUNAVAILABLE_TOOMANYCREATINGSNAPSHOT = "ResourceUnavailable.TooManyCreatingSnapshot"
//  UNSUPPORTEDOPERATION_DISKENCRYPT = "UnsupportedOperation.DiskEncrypt"
func (c *Client) CreateSnapshotWithContext(ctx context.Context, request *CreateSnapshotRequest) (response *CreateSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAutoSnapshotPoliciesRequest() (request *DeleteAutoSnapshotPoliciesRequest) {
    request = &DeleteAutoSnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "DeleteAutoSnapshotPolicies")
    
    
    return
}

func NewDeleteAutoSnapshotPoliciesResponse() (response *DeleteAutoSnapshotPoliciesResponse) {
    response = &DeleteAutoSnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAutoSnapshotPolicies
// This API (DeleteAutoSnapshotPolicies) is used to delete scheduled snapshot policies.
//
// 
//
// * Batch operations are supported. If one of the scheduled snapshot policies in a batch cannot be deleted, the operation is not performed and a specific error code is returned.
//
// error code that may be returned:
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAutoSnapshotPolicies(request *DeleteAutoSnapshotPoliciesRequest) (response *DeleteAutoSnapshotPoliciesResponse, err error) {
    return c.DeleteAutoSnapshotPoliciesWithContext(context.Background(), request)
}

// DeleteAutoSnapshotPolicies
// This API (DeleteAutoSnapshotPolicies) is used to delete scheduled snapshot policies.
//
// 
//
// * Batch operations are supported. If one of the scheduled snapshot policies in a batch cannot be deleted, the operation is not performed and a specific error code is returned.
//
// error code that may be returned:
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteAutoSnapshotPoliciesWithContext(ctx context.Context, request *DeleteAutoSnapshotPoliciesRequest) (response *DeleteAutoSnapshotPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteAutoSnapshotPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAutoSnapshotPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAutoSnapshotPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
    request = &DeleteSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "DeleteSnapshots")
    
    
    return
}

func NewDeleteSnapshotsResponse() (response *DeleteSnapshotsResponse) {
    response = &DeleteSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSnapshots
// This API is used to delete snapshots.
//
// 
//
// * Only snapshots in the `NORMAL` state can be deleted. To query the state of a snapshot, you can call the [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1) API and check the `SnapshotState` field in the response.
//
// * Batch operations are supported. If there is any snapshot that cannot be deleted, the operation will not be performed and a specific error code will be returned.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_SNAPHASSHARED = "UnsupportedOperation.SnapHasShared"
//  UNSUPPORTEDOPERATION_SNAPSHOTHASBINDEDIMAGE = "UnsupportedOperation.SnapshotHasBindedImage"
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    return c.DeleteSnapshotsWithContext(context.Background(), request)
}

// DeleteSnapshots
// This API is used to delete snapshots.
//
// 
//
// * Only snapshots in the `NORMAL` state can be deleted. To query the state of a snapshot, you can call the [DescribeSnapshots](https://intl.cloud.tencent.com/document/product/362/15647?from_cn_redirect=1) API and check the `SnapshotState` field in the response.
//
// * Batch operations are supported. If there is any snapshot that cannot be deleted, the operation will not be performed and a specific error code will be returned.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_SNAPHASSHARED = "UnsupportedOperation.SnapHasShared"
//  UNSUPPORTEDOPERATION_SNAPSHOTHASBINDEDIMAGE = "UnsupportedOperation.SnapshotHasBindedImage"
func (c *Client) DeleteSnapshotsWithContext(ctx context.Context, request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoSnapshotPoliciesRequest() (request *DescribeAutoSnapshotPoliciesRequest) {
    request = &DescribeAutoSnapshotPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeAutoSnapshotPolicies")
    
    
    return
}

func NewDescribeAutoSnapshotPoliciesResponse() (response *DescribeAutoSnapshotPoliciesResponse) {
    response = &DescribeAutoSnapshotPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoSnapshotPolicies
// This API (DescribeAutoSnapshotPolicies) is used to query scheduled snapshot policies.
//
// 
//
// * You can query the detailed information of scheduled snapshot policies by ID, name, or status. Insert `AND` between different values. For details on filtering information, see `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit`; the default is 20) of the scheduled snapshot policy lists are returned to the current user.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAutoSnapshotPolicies(request *DescribeAutoSnapshotPoliciesRequest) (response *DescribeAutoSnapshotPoliciesResponse, err error) {
    return c.DescribeAutoSnapshotPoliciesWithContext(context.Background(), request)
}

// DescribeAutoSnapshotPolicies
// This API (DescribeAutoSnapshotPolicies) is used to query scheduled snapshot policies.
//
// 
//
// * You can query the detailed information of scheduled snapshot policies by ID, name, or status. Insert `AND` between different values. For details on filtering information, see `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit`; the default is 20) of the scheduled snapshot policy lists are returned to the current user.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeAutoSnapshotPoliciesWithContext(ctx context.Context, request *DescribeAutoSnapshotPoliciesRequest) (response *DescribeAutoSnapshotPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoSnapshotPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoSnapshotPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoSnapshotPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiskAssociatedAutoSnapshotPolicyRequest() (request *DescribeDiskAssociatedAutoSnapshotPolicyRequest) {
    request = &DescribeDiskAssociatedAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskAssociatedAutoSnapshotPolicy")
    
    
    return
}

func NewDescribeDiskAssociatedAutoSnapshotPolicyResponse() (response *DescribeDiskAssociatedAutoSnapshotPolicyResponse) {
    response = &DescribeDiskAssociatedAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskAssociatedAutoSnapshotPolicy
// This API (DescribeDiskAssociatedAutoSnapshotPolicy) is used to query the scheduled snapshot policy bound to a cloud disk.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDiskAssociatedAutoSnapshotPolicy(request *DescribeDiskAssociatedAutoSnapshotPolicyRequest) (response *DescribeDiskAssociatedAutoSnapshotPolicyResponse, err error) {
    return c.DescribeDiskAssociatedAutoSnapshotPolicyWithContext(context.Background(), request)
}

// DescribeDiskAssociatedAutoSnapshotPolicy
// This API (DescribeDiskAssociatedAutoSnapshotPolicy) is used to query the scheduled snapshot policy bound to a cloud disk.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDiskAssociatedAutoSnapshotPolicyWithContext(ctx context.Context, request *DescribeDiskAssociatedAutoSnapshotPolicyRequest) (response *DescribeDiskAssociatedAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeDiskAssociatedAutoSnapshotPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskAssociatedAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskAssociatedAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiskConfigQuotaRequest() (request *DescribeDiskConfigQuotaRequest) {
    request = &DescribeDiskConfigQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskConfigQuota")
    
    
    return
}

func NewDescribeDiskConfigQuotaResponse() (response *DescribeDiskConfigQuotaResponse) {
    response = &DescribeDiskConfigQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskConfigQuota
// This API (DescribeDiskConfigQuota) is used to query the cloud disk quota.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDiskConfigQuota(request *DescribeDiskConfigQuotaRequest) (response *DescribeDiskConfigQuotaResponse, err error) {
    return c.DescribeDiskConfigQuotaWithContext(context.Background(), request)
}

// DescribeDiskConfigQuota
// This API (DescribeDiskConfigQuota) is used to query the cloud disk quota.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDiskConfigQuotaWithContext(ctx context.Context, request *DescribeDiskConfigQuotaRequest) (response *DescribeDiskConfigQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeDiskConfigQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskConfigQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskConfigQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiskOperationLogsRequest() (request *DescribeDiskOperationLogsRequest) {
    request = &DescribeDiskOperationLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskOperationLogs")
    
    
    return
}

func NewDescribeDiskOperationLogsResponse() (response *DescribeDiskOperationLogsResponse) {
    response = &DescribeDiskOperationLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskOperationLogs
// This API (DescribeDiskOperationLogs) is used to query a list of cloud disk operation logs.
//
// 
//
// This can be filtered according to the cloud disk ID. The format of cloud disk IDs is as follows: disk-a1kmcp13.
//
// error code that may be returned:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDiskOperationLogs(request *DescribeDiskOperationLogsRequest) (response *DescribeDiskOperationLogsResponse, err error) {
    return c.DescribeDiskOperationLogsWithContext(context.Background(), request)
}

// DescribeDiskOperationLogs
// This API (DescribeDiskOperationLogs) is used to query a list of cloud disk operation logs.
//
// 
//
// This can be filtered according to the cloud disk ID. The format of cloud disk IDs is as follows: disk-a1kmcp13.
//
// error code that may be returned:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDiskOperationLogsWithContext(ctx context.Context, request *DescribeDiskOperationLogsRequest) (response *DescribeDiskOperationLogsResponse, err error) {
    if request == nil {
        request = NewDescribeDiskOperationLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskOperationLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskOperationLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
    request = &DescribeDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeDisks")
    
    
    return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
    response = &DescribeDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDisks
// This API (DescribeDisks) is used to query the list of cloud disks.
//
// 
//
// * The details of the cloud disk can be queried based on the ID, type or status of the cloud disk. The relationship between different conditions is AND. For more information about filtering, please see the `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit`; the default is 20) of cloud disk lists are returned to the current user.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
    return c.DescribeDisksWithContext(context.Background(), request)
}

// DescribeDisks
// This API (DescribeDisks) is used to query the list of cloud disks.
//
// 
//
// * The details of the cloud disk can be queried based on the ID, type or status of the cloud disk. The relationship between different conditions is AND. For more information about filtering, please see the `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit`; the default is 20) of cloud disk lists are returned to the current user.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDisksWithContext(ctx context.Context, request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
    if request == nil {
        request = NewDescribeDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesDiskNumRequest() (request *DescribeInstancesDiskNumRequest) {
    request = &DescribeInstancesDiskNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeInstancesDiskNum")
    
    
    return
}

func NewDescribeInstancesDiskNumResponse() (response *DescribeInstancesDiskNumResponse) {
    response = &DescribeInstancesDiskNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesDiskNum
// This API (DescribeInstancesDiskNum) is used to query the number of cloud disks mounted in the instance.
//
// 
//
// * Batch operations are supported. If multiple CVM instance IDs are specified, the returned results will list the number of cloud disks mounted on each CVM.
//
// error code that may be returned:
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeInstancesDiskNum(request *DescribeInstancesDiskNumRequest) (response *DescribeInstancesDiskNumResponse, err error) {
    return c.DescribeInstancesDiskNumWithContext(context.Background(), request)
}

// DescribeInstancesDiskNum
// This API (DescribeInstancesDiskNum) is used to query the number of cloud disks mounted in the instance.
//
// 
//
// * Batch operations are supported. If multiple CVM instance IDs are specified, the returned results will list the number of cloud disks mounted on each CVM.
//
// error code that may be returned:
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeInstancesDiskNumWithContext(ctx context.Context, request *DescribeInstancesDiskNumRequest) (response *DescribeInstancesDiskNumResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesDiskNumRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesDiskNum require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesDiskNumResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotOperationLogsRequest() (request *DescribeSnapshotOperationLogsRequest) {
    request = &DescribeSnapshotOperationLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotOperationLogs")
    
    
    return
}

func NewDescribeSnapshotOperationLogsResponse() (response *DescribeSnapshotOperationLogsResponse) {
    response = &DescribeSnapshotOperationLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshotOperationLogs
// This API (DescribeSnapshotOperationLogs) is used to query a list of snapshot operation logs.
//
// 
//
// You can filter according to the snapshot ID. The snapshot ID format is as follows: snap-a1kmcp13.
//
// error code that may be returned:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshotOperationLogs(request *DescribeSnapshotOperationLogsRequest) (response *DescribeSnapshotOperationLogsResponse, err error) {
    return c.DescribeSnapshotOperationLogsWithContext(context.Background(), request)
}

// DescribeSnapshotOperationLogs
// This API (DescribeSnapshotOperationLogs) is used to query a list of snapshot operation logs.
//
// 
//
// You can filter according to the snapshot ID. The snapshot ID format is as follows: snap-a1kmcp13.
//
// error code that may be returned:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshotOperationLogsWithContext(ctx context.Context, request *DescribeSnapshotOperationLogsRequest) (response *DescribeSnapshotOperationLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotOperationLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotOperationLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotOperationLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotSharePermissionRequest() (request *DescribeSnapshotSharePermissionRequest) {
    request = &DescribeSnapshotSharePermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshotSharePermission")
    
    
    return
}

func NewDescribeSnapshotSharePermissionResponse() (response *DescribeSnapshotSharePermissionResponse) {
    response = &DescribeSnapshotSharePermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshotSharePermission
// This API is used to query the sharing information of snapshots.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshotSharePermission(request *DescribeSnapshotSharePermissionRequest) (response *DescribeSnapshotSharePermissionResponse, err error) {
    return c.DescribeSnapshotSharePermissionWithContext(context.Background(), request)
}

// DescribeSnapshotSharePermission
// This API is used to query the sharing information of snapshots.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshotSharePermissionWithContext(ctx context.Context, request *DescribeSnapshotSharePermissionRequest) (response *DescribeSnapshotSharePermissionResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotSharePermissionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotSharePermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotSharePermissionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshots")
    
    
    return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
    response = &DescribeSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshots
// This API (DescribeSnapshots) is used to query the details of snapshots.
//
// 
//
// * Filter the results by the snapshot ID, the ID of cloud disk, for which the snapshot is created, and the type of cloud disk, for which the snapshot is created. The relationship between different conditions is AND. For more information about filtering, please see `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit`; the default is 20) of snapshot lists are returned to the current user.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    return c.DescribeSnapshotsWithContext(context.Background(), request)
}

// DescribeSnapshots
// This API (DescribeSnapshots) is used to query the details of snapshots.
//
// 
//
// * Filter the results by the snapshot ID, the ID of cloud disk, for which the snapshot is created, and the type of cloud disk, for which the snapshot is created. The relationship between different conditions is AND. For more information about filtering, please see `Filter`.
//
// * If the parameter is empty, a certain number (specified by `Limit`; the default is 20) of snapshot lists are returned to the current user.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDetachDisksRequest() (request *DetachDisksRequest) {
    request = &DetachDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "DetachDisks")
    
    
    return
}

func NewDetachDisksResponse() (response *DetachDisksResponse) {
    response = &DetachDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachDisks
// This API is used to unmount one or more cloud disks.
//
// 
//
// * Batch operation is supported. You can unmount multiple cloud disks from the same CVM in a single request. If any of these cloud disks cannot be unmounted, the operation fails and a specific error code returns.
//
// * This is an async API. A successful request does not mean that the cloud disks have been unmounted successfully. You can call the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API to query the status of cloud disks. When the status changes from `ATTACHED` to `UNATTACHED`, the unmounting is successful.
//
// error code that may be returned:
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_DISKMIGRATING = "ResourceInUse.DiskMigrating"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_DETACHPOD = "UnsupportedOperation.DetachPod"
func (c *Client) DetachDisks(request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
    return c.DetachDisksWithContext(context.Background(), request)
}

// DetachDisks
// This API is used to unmount one or more cloud disks.
//
// 
//
// * Batch operation is supported. You can unmount multiple cloud disks from the same CVM in a single request. If any of these cloud disks cannot be unmounted, the operation fails and a specific error code returns.
//
// * This is an async API. A successful request does not mean that the cloud disks have been unmounted successfully. You can call the [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1) API to query the status of cloud disks. When the status changes from `ATTACHED` to `UNATTACHED`, the unmounting is successful.
//
// error code that may be returned:
//  INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINUSE_DISKMIGRATING = "ResourceInUse.DiskMigrating"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNSUPPORTEDOPERATION_DETACHPOD = "UnsupportedOperation.DetachPod"
func (c *Client) DetachDisksWithContext(ctx context.Context, request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
    if request == nil {
        request = NewDetachDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachDisksResponse()
    err = c.Send(request, response)
    return
}

func NewGetSnapOverviewRequest() (request *GetSnapOverviewRequest) {
    request = &GetSnapOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "GetSnapOverview")
    
    
    return
}

func NewGetSnapOverviewResponse() (response *GetSnapOverviewResponse) {
    response = &GetSnapOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetSnapOverview
// This API is used to get snapshot overview information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetSnapOverview(request *GetSnapOverviewRequest) (response *GetSnapOverviewResponse, err error) {
    return c.GetSnapOverviewWithContext(context.Background(), request)
}

// GetSnapOverview
// This API is used to get snapshot overview information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetSnapOverviewWithContext(ctx context.Context, request *GetSnapOverviewRequest) (response *GetSnapOverviewResponse, err error) {
    if request == nil {
        request = NewGetSnapOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSnapOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSnapOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewInitializeDisksRequest() (request *InitializeDisksRequest) {
    request = &InitializeDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "InitializeDisks")
    
    
    return
}

func NewInitializeDisksResponse() (response *InitializeDisksResponse) {
    response = &InitializeDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InitializeDisks
// This API is used to reinitialize the cloud disks. Note the following when reinitializing the cloud disks:
//
// 1. For a cloud disk created from a snapshot, it is rolled back to the state of the snapshot;
//
// 2. For a cloud disk created from the scratch, all data are cleared. Please check and back up the necessary data before the reinitialization;
//
// 3. Currently, you can only re-initialize a cloud disk when it’s not attached to a resource and not shared by others;
//
// 4. For a cloud disk created from a snapshot, if the snapshot has been deleted, it cannot be reinitialized.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) InitializeDisks(request *InitializeDisksRequest) (response *InitializeDisksResponse, err error) {
    return c.InitializeDisksWithContext(context.Background(), request)
}

// InitializeDisks
// This API is used to reinitialize the cloud disks. Note the following when reinitializing the cloud disks:
//
// 1. For a cloud disk created from a snapshot, it is rolled back to the state of the snapshot;
//
// 2. For a cloud disk created from the scratch, all data are cleared. Please check and back up the necessary data before the reinitialization;
//
// 3. Currently, you can only re-initialize a cloud disk when it’s not attached to a resource and not shared by others;
//
// 4. For a cloud disk created from a snapshot, if the snapshot has been deleted, it cannot be reinitialized.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) InitializeDisksWithContext(ctx context.Context, request *InitializeDisksRequest) (response *InitializeDisksResponse, err error) {
    if request == nil {
        request = NewInitializeDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InitializeDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewInitializeDisksResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceModifyDiskExtraPerformanceRequest() (request *InquirePriceModifyDiskExtraPerformanceRequest) {
    request = &InquirePriceModifyDiskExtraPerformanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "InquirePriceModifyDiskExtraPerformance")
    
    
    return
}

func NewInquirePriceModifyDiskExtraPerformanceResponse() (response *InquirePriceModifyDiskExtraPerformanceResponse) {
    response = &InquirePriceModifyDiskExtraPerformanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceModifyDiskExtraPerformance
// This API is used to query the price for adjusting the cloud disk’s extra performance.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
func (c *Client) InquirePriceModifyDiskExtraPerformance(request *InquirePriceModifyDiskExtraPerformanceRequest) (response *InquirePriceModifyDiskExtraPerformanceResponse, err error) {
    return c.InquirePriceModifyDiskExtraPerformanceWithContext(context.Background(), request)
}

// InquirePriceModifyDiskExtraPerformance
// This API is used to query the price for adjusting the cloud disk’s extra performance.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
func (c *Client) InquirePriceModifyDiskExtraPerformanceWithContext(ctx context.Context, request *InquirePriceModifyDiskExtraPerformanceRequest) (response *InquirePriceModifyDiskExtraPerformanceResponse, err error) {
    if request == nil {
        request = NewInquirePriceModifyDiskExtraPerformanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceModifyDiskExtraPerformance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceModifyDiskExtraPerformanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateDisksRequest() (request *InquiryPriceCreateDisksRequest) {
    request = &InquiryPriceCreateDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "InquiryPriceCreateDisks")
    
    
    return
}

func NewInquiryPriceCreateDisksResponse() (response *InquiryPriceCreateDisksResponse) {
    response = &InquiryPriceCreateDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceCreateDisks
// This API (InquiryPriceCreateDisks) is used to inquire the price for cloud disk creation.
//
// 
//
// * It supports inquiring the price for the creation of multiple cloud disks. The total price for the creation is returned.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) InquiryPriceCreateDisks(request *InquiryPriceCreateDisksRequest) (response *InquiryPriceCreateDisksResponse, err error) {
    return c.InquiryPriceCreateDisksWithContext(context.Background(), request)
}

// InquiryPriceCreateDisks
// This API (InquiryPriceCreateDisks) is used to inquire the price for cloud disk creation.
//
// 
//
// * It supports inquiring the price for the creation of multiple cloud disks. The total price for the creation is returned.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) InquiryPriceCreateDisksWithContext(ctx context.Context, request *InquiryPriceCreateDisksRequest) (response *InquiryPriceCreateDisksResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceCreateDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceCreateDisksResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceResizeDiskRequest() (request *InquiryPriceResizeDiskRequest) {
    request = &InquiryPriceResizeDiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "InquiryPriceResizeDisk")
    
    
    return
}

func NewInquiryPriceResizeDiskResponse() (response *InquiryPriceResizeDiskResponse) {
    response = &InquiryPriceResizeDiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceResizeDisk
// This API is used to query the price for expanding cloud disks.
//
// error code that may be returned:
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
func (c *Client) InquiryPriceResizeDisk(request *InquiryPriceResizeDiskRequest) (response *InquiryPriceResizeDiskResponse, err error) {
    return c.InquiryPriceResizeDiskWithContext(context.Background(), request)
}

// InquiryPriceResizeDisk
// This API is used to query the price for expanding cloud disks.
//
// error code that may be returned:
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
func (c *Client) InquiryPriceResizeDiskWithContext(ctx context.Context, request *InquiryPriceResizeDiskRequest) (response *InquiryPriceResizeDiskResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResizeDiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceResizeDisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceResizeDiskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoSnapshotPolicyAttributeRequest() (request *ModifyAutoSnapshotPolicyAttributeRequest) {
    request = &ModifyAutoSnapshotPolicyAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "ModifyAutoSnapshotPolicyAttribute")
    
    
    return
}

func NewModifyAutoSnapshotPolicyAttributeResponse() (response *ModifyAutoSnapshotPolicyAttributeResponse) {
    response = &ModifyAutoSnapshotPolicyAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAutoSnapshotPolicyAttribute
// This API (ModifyAutoSnapshotPolicyAttribute) is used to modify the attributes of an automatic snapshot policy.
//
// 
//
// * You can use this API to modify the attributes of a scheduled snapshot policy, including the execution policy, name, and activation.
//
// * When modifying the number of days for retention, you must ensure that there is no clash with the permanent retention attribute. Otherwise, the entire operation will fail and a specific error code will be returned.
//
// error code that may be returned:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) ModifyAutoSnapshotPolicyAttribute(request *ModifyAutoSnapshotPolicyAttributeRequest) (response *ModifyAutoSnapshotPolicyAttributeResponse, err error) {
    return c.ModifyAutoSnapshotPolicyAttributeWithContext(context.Background(), request)
}

// ModifyAutoSnapshotPolicyAttribute
// This API (ModifyAutoSnapshotPolicyAttribute) is used to modify the attributes of an automatic snapshot policy.
//
// 
//
// * You can use this API to modify the attributes of a scheduled snapshot policy, including the execution policy, name, and activation.
//
// * When modifying the number of days for retention, you must ensure that there is no clash with the permanent retention attribute. Otherwise, the entire operation will fail and a specific error code will be returned.
//
// error code that may be returned:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
func (c *Client) ModifyAutoSnapshotPolicyAttributeWithContext(ctx context.Context, request *ModifyAutoSnapshotPolicyAttributeRequest) (response *ModifyAutoSnapshotPolicyAttributeResponse, err error) {
    if request == nil {
        request = NewModifyAutoSnapshotPolicyAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoSnapshotPolicyAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoSnapshotPolicyAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDiskAttributesRequest() (request *ModifyDiskAttributesRequest) {
    request = &ModifyDiskAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "ModifyDiskAttributes")
    
    
    return
}

func NewModifyDiskAttributesResponse() (response *ModifyDiskAttributesResponse) {
    response = &ModifyDiskAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDiskAttributes
// * Only the project ID of elastic cloud disk can be modified. The project ID of the cloud disk created with the CVM is linked with the CVM. The project ID can be can be queried in the Portable field in the output parameters through the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1).
//
// * "Cloud disk name" is only used by users for their management. Tencent Cloud does not use the name as the basis for ticket submission or cloud disk management.
//
// * Batch operations are supported. If multiple cloud disk IDs are specified, all the specified cloud disks must have the same attribute. If there is a cloud disk that does not allow this operation, the operation is not performed and a specific error code is returned.
//
// error code that may be returned:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_DISKMIGRATING = "ResourceInUse.DiskMigrating"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) ModifyDiskAttributes(request *ModifyDiskAttributesRequest) (response *ModifyDiskAttributesResponse, err error) {
    return c.ModifyDiskAttributesWithContext(context.Background(), request)
}

// ModifyDiskAttributes
// * Only the project ID of elastic cloud disk can be modified. The project ID of the cloud disk created with the CVM is linked with the CVM. The project ID can be can be queried in the Portable field in the output parameters through the API [DescribeDisks](https://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1).
//
// * "Cloud disk name" is only used by users for their management. Tencent Cloud does not use the name as the basis for ticket submission or cloud disk management.
//
// * Batch operations are supported. If multiple cloud disk IDs are specified, all the specified cloud disks must have the same attribute. If there is a cloud disk that does not allow this operation, the operation is not performed and a specific error code is returned.
//
// error code that may be returned:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE_DISKMIGRATING = "ResourceInUse.DiskMigrating"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
func (c *Client) ModifyDiskAttributesWithContext(ctx context.Context, request *ModifyDiskAttributesRequest) (response *ModifyDiskAttributesResponse, err error) {
    if request == nil {
        request = NewModifyDiskAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDiskAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDiskAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDiskExtraPerformanceRequest() (request *ModifyDiskExtraPerformanceRequest) {
    request = &ModifyDiskExtraPerformanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "ModifyDiskExtraPerformance")
    
    
    return
}

func NewModifyDiskExtraPerformanceResponse() (response *ModifyDiskExtraPerformanceResponse) {
    response = &ModifyDiskExtraPerformanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDiskExtraPerformance
// This API is used to adjust the cloud disk’s extra performance.
//
// 
//
// * Currently, only Tremendous SSD (CLOUD_TSSD) and Enhanced SSD (CLOUD_HSSD) support extra performance adjustment.
//
// error code that may be returned:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDiskExtraPerformance(request *ModifyDiskExtraPerformanceRequest) (response *ModifyDiskExtraPerformanceResponse, err error) {
    return c.ModifyDiskExtraPerformanceWithContext(context.Background(), request)
}

// ModifyDiskExtraPerformance
// This API is used to adjust the cloud disk’s extra performance.
//
// 
//
// * Currently, only Tremendous SSD (CLOUD_TSSD) and Enhanced SSD (CLOUD_HSSD) support extra performance adjustment.
//
// error code that may be returned:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDiskExtraPerformanceWithContext(ctx context.Context, request *ModifyDiskExtraPerformanceRequest) (response *ModifyDiskExtraPerformanceResponse, err error) {
    if request == nil {
        request = NewModifyDiskExtraPerformanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDiskExtraPerformance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDiskExtraPerformanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotAttributeRequest() (request *ModifySnapshotAttributeRequest) {
    request = &ModifySnapshotAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "ModifySnapshotAttribute")
    
    
    return
}

func NewModifySnapshotAttributeResponse() (response *ModifySnapshotAttributeResponse) {
    response = &ModifySnapshotAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySnapshotAttribute
// This API (ModifySnapshotAttribute) is used to modify the attributes of a specified snapshot.
//
// 
//
// * Currently, you can only modify snapshot name and change non-permanent snapshots into permanent snapshots.
//
// * "Snapshot name" is only used by users for their management. Tencent Cloud does not use the name as the basis for ticket submission or snapshot management.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifySnapshotAttribute(request *ModifySnapshotAttributeRequest) (response *ModifySnapshotAttributeResponse, err error) {
    return c.ModifySnapshotAttributeWithContext(context.Background(), request)
}

// ModifySnapshotAttribute
// This API (ModifySnapshotAttribute) is used to modify the attributes of a specified snapshot.
//
// 
//
// * Currently, you can only modify snapshot name and change non-permanent snapshots into permanent snapshots.
//
// * "Snapshot name" is only used by users for their management. Tencent Cloud does not use the name as the basis for ticket submission or snapshot management.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifySnapshotAttributeWithContext(ctx context.Context, request *ModifySnapshotAttributeRequest) (response *ModifySnapshotAttributeResponse, err error) {
    if request == nil {
        request = NewModifySnapshotAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySnapshotAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySnapshotAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotsSharePermissionRequest() (request *ModifySnapshotsSharePermissionRequest) {
    request = &ModifySnapshotsSharePermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "ModifySnapshotsSharePermission")
    
    
    return
}

func NewModifySnapshotsSharePermissionResponse() (response *ModifySnapshotsSharePermissionResponse) {
    response = &ModifySnapshotsSharePermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySnapshotsSharePermission
// This API is used to modify snapshot sharing information.
//
// 
//
// After snapshots are shared, the accounts they are shared to can use the snapshot to create cloud disks.
//
// * Each snapshot can be shared to at most 50 accounts.
//
// * You can use a shared snapshot to create cloud disks, but you cannot change its name or description.
//
// * Snapshots can only be shared with accounts in the same region.
//
// * Only data disk snapshots can be shared.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_OVERQUOTA = "ResourceInsufficient.OverQuota"
func (c *Client) ModifySnapshotsSharePermission(request *ModifySnapshotsSharePermissionRequest) (response *ModifySnapshotsSharePermissionResponse, err error) {
    return c.ModifySnapshotsSharePermissionWithContext(context.Background(), request)
}

// ModifySnapshotsSharePermission
// This API is used to modify snapshot sharing information.
//
// 
//
// After snapshots are shared, the accounts they are shared to can use the snapshot to create cloud disks.
//
// * Each snapshot can be shared to at most 50 accounts.
//
// * You can use a shared snapshot to create cloud disks, but you cannot change its name or description.
//
// * Snapshots can only be shared with accounts in the same region.
//
// * Only data disk snapshots can be shared.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"
//  INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_OVERQUOTA = "ResourceInsufficient.OverQuota"
func (c *Client) ModifySnapshotsSharePermissionWithContext(ctx context.Context, request *ModifySnapshotsSharePermissionRequest) (response *ModifySnapshotsSharePermissionResponse, err error) {
    if request == nil {
        request = NewModifySnapshotsSharePermissionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySnapshotsSharePermission require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySnapshotsSharePermissionResponse()
    err = c.Send(request, response)
    return
}

func NewResizeDiskRequest() (request *ResizeDiskRequest) {
    request = &ResizeDiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "ResizeDisk")
    
    
    return
}

func NewResizeDiskResponse() (response *ResizeDiskResponse) {
    response = &ResizeDiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResizeDisk
// This API is used to expand the capacity of a cloud disk.
//
// 
//
// * This API supports only the expansion of elastic cloud disks. To query the type of a cloud disk, you can call the [DescribeDisks](https://intl.cloud.tencent.comhttps://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1?from_cn_redirect=1) API and check the `Portable` field in the response. To expand non-elastic cloud disks, you can call the [ResizeInstanceDisks](https://intl.cloud.tencent.com/document/product/213/15731?from_cn_redirect=1) API.
//
// * This is an async API. A successful return of this API does not mean that the cloud disk has been expanded successfully. You can call the [DescribeDisks](https://intl.cloud.tencent.comhttps://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1?from_cn_redirect=1) API to query the status of a cloud disk. `EXPANDING` indicates that the expansion is in process. 
//
// error code that may be returned:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
//  UNSUPPORTEDOPERATION_INSTANCENOTSTOPPED = "UnsupportedOperation.InstanceNotStopped"
func (c *Client) ResizeDisk(request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
    return c.ResizeDiskWithContext(context.Background(), request)
}

// ResizeDisk
// This API is used to expand the capacity of a cloud disk.
//
// 
//
// * This API supports only the expansion of elastic cloud disks. To query the type of a cloud disk, you can call the [DescribeDisks](https://intl.cloud.tencent.comhttps://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1?from_cn_redirect=1) API and check the `Portable` field in the response. To expand non-elastic cloud disks, you can call the [ResizeInstanceDisks](https://intl.cloud.tencent.com/document/product/213/15731?from_cn_redirect=1) API.
//
// * This is an async API. A successful return of this API does not mean that the cloud disk has been expanded successfully. You can call the [DescribeDisks](https://intl.cloud.tencent.comhttps://intl.cloud.tencent.com/document/product/362/16315?from_cn_redirect=1?from_cn_redirect=1) API to query the status of a cloud disk. `EXPANDING` indicates that the expansion is in process. 
//
// error code that may be returned:
//  INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDDISK_BUSY = "InvalidDisk.Busy"
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  TRADEDEALCONFLICT = "TradeDealConflict"
//  UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"
//  UNSUPPORTEDOPERATION_INSTANCENOTSTOPPED = "UnsupportedOperation.InstanceNotStopped"
func (c *Client) ResizeDiskWithContext(ctx context.Context, request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
    if request == nil {
        request = NewResizeDiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResizeDisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewResizeDiskResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDisksRequest() (request *TerminateDisksRequest) {
    request = &TerminateDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "TerminateDisks")
    
    
    return
}

func NewTerminateDisksResponse() (response *TerminateDisksResponse) {
    response = &TerminateDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateDisks
// This API is used to return cloud disks.
//
// 
//
// * You can use this API to return cloud disks you no longer need.
//
// * This API can be used to return pay-as-you-go cloud disks billed on hourly basis. 
//
// * Batch operations are supported. The maximum number of cloud disks in each request is 50. If there is any specified cloud disk that cannot be returned, an error code will be returned.
//
// error code that may be returned:
//  INSUFFICIENTREFUNDQUOTA = "InsufficientRefundQuota"
//  INTERNALERROR_FAILQUERYRESOURCE = "InternalError.FailQueryResource"
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINSUFFICIENT_OVERREFUNDQUOTA = "ResourceInsufficient.OverRefundQuota"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
//  RESOURCEUNAVAILABLE_NOTSUPPORTREFUND = "ResourceUnavailable.NotSupportRefund"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_REPEATREFUND = "ResourceUnavailable.RepeatRefund"
//  TRADEDEALCONFLICT = "TradeDealConflict"
func (c *Client) TerminateDisks(request *TerminateDisksRequest) (response *TerminateDisksResponse, err error) {
    return c.TerminateDisksWithContext(context.Background(), request)
}

// TerminateDisks
// This API is used to return cloud disks.
//
// 
//
// * You can use this API to return cloud disks you no longer need.
//
// * This API can be used to return pay-as-you-go cloud disks billed on hourly basis. 
//
// * Batch operations are supported. The maximum number of cloud disks in each request is 50. If there is any specified cloud disk that cannot be returned, an error code will be returned.
//
// error code that may be returned:
//  INSUFFICIENTREFUNDQUOTA = "InsufficientRefundQuota"
//  INTERNALERROR_FAILQUERYRESOURCE = "InternalError.FailQueryResource"
//  INVALIDDISK_EXPIRE = "InvalidDisk.Expire"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEBUSY = "ResourceBusy"
//  RESOURCEINSUFFICIENT_OVERREFUNDQUOTA = "ResourceInsufficient.OverRefundQuota"
//  RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"
//  RESOURCEUNAVAILABLE_NOTSUPPORTREFUND = "ResourceUnavailable.NotSupportRefund"
//  RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"
//  RESOURCEUNAVAILABLE_REPEATREFUND = "ResourceUnavailable.RepeatRefund"
//  TRADEDEALCONFLICT = "TradeDealConflict"
func (c *Client) TerminateDisksWithContext(ctx context.Context, request *TerminateDisksRequest) (response *TerminateDisksResponse, err error) {
    if request == nil {
        request = NewTerminateDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateDisksResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindAutoSnapshotPolicyRequest() (request *UnbindAutoSnapshotPolicyRequest) {
    request = &UnbindAutoSnapshotPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cbs", APIVersion, "UnbindAutoSnapshotPolicy")
    
    
    return
}

func NewUnbindAutoSnapshotPolicyResponse() (response *UnbindAutoSnapshotPolicyResponse) {
    response = &UnbindAutoSnapshotPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindAutoSnapshotPolicy
// This API (UnbindAutoSnapshotPolicy) is used to unbind the cloud disk from the specified scheduled snapshot policy.
//
// 
//
// * Batch operations are supported. Multiple cloud disks can be unbound from a snapshot policy at one time. 
//
// * If the passed-in cloud disks are not bound to the current scheduled snapshot policy, they will be skipped. Only cloud disks that are bound to the current scheduled snapshot policy are processed.
//
// error code that may be returned:
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UnbindAutoSnapshotPolicy(request *UnbindAutoSnapshotPolicyRequest) (response *UnbindAutoSnapshotPolicyResponse, err error) {
    return c.UnbindAutoSnapshotPolicyWithContext(context.Background(), request)
}

// UnbindAutoSnapshotPolicy
// This API (UnbindAutoSnapshotPolicy) is used to unbind the cloud disk from the specified scheduled snapshot policy.
//
// 
//
// * Batch operations are supported. Multiple cloud disks can be unbound from a snapshot policy at one time. 
//
// * If the passed-in cloud disks are not bound to the current scheduled snapshot policy, they will be skipped. Only cloud disks that are bound to the current scheduled snapshot policy are processed.
//
// error code that may be returned:
//  INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) UnbindAutoSnapshotPolicyWithContext(ctx context.Context, request *UnbindAutoSnapshotPolicyRequest) (response *UnbindAutoSnapshotPolicyResponse, err error) {
    if request == nil {
        request = NewUnbindAutoSnapshotPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindAutoSnapshotPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindAutoSnapshotPolicyResponse()
    err = c.Send(request, response)
    return
}
