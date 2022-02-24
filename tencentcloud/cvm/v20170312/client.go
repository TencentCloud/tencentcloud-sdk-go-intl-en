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
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
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


func NewAllocateHostsRequest() (request *AllocateHostsRequest) {
    request = &AllocateHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "AllocateHosts")
    
    
    return
}

func NewAllocateHostsResponse() (response *AllocateHostsResponse) {
    response = &AllocateHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AllocateHosts
// This API is used to create CDH instances with specified configuration.
//
// * When HostChargeType is PREPAID, the HostChargePrepaid parameter must be specified.
//
// error code that may be returned:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDPERIOD = "InvalidPeriod"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  INVALIDREGION_NOTFOUND = "InvalidRegion.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  RESOURCEINSUFFICIENT_ZONESOLDOUTFORSPECIFIEDINSTANCE = "ResourceInsufficient.ZoneSoldOutForSpecifiedInstance"
func (c *Client) AllocateHosts(request *AllocateHostsRequest) (response *AllocateHostsResponse, err error) {
    if request == nil {
        request = NewAllocateHostsRequest()
    }
    
    response = NewAllocateHostsResponse()
    err = c.Send(request, response)
    return
}

// AllocateHosts
// This API is used to create CDH instances with specified configuration.
//
// * When HostChargeType is PREPAID, the HostChargePrepaid parameter must be specified.
//
// error code that may be returned:
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDPERIOD = "InvalidPeriod"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  INVALIDREGION_NOTFOUND = "InvalidRegion.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  RESOURCEINSUFFICIENT_ZONESOLDOUTFORSPECIFIEDINSTANCE = "ResourceInsufficient.ZoneSoldOutForSpecifiedInstance"
func (c *Client) AllocateHostsWithContext(ctx context.Context, request *AllocateHostsRequest) (response *AllocateHostsResponse, err error) {
    if request == nil {
        request = NewAllocateHostsRequest()
    }
    request.SetContext(ctx)
    
    response = NewAllocateHostsResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateInstancesKeyPairsRequest() (request *AssociateInstancesKeyPairsRequest) {
    request = &AssociateInstancesKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "AssociateInstancesKeyPairs")
    
    
    return
}

func NewAssociateInstancesKeyPairsResponse() (response *AssociateInstancesKeyPairsResponse) {
    response = &AssociateInstancesKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateInstancesKeyPairs
// This API is used to associate key pairs with instances.
//
// 
//
// * If the public key of a key pair is written to the `SSH` configuration of the instance, users will be able to log in to the instance with the private key of the key pair.
//
// * If the instance is already associated with a key, the old key will become invalid.
//
// If you currently use a password to log in, you will no longer be able to do so after you associate the instance with a key.
//
// * Batch operations are supported. The maximum number of instances in each request is 100. If any instance in the request cannot be associated with a key, you will get an error code.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDKEYPAIRID_MALFORMED = "InvalidKeyPairId.Malformed"
//  INVALIDKEYPAIRID_NOTFOUND = "InvalidKeyPairId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYPAIRNOTSUPPORTED = "InvalidParameterValue.KeyPairNotSupported"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCEOSWINDOWS = "UnsupportedOperation.InstanceOsWindows"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) AssociateInstancesKeyPairs(request *AssociateInstancesKeyPairsRequest) (response *AssociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewAssociateInstancesKeyPairsRequest()
    }
    
    response = NewAssociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

// AssociateInstancesKeyPairs
// This API is used to associate key pairs with instances.
//
// 
//
// * If the public key of a key pair is written to the `SSH` configuration of the instance, users will be able to log in to the instance with the private key of the key pair.
//
// * If the instance is already associated with a key, the old key will become invalid.
//
// If you currently use a password to log in, you will no longer be able to do so after you associate the instance with a key.
//
// * Batch operations are supported. The maximum number of instances in each request is 100. If any instance in the request cannot be associated with a key, you will get an error code.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDKEYPAIRID_MALFORMED = "InvalidKeyPairId.Malformed"
//  INVALIDKEYPAIRID_NOTFOUND = "InvalidKeyPairId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYPAIRNOTSUPPORTED = "InvalidParameterValue.KeyPairNotSupported"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCEOSWINDOWS = "UnsupportedOperation.InstanceOsWindows"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) AssociateInstancesKeyPairsWithContext(ctx context.Context, request *AssociateInstancesKeyPairsRequest) (response *AssociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewAssociateInstancesKeyPairsRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "AssociateSecurityGroups")
    
    
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateSecurityGroups
// This API is used to associate security groups with specified instances.
//
// error code that may be returned:
//  FAILEDOPERATION_SECURITYGROUPACTIONFAILED = "FailedOperation.SecurityGroupActionFailed"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDSGID_MALFORMED = "InvalidSgId.Malformed"
//  LIMITEXCEEDED_ASSOCIATEUSGLIMITEXCEEDED = "LimitExceeded.AssociateUSGLimitExceeded"
//  LIMITEXCEEDED_CVMSVIFSPERSECGROUPLIMITEXCEEDED = "LimitExceeded.CvmsVifsPerSecGroupLimitExceeded"
//  LIMITEXCEEDED_SINGLEUSGQUOTA = "LimitExceeded.SingleUSGQuota"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  SECGROUPACTIONFAILURE = "SecGroupActionFailure"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

// AssociateSecurityGroups
// This API is used to associate security groups with specified instances.
//
// error code that may be returned:
//  FAILEDOPERATION_SECURITYGROUPACTIONFAILED = "FailedOperation.SecurityGroupActionFailed"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDSGID_MALFORMED = "InvalidSgId.Malformed"
//  LIMITEXCEEDED_ASSOCIATEUSGLIMITEXCEEDED = "LimitExceeded.AssociateUSGLimitExceeded"
//  LIMITEXCEEDED_CVMSVIFSPERSECGROUPLIMITEXCEEDED = "LimitExceeded.CvmsVifsPerSecGroupLimitExceeded"
//  LIMITEXCEEDED_SINGLEUSGQUOTA = "LimitExceeded.SingleUSGQuota"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  SECGROUPACTIONFAILURE = "SecGroupActionFailure"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
func (c *Client) AssociateSecurityGroupsWithContext(ctx context.Context, request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDisasterRecoverGroupRequest() (request *CreateDisasterRecoverGroupRequest) {
    request = &CreateDisasterRecoverGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "CreateDisasterRecoverGroup")
    
    
    return
}

func NewCreateDisasterRecoverGroupResponse() (response *CreateDisasterRecoverGroupResponse) {
    response = &CreateDisasterRecoverGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDisasterRecoverGroup
// This API is used to create a [spread placement group](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1). After you create one, you can specify it for an instance when you [create the instance](https://intl.cloud.tencent.com/document/api/213/15730?from_cn_redirect=1), 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
func (c *Client) CreateDisasterRecoverGroup(request *CreateDisasterRecoverGroupRequest) (response *CreateDisasterRecoverGroupResponse, err error) {
    if request == nil {
        request = NewCreateDisasterRecoverGroupRequest()
    }
    
    response = NewCreateDisasterRecoverGroupResponse()
    err = c.Send(request, response)
    return
}

// CreateDisasterRecoverGroup
// This API is used to create a [spread placement group](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1). After you create one, you can specify it for an instance when you [create the instance](https://intl.cloud.tencent.com/document/api/213/15730?from_cn_redirect=1), 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
func (c *Client) CreateDisasterRecoverGroupWithContext(ctx context.Context, request *CreateDisasterRecoverGroupRequest) (response *CreateDisasterRecoverGroupResponse, err error) {
    if request == nil {
        request = NewCreateDisasterRecoverGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateDisasterRecoverGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageRequest() (request *CreateImageRequest) {
    request = &CreateImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "CreateImage")
    
    
    return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
    response = &CreateImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateImage
// This API is used to create a new image with the system disk of an instance. The image can be used to create new instances.
//
// error code that may be returned:
//  IMAGEQUOTALIMITEXCEEDED = "ImageQuotaLimitExceeded"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDIMAGENAME_DUPLICATE = "InvalidImageName.Duplicate"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER_DATADISKIDCONTAINSROOTDISK = "InvalidParameter.DataDiskIdContainsRootDisk"
//  INVALIDPARAMETER_DATADISKNOTBELONGSPECIFIEDINSTANCE = "InvalidParameter.DataDiskNotBelongSpecifiedInstance"
//  INVALIDPARAMETER_DUPLICATESYSTEMSNAPSHOTS = "InvalidParameter.DuplicateSystemSnapshots"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETER_INVALIDDEPENDENCE = "InvalidParameter.InvalidDependence"
//  INVALIDPARAMETER_LOCALDATADISKNOTSUPPORT = "InvalidParameter.LocalDataDiskNotSupport"
//  INVALIDPARAMETER_SNAPSHOTNOTFOUND = "InvalidParameter.SnapshotNotFound"
//  INVALIDPARAMETER_SPECIFYONEPARAMETER = "InvalidParameter.SpecifyOneParameter"
//  INVALIDPARAMETER_SWAPDISKNOTSUPPORT = "InvalidParameter.SwapDiskNotSupport"
//  INVALIDPARAMETER_SYSTEMSNAPSHOTNOTFOUND = "InvalidParameter.SystemSnapshotNotFound"
//  INVALIDPARAMETER_VALUETOOLARGE = "InvalidParameter.ValueTooLarge"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_NOTSUPPORTINSTANCEIMAGE = "UnsupportedOperation.NotSupportInstanceImage"
//  UNSUPPORTEDOPERATION_SPECIALINSTANCETYPE = "UnsupportedOperation.SpecialInstanceType"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
    if request == nil {
        request = NewCreateImageRequest()
    }
    
    response = NewCreateImageResponse()
    err = c.Send(request, response)
    return
}

// CreateImage
// This API is used to create a new image with the system disk of an instance. The image can be used to create new instances.
//
// error code that may be returned:
//  IMAGEQUOTALIMITEXCEEDED = "ImageQuotaLimitExceeded"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDIMAGENAME_DUPLICATE = "InvalidImageName.Duplicate"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER_DATADISKIDCONTAINSROOTDISK = "InvalidParameter.DataDiskIdContainsRootDisk"
//  INVALIDPARAMETER_DATADISKNOTBELONGSPECIFIEDINSTANCE = "InvalidParameter.DataDiskNotBelongSpecifiedInstance"
//  INVALIDPARAMETER_DUPLICATESYSTEMSNAPSHOTS = "InvalidParameter.DuplicateSystemSnapshots"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETER_INVALIDDEPENDENCE = "InvalidParameter.InvalidDependence"
//  INVALIDPARAMETER_LOCALDATADISKNOTSUPPORT = "InvalidParameter.LocalDataDiskNotSupport"
//  INVALIDPARAMETER_SNAPSHOTNOTFOUND = "InvalidParameter.SnapshotNotFound"
//  INVALIDPARAMETER_SPECIFYONEPARAMETER = "InvalidParameter.SpecifyOneParameter"
//  INVALIDPARAMETER_SWAPDISKNOTSUPPORT = "InvalidParameter.SwapDiskNotSupport"
//  INVALIDPARAMETER_SYSTEMSNAPSHOTNOTFOUND = "InvalidParameter.SystemSnapshotNotFound"
//  INVALIDPARAMETER_VALUETOOLARGE = "InvalidParameter.ValueTooLarge"
//  INVALIDPARAMETERCONFLICT = "InvalidParameterConflict"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_NOTSUPPORTINSTANCEIMAGE = "UnsupportedOperation.NotSupportInstanceImage"
//  UNSUPPORTEDOPERATION_SPECIALINSTANCETYPE = "UnsupportedOperation.SpecialInstanceType"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest) (response *CreateImageResponse, err error) {
    if request == nil {
        request = NewCreateImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKeyPairRequest() (request *CreateKeyPairRequest) {
    request = &CreateKeyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "CreateKeyPair")
    
    
    return
}

func NewCreateKeyPairResponse() (response *CreateKeyPairResponse) {
    response = &CreateKeyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateKeyPair
// This API is used to create an `OpenSSH RSA` key pair, which you can use to log in to a `Linux` instance.
//
// 
//
// * You only need to specify a name, and the system will automatically create a key pair and return its `ID` and the public and private keys.
//
// * The name of the key pair must be unique.
//
// * You can save the private key to a file and use it as an authentication method for `SSH`.
//
// * Tencent Cloud does not save users' private keys. Be sure to save it yourself.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDKEYPAIR_LIMITEXCEEDED = "InvalidKeyPair.LimitExceeded"
//  INVALIDKEYPAIRNAME_DUPLICATE = "InvalidKeyPairName.Duplicate"
//  INVALIDKEYPAIRNAMEEMPTY = "InvalidKeyPairNameEmpty"
//  INVALIDKEYPAIRNAMEINCLUDEILLEGALCHAR = "InvalidKeyPairNameIncludeIllegalChar"
//  INVALIDKEYPAIRNAMETOOLONG = "InvalidKeyPairNameTooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateKeyPair(request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
    if request == nil {
        request = NewCreateKeyPairRequest()
    }
    
    response = NewCreateKeyPairResponse()
    err = c.Send(request, response)
    return
}

// CreateKeyPair
// This API is used to create an `OpenSSH RSA` key pair, which you can use to log in to a `Linux` instance.
//
// 
//
// * You only need to specify a name, and the system will automatically create a key pair and return its `ID` and the public and private keys.
//
// * The name of the key pair must be unique.
//
// * You can save the private key to a file and use it as an authentication method for `SSH`.
//
// * Tencent Cloud does not save users' private keys. Be sure to save it yourself.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDKEYPAIR_LIMITEXCEEDED = "InvalidKeyPair.LimitExceeded"
//  INVALIDKEYPAIRNAME_DUPLICATE = "InvalidKeyPairName.Duplicate"
//  INVALIDKEYPAIRNAMEEMPTY = "InvalidKeyPairNameEmpty"
//  INVALIDKEYPAIRNAMEINCLUDEILLEGALCHAR = "InvalidKeyPairNameIncludeIllegalChar"
//  INVALIDKEYPAIRNAMETOOLONG = "InvalidKeyPairNameTooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateKeyPairWithContext(ctx context.Context, request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
    if request == nil {
        request = NewCreateKeyPairRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateKeyPairResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLaunchTemplateRequest() (request *CreateLaunchTemplateRequest) {
    request = &CreateLaunchTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "CreateLaunchTemplate")
    
    
    return
}

func NewCreateLaunchTemplateResponse() (response *CreateLaunchTemplateResponse) {
    response = &CreateLaunchTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLaunchTemplate
// This API is used to create an instance launch template.
//
// 
//
// An instance launch template contains the configuration information required to create an instance, including instance type, data/system disk type and size, and security group, etc.
//
// 
//
// When a template is created, it defaults to Version 1. You can use `CreateLaunchTemplateVersion` to create new versions of this template, with the version number increasing. When you run `RunInstances` to create instances, you can specify the instance launch template version. If it’s not specified, the default template version is used.
//
// error code that may be returned:
//  ACCOUNTQUALIFICATIONRESTRICTIONS = "AccountQualificationRestrictions"
//  AUTHFAILURE_CAMROLENAMEAUTHENTICATEFAILED = "AuthFailure.CamRoleNameAuthenticateFailed"
//  FAILEDOPERATION_DISASTERRECOVERGROUPNOTFOUND = "FailedOperation.DisasterRecoverGroupNotFound"
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  FAILEDOPERATION_TAGKEYRESERVED = "FailedOperation.TagKeyReserved"
//  INSTANCESQUOTALIMITEXCEEDED = "InstancesQuotaLimitExceeded"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDCLIENTTOKEN_TOOLONG = "InvalidClientToken.TooLong"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDHOSTID_NOTFOUND = "InvalidHostId.NotFound"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDINSTANCENAME_TOOLONG = "InvalidInstanceName.TooLong"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETER_INVALIDIPFORMAT = "InvalidParameter.InvalidIpFormat"
//  INVALIDPARAMETER_LACKCORECOUNTORTHREADPERCORE = "InvalidParameter.LackCoreCountOrThreadPerCore"
//  INVALIDPARAMETER_PASSWORDNOTSUPPORTED = "InvalidParameter.PasswordNotSupported"
//  INVALIDPARAMETER_SNAPSHOTNOTFOUND = "InvalidParameter.SnapshotNotFound"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLOUDSSDDATADISKSIZETOOSMALL = "InvalidParameterValue.CloudSsdDataDiskSizeTooSmall"
//  INVALIDPARAMETERVALUE_CORECOUNTVALUE = "InvalidParameterValue.CoreCountValue"
//  INVALIDPARAMETERVALUE_ILLEGALHOSTNAME = "InvalidParameterValue.IllegalHostName"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTSUPPORTHPCCLUSTER = "InvalidParameterValue.InstanceTypeNotSupportHpcCluster"
//  INVALIDPARAMETERVALUE_INSTANCETYPEREQUIREDHPCCLUSTER = "InvalidParameterValue.InstanceTypeRequiredHpcCluster"
//  INVALIDPARAMETERVALUE_INSUFFICIENTOFFERING = "InvalidParameterValue.InsufficientOffering"
//  INVALIDPARAMETERVALUE_INSUFFICIENTPRICE = "InvalidParameterValue.InsufficientPrice"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHTEMPLATEDESCRIPTION = "InvalidParameterValue.InvalidLaunchTemplateDescription"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHTEMPLATENAME = "InvalidParameterValue.InvalidLaunchTemplateName"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHTEMPLATEVERSIONDESCRIPTION = "InvalidParameterValue.InvalidLaunchTemplateVersionDescription"
//  INVALIDPARAMETERVALUE_INVALIDUSERDATAFORMAT = "InvalidParameterValue.InvalidUserDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MUSTDHCPENABLEDVPC = "InvalidParameterValue.MustDhcpEnabledVpc"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SUBNETNOTEXIST = "InvalidParameterValue.SubnetNotExist"
//  INVALIDPARAMETERVALUE_THREADPERCOREVALUE = "InvalidParameterValue.ThreadPerCoreValue"
//  INVALIDPARAMETERVALUE_VPCIDZONEIDNOTMATCH = "InvalidParameterValue.VpcIdZoneIdNotMatch"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDPASSWORD = "InvalidPassword"
//  INVALIDPERIOD = "InvalidPeriod"
//  INVALIDPERMISSION = "InvalidPermission"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"
//  LIMITEXCEEDED_LAUNCHTEMPLATEQUOTA = "LimitExceeded.LaunchTemplateQuota"
//  LIMITEXCEEDED_SINGLEUSGQUOTA = "LimitExceeded.SingleUSGQuota"
//  LIMITEXCEEDED_SPOTQUOTA = "LimitExceeded.SpotQuota"
//  LIMITEXCEEDED_USERSPOTQUOTA = "LimitExceeded.UserSpotQuota"
//  LIMITEXCEEDED_VPCSUBNETNUM = "LimitExceeded.VpcSubnetNum"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DPDKINSTANCETYPEREQUIREDVPC = "MissingParameter.DPDKInstanceTypeRequiredVPC"
//  MISSINGPARAMETER_MONITORSERVICE = "MissingParameter.MonitorService"
//  RESOURCEINSUFFICIENT_CLOUDDISKSOLDOUT = "ResourceInsufficient.CloudDiskSoldOut"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCEINSUFFICIENT_DISASTERRECOVERGROUPCVMQUOTA = "ResourceInsufficient.DisasterRecoverGroupCvmQuota"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  RESOURCENOTFOUND_HPCCLUSTER = "ResourceNotFound.HpcCluster"
//  RESOURCENOTFOUND_NODEFAULTCBS = "ResourceNotFound.NoDefaultCbs"
//  RESOURCENOTFOUND_NODEFAULTCBSWITHREASON = "ResourceNotFound.NoDefaultCbsWithReason"
//  RESOURCEUNAVAILABLE_INSTANCETYPE = "ResourceUnavailable.InstanceType"
//  RESOURCESSOLDOUT_EIPINSUFFICIENT = "ResourcesSoldOut.EipInsufficient"
//  RESOURCESSOLDOUT_SPECIFIEDINSTANCETYPE = "ResourcesSoldOut.SpecifiedInstanceType"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDDISK = "UnsupportedOperation.InvalidDisk"
//  UNSUPPORTEDOPERATION_KEYPAIRUNSUPPORTEDWINDOWS = "UnsupportedOperation.KeyPairUnsupportedWindows"
//  UNSUPPORTEDOPERATION_NOINSTANCETYPESUPPORTSPOT = "UnsupportedOperation.NoInstanceTypeSupportSpot"
//  UNSUPPORTEDOPERATION_NOTSUPPORTIMPORTINSTANCESACTIONTIMER = "UnsupportedOperation.NotSupportImportInstancesActionTimer"
//  UNSUPPORTEDOPERATION_ONLYFORPREPAIDACCOUNT = "UnsupportedOperation.OnlyForPrepaidAccount"
//  VPCADDRNOTINSUBNET = "VpcAddrNotInSubNet"
//  VPCIPISUSED = "VpcIpIsUsed"
func (c *Client) CreateLaunchTemplate(request *CreateLaunchTemplateRequest) (response *CreateLaunchTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLaunchTemplateRequest()
    }
    
    response = NewCreateLaunchTemplateResponse()
    err = c.Send(request, response)
    return
}

// CreateLaunchTemplate
// This API is used to create an instance launch template.
//
// 
//
// An instance launch template contains the configuration information required to create an instance, including instance type, data/system disk type and size, and security group, etc.
//
// 
//
// When a template is created, it defaults to Version 1. You can use `CreateLaunchTemplateVersion` to create new versions of this template, with the version number increasing. When you run `RunInstances` to create instances, you can specify the instance launch template version. If it’s not specified, the default template version is used.
//
// error code that may be returned:
//  ACCOUNTQUALIFICATIONRESTRICTIONS = "AccountQualificationRestrictions"
//  AUTHFAILURE_CAMROLENAMEAUTHENTICATEFAILED = "AuthFailure.CamRoleNameAuthenticateFailed"
//  FAILEDOPERATION_DISASTERRECOVERGROUPNOTFOUND = "FailedOperation.DisasterRecoverGroupNotFound"
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  FAILEDOPERATION_TAGKEYRESERVED = "FailedOperation.TagKeyReserved"
//  INSTANCESQUOTALIMITEXCEEDED = "InstancesQuotaLimitExceeded"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDCLIENTTOKEN_TOOLONG = "InvalidClientToken.TooLong"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDHOSTID_NOTFOUND = "InvalidHostId.NotFound"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDINSTANCENAME_TOOLONG = "InvalidInstanceName.TooLong"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETER_INVALIDIPFORMAT = "InvalidParameter.InvalidIpFormat"
//  INVALIDPARAMETER_LACKCORECOUNTORTHREADPERCORE = "InvalidParameter.LackCoreCountOrThreadPerCore"
//  INVALIDPARAMETER_PASSWORDNOTSUPPORTED = "InvalidParameter.PasswordNotSupported"
//  INVALIDPARAMETER_SNAPSHOTNOTFOUND = "InvalidParameter.SnapshotNotFound"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLOUDSSDDATADISKSIZETOOSMALL = "InvalidParameterValue.CloudSsdDataDiskSizeTooSmall"
//  INVALIDPARAMETERVALUE_CORECOUNTVALUE = "InvalidParameterValue.CoreCountValue"
//  INVALIDPARAMETERVALUE_ILLEGALHOSTNAME = "InvalidParameterValue.IllegalHostName"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTSUPPORTHPCCLUSTER = "InvalidParameterValue.InstanceTypeNotSupportHpcCluster"
//  INVALIDPARAMETERVALUE_INSTANCETYPEREQUIREDHPCCLUSTER = "InvalidParameterValue.InstanceTypeRequiredHpcCluster"
//  INVALIDPARAMETERVALUE_INSUFFICIENTOFFERING = "InvalidParameterValue.InsufficientOffering"
//  INVALIDPARAMETERVALUE_INSUFFICIENTPRICE = "InvalidParameterValue.InsufficientPrice"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHTEMPLATEDESCRIPTION = "InvalidParameterValue.InvalidLaunchTemplateDescription"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHTEMPLATENAME = "InvalidParameterValue.InvalidLaunchTemplateName"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHTEMPLATEVERSIONDESCRIPTION = "InvalidParameterValue.InvalidLaunchTemplateVersionDescription"
//  INVALIDPARAMETERVALUE_INVALIDUSERDATAFORMAT = "InvalidParameterValue.InvalidUserDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MUSTDHCPENABLEDVPC = "InvalidParameterValue.MustDhcpEnabledVpc"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SUBNETNOTEXIST = "InvalidParameterValue.SubnetNotExist"
//  INVALIDPARAMETERVALUE_THREADPERCOREVALUE = "InvalidParameterValue.ThreadPerCoreValue"
//  INVALIDPARAMETERVALUE_VPCIDZONEIDNOTMATCH = "InvalidParameterValue.VpcIdZoneIdNotMatch"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDPASSWORD = "InvalidPassword"
//  INVALIDPERIOD = "InvalidPeriod"
//  INVALIDPERMISSION = "InvalidPermission"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"
//  LIMITEXCEEDED_LAUNCHTEMPLATEQUOTA = "LimitExceeded.LaunchTemplateQuota"
//  LIMITEXCEEDED_SINGLEUSGQUOTA = "LimitExceeded.SingleUSGQuota"
//  LIMITEXCEEDED_SPOTQUOTA = "LimitExceeded.SpotQuota"
//  LIMITEXCEEDED_USERSPOTQUOTA = "LimitExceeded.UserSpotQuota"
//  LIMITEXCEEDED_VPCSUBNETNUM = "LimitExceeded.VpcSubnetNum"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DPDKINSTANCETYPEREQUIREDVPC = "MissingParameter.DPDKInstanceTypeRequiredVPC"
//  MISSINGPARAMETER_MONITORSERVICE = "MissingParameter.MonitorService"
//  RESOURCEINSUFFICIENT_CLOUDDISKSOLDOUT = "ResourceInsufficient.CloudDiskSoldOut"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCEINSUFFICIENT_DISASTERRECOVERGROUPCVMQUOTA = "ResourceInsufficient.DisasterRecoverGroupCvmQuota"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  RESOURCENOTFOUND_HPCCLUSTER = "ResourceNotFound.HpcCluster"
//  RESOURCENOTFOUND_NODEFAULTCBS = "ResourceNotFound.NoDefaultCbs"
//  RESOURCENOTFOUND_NODEFAULTCBSWITHREASON = "ResourceNotFound.NoDefaultCbsWithReason"
//  RESOURCEUNAVAILABLE_INSTANCETYPE = "ResourceUnavailable.InstanceType"
//  RESOURCESSOLDOUT_EIPINSUFFICIENT = "ResourcesSoldOut.EipInsufficient"
//  RESOURCESSOLDOUT_SPECIFIEDINSTANCETYPE = "ResourcesSoldOut.SpecifiedInstanceType"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDDISK = "UnsupportedOperation.InvalidDisk"
//  UNSUPPORTEDOPERATION_KEYPAIRUNSUPPORTEDWINDOWS = "UnsupportedOperation.KeyPairUnsupportedWindows"
//  UNSUPPORTEDOPERATION_NOINSTANCETYPESUPPORTSPOT = "UnsupportedOperation.NoInstanceTypeSupportSpot"
//  UNSUPPORTEDOPERATION_NOTSUPPORTIMPORTINSTANCESACTIONTIMER = "UnsupportedOperation.NotSupportImportInstancesActionTimer"
//  UNSUPPORTEDOPERATION_ONLYFORPREPAIDACCOUNT = "UnsupportedOperation.OnlyForPrepaidAccount"
//  VPCADDRNOTINSUBNET = "VpcAddrNotInSubNet"
//  VPCIPISUSED = "VpcIpIsUsed"
func (c *Client) CreateLaunchTemplateWithContext(ctx context.Context, request *CreateLaunchTemplateRequest) (response *CreateLaunchTemplateResponse, err error) {
    if request == nil {
        request = NewCreateLaunchTemplateRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateLaunchTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLaunchTemplateVersionRequest() (request *CreateLaunchTemplateVersionRequest) {
    request = &CreateLaunchTemplateVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "CreateLaunchTemplateVersion")
    
    
    return
}

func NewCreateLaunchTemplateVersionResponse() (response *CreateLaunchTemplateVersionResponse) {
    response = &CreateLaunchTemplateVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLaunchTemplateVersion
// This API is used to create an instance launch template based on the specified template ID and the corresponding template version number. The default version number will be used when no template version numbers are specified. Each instance launch template can have up to 30 version numbers.
//
// error code that may be returned:
//  AUTHFAILURE_CAMROLENAMEAUTHENTICATEFAILED = "AuthFailure.CamRoleNameAuthenticateFailed"
//  FAILEDOPERATION_DISASTERRECOVERGROUPNOTFOUND = "FailedOperation.DisasterRecoverGroupNotFound"
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  FAILEDOPERATION_TAGKEYRESERVED = "FailedOperation.TagKeyReserved"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INVALIDCLIENTTOKEN_TOOLONG = "InvalidClientToken.TooLong"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDHOSTID_NOTFOUND = "InvalidHostId.NotFound"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDINSTANCENAME_TOOLONG = "InvalidInstanceName.TooLong"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETER_INVALIDIPFORMAT = "InvalidParameter.InvalidIpFormat"
//  INVALIDPARAMETER_PASSWORDNOTSUPPORTED = "InvalidParameter.PasswordNotSupported"
//  INVALIDPARAMETER_SNAPSHOTNOTFOUND = "InvalidParameter.SnapshotNotFound"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLOUDSSDDATADISKSIZETOOSMALL = "InvalidParameterValue.CloudSsdDataDiskSizeTooSmall"
//  INVALIDPARAMETERVALUE_ILLEGALHOSTNAME = "InvalidParameterValue.IllegalHostName"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTSUPPORTHPCCLUSTER = "InvalidParameterValue.InstanceTypeNotSupportHpcCluster"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHTEMPLATEVERSIONDESCRIPTION = "InvalidParameterValue.InvalidLaunchTemplateVersionDescription"
//  INVALIDPARAMETERVALUE_INVALIDUSERDATAFORMAT = "InvalidParameterValue.InvalidUserDataFormat"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MUSTDHCPENABLEDVPC = "InvalidParameterValue.MustDhcpEnabledVpc"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SUBNETNOTEXIST = "InvalidParameterValue.SubnetNotExist"
//  INVALIDPARAMETERVALUE_THREADPERCOREVALUE = "InvalidParameterValue.ThreadPerCoreValue"
//  INVALIDPARAMETERVALUE_VPCIDZONEIDNOTMATCH = "InvalidParameterValue.VpcIdZoneIdNotMatch"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDPASSWORD = "InvalidPassword"
//  INVALIDPERIOD = "InvalidPeriod"
//  INVALIDPERMISSION = "InvalidPermission"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"
//  LIMITEXCEEDED_LAUNCHTEMPLATEQUOTA = "LimitExceeded.LaunchTemplateQuota"
//  LIMITEXCEEDED_LAUNCHTEMPLATEVERSIONQUOTA = "LimitExceeded.LaunchTemplateVersionQuota"
//  LIMITEXCEEDED_SINGLEUSGQUOTA = "LimitExceeded.SingleUSGQuota"
//  LIMITEXCEEDED_SPOTQUOTA = "LimitExceeded.SpotQuota"
//  LIMITEXCEEDED_USERSPOTQUOTA = "LimitExceeded.UserSpotQuota"
//  LIMITEXCEEDED_VPCSUBNETNUM = "LimitExceeded.VpcSubnetNum"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DPDKINSTANCETYPEREQUIREDVPC = "MissingParameter.DPDKInstanceTypeRequiredVPC"
//  MISSINGPARAMETER_MONITORSERVICE = "MissingParameter.MonitorService"
//  RESOURCEINSUFFICIENT_CLOUDDISKSOLDOUT = "ResourceInsufficient.CloudDiskSoldOut"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCEINSUFFICIENT_DISASTERRECOVERGROUPCVMQUOTA = "ResourceInsufficient.DisasterRecoverGroupCvmQuota"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  RESOURCENOTFOUND_HPCCLUSTER = "ResourceNotFound.HpcCluster"
//  RESOURCENOTFOUND_NODEFAULTCBS = "ResourceNotFound.NoDefaultCbs"
//  RESOURCENOTFOUND_NODEFAULTCBSWITHREASON = "ResourceNotFound.NoDefaultCbsWithReason"
//  RESOURCEUNAVAILABLE_INSTANCETYPE = "ResourceUnavailable.InstanceType"
//  RESOURCESSOLDOUT_EIPINSUFFICIENT = "ResourcesSoldOut.EipInsufficient"
//  RESOURCESSOLDOUT_SPECIFIEDINSTANCETYPE = "ResourcesSoldOut.SpecifiedInstanceType"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDDISK = "UnsupportedOperation.InvalidDisk"
//  UNSUPPORTEDOPERATION_KEYPAIRUNSUPPORTEDWINDOWS = "UnsupportedOperation.KeyPairUnsupportedWindows"
//  UNSUPPORTEDOPERATION_NOINSTANCETYPESUPPORTSPOT = "UnsupportedOperation.NoInstanceTypeSupportSpot"
//  UNSUPPORTEDOPERATION_ONLYFORPREPAIDACCOUNT = "UnsupportedOperation.OnlyForPrepaidAccount"
//  VPCADDRNOTINSUBNET = "VpcAddrNotInSubNet"
//  VPCIPISUSED = "VpcIpIsUsed"
func (c *Client) CreateLaunchTemplateVersion(request *CreateLaunchTemplateVersionRequest) (response *CreateLaunchTemplateVersionResponse, err error) {
    if request == nil {
        request = NewCreateLaunchTemplateVersionRequest()
    }
    
    response = NewCreateLaunchTemplateVersionResponse()
    err = c.Send(request, response)
    return
}

// CreateLaunchTemplateVersion
// This API is used to create an instance launch template based on the specified template ID and the corresponding template version number. The default version number will be used when no template version numbers are specified. Each instance launch template can have up to 30 version numbers.
//
// error code that may be returned:
//  AUTHFAILURE_CAMROLENAMEAUTHENTICATEFAILED = "AuthFailure.CamRoleNameAuthenticateFailed"
//  FAILEDOPERATION_DISASTERRECOVERGROUPNOTFOUND = "FailedOperation.DisasterRecoverGroupNotFound"
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  FAILEDOPERATION_TAGKEYRESERVED = "FailedOperation.TagKeyReserved"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INVALIDCLIENTTOKEN_TOOLONG = "InvalidClientToken.TooLong"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDHOSTID_NOTFOUND = "InvalidHostId.NotFound"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDINSTANCENAME_TOOLONG = "InvalidInstanceName.TooLong"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETER_INVALIDIPFORMAT = "InvalidParameter.InvalidIpFormat"
//  INVALIDPARAMETER_PASSWORDNOTSUPPORTED = "InvalidParameter.PasswordNotSupported"
//  INVALIDPARAMETER_SNAPSHOTNOTFOUND = "InvalidParameter.SnapshotNotFound"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLOUDSSDDATADISKSIZETOOSMALL = "InvalidParameterValue.CloudSsdDataDiskSizeTooSmall"
//  INVALIDPARAMETERVALUE_ILLEGALHOSTNAME = "InvalidParameterValue.IllegalHostName"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTSUPPORTHPCCLUSTER = "InvalidParameterValue.InstanceTypeNotSupportHpcCluster"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHTEMPLATEVERSIONDESCRIPTION = "InvalidParameterValue.InvalidLaunchTemplateVersionDescription"
//  INVALIDPARAMETERVALUE_INVALIDUSERDATAFORMAT = "InvalidParameterValue.InvalidUserDataFormat"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MUSTDHCPENABLEDVPC = "InvalidParameterValue.MustDhcpEnabledVpc"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SUBNETNOTEXIST = "InvalidParameterValue.SubnetNotExist"
//  INVALIDPARAMETERVALUE_THREADPERCOREVALUE = "InvalidParameterValue.ThreadPerCoreValue"
//  INVALIDPARAMETERVALUE_VPCIDZONEIDNOTMATCH = "InvalidParameterValue.VpcIdZoneIdNotMatch"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDPASSWORD = "InvalidPassword"
//  INVALIDPERIOD = "InvalidPeriod"
//  INVALIDPERMISSION = "InvalidPermission"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"
//  LIMITEXCEEDED_LAUNCHTEMPLATEQUOTA = "LimitExceeded.LaunchTemplateQuota"
//  LIMITEXCEEDED_LAUNCHTEMPLATEVERSIONQUOTA = "LimitExceeded.LaunchTemplateVersionQuota"
//  LIMITEXCEEDED_SINGLEUSGQUOTA = "LimitExceeded.SingleUSGQuota"
//  LIMITEXCEEDED_SPOTQUOTA = "LimitExceeded.SpotQuota"
//  LIMITEXCEEDED_USERSPOTQUOTA = "LimitExceeded.UserSpotQuota"
//  LIMITEXCEEDED_VPCSUBNETNUM = "LimitExceeded.VpcSubnetNum"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DPDKINSTANCETYPEREQUIREDVPC = "MissingParameter.DPDKInstanceTypeRequiredVPC"
//  MISSINGPARAMETER_MONITORSERVICE = "MissingParameter.MonitorService"
//  RESOURCEINSUFFICIENT_CLOUDDISKSOLDOUT = "ResourceInsufficient.CloudDiskSoldOut"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCEINSUFFICIENT_DISASTERRECOVERGROUPCVMQUOTA = "ResourceInsufficient.DisasterRecoverGroupCvmQuota"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  RESOURCENOTFOUND_HPCCLUSTER = "ResourceNotFound.HpcCluster"
//  RESOURCENOTFOUND_NODEFAULTCBS = "ResourceNotFound.NoDefaultCbs"
//  RESOURCENOTFOUND_NODEFAULTCBSWITHREASON = "ResourceNotFound.NoDefaultCbsWithReason"
//  RESOURCEUNAVAILABLE_INSTANCETYPE = "ResourceUnavailable.InstanceType"
//  RESOURCESSOLDOUT_EIPINSUFFICIENT = "ResourcesSoldOut.EipInsufficient"
//  RESOURCESSOLDOUT_SPECIFIEDINSTANCETYPE = "ResourcesSoldOut.SpecifiedInstanceType"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDDISK = "UnsupportedOperation.InvalidDisk"
//  UNSUPPORTEDOPERATION_KEYPAIRUNSUPPORTEDWINDOWS = "UnsupportedOperation.KeyPairUnsupportedWindows"
//  UNSUPPORTEDOPERATION_NOINSTANCETYPESUPPORTSPOT = "UnsupportedOperation.NoInstanceTypeSupportSpot"
//  UNSUPPORTEDOPERATION_ONLYFORPREPAIDACCOUNT = "UnsupportedOperation.OnlyForPrepaidAccount"
//  VPCADDRNOTINSUBNET = "VpcAddrNotInSubNet"
//  VPCIPISUSED = "VpcIpIsUsed"
func (c *Client) CreateLaunchTemplateVersionWithContext(ctx context.Context, request *CreateLaunchTemplateVersionRequest) (response *CreateLaunchTemplateVersionResponse, err error) {
    if request == nil {
        request = NewCreateLaunchTemplateVersionRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateLaunchTemplateVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDisasterRecoverGroupsRequest() (request *DeleteDisasterRecoverGroupsRequest) {
    request = &DeleteDisasterRecoverGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DeleteDisasterRecoverGroups")
    
    
    return
}

func NewDeleteDisasterRecoverGroupsResponse() (response *DeleteDisasterRecoverGroupsResponse) {
    response = &DeleteDisasterRecoverGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDisasterRecoverGroups
// This API is used to delete a [spread placement group](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1). Only empty placement groups can be deleted. To delete a non-empty group, you need to terminate all the CVM instances in it first. Otherwise, the deletion will fail.
//
// error code that may be returned:
//  FAILEDOPERATION_PLACEMENTSETNOTEMPTY = "FailedOperation.PlacementSetNotEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCEINSUFFICIENT_DISASTERRECOVERGROUPCVMQUOTA = "ResourceInsufficient.DisasterRecoverGroupCvmQuota"
//  RESOURCENOTFOUND_INVALIDPLACEMENTSET = "ResourceNotFound.InvalidPlacementSet"
func (c *Client) DeleteDisasterRecoverGroups(request *DeleteDisasterRecoverGroupsRequest) (response *DeleteDisasterRecoverGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteDisasterRecoverGroupsRequest()
    }
    
    response = NewDeleteDisasterRecoverGroupsResponse()
    err = c.Send(request, response)
    return
}

// DeleteDisasterRecoverGroups
// This API is used to delete a [spread placement group](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1). Only empty placement groups can be deleted. To delete a non-empty group, you need to terminate all the CVM instances in it first. Otherwise, the deletion will fail.
//
// error code that may be returned:
//  FAILEDOPERATION_PLACEMENTSETNOTEMPTY = "FailedOperation.PlacementSetNotEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCEINSUFFICIENT_DISASTERRECOVERGROUPCVMQUOTA = "ResourceInsufficient.DisasterRecoverGroupCvmQuota"
//  RESOURCENOTFOUND_INVALIDPLACEMENTSET = "ResourceNotFound.InvalidPlacementSet"
func (c *Client) DeleteDisasterRecoverGroupsWithContext(ctx context.Context, request *DeleteDisasterRecoverGroupsRequest) (response *DeleteDisasterRecoverGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteDisasterRecoverGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteDisasterRecoverGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImagesRequest() (request *DeleteImagesRequest) {
    request = &DeleteImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DeleteImages")
    
    
    return
}

func NewDeleteImagesResponse() (response *DeleteImagesResponse) {
    response = &DeleteImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImages
// This API is used to delete images.
//
// 
//
// * If the [ImageState](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#image_state) of an image is `Creating` or `In Use`, it cannot be deleted. Use [DescribeImages](https://intl.cloud.tencent.com/document/api/213/9418?from_cn_redirect=1) to query the image state.
//
// * You can only create up to 10 custom images in each region. If you have used up the quota, you can delete images to create new ones.
//
// * A shared image cannot be deleted.
//
// error code that may be returned:
//  INVALIDIMAGEID_INSHARED = "InvalidImageId.InShared"
//  INVALIDIMAGEID_INCORRECTSTATE = "InvalidImageId.IncorrectState"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
func (c *Client) DeleteImages(request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
    if request == nil {
        request = NewDeleteImagesRequest()
    }
    
    response = NewDeleteImagesResponse()
    err = c.Send(request, response)
    return
}

// DeleteImages
// This API is used to delete images.
//
// 
//
// * If the [ImageState](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#image_state) of an image is `Creating` or `In Use`, it cannot be deleted. Use [DescribeImages](https://intl.cloud.tencent.com/document/api/213/9418?from_cn_redirect=1) to query the image state.
//
// * You can only create up to 10 custom images in each region. If you have used up the quota, you can delete images to create new ones.
//
// * A shared image cannot be deleted.
//
// error code that may be returned:
//  INVALIDIMAGEID_INSHARED = "InvalidImageId.InShared"
//  INVALIDIMAGEID_INCORRECTSTATE = "InvalidImageId.IncorrectState"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageId"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
func (c *Client) DeleteImagesWithContext(ctx context.Context, request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
    if request == nil {
        request = NewDeleteImagesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKeyPairsRequest() (request *DeleteKeyPairsRequest) {
    request = &DeleteKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DeleteKeyPairs")
    
    
    return
}

func NewDeleteKeyPairsResponse() (response *DeleteKeyPairsResponse) {
    response = &DeleteKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteKeyPairs
// This API is used to delete the key pairs hosted in Tencent Cloud.
//
// 
//
// * You can delete multiple key pairs at the same time.
//
// * A key pair used by an instance or image cannot be deleted. Therefore, you need to verify whether all the key pairs have been deleted successfully.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDKEYPAIR_LIMITEXCEEDED = "InvalidKeyPair.LimitExceeded"
//  INVALIDKEYPAIRID_MALFORMED = "InvalidKeyPairId.Malformed"
//  INVALIDKEYPAIRID_NOTFOUND = "InvalidKeyPairId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYPAIRNOTSUPPORTED = "InvalidParameterValue.KeyPairNotSupported"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (response *DeleteKeyPairsResponse, err error) {
    if request == nil {
        request = NewDeleteKeyPairsRequest()
    }
    
    response = NewDeleteKeyPairsResponse()
    err = c.Send(request, response)
    return
}

// DeleteKeyPairs
// This API is used to delete the key pairs hosted in Tencent Cloud.
//
// 
//
// * You can delete multiple key pairs at the same time.
//
// * A key pair used by an instance or image cannot be deleted. Therefore, you need to verify whether all the key pairs have been deleted successfully.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDKEYPAIR_LIMITEXCEEDED = "InvalidKeyPair.LimitExceeded"
//  INVALIDKEYPAIRID_MALFORMED = "InvalidKeyPairId.Malformed"
//  INVALIDKEYPAIRID_NOTFOUND = "InvalidKeyPairId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KEYPAIRNOTSUPPORTED = "InvalidParameterValue.KeyPairNotSupported"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteKeyPairsWithContext(ctx context.Context, request *DeleteKeyPairsRequest) (response *DeleteKeyPairsResponse, err error) {
    if request == nil {
        request = NewDeleteKeyPairsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLaunchTemplateRequest() (request *DeleteLaunchTemplateRequest) {
    request = &DeleteLaunchTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DeleteLaunchTemplate")
    
    
    return
}

func NewDeleteLaunchTemplateResponse() (response *DeleteLaunchTemplateResponse) {
    response = &DeleteLaunchTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLaunchTemplate
// This API is used to delete an instance launch template.
//
// error code that may be returned:
//  ACCOUNTQUALIFICATIONRESTRICTIONS = "AccountQualificationRestrictions"
//  AUTHFAILURE_CAMROLENAMEAUTHENTICATEFAILED = "AuthFailure.CamRoleNameAuthenticateFailed"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
func (c *Client) DeleteLaunchTemplate(request *DeleteLaunchTemplateRequest) (response *DeleteLaunchTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLaunchTemplateRequest()
    }
    
    response = NewDeleteLaunchTemplateResponse()
    err = c.Send(request, response)
    return
}

// DeleteLaunchTemplate
// This API is used to delete an instance launch template.
//
// error code that may be returned:
//  ACCOUNTQUALIFICATIONRESTRICTIONS = "AccountQualificationRestrictions"
//  AUTHFAILURE_CAMROLENAMEAUTHENTICATEFAILED = "AuthFailure.CamRoleNameAuthenticateFailed"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
func (c *Client) DeleteLaunchTemplateWithContext(ctx context.Context, request *DeleteLaunchTemplateRequest) (response *DeleteLaunchTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteLaunchTemplateRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteLaunchTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLaunchTemplateVersionsRequest() (request *DeleteLaunchTemplateVersionsRequest) {
    request = &DeleteLaunchTemplateVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DeleteLaunchTemplateVersions")
    
    
    return
}

func NewDeleteLaunchTemplateVersionsResponse() (response *DeleteLaunchTemplateVersionsResponse) {
    response = &DeleteLaunchTemplateVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLaunchTemplateVersions
// This API is used to delete one or more instance launch template versions.
//
// error code that may be returned:
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEDEFAULTVERSION = "InvalidParameterValue.LaunchTemplateDefaultVersion"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteLaunchTemplateVersions(request *DeleteLaunchTemplateVersionsRequest) (response *DeleteLaunchTemplateVersionsResponse, err error) {
    if request == nil {
        request = NewDeleteLaunchTemplateVersionsRequest()
    }
    
    response = NewDeleteLaunchTemplateVersionsResponse()
    err = c.Send(request, response)
    return
}

// DeleteLaunchTemplateVersions
// This API is used to delete one or more instance launch template versions.
//
// error code that may be returned:
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEDEFAULTVERSION = "InvalidParameterValue.LaunchTemplateDefaultVersion"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteLaunchTemplateVersionsWithContext(ctx context.Context, request *DeleteLaunchTemplateVersionsRequest) (response *DeleteLaunchTemplateVersionsResponse, err error) {
    if request == nil {
        request = NewDeleteLaunchTemplateVersionsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteLaunchTemplateVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisasterRecoverGroupQuotaRequest() (request *DescribeDisasterRecoverGroupQuotaRequest) {
    request = &DescribeDisasterRecoverGroupQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeDisasterRecoverGroupQuota")
    
    
    return
}

func NewDescribeDisasterRecoverGroupQuotaResponse() (response *DescribeDisasterRecoverGroupQuotaResponse) {
    response = &DescribeDisasterRecoverGroupQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDisasterRecoverGroupQuota
// This API is used to query the quota of [spread placement groups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEDEFAULTVERSION = "InvalidParameterValue.LaunchTemplateDefaultVersion"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDisasterRecoverGroupQuota(request *DescribeDisasterRecoverGroupQuotaRequest) (response *DescribeDisasterRecoverGroupQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeDisasterRecoverGroupQuotaRequest()
    }
    
    response = NewDescribeDisasterRecoverGroupQuotaResponse()
    err = c.Send(request, response)
    return
}

// DescribeDisasterRecoverGroupQuota
// This API is used to query the quota of [spread placement groups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEDEFAULTVERSION = "InvalidParameterValue.LaunchTemplateDefaultVersion"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDisasterRecoverGroupQuotaWithContext(ctx context.Context, request *DescribeDisasterRecoverGroupQuotaRequest) (response *DescribeDisasterRecoverGroupQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeDisasterRecoverGroupQuotaRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDisasterRecoverGroupQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisasterRecoverGroupsRequest() (request *DescribeDisasterRecoverGroupsRequest) {
    request = &DescribeDisasterRecoverGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeDisasterRecoverGroups")
    
    
    return
}

func NewDescribeDisasterRecoverGroupsResponse() (response *DescribeDisasterRecoverGroupsResponse) {
    response = &DescribeDisasterRecoverGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDisasterRecoverGroups
// This API is used to query the information on [spread placement groups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DISASTERRECOVERGROUPIDMALFORMED = "InvalidParameterValue.DisasterRecoverGroupIdMalformed"
func (c *Client) DescribeDisasterRecoverGroups(request *DescribeDisasterRecoverGroupsRequest) (response *DescribeDisasterRecoverGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDisasterRecoverGroupsRequest()
    }
    
    response = NewDescribeDisasterRecoverGroupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeDisasterRecoverGroups
// This API is used to query the information on [spread placement groups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DISASTERRECOVERGROUPIDMALFORMED = "InvalidParameterValue.DisasterRecoverGroupIdMalformed"
func (c *Client) DescribeDisasterRecoverGroupsWithContext(ctx context.Context, request *DescribeDisasterRecoverGroupsRequest) (response *DescribeDisasterRecoverGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDisasterRecoverGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDisasterRecoverGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostsRequest() (request *DescribeHostsRequest) {
    request = &DescribeHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeHosts")
    
    
    return
}

func NewDescribeHostsResponse() (response *DescribeHostsResponse) {
    response = &DescribeHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHosts
// This API is used to query the details of CDH instances.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeHosts(request *DescribeHostsRequest) (response *DescribeHostsResponse, err error) {
    if request == nil {
        request = NewDescribeHostsRequest()
    }
    
    response = NewDescribeHostsResponse()
    err = c.Send(request, response)
    return
}

// DescribeHosts
// This API is used to query the details of CDH instances.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeHostsWithContext(ctx context.Context, request *DescribeHostsRequest) (response *DescribeHostsResponse, err error) {
    if request == nil {
        request = NewDescribeHostsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeHostsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageQuotaRequest() (request *DescribeImageQuotaRequest) {
    request = &DescribeImageQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeImageQuota")
    
    
    return
}

func NewDescribeImageQuotaResponse() (response *DescribeImageQuotaResponse) {
    response = &DescribeImageQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageQuota
// This API is used to query the image quota of an user account.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageQuota(request *DescribeImageQuotaRequest) (response *DescribeImageQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeImageQuotaRequest()
    }
    
    response = NewDescribeImageQuotaResponse()
    err = c.Send(request, response)
    return
}

// DescribeImageQuota
// This API is used to query the image quota of an user account.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageQuotaWithContext(ctx context.Context, request *DescribeImageQuotaRequest) (response *DescribeImageQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeImageQuotaRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeImageQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageSharePermissionRequest() (request *DescribeImageSharePermissionRequest) {
    request = &DescribeImageSharePermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeImageSharePermission")
    
    
    return
}

func NewDescribeImageSharePermissionResponse() (response *DescribeImageSharePermissionResponse) {
    response = &DescribeImageSharePermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageSharePermission
// This API is used to query image sharing information.
//
// error code that may be returned:
//  INVALIDACCOUNTID_NOTFOUND = "InvalidAccountId.NotFound"
//  INVALIDACCOUNTIS_YOURSELF = "InvalidAccountIs.YourSelf"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  OVERQUOTA = "OverQuota"
//  UNAUTHORIZEDOPERATION_IMAGENOTBELONGTOACCOUNT = "UnauthorizedOperation.ImageNotBelongToAccount"
func (c *Client) DescribeImageSharePermission(request *DescribeImageSharePermissionRequest) (response *DescribeImageSharePermissionResponse, err error) {
    if request == nil {
        request = NewDescribeImageSharePermissionRequest()
    }
    
    response = NewDescribeImageSharePermissionResponse()
    err = c.Send(request, response)
    return
}

// DescribeImageSharePermission
// This API is used to query image sharing information.
//
// error code that may be returned:
//  INVALIDACCOUNTID_NOTFOUND = "InvalidAccountId.NotFound"
//  INVALIDACCOUNTIS_YOURSELF = "InvalidAccountIs.YourSelf"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  OVERQUOTA = "OverQuota"
//  UNAUTHORIZEDOPERATION_IMAGENOTBELONGTOACCOUNT = "UnauthorizedOperation.ImageNotBelongToAccount"
func (c *Client) DescribeImageSharePermissionWithContext(ctx context.Context, request *DescribeImageSharePermissionRequest) (response *DescribeImageSharePermissionResponse, err error) {
    if request == nil {
        request = NewDescribeImageSharePermissionRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeImageSharePermissionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeImages")
    
    
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImages
// This API is used to view the list of images.
//
// 
//
// * You specify the image ID or set filters to query the details of certain images.
//
// * You can specify `Offset` and `Limit` to select a certain part of the results. By default, the information on the first 20 matching results is returned.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOEXISTIMAGEIDSFILTERS = "InvalidParameter.InvalidParameterCoexistImageIdsFilters"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTFOUND = "InvalidParameterValue.InstanceTypeNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

// DescribeImages
// This API is used to view the list of images.
//
// 
//
// * You specify the image ID or set filters to query the details of certain images.
//
// * You can specify `Offset` and `Limit` to select a certain part of the results. By default, the information on the first 20 matching results is returned.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOEXISTIMAGEIDSFILTERS = "InvalidParameter.InvalidParameterCoexistImageIdsFilters"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTFOUND = "InvalidParameterValue.InstanceTypeNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImportImageOsRequest() (request *DescribeImportImageOsRequest) {
    request = &DescribeImportImageOsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeImportImageOs")
    
    
    return
}

func NewDescribeImportImageOsResponse() (response *DescribeImportImageOsResponse) {
    response = &DescribeImportImageOsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImportImageOs
// This API is used to query the list of supported operating systems of imported images.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOEXISTIMAGEIDSFILTERS = "InvalidParameter.InvalidParameterCoexistImageIdsFilters"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTFOUND = "InvalidParameterValue.InstanceTypeNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeImportImageOs(request *DescribeImportImageOsRequest) (response *DescribeImportImageOsResponse, err error) {
    if request == nil {
        request = NewDescribeImportImageOsRequest()
    }
    
    response = NewDescribeImportImageOsResponse()
    err = c.Send(request, response)
    return
}

// DescribeImportImageOs
// This API is used to query the list of supported operating systems of imported images.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOEXISTIMAGEIDSFILTERS = "InvalidParameter.InvalidParameterCoexistImageIdsFilters"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTFOUND = "InvalidParameterValue.InstanceTypeNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeImportImageOsWithContext(ctx context.Context, request *DescribeImportImageOsRequest) (response *DescribeImportImageOsResponse, err error) {
    if request == nil {
        request = NewDescribeImportImageOsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeImportImageOsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceFamilyConfigsRequest() (request *DescribeInstanceFamilyConfigsRequest) {
    request = &DescribeInstanceFamilyConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceFamilyConfigs")
    
    
    return
}

func NewDescribeInstanceFamilyConfigsResponse() (response *DescribeInstanceFamilyConfigsResponse) {
    response = &DescribeInstanceFamilyConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceFamilyConfigs
// This API is used to query a list of model families available to the current user in the current region.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDREGION_NOTFOUND = "InvalidRegion.NotFound"
func (c *Client) DescribeInstanceFamilyConfigs(request *DescribeInstanceFamilyConfigsRequest) (response *DescribeInstanceFamilyConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceFamilyConfigsRequest()
    }
    
    response = NewDescribeInstanceFamilyConfigsResponse()
    err = c.Send(request, response)
    return
}

// DescribeInstanceFamilyConfigs
// This API is used to query a list of model families available to the current user in the current region.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDREGION_NOTFOUND = "InvalidRegion.NotFound"
func (c *Client) DescribeInstanceFamilyConfigsWithContext(ctx context.Context, request *DescribeInstanceFamilyConfigsRequest) (response *DescribeInstanceFamilyConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceFamilyConfigsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeInstanceFamilyConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTypeConfigsRequest() (request *DescribeInstanceTypeConfigsRequest) {
    request = &DescribeInstanceTypeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceTypeConfigs")
    
    
    return
}

func NewDescribeInstanceTypeConfigsResponse() (response *DescribeInstanceTypeConfigsResponse) {
    response = &DescribeInstanceTypeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceTypeConfigs
// This API is used to query the model configuration of an instance.
//
// 
//
// * You can filter the query results with `zone` or `instance-family`. For more information on filtering conditions, see [`Filter`](https://intl.cloud.tencent.com/document/api/213/15753?from_cn_redirect=1#Filter).
//
// * If no parameter is defined, the model configuration of all the instances in the specified region will be returned.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeInstanceTypeConfigs(request *DescribeInstanceTypeConfigsRequest) (response *DescribeInstanceTypeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypeConfigsRequest()
    }
    
    response = NewDescribeInstanceTypeConfigsResponse()
    err = c.Send(request, response)
    return
}

// DescribeInstanceTypeConfigs
// This API is used to query the model configuration of an instance.
//
// 
//
// * You can filter the query results with `zone` or `instance-family`. For more information on filtering conditions, see [`Filter`](https://intl.cloud.tencent.com/document/api/213/15753?from_cn_redirect=1#Filter).
//
// * If no parameter is defined, the model configuration of all the instances in the specified region will be returned.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeInstanceTypeConfigsWithContext(ctx context.Context, request *DescribeInstanceTypeConfigsRequest) (response *DescribeInstanceTypeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypeConfigsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeInstanceTypeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceVncUrlRequest() (request *DescribeInstanceVncUrlRequest) {
    request = &DescribeInstanceVncUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstanceVncUrl")
    
    
    return
}

func NewDescribeInstanceVncUrlResponse() (response *DescribeInstanceVncUrlResponse) {
    response = &DescribeInstanceVncUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceVncUrl
// This API is used to query the Virtual Network Console (VNC) URL of an instance for its login to the VNC.
//
// 
//
// * It does not support `STOPPED` CVMs.
//
// * A VNC URL is only valid for 15 seconds. If you do not access the URL within 15 seconds, it will become invalid and you have to query a URL again.
//
// * Once the VNC URL is accessed, it will become invalid and you have to query a URL again if needed.
//
// * If the connection is interrupted, you can make up to 30 reconnection attempts per minute.
//
// * After getting the value `InstanceVncUrl`, you need to append `InstanceVncUrl=xxxx` to the end of the link <https://img.qcloud.com/qcloud/app/active_vnc/index.html?>.
//
// 
//
//   - `InstanceVncUrl`: its value will be returned after the API is successfully called.
//
// 
//
//     The final URL is in the following format:
//
// 
//
// ```
//
// https://img.qcloud.com/qcloud/app/active_vnc/index.html?InstanceVncUrl=wss%3A%2F%2Fbjvnc.qcloud.com%3A26789%2Fvnc%3Fs%3DaHpjWnRVMFNhYmxKdDM5MjRHNlVTSVQwajNUSW0wb2tBbmFtREFCTmFrcy8vUUNPMG0wSHZNOUUxRm5PMmUzWmFDcWlOdDJIbUJxSTZDL0RXcHZxYnZZMmRkWWZWcEZia2lyb09XMzdKNmM9
//
// ```
//
// error code that may be returned:
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCESTATE = "InvalidInstanceState"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateEnterServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateExitServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPED = "UnsupportedOperation.InstanceStateStopped"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_SPECIALINSTANCETYPE = "UnsupportedOperation.SpecialInstanceType"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceVncUrlRequest()
    }
    
    response = NewDescribeInstanceVncUrlResponse()
    err = c.Send(request, response)
    return
}

// DescribeInstanceVncUrl
// This API is used to query the Virtual Network Console (VNC) URL of an instance for its login to the VNC.
//
// 
//
// * It does not support `STOPPED` CVMs.
//
// * A VNC URL is only valid for 15 seconds. If you do not access the URL within 15 seconds, it will become invalid and you have to query a URL again.
//
// * Once the VNC URL is accessed, it will become invalid and you have to query a URL again if needed.
//
// * If the connection is interrupted, you can make up to 30 reconnection attempts per minute.
//
// * After getting the value `InstanceVncUrl`, you need to append `InstanceVncUrl=xxxx` to the end of the link <https://img.qcloud.com/qcloud/app/active_vnc/index.html?>.
//
// 
//
//   - `InstanceVncUrl`: its value will be returned after the API is successfully called.
//
// 
//
//     The final URL is in the following format:
//
// 
//
// ```
//
// https://img.qcloud.com/qcloud/app/active_vnc/index.html?InstanceVncUrl=wss%3A%2F%2Fbjvnc.qcloud.com%3A26789%2Fvnc%3Fs%3DaHpjWnRVMFNhYmxKdDM5MjRHNlVTSVQwajNUSW0wb2tBbmFtREFCTmFrcy8vUUNPMG0wSHZNOUUxRm5PMmUzWmFDcWlOdDJIbUJxSTZDL0RXcHZxYnZZMmRkWWZWcEZia2lyb09XMzdKNmM9
//
// ```
//
// error code that may be returned:
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCESTATE = "InvalidInstanceState"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateEnterServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateExitServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPED = "UnsupportedOperation.InstanceStateStopped"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_SPECIALINSTANCETYPE = "UnsupportedOperation.SpecialInstanceType"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) DescribeInstanceVncUrlWithContext(ctx context.Context, request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceVncUrlRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeInstanceVncUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// This API is used to query the details of instances.
//
// 
//
// * You can filter the query results with the instance `ID`, name, or billing method. See `Filter` for more information.
//
// * If no parameter is defined, a certain number of instances under the current account will be returned. The number is specified by `Limit` and is 20 by default.
//
// error code that may be returned:
//  FAILEDOPERATION_ILLEGALTAGKEY = "FailedOperation.IllegalTagKey"
//  FAILEDOPERATION_ILLEGALTAGVALUE = "FailedOperation.IllegalTagValue"
//  FAILEDOPERATION_TAGKEYRESERVED = "FailedOperation.TagKeyReserved"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IPADDRESSMALFORMED = "InvalidParameterValue.IPAddressMalformed"
//  INVALIDPARAMETERVALUE_IPV6ADDRESSMALFORMED = "InvalidParameterValue.IPv6AddressMalformed"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDVAGUENAME = "InvalidParameterValue.InvalidVagueName"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SUBNETIDMALFORMED = "InvalidParameterValue.SubnetIdMalformed"
//  INVALIDPARAMETERVALUE_TAGKEYNOTFOUND = "InvalidParameterValue.TagKeyNotFound"
//  INVALIDPARAMETERVALUE_VPCIDMALFORMED = "InvalidParameterValue.VpcIdMalformed"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDSGID_MALFORMED = "InvalidSgId.Malformed"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  RESOURCENOTFOUND_HPCCLUSTER = "ResourceNotFound.HpcCluster"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribeInstances
// This API is used to query the details of instances.
//
// 
//
// * You can filter the query results with the instance `ID`, name, or billing method. See `Filter` for more information.
//
// * If no parameter is defined, a certain number of instances under the current account will be returned. The number is specified by `Limit` and is 20 by default.
//
// error code that may be returned:
//  FAILEDOPERATION_ILLEGALTAGKEY = "FailedOperation.IllegalTagKey"
//  FAILEDOPERATION_ILLEGALTAGVALUE = "FailedOperation.IllegalTagValue"
//  FAILEDOPERATION_TAGKEYRESERVED = "FailedOperation.TagKeyReserved"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IPADDRESSMALFORMED = "InvalidParameterValue.IPAddressMalformed"
//  INVALIDPARAMETERVALUE_IPV6ADDRESSMALFORMED = "InvalidParameterValue.IPv6AddressMalformed"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDVAGUENAME = "InvalidParameterValue.InvalidVagueName"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SUBNETIDMALFORMED = "InvalidParameterValue.SubnetIdMalformed"
//  INVALIDPARAMETERVALUE_TAGKEYNOTFOUND = "InvalidParameterValue.TagKeyNotFound"
//  INVALIDPARAMETERVALUE_VPCIDMALFORMED = "InvalidParameterValue.VpcIdMalformed"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDSGID_MALFORMED = "InvalidSgId.Malformed"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  RESOURCENOTFOUND_HPCCLUSTER = "ResourceNotFound.HpcCluster"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesOperationLimitRequest() (request *DescribeInstancesOperationLimitRequest) {
    request = &DescribeInstancesOperationLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesOperationLimit")
    
    
    return
}

func NewDescribeInstancesOperationLimitResponse() (response *DescribeInstancesOperationLimitResponse) {
    response = &DescribeInstancesOperationLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesOperationLimit
// This API is used to query limitations on operations on an instance.
//
// 
//
// * Currently you can use this API to query the maximum number of times you can modify the configuration of an instance.
//
// error code that may be returned:
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
func (c *Client) DescribeInstancesOperationLimit(request *DescribeInstancesOperationLimitRequest) (response *DescribeInstancesOperationLimitResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesOperationLimitRequest()
    }
    
    response = NewDescribeInstancesOperationLimitResponse()
    err = c.Send(request, response)
    return
}

// DescribeInstancesOperationLimit
// This API is used to query limitations on operations on an instance.
//
// 
//
// * Currently you can use this API to query the maximum number of times you can modify the configuration of an instance.
//
// error code that may be returned:
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
func (c *Client) DescribeInstancesOperationLimitWithContext(ctx context.Context, request *DescribeInstancesOperationLimitRequest) (response *DescribeInstancesOperationLimitResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesOperationLimitRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeInstancesOperationLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesStatusRequest() (request *DescribeInstancesStatusRequest) {
    request = &DescribeInstancesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInstancesStatus")
    
    
    return
}

func NewDescribeInstancesStatusResponse() (response *DescribeInstancesStatusResponse) {
    response = &DescribeInstancesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesStatus
// This API is used to query the status of instances.
//
// 
//
// * You can query the status of an instance with its `ID`.
//
// * If no parameter is defined, the status of a certain number of instances under the current account will be returned. The number is specified by `Limit` and is 20 by default.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
func (c *Client) DescribeInstancesStatus(request *DescribeInstancesStatusRequest) (response *DescribeInstancesStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesStatusRequest()
    }
    
    response = NewDescribeInstancesStatusResponse()
    err = c.Send(request, response)
    return
}

// DescribeInstancesStatus
// This API is used to query the status of instances.
//
// 
//
// * You can query the status of an instance with its `ID`.
//
// * If no parameter is defined, the status of a certain number of instances under the current account will be returned. The number is specified by `Limit` and is 20 by default.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
func (c *Client) DescribeInstancesStatusWithContext(ctx context.Context, request *DescribeInstancesStatusRequest) (response *DescribeInstancesStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeInstancesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInternetChargeTypeConfigsRequest() (request *DescribeInternetChargeTypeConfigsRequest) {
    request = &DescribeInternetChargeTypeConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeInternetChargeTypeConfigs")
    
    
    return
}

func NewDescribeInternetChargeTypeConfigsResponse() (response *DescribeInternetChargeTypeConfigsResponse) {
    response = &DescribeInternetChargeTypeConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInternetChargeTypeConfigs
// This API is used to query the network billing methods.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
func (c *Client) DescribeInternetChargeTypeConfigs(request *DescribeInternetChargeTypeConfigsRequest) (response *DescribeInternetChargeTypeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInternetChargeTypeConfigsRequest()
    }
    
    response = NewDescribeInternetChargeTypeConfigsResponse()
    err = c.Send(request, response)
    return
}

// DescribeInternetChargeTypeConfigs
// This API is used to query the network billing methods.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
func (c *Client) DescribeInternetChargeTypeConfigsWithContext(ctx context.Context, request *DescribeInternetChargeTypeConfigsRequest) (response *DescribeInternetChargeTypeConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInternetChargeTypeConfigsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeInternetChargeTypeConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeyPairsRequest() (request *DescribeKeyPairsRequest) {
    request = &DescribeKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeKeyPairs")
    
    
    return
}

func NewDescribeKeyPairsResponse() (response *DescribeKeyPairsResponse) {
    response = &DescribeKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKeyPairs
// This API is used to query key pairs.
//
// 
//
// * A key pair is a pair of keys generated by an algorithm in which the public key is available to the public and the private key is available only to the user. You can use this API to query the public key but not the private key.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDKEYPAIR_LIMITEXCEEDED = "InvalidKeyPair.LimitExceeded"
//  INVALIDKEYPAIRID_MALFORMED = "InvalidKeyPairId.Malformed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUELIMIT = "InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUEOFFSET = "InvalidParameterValueOffset"
func (c *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
    if request == nil {
        request = NewDescribeKeyPairsRequest()
    }
    
    response = NewDescribeKeyPairsResponse()
    err = c.Send(request, response)
    return
}

// DescribeKeyPairs
// This API is used to query key pairs.
//
// 
//
// * A key pair is a pair of keys generated by an algorithm in which the public key is available to the public and the private key is available only to the user. You can use this API to query the public key but not the private key.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDKEYPAIR_LIMITEXCEEDED = "InvalidKeyPair.LimitExceeded"
//  INVALIDKEYPAIRID_MALFORMED = "InvalidKeyPairId.Malformed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUELIMIT = "InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUEOFFSET = "InvalidParameterValueOffset"
func (c *Client) DescribeKeyPairsWithContext(ctx context.Context, request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
    if request == nil {
        request = NewDescribeKeyPairsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLaunchTemplateVersionsRequest() (request *DescribeLaunchTemplateVersionsRequest) {
    request = &DescribeLaunchTemplateVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeLaunchTemplateVersions")
    
    
    return
}

func NewDescribeLaunchTemplateVersionsResponse() (response *DescribeLaunchTemplateVersionsResponse) {
    response = &DescribeLaunchTemplateVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLaunchTemplateVersions
// This API is used to query the information of instance launch template versions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLaunchTemplateVersions(request *DescribeLaunchTemplateVersionsRequest) (response *DescribeLaunchTemplateVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeLaunchTemplateVersionsRequest()
    }
    
    response = NewDescribeLaunchTemplateVersionsResponse()
    err = c.Send(request, response)
    return
}

// DescribeLaunchTemplateVersions
// This API is used to query the information of instance launch template versions.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLaunchTemplateVersionsWithContext(ctx context.Context, request *DescribeLaunchTemplateVersionsRequest) (response *DescribeLaunchTemplateVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeLaunchTemplateVersionsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeLaunchTemplateVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLaunchTemplatesRequest() (request *DescribeLaunchTemplatesRequest) {
    request = &DescribeLaunchTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeLaunchTemplates")
    
    
    return
}

func NewDescribeLaunchTemplatesResponse() (response *DescribeLaunchTemplatesResponse) {
    response = &DescribeLaunchTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLaunchTemplates
// This API is used to query one or more instance launch templates.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHTEMPLATENAME = "InvalidParameterValue.InvalidLaunchTemplateName"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLaunchTemplates(request *DescribeLaunchTemplatesRequest) (response *DescribeLaunchTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLaunchTemplatesRequest()
    }
    
    response = NewDescribeLaunchTemplatesResponse()
    err = c.Send(request, response)
    return
}

// DescribeLaunchTemplates
// This API is used to query one or more instance launch templates.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHTEMPLATENAME = "InvalidParameterValue.InvalidLaunchTemplateName"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLaunchTemplatesWithContext(ctx context.Context, request *DescribeLaunchTemplatesRequest) (response *DescribeLaunchTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeLaunchTemplatesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeLaunchTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegions
// (Suspended) This API is used to query the information of regions. Due to platform policy, the update of this API has been temporarily stopped. Please try the new one as described in https://intl.cloud.tencent.com/document/product/1278/55255?from_cn_redirect=1
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHTEMPLATENAME = "InvalidParameterValue.InvalidLaunchTemplateName"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

// DescribeRegions
// (Suspended) This API is used to query the information of regions. Due to platform policy, the update of this API has been temporarily stopped. Please try the new one as described in https://intl.cloud.tencent.com/document/product/1278/55255?from_cn_redirect=1
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDLAUNCHTEMPLATENAME = "InvalidParameterValue.InvalidLaunchTemplateName"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReservedInstancesRequest() (request *DescribeReservedInstancesRequest) {
    request = &DescribeReservedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeReservedInstances")
    
    
    return
}

func NewDescribeReservedInstancesResponse() (response *DescribeReservedInstancesResponse) {
    response = &DescribeReservedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReservedInstances
// This API is used to list reserved instances the user has purchased.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNSUPPORTEDOPERATION_INVALIDPERMISSIONNONINTERNATIONALACCOUNT = "UnsupportedOperation.InvalidPermissionNonInternationalAccount"
//  UNSUPPORTEDOPERATION_RESERVEDINSTANCEINVISIBLEFORUSER = "UnsupportedOperation.ReservedInstanceInvisibleForUser"
func (c *Client) DescribeReservedInstances(request *DescribeReservedInstancesRequest) (response *DescribeReservedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeReservedInstancesRequest()
    }
    
    response = NewDescribeReservedInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribeReservedInstances
// This API is used to list reserved instances the user has purchased.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNSUPPORTEDOPERATION_INVALIDPERMISSIONNONINTERNATIONALACCOUNT = "UnsupportedOperation.InvalidPermissionNonInternationalAccount"
//  UNSUPPORTEDOPERATION_RESERVEDINSTANCEINVISIBLEFORUSER = "UnsupportedOperation.ReservedInstanceInvisibleForUser"
func (c *Client) DescribeReservedInstancesWithContext(ctx context.Context, request *DescribeReservedInstancesRequest) (response *DescribeReservedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeReservedInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeReservedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReservedInstancesConfigInfosRequest() (request *DescribeReservedInstancesConfigInfosRequest) {
    request = &DescribeReservedInstancesConfigInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeReservedInstancesConfigInfos")
    
    
    return
}

func NewDescribeReservedInstancesConfigInfosResponse() (response *DescribeReservedInstancesConfigInfosResponse) {
    response = &DescribeReservedInstancesConfigInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReservedInstancesConfigInfos
// This API is used to describe reserved instance (RI) offerings. Currently, RIs are only offered to beta users.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNSUPPORTEDOPERATION_INVALIDPERMISSIONNONINTERNATIONALACCOUNT = "UnsupportedOperation.InvalidPermissionNonInternationalAccount"
//  UNSUPPORTEDOPERATION_RESERVEDINSTANCEINVISIBLEFORUSER = "UnsupportedOperation.ReservedInstanceInvisibleForUser"
func (c *Client) DescribeReservedInstancesConfigInfos(request *DescribeReservedInstancesConfigInfosRequest) (response *DescribeReservedInstancesConfigInfosResponse, err error) {
    if request == nil {
        request = NewDescribeReservedInstancesConfigInfosRequest()
    }
    
    response = NewDescribeReservedInstancesConfigInfosResponse()
    err = c.Send(request, response)
    return
}

// DescribeReservedInstancesConfigInfos
// This API is used to describe reserved instance (RI) offerings. Currently, RIs are only offered to beta users.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNSUPPORTEDOPERATION_INVALIDPERMISSIONNONINTERNATIONALACCOUNT = "UnsupportedOperation.InvalidPermissionNonInternationalAccount"
//  UNSUPPORTEDOPERATION_RESERVEDINSTANCEINVISIBLEFORUSER = "UnsupportedOperation.ReservedInstanceInvisibleForUser"
func (c *Client) DescribeReservedInstancesConfigInfosWithContext(ctx context.Context, request *DescribeReservedInstancesConfigInfosRequest) (response *DescribeReservedInstancesConfigInfosResponse, err error) {
    if request == nil {
        request = NewDescribeReservedInstancesConfigInfosRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeReservedInstancesConfigInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReservedInstancesOfferingsRequest() (request *DescribeReservedInstancesOfferingsRequest) {
    request = &DescribeReservedInstancesOfferingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeReservedInstancesOfferings")
    
    
    return
}

func NewDescribeReservedInstancesOfferingsResponse() (response *DescribeReservedInstancesOfferingsResponse) {
    response = &DescribeReservedInstancesOfferingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReservedInstancesOfferings
// This API is used to describe Reserved Instance offerings that are available for purchase.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNSUPPORTEDOPERATION_INVALIDPERMISSIONNONINTERNATIONALACCOUNT = "UnsupportedOperation.InvalidPermissionNonInternationalAccount"
//  UNSUPPORTEDOPERATION_RESERVEDINSTANCEINVISIBLEFORUSER = "UnsupportedOperation.ReservedInstanceInvisibleForUser"
func (c *Client) DescribeReservedInstancesOfferings(request *DescribeReservedInstancesOfferingsRequest) (response *DescribeReservedInstancesOfferingsResponse, err error) {
    if request == nil {
        request = NewDescribeReservedInstancesOfferingsRequest()
    }
    
    response = NewDescribeReservedInstancesOfferingsResponse()
    err = c.Send(request, response)
    return
}

// DescribeReservedInstancesOfferings
// This API is used to describe Reserved Instance offerings that are available for purchase.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  UNSUPPORTEDOPERATION_INVALIDPERMISSIONNONINTERNATIONALACCOUNT = "UnsupportedOperation.InvalidPermissionNonInternationalAccount"
//  UNSUPPORTEDOPERATION_RESERVEDINSTANCEINVISIBLEFORUSER = "UnsupportedOperation.ReservedInstanceInvisibleForUser"
func (c *Client) DescribeReservedInstancesOfferingsWithContext(ctx context.Context, request *DescribeReservedInstancesOfferingsRequest) (response *DescribeReservedInstancesOfferingsResponse, err error) {
    if request == nil {
        request = NewDescribeReservedInstancesOfferingsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeReservedInstancesOfferingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneInstanceConfigInfosRequest() (request *DescribeZoneInstanceConfigInfosRequest) {
    request = &DescribeZoneInstanceConfigInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeZoneInstanceConfigInfos")
    
    
    return
}

func NewDescribeZoneInstanceConfigInfosResponse() (response *DescribeZoneInstanceConfigInfosResponse) {
    response = &DescribeZoneInstanceConfigInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneInstanceConfigInfos
// This API is used to query the configurations of models in an availability zone.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDREGION_NOTFOUND = "InvalidRegion.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  RESOURCEINSUFFICIENT_AVAILABILITYZONESOLDOUT = "ResourceInsufficient.AvailabilityZoneSoldOut"
func (c *Client) DescribeZoneInstanceConfigInfos(request *DescribeZoneInstanceConfigInfosRequest) (response *DescribeZoneInstanceConfigInfosResponse, err error) {
    if request == nil {
        request = NewDescribeZoneInstanceConfigInfosRequest()
    }
    
    response = NewDescribeZoneInstanceConfigInfosResponse()
    err = c.Send(request, response)
    return
}

// DescribeZoneInstanceConfigInfos
// This API is used to query the configurations of models in an availability zone.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDREGION_NOTFOUND = "InvalidRegion.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  RESOURCEINSUFFICIENT_AVAILABILITYZONESOLDOUT = "ResourceInsufficient.AvailabilityZoneSoldOut"
func (c *Client) DescribeZoneInstanceConfigInfosWithContext(ctx context.Context, request *DescribeZoneInstanceConfigInfosRequest) (response *DescribeZoneInstanceConfigInfosResponse, err error) {
    if request == nil {
        request = NewDescribeZoneInstanceConfigInfosRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeZoneInstanceConfigInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZones
// This API is used to query availability zones.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDREGION_NOTFOUND = "InvalidRegion.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  RESOURCEINSUFFICIENT_AVAILABILITYZONESOLDOUT = "ResourceInsufficient.AvailabilityZoneSoldOut"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

// DescribeZones
// This API is used to query availability zones.
//
// error code that may be returned:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDREGION_NOTFOUND = "InvalidRegion.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  RESOURCEINSUFFICIENT_AVAILABILITYZONESOLDOUT = "ResourceInsufficient.AvailabilityZoneSoldOut"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateInstancesKeyPairsRequest() (request *DisassociateInstancesKeyPairsRequest) {
    request = &DisassociateInstancesKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DisassociateInstancesKeyPairs")
    
    
    return
}

func NewDisassociateInstancesKeyPairsResponse() (response *DisassociateInstancesKeyPairsResponse) {
    response = &DisassociateInstancesKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateInstancesKeyPairs
// This API is used to unbind one or more key pairs from one or more instances.
//
// 
//
// * It only supports [`STOPPED`](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#InstanceStatus) Linux instances.
//
// * After a key pair is disassociated from an instance, you can log in to the instance with password.
//
// * If you did not set a password for the instance, you will not be able to log in via SSH after the unbinding. In this case, you can call [ResetInstancesPassword](https://intl.cloud.tencent.com/document/api/213/15736?from_cn_redirect=1) to set a login password.
//
// * Batch operations are supported. The maximum number of instances in each request is 100. If instances not available for the operation are selected, you will get an error code.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDKEYPAIRID_MALFORMED = "InvalidKeyPairId.Malformed"
//  INVALIDKEYPAIRID_NOTFOUND = "InvalidKeyPairId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_INSTANCEOSWINDOWS = "UnsupportedOperation.InstanceOsWindows"
//  UNSUPPORTEDOPERATION_INSTANCESTATEBANNING = "UnsupportedOperation.InstanceStateBanning"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) DisassociateInstancesKeyPairs(request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewDisassociateInstancesKeyPairsRequest()
    }
    
    response = NewDisassociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

// DisassociateInstancesKeyPairs
// This API is used to unbind one or more key pairs from one or more instances.
//
// 
//
// * It only supports [`STOPPED`](https://intl.cloud.tencent.com/document/product/213/15753?from_cn_redirect=1#InstanceStatus) Linux instances.
//
// * After a key pair is disassociated from an instance, you can log in to the instance with password.
//
// * If you did not set a password for the instance, you will not be able to log in via SSH after the unbinding. In this case, you can call [ResetInstancesPassword](https://intl.cloud.tencent.com/document/api/213/15736?from_cn_redirect=1) to set a login password.
//
// * Batch operations are supported. The maximum number of instances in each request is 100. If instances not available for the operation are selected, you will get an error code.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDKEYPAIRID_MALFORMED = "InvalidKeyPairId.Malformed"
//  INVALIDKEYPAIRID_NOTFOUND = "InvalidKeyPairId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_INSTANCEOSWINDOWS = "UnsupportedOperation.InstanceOsWindows"
//  UNSUPPORTEDOPERATION_INSTANCESTATEBANNING = "UnsupportedOperation.InstanceStateBanning"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) DisassociateInstancesKeyPairsWithContext(ctx context.Context, request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewDisassociateInstancesKeyPairsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisassociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "DisassociateSecurityGroups")
    
    
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateSecurityGroups
// This API is used to disassociate security groups from instances.
//
// error code that may be returned:
//  FAILEDOPERATION_SECURITYGROUPACTIONFAILED = "FailedOperation.SecurityGroupActionFailed"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDSGID_MALFORMED = "InvalidSgId.Malformed"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  SECGROUPACTIONFAILURE = "SecGroupActionFailure"
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

// DisassociateSecurityGroups
// This API is used to disassociate security groups from instances.
//
// error code that may be returned:
//  FAILEDOPERATION_SECURITYGROUPACTIONFAILED = "FailedOperation.SecurityGroupActionFailed"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDSGID_MALFORMED = "InvalidSgId.Malformed"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  SECGROUPACTIONFAILURE = "SecGroupActionFailure"
func (c *Client) DisassociateSecurityGroupsWithContext(ctx context.Context, request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewImportImageRequest() (request *ImportImageRequest) {
    request = &ImportImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ImportImage")
    
    
    return
}

func NewImportImageResponse() (response *ImportImageResponse) {
    response = &ImportImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportImage
// The API is used to import an image. The image imported can be used to create instances. Currently, this API can import images in formats like RAW, VHD, QCOW2, and VMDK.
//
// error code that may be returned:
//  IMAGEQUOTALIMITEXCEEDED = "ImageQuotaLimitExceeded"
//  INVALIDIMAGENAME_DUPLICATE = "InvalidImageName.Duplicate"
//  INVALIDIMAGEOSTYPE_UNSUPPORTED = "InvalidImageOsType.Unsupported"
//  INVALIDIMAGEOSVERSION_UNSUPPORTED = "InvalidImageOsVersion.Unsupported"
//  INVALIDPARAMETER_INVALIDPARAMETERURLERROR = "InvalidParameter.InvalidParameterUrlError"
//  INVALIDPARAMETERVALUE_TAGKEYNOTFOUND = "InvalidParameterValue.TagKeyNotFound"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  REGIONABILITYLIMIT_UNSUPPORTEDTOIMPORTIMAGE = "RegionAbilityLimit.UnsupportedToImportImage"
func (c *Client) ImportImage(request *ImportImageRequest) (response *ImportImageResponse, err error) {
    if request == nil {
        request = NewImportImageRequest()
    }
    
    response = NewImportImageResponse()
    err = c.Send(request, response)
    return
}

// ImportImage
// The API is used to import an image. The image imported can be used to create instances. Currently, this API can import images in formats like RAW, VHD, QCOW2, and VMDK.
//
// error code that may be returned:
//  IMAGEQUOTALIMITEXCEEDED = "ImageQuotaLimitExceeded"
//  INVALIDIMAGENAME_DUPLICATE = "InvalidImageName.Duplicate"
//  INVALIDIMAGEOSTYPE_UNSUPPORTED = "InvalidImageOsType.Unsupported"
//  INVALIDIMAGEOSVERSION_UNSUPPORTED = "InvalidImageOsVersion.Unsupported"
//  INVALIDPARAMETER_INVALIDPARAMETERURLERROR = "InvalidParameter.InvalidParameterUrlError"
//  INVALIDPARAMETERVALUE_TAGKEYNOTFOUND = "InvalidParameterValue.TagKeyNotFound"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
//  REGIONABILITYLIMIT_UNSUPPORTEDTOIMPORTIMAGE = "RegionAbilityLimit.UnsupportedToImportImage"
func (c *Client) ImportImageWithContext(ctx context.Context, request *ImportImageRequest) (response *ImportImageResponse, err error) {
    if request == nil {
        request = NewImportImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewImportImageResponse()
    err = c.Send(request, response)
    return
}

func NewImportKeyPairRequest() (request *ImportKeyPairRequest) {
    request = &ImportKeyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ImportKeyPair")
    
    
    return
}

func NewImportKeyPairResponse() (response *ImportKeyPairResponse) {
    response = &ImportKeyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportKeyPair
// This API is used to import key pairs.
//
// 
//
// * You can use this API to import key pairs to a user account, but the key pairs will not be automatically associated with any instance. You may use [AssociasteInstancesKeyPair](https://intl.cloud.tencent.com/document/api/213/9404?from_cn_redirect=1) to associate key pairs with instances.
//
// * You need to specify the names of the key pairs and the content of the public keys.
//
// * If you only have private keys, you can convert them to public keys with the `SSL` tool before importing them.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDKEYPAIR_LIMITEXCEEDED = "InvalidKeyPair.LimitExceeded"
//  INVALIDKEYPAIRNAME_DUPLICATE = "InvalidKeyPairName.Duplicate"
//  INVALIDKEYPAIRNAMEEMPTY = "InvalidKeyPairNameEmpty"
//  INVALIDKEYPAIRNAMEINCLUDEILLEGALCHAR = "InvalidKeyPairNameIncludeIllegalChar"
//  INVALIDKEYPAIRNAMETOOLONG = "InvalidKeyPairNameTooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  INVALIDPUBLICKEY_DUPLICATE = "InvalidPublicKey.Duplicate"
//  INVALIDPUBLICKEY_MALFORMED = "InvalidPublicKey.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ImportKeyPair(request *ImportKeyPairRequest) (response *ImportKeyPairResponse, err error) {
    if request == nil {
        request = NewImportKeyPairRequest()
    }
    
    response = NewImportKeyPairResponse()
    err = c.Send(request, response)
    return
}

// ImportKeyPair
// This API is used to import key pairs.
//
// 
//
// * You can use this API to import key pairs to a user account, but the key pairs will not be automatically associated with any instance. You may use [AssociasteInstancesKeyPair](https://intl.cloud.tencent.com/document/api/213/9404?from_cn_redirect=1) to associate key pairs with instances.
//
// * You need to specify the names of the key pairs and the content of the public keys.
//
// * If you only have private keys, you can convert them to public keys with the `SSL` tool before importing them.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDKEYPAIR_LIMITEXCEEDED = "InvalidKeyPair.LimitExceeded"
//  INVALIDKEYPAIRNAME_DUPLICATE = "InvalidKeyPairName.Duplicate"
//  INVALIDKEYPAIRNAMEEMPTY = "InvalidKeyPairNameEmpty"
//  INVALIDKEYPAIRNAMEINCLUDEILLEGALCHAR = "InvalidKeyPairNameIncludeIllegalChar"
//  INVALIDKEYPAIRNAMETOOLONG = "InvalidKeyPairNameTooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  INVALIDPUBLICKEY_DUPLICATE = "InvalidPublicKey.Duplicate"
//  INVALIDPUBLICKEY_MALFORMED = "InvalidPublicKey.Malformed"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ImportKeyPairWithContext(ctx context.Context, request *ImportKeyPairRequest) (response *ImportKeyPairResponse, err error) {
    if request == nil {
        request = NewImportKeyPairRequest()
    }
    request.SetContext(ctx)
    
    response = NewImportKeyPairResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePricePurchaseReservedInstancesOfferingRequest() (request *InquirePricePurchaseReservedInstancesOfferingRequest) {
    request = &InquirePricePurchaseReservedInstancesOfferingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "InquirePricePurchaseReservedInstancesOffering")
    
    
    return
}

func NewInquirePricePurchaseReservedInstancesOfferingResponse() (response *InquirePricePurchaseReservedInstancesOfferingResponse) {
    response = &InquirePricePurchaseReservedInstancesOfferingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePricePurchaseReservedInstancesOffering
// This API is used to query the price of reserved instances. It only supports querying purchasable reserved instance offerings. Currently, RIs are only offered to beta users.
//
// error code that may be returned:
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_INVALIDPERMISSIONNONINTERNATIONALACCOUNT = "UnsupportedOperation.InvalidPermissionNonInternationalAccount"
func (c *Client) InquirePricePurchaseReservedInstancesOffering(request *InquirePricePurchaseReservedInstancesOfferingRequest) (response *InquirePricePurchaseReservedInstancesOfferingResponse, err error) {
    if request == nil {
        request = NewInquirePricePurchaseReservedInstancesOfferingRequest()
    }
    
    response = NewInquirePricePurchaseReservedInstancesOfferingResponse()
    err = c.Send(request, response)
    return
}

// InquirePricePurchaseReservedInstancesOffering
// This API is used to query the price of reserved instances. It only supports querying purchasable reserved instance offerings. Currently, RIs are only offered to beta users.
//
// error code that may be returned:
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_INVALIDPERMISSIONNONINTERNATIONALACCOUNT = "UnsupportedOperation.InvalidPermissionNonInternationalAccount"
func (c *Client) InquirePricePurchaseReservedInstancesOfferingWithContext(ctx context.Context, request *InquirePricePurchaseReservedInstancesOfferingRequest) (response *InquirePricePurchaseReservedInstancesOfferingResponse, err error) {
    if request == nil {
        request = NewInquirePricePurchaseReservedInstancesOfferingRequest()
    }
    request.SetContext(ctx)
    
    response = NewInquirePricePurchaseReservedInstancesOfferingResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceResetInstanceRequest() (request *InquiryPriceResetInstanceRequest) {
    request = &InquiryPriceResetInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceResetInstance")
    
    
    return
}

func NewInquiryPriceResetInstanceResponse() (response *InquiryPriceResetInstanceResponse) {
    response = &InquiryPriceResetInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceResetInstance
// This API is used to query the price for reinstalling an instance.
//
// 
//
// * If you have specified the `ImageId` parameter, the price query is performed with the specified image. Otherwise, the image used by the current instance is used.
//
// * You can only query the price for reinstallation caused by switching between Linux and Windows OS. And the [system disk type](https://intl.cloud.tencent.com/document/api/213/15753?from_cn_redirect=1#SystemDisk) of the instance must be `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`.
//
// * Currently, this API only supports instances in Mainland China regions.
//
// error code that may be returned:
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEIDFORRETSETINSTANCE = "InvalidParameterValue.InvalidImageIdForRetsetInstance"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEOSNAME = "InvalidParameterValue.InvalidImageOsName"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MONITORSERVICE = "MissingParameter.MonitorService"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  UNSUPPORTEDOPERATION_RAWLOCALDISKINSREINSTALLTOQCOW2 = "UnsupportedOperation.RawLocalDiskInsReinstalltoQcow2"
func (c *Client) InquiryPriceResetInstance(request *InquiryPriceResetInstanceRequest) (response *InquiryPriceResetInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResetInstanceRequest()
    }
    
    response = NewInquiryPriceResetInstanceResponse()
    err = c.Send(request, response)
    return
}

// InquiryPriceResetInstance
// This API is used to query the price for reinstalling an instance.
//
// 
//
// * If you have specified the `ImageId` parameter, the price query is performed with the specified image. Otherwise, the image used by the current instance is used.
//
// * You can only query the price for reinstallation caused by switching between Linux and Windows OS. And the [system disk type](https://intl.cloud.tencent.com/document/api/213/15753?from_cn_redirect=1#SystemDisk) of the instance must be `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`.
//
// * Currently, this API only supports instances in Mainland China regions.
//
// error code that may be returned:
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEIDFORRETSETINSTANCE = "InvalidParameterValue.InvalidImageIdForRetsetInstance"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEOSNAME = "InvalidParameterValue.InvalidImageOsName"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MONITORSERVICE = "MissingParameter.MonitorService"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  UNSUPPORTEDOPERATION_RAWLOCALDISKINSREINSTALLTOQCOW2 = "UnsupportedOperation.RawLocalDiskInsReinstalltoQcow2"
func (c *Client) InquiryPriceResetInstanceWithContext(ctx context.Context, request *InquiryPriceResetInstanceRequest) (response *InquiryPriceResetInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResetInstanceRequest()
    }
    request.SetContext(ctx)
    
    response = NewInquiryPriceResetInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceResetInstancesInternetMaxBandwidthRequest() (request *InquiryPriceResetInstancesInternetMaxBandwidthRequest) {
    request = &InquiryPriceResetInstancesInternetMaxBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceResetInstancesInternetMaxBandwidth")
    
    
    return
}

func NewInquiryPriceResetInstancesInternetMaxBandwidthResponse() (response *InquiryPriceResetInstancesInternetMaxBandwidthResponse) {
    response = &InquiryPriceResetInstancesInternetMaxBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceResetInstancesInternetMaxBandwidth
// This API is used to query the price for upgrading the public bandwidth cap of an instance.
//
// 
//
// * The allowed bandwidth cap varies for different models. For details, see [Purchasing Network Bandwidth](https://intl.cloud.tencent.com/document/product/213/509?from_cn_redirect=1).
//
// * For bandwidth billed by the `TRAFFIC_POSTPAID_BY_HOUR` method, changing the bandwidth cap through this API takes effect in real time. You can increase or reduce bandwidth within applicable limits.
//
// error code that may be returned:
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  FAILEDOPERATION_NOTFOUNDEIP = "FailedOperation.NotFoundEIP"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPERMISSION = "InvalidPermission"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) InquiryPriceResetInstancesInternetMaxBandwidth(request *InquiryPriceResetInstancesInternetMaxBandwidthRequest) (response *InquiryPriceResetInstancesInternetMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResetInstancesInternetMaxBandwidthRequest()
    }
    
    response = NewInquiryPriceResetInstancesInternetMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

// InquiryPriceResetInstancesInternetMaxBandwidth
// This API is used to query the price for upgrading the public bandwidth cap of an instance.
//
// 
//
// * The allowed bandwidth cap varies for different models. For details, see [Purchasing Network Bandwidth](https://intl.cloud.tencent.com/document/product/213/509?from_cn_redirect=1).
//
// * For bandwidth billed by the `TRAFFIC_POSTPAID_BY_HOUR` method, changing the bandwidth cap through this API takes effect in real time. You can increase or reduce bandwidth within applicable limits.
//
// error code that may be returned:
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  FAILEDOPERATION_NOTFOUNDEIP = "FailedOperation.NotFoundEIP"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPERMISSION = "InvalidPermission"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) InquiryPriceResetInstancesInternetMaxBandwidthWithContext(ctx context.Context, request *InquiryPriceResetInstancesInternetMaxBandwidthRequest) (response *InquiryPriceResetInstancesInternetMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResetInstancesInternetMaxBandwidthRequest()
    }
    request.SetContext(ctx)
    
    response = NewInquiryPriceResetInstancesInternetMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceResetInstancesTypeRequest() (request *InquiryPriceResetInstancesTypeRequest) {
    request = &InquiryPriceResetInstancesTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceResetInstancesType")
    
    
    return
}

func NewInquiryPriceResetInstancesTypeResponse() (response *InquiryPriceResetInstancesTypeResponse) {
    response = &InquiryPriceResetInstancesTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceResetInstancesType
// This API is used to query the price for adjusting the instance model.
//
// 
//
// * Currently, you can only use this API to query the prices of instances whose [system disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`.
//
// * Currently, you cannot use this API to query the prices of [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances.
//
// error code that may be returned:
//  FAILEDOPERATION_INQUIRYREFUNDPRICEFAILED = "FailedOperation.InquiryRefundPriceFailed"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GPUINSTANCEFAMILY = "InvalidParameterValue.GPUInstanceFamily"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPERMISSION = "InvalidPermission"
//  LIMITEXCEEDED_ENINUMLIMIT = "LimitExceeded.EniNumLimit"
//  LIMITEXCEEDED_INSTANCETYPEBANDWIDTH = "LimitExceeded.InstanceTypeBandwidth"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCEUNAVAILABLE_INSTANCETYPE = "ResourceUnavailable.InstanceType"
//  UNSUPPORTEDOPERATION_HETEROGENEOUSCHANGEINSTANCEFAMILY = "UnsupportedOperation.HeterogeneousChangeInstanceFamily"
//  UNSUPPORTEDOPERATION_LOCALDATADISKCHANGEINSTANCEFAMILY = "UnsupportedOperation.LocalDataDiskChangeInstanceFamily"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDCHANGEINSTANCEFAMILYTOSA3 = "UnsupportedOperation.UnsupportedChangeInstanceFamilyToSA3"
func (c *Client) InquiryPriceResetInstancesType(request *InquiryPriceResetInstancesTypeRequest) (response *InquiryPriceResetInstancesTypeResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResetInstancesTypeRequest()
    }
    
    response = NewInquiryPriceResetInstancesTypeResponse()
    err = c.Send(request, response)
    return
}

// InquiryPriceResetInstancesType
// This API is used to query the price for adjusting the instance model.
//
// 
//
// * Currently, you can only use this API to query the prices of instances whose [system disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`.
//
// * Currently, you cannot use this API to query the prices of [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances.
//
// error code that may be returned:
//  FAILEDOPERATION_INQUIRYREFUNDPRICEFAILED = "FailedOperation.InquiryRefundPriceFailed"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GPUINSTANCEFAMILY = "InvalidParameterValue.GPUInstanceFamily"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPERMISSION = "InvalidPermission"
//  LIMITEXCEEDED_ENINUMLIMIT = "LimitExceeded.EniNumLimit"
//  LIMITEXCEEDED_INSTANCETYPEBANDWIDTH = "LimitExceeded.InstanceTypeBandwidth"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCEUNAVAILABLE_INSTANCETYPE = "ResourceUnavailable.InstanceType"
//  UNSUPPORTEDOPERATION_HETEROGENEOUSCHANGEINSTANCEFAMILY = "UnsupportedOperation.HeterogeneousChangeInstanceFamily"
//  UNSUPPORTEDOPERATION_LOCALDATADISKCHANGEINSTANCEFAMILY = "UnsupportedOperation.LocalDataDiskChangeInstanceFamily"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDCHANGEINSTANCEFAMILYTOSA3 = "UnsupportedOperation.UnsupportedChangeInstanceFamilyToSA3"
func (c *Client) InquiryPriceResetInstancesTypeWithContext(ctx context.Context, request *InquiryPriceResetInstancesTypeRequest) (response *InquiryPriceResetInstancesTypeResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResetInstancesTypeRequest()
    }
    request.SetContext(ctx)
    
    response = NewInquiryPriceResetInstancesTypeResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceResizeInstanceDisksRequest() (request *InquiryPriceResizeInstanceDisksRequest) {
    request = &InquiryPriceResizeInstanceDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceResizeInstanceDisks")
    
    
    return
}

func NewInquiryPriceResizeInstanceDisksResponse() (response *InquiryPriceResizeInstanceDisksResponse) {
    response = &InquiryPriceResizeInstanceDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceResizeInstanceDisks
// This API is used to query the price for expanding data disks of an instance.
//
// 
//
// * Currently, you can only use this API to query the price of non-elastic data disks whose [disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`. You can use [`DescribeDisks`](https://intl.cloud.tencent.com/document/api/362/16315?from_cn_redirect=1) to check whether a disk is elastic. If the `Portable` field in the response is `false`, it means that the disk is non-elastic.
//
// * Currently, you cannot use this API to query the price for [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances. *Also, you can only query the price of expanding one data disk at a time.
//
// error code that may be returned:
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ATLEASTONE = "MissingParameter.AtLeastOne"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  UNSUPPORTEDOPERATION_INSTANCESTATECORRUPTED = "UnsupportedOperation.InstanceStateCorrupted"
//  UNSUPPORTEDOPERATION_INVALIDDISK = "UnsupportedOperation.InvalidDisk"
//  UNSUPPORTEDOPERATION_LOCALDISKMIGRATINGTOCLOUDDISK = "UnsupportedOperation.LocalDiskMigratingToCloudDisk"
func (c *Client) InquiryPriceResizeInstanceDisks(request *InquiryPriceResizeInstanceDisksRequest) (response *InquiryPriceResizeInstanceDisksResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResizeInstanceDisksRequest()
    }
    
    response = NewInquiryPriceResizeInstanceDisksResponse()
    err = c.Send(request, response)
    return
}

// InquiryPriceResizeInstanceDisks
// This API is used to query the price for expanding data disks of an instance.
//
// 
//
// * Currently, you can only use this API to query the price of non-elastic data disks whose [disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`. You can use [`DescribeDisks`](https://intl.cloud.tencent.com/document/api/362/16315?from_cn_redirect=1) to check whether a disk is elastic. If the `Portable` field in the response is `false`, it means that the disk is non-elastic.
//
// * Currently, you cannot use this API to query the price for [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances. *Also, you can only query the price of expanding one data disk at a time.
//
// error code that may be returned:
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ATLEASTONE = "MissingParameter.AtLeastOne"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  UNSUPPORTEDOPERATION_INSTANCESTATECORRUPTED = "UnsupportedOperation.InstanceStateCorrupted"
//  UNSUPPORTEDOPERATION_INVALIDDISK = "UnsupportedOperation.InvalidDisk"
//  UNSUPPORTEDOPERATION_LOCALDISKMIGRATINGTOCLOUDDISK = "UnsupportedOperation.LocalDiskMigratingToCloudDisk"
func (c *Client) InquiryPriceResizeInstanceDisksWithContext(ctx context.Context, request *InquiryPriceResizeInstanceDisksRequest) (response *InquiryPriceResizeInstanceDisksResponse, err error) {
    if request == nil {
        request = NewInquiryPriceResizeInstanceDisksRequest()
    }
    request.SetContext(ctx)
    
    response = NewInquiryPriceResizeInstanceDisksResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRunInstancesRequest() (request *InquiryPriceRunInstancesRequest) {
    request = &InquiryPriceRunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "InquiryPriceRunInstances")
    
    
    return
}

func NewInquiryPriceRunInstancesResponse() (response *InquiryPriceRunInstancesResponse) {
    response = &InquiryPriceRunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceRunInstances
// This API is used to query the price of creating instances. You can only use this API for instances whose configuration is within the purchase limit. For more information, see [RunInstances](https://intl.cloud.tencent.com/document/api/213/15730?from_cn_redirect=1).
//
// error code that may be returned:
//  ACCOUNTQUALIFICATIONRESTRICTIONS = "AccountQualificationRestrictions"
//  AUTHFAILURE_CAMROLENAMEAUTHENTICATEFAILED = "AuthFailure.CamRoleNameAuthenticateFailed"
//  FAILEDOPERATION_DISASTERRECOVERGROUPNOTFOUND = "FailedOperation.DisasterRecoverGroupNotFound"
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  FAILEDOPERATION_SNAPSHOTSIZELARGERTHANDATASIZE = "FailedOperation.SnapshotSizeLargerThanDataSize"
//  FAILEDOPERATION_TAGKEYRESERVED = "FailedOperation.TagKeyReserved"
//  FAILEDOPERATION_TATAGENTNOTSUPPORT = "FailedOperation.TatAgentNotSupport"
//  INSTANCESQUOTALIMITEXCEEDED = "InstancesQuotaLimitExceeded"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INVALIDCLIENTTOKEN_TOOLONG = "InvalidClientToken.TooLong"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDHOSTID_NOTFOUND = "InvalidHostId.NotFound"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDINSTANCENAME_TOOLONG = "InvalidInstanceName.TooLong"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETER_SNAPSHOTNOTFOUND = "InvalidParameter.SnapshotNotFound"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDNOTFOUND = "InvalidParameterValue.BandwidthPackageIdNotFound"
//  INVALIDPARAMETERVALUE_CLOUDSSDDATADISKSIZETOOSMALL = "InvalidParameterValue.CloudSsdDataDiskSizeTooSmall"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTFOUND = "InvalidParameterValue.InstanceTypeNotFound"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTSUPPORTHPCCLUSTER = "InvalidParameterValue.InstanceTypeNotSupportHpcCluster"
//  INVALIDPARAMETERVALUE_INSTANCETYPEREQUIREDHPCCLUSTER = "InvalidParameterValue.InstanceTypeRequiredHpcCluster"
//  INVALIDPARAMETERVALUE_INSUFFICIENTPRICE = "InvalidParameterValue.InsufficientPrice"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageId"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEOSNAME = "InvalidParameterValue.InvalidImageOsName"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_INVALIDUSERDATAFORMAT = "InvalidParameterValue.InvalidUserDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPASSWORD = "InvalidPassword"
//  INVALIDPERIOD = "InvalidPeriod"
//  INVALIDPERMISSION = "InvalidPermission"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_DISASTERRECOVERGROUP = "LimitExceeded.DisasterRecoverGroup"
//  LIMITEXCEEDED_INSTANCEENINUMLIMIT = "LimitExceeded.InstanceEniNumLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MONITORSERVICE = "MissingParameter.MonitorService"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCEINSUFFICIENT_DISASTERRECOVERGROUPCVMQUOTA = "ResourceInsufficient.DisasterRecoverGroupCvmQuota"
//  RESOURCENOTFOUND_HPCCLUSTER = "ResourceNotFound.HpcCluster"
//  RESOURCENOTFOUND_NODEFAULTCBS = "ResourceNotFound.NoDefaultCbs"
//  RESOURCENOTFOUND_NODEFAULTCBSWITHREASON = "ResourceNotFound.NoDefaultCbsWithReason"
//  RESOURCEUNAVAILABLE_INSTANCETYPE = "ResourceUnavailable.InstanceType"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDDISK = "UnsupportedOperation.InvalidDisk"
//  UNSUPPORTEDOPERATION_NOINSTANCETYPESUPPORTSPOT = "UnsupportedOperation.NoInstanceTypeSupportSpot"
//  UNSUPPORTEDOPERATION_ONLYFORPREPAIDACCOUNT = "UnsupportedOperation.OnlyForPrepaidAccount"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDINTERNATIONALUSER = "UnsupportedOperation.UnsupportedInternationalUser"
func (c *Client) InquiryPriceRunInstances(request *InquiryPriceRunInstancesRequest) (response *InquiryPriceRunInstancesResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRunInstancesRequest()
    }
    
    response = NewInquiryPriceRunInstancesResponse()
    err = c.Send(request, response)
    return
}

// InquiryPriceRunInstances
// This API is used to query the price of creating instances. You can only use this API for instances whose configuration is within the purchase limit. For more information, see [RunInstances](https://intl.cloud.tencent.com/document/api/213/15730?from_cn_redirect=1).
//
// error code that may be returned:
//  ACCOUNTQUALIFICATIONRESTRICTIONS = "AccountQualificationRestrictions"
//  AUTHFAILURE_CAMROLENAMEAUTHENTICATEFAILED = "AuthFailure.CamRoleNameAuthenticateFailed"
//  FAILEDOPERATION_DISASTERRECOVERGROUPNOTFOUND = "FailedOperation.DisasterRecoverGroupNotFound"
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  FAILEDOPERATION_SNAPSHOTSIZELARGERTHANDATASIZE = "FailedOperation.SnapshotSizeLargerThanDataSize"
//  FAILEDOPERATION_TAGKEYRESERVED = "FailedOperation.TagKeyReserved"
//  FAILEDOPERATION_TATAGENTNOTSUPPORT = "FailedOperation.TatAgentNotSupport"
//  INSTANCESQUOTALIMITEXCEEDED = "InstancesQuotaLimitExceeded"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INVALIDCLIENTTOKEN_TOOLONG = "InvalidClientToken.TooLong"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDHOSTID_NOTFOUND = "InvalidHostId.NotFound"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDINSTANCENAME_TOOLONG = "InvalidInstanceName.TooLong"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETER_SNAPSHOTNOTFOUND = "InvalidParameter.SnapshotNotFound"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDNOTFOUND = "InvalidParameterValue.BandwidthPackageIdNotFound"
//  INVALIDPARAMETERVALUE_CLOUDSSDDATADISKSIZETOOSMALL = "InvalidParameterValue.CloudSsdDataDiskSizeTooSmall"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTFOUND = "InvalidParameterValue.InstanceTypeNotFound"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTSUPPORTHPCCLUSTER = "InvalidParameterValue.InstanceTypeNotSupportHpcCluster"
//  INVALIDPARAMETERVALUE_INSTANCETYPEREQUIREDHPCCLUSTER = "InvalidParameterValue.InstanceTypeRequiredHpcCluster"
//  INVALIDPARAMETERVALUE_INSUFFICIENTPRICE = "InvalidParameterValue.InsufficientPrice"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageId"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEOSNAME = "InvalidParameterValue.InvalidImageOsName"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_INVALIDUSERDATAFORMAT = "InvalidParameterValue.InvalidUserDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPASSWORD = "InvalidPassword"
//  INVALIDPERIOD = "InvalidPeriod"
//  INVALIDPERMISSION = "InvalidPermission"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_DISASTERRECOVERGROUP = "LimitExceeded.DisasterRecoverGroup"
//  LIMITEXCEEDED_INSTANCEENINUMLIMIT = "LimitExceeded.InstanceEniNumLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MONITORSERVICE = "MissingParameter.MonitorService"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCEINSUFFICIENT_DISASTERRECOVERGROUPCVMQUOTA = "ResourceInsufficient.DisasterRecoverGroupCvmQuota"
//  RESOURCENOTFOUND_HPCCLUSTER = "ResourceNotFound.HpcCluster"
//  RESOURCENOTFOUND_NODEFAULTCBS = "ResourceNotFound.NoDefaultCbs"
//  RESOURCENOTFOUND_NODEFAULTCBSWITHREASON = "ResourceNotFound.NoDefaultCbsWithReason"
//  RESOURCEUNAVAILABLE_INSTANCETYPE = "ResourceUnavailable.InstanceType"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INVALIDDISK = "UnsupportedOperation.InvalidDisk"
//  UNSUPPORTEDOPERATION_NOINSTANCETYPESUPPORTSPOT = "UnsupportedOperation.NoInstanceTypeSupportSpot"
//  UNSUPPORTEDOPERATION_ONLYFORPREPAIDACCOUNT = "UnsupportedOperation.OnlyForPrepaidAccount"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDINTERNATIONALUSER = "UnsupportedOperation.UnsupportedInternationalUser"
func (c *Client) InquiryPriceRunInstancesWithContext(ctx context.Context, request *InquiryPriceRunInstancesRequest) (response *InquiryPriceRunInstancesResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRunInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewInquiryPriceRunInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDisasterRecoverGroupAttributeRequest() (request *ModifyDisasterRecoverGroupAttributeRequest) {
    request = &ModifyDisasterRecoverGroupAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyDisasterRecoverGroupAttribute")
    
    
    return
}

func NewModifyDisasterRecoverGroupAttributeResponse() (response *ModifyDisasterRecoverGroupAttributeResponse) {
    response = &ModifyDisasterRecoverGroupAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDisasterRecoverGroupAttribute
// This API is used to modify the attributes of [spread placement groups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1).
//
// error code that may be returned:
//  RESOURCEINSUFFICIENT_DISASTERRECOVERGROUPCVMQUOTA = "ResourceInsufficient.DisasterRecoverGroupCvmQuota"
//  RESOURCENOTFOUND_INVALIDPLACEMENTSET = "ResourceNotFound.InvalidPlacementSet"
func (c *Client) ModifyDisasterRecoverGroupAttribute(request *ModifyDisasterRecoverGroupAttributeRequest) (response *ModifyDisasterRecoverGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDisasterRecoverGroupAttributeRequest()
    }
    
    response = NewModifyDisasterRecoverGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyDisasterRecoverGroupAttribute
// This API is used to modify the attributes of [spread placement groups](https://intl.cloud.tencent.com/document/product/213/15486?from_cn_redirect=1).
//
// error code that may be returned:
//  RESOURCEINSUFFICIENT_DISASTERRECOVERGROUPCVMQUOTA = "ResourceInsufficient.DisasterRecoverGroupCvmQuota"
//  RESOURCENOTFOUND_INVALIDPLACEMENTSET = "ResourceNotFound.InvalidPlacementSet"
func (c *Client) ModifyDisasterRecoverGroupAttributeWithContext(ctx context.Context, request *ModifyDisasterRecoverGroupAttributeRequest) (response *ModifyDisasterRecoverGroupAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDisasterRecoverGroupAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyDisasterRecoverGroupAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostsAttributeRequest() (request *ModifyHostsAttributeRequest) {
    request = &ModifyHostsAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyHostsAttribute")
    
    
    return
}

func NewModifyHostsAttributeResponse() (response *ModifyHostsAttributeResponse) {
    response = &ModifyHostsAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyHostsAttribute
// This API is used to modify the attributes of a CDH instance, such as instance name and renewal flag. One of the two parameters, HostName and RenewFlag, must be set, but you cannot set both of them at the same time.
//
// error code that may be returned:
//  INVALIDHOST_NOTSUPPORTED = "InvalidHost.NotSupported"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDHOSTID_NOTFOUND = "InvalidHostId.NotFound"
func (c *Client) ModifyHostsAttribute(request *ModifyHostsAttributeRequest) (response *ModifyHostsAttributeResponse, err error) {
    if request == nil {
        request = NewModifyHostsAttributeRequest()
    }
    
    response = NewModifyHostsAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyHostsAttribute
// This API is used to modify the attributes of a CDH instance, such as instance name and renewal flag. One of the two parameters, HostName and RenewFlag, must be set, but you cannot set both of them at the same time.
//
// error code that may be returned:
//  INVALIDHOST_NOTSUPPORTED = "InvalidHost.NotSupported"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDHOSTID_NOTFOUND = "InvalidHostId.NotFound"
func (c *Client) ModifyHostsAttributeWithContext(ctx context.Context, request *ModifyHostsAttributeRequest) (response *ModifyHostsAttributeResponse, err error) {
    if request == nil {
        request = NewModifyHostsAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyHostsAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImageAttributeRequest() (request *ModifyImageAttributeRequest) {
    request = &ModifyImageAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyImageAttribute")
    
    
    return
}

func NewModifyImageAttributeResponse() (response *ModifyImageAttributeResponse) {
    response = &ModifyImageAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyImageAttribute
// This API is used to modify image attributes.
//
// 
//
// * Attributes of shared images cannot be modified.
//
// error code that may be returned:
//  INVALIDIMAGEID_INCORRECTSTATE = "InvalidImageId.IncorrectState"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDIMAGENAME_DUPLICATE = "InvalidImageName.Duplicate"
//  INVALIDPARAMETER_VALUETOOLARGE = "InvalidParameter.ValueTooLarge"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
func (c *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (response *ModifyImageAttributeResponse, err error) {
    if request == nil {
        request = NewModifyImageAttributeRequest()
    }
    
    response = NewModifyImageAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyImageAttribute
// This API is used to modify image attributes.
//
// 
//
// * Attributes of shared images cannot be modified.
//
// error code that may be returned:
//  INVALIDIMAGEID_INCORRECTSTATE = "InvalidImageId.IncorrectState"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDIMAGENAME_DUPLICATE = "InvalidImageName.Duplicate"
//  INVALIDPARAMETER_VALUETOOLARGE = "InvalidParameter.ValueTooLarge"
//  INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"
func (c *Client) ModifyImageAttributeWithContext(ctx context.Context, request *ModifyImageAttributeRequest) (response *ModifyImageAttributeResponse, err error) {
    if request == nil {
        request = NewModifyImageAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyImageAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyImageSharePermissionRequest() (request *ModifyImageSharePermissionRequest) {
    request = &ModifyImageSharePermissionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyImageSharePermission")
    
    
    return
}

func NewModifyImageSharePermissionResponse() (response *ModifyImageSharePermissionResponse) {
    response = &ModifyImageSharePermissionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyImageSharePermission
// This API is used to modify image sharing information.
//
// 
//
// * The accounts with which an image is shared can use the shared image to create instances.
//
// * Each custom image can be shared with up to 50 accounts.
//
// * You can use a shared image to create instances, but you cannot change its name and description.
//
// * If an image is shared with another account, the shared image will be in the same region as the original image.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTISYOURSELF = "FailedOperation.AccountIsYourSelf"
//  FAILEDOPERATION_NOTMASTERACCOUNT = "FailedOperation.NotMasterAccount"
//  FAILEDOPERATION_QIMAGESHAREFAILED = "FailedOperation.QImageShareFailed"
//  FAILEDOPERATION_RIMAGESHAREFAILED = "FailedOperation.RImageShareFailed"
//  INVALIDACCOUNTID_NOTFOUND = "InvalidAccountId.NotFound"
//  INVALIDACCOUNTIS_YOURSELF = "InvalidAccountIs.YourSelf"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageId"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  OVERQUOTA = "OverQuota"
//  UNAUTHORIZEDOPERATION_IMAGENOTBELONGTOACCOUNT = "UnauthorizedOperation.ImageNotBelongToAccount"
func (c *Client) ModifyImageSharePermission(request *ModifyImageSharePermissionRequest) (response *ModifyImageSharePermissionResponse, err error) {
    if request == nil {
        request = NewModifyImageSharePermissionRequest()
    }
    
    response = NewModifyImageSharePermissionResponse()
    err = c.Send(request, response)
    return
}

// ModifyImageSharePermission
// This API is used to modify image sharing information.
//
// 
//
// * The accounts with which an image is shared can use the shared image to create instances.
//
// * Each custom image can be shared with up to 50 accounts.
//
// * You can use a shared image to create instances, but you cannot change its name and description.
//
// * If an image is shared with another account, the shared image will be in the same region as the original image.
//
// error code that may be returned:
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTISYOURSELF = "FailedOperation.AccountIsYourSelf"
//  FAILEDOPERATION_NOTMASTERACCOUNT = "FailedOperation.NotMasterAccount"
//  FAILEDOPERATION_QIMAGESHAREFAILED = "FailedOperation.QImageShareFailed"
//  FAILEDOPERATION_RIMAGESHAREFAILED = "FailedOperation.RImageShareFailed"
//  INVALIDACCOUNTID_NOTFOUND = "InvalidAccountId.NotFound"
//  INVALIDACCOUNTIS_YOURSELF = "InvalidAccountIs.YourSelf"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageId"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  OVERQUOTA = "OverQuota"
//  UNAUTHORIZEDOPERATION_IMAGENOTBELONGTOACCOUNT = "UnauthorizedOperation.ImageNotBelongToAccount"
func (c *Client) ModifyImageSharePermissionWithContext(ctx context.Context, request *ModifyImageSharePermissionRequest) (response *ModifyImageSharePermissionResponse, err error) {
    if request == nil {
        request = NewModifyImageSharePermissionRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyImageSharePermissionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
    request = &ModifyInstancesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesAttribute")
    
    
    return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
    response = &ModifyInstancesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesAttribute
// The API is used to modify the attributes of an instance. Only the name and the associated security groups can be modified for now.
//
// 
//
// * An attribute must be specified in the request.
//
// * "Instance name" is a custom name for easier management. Tencent Cloud does not use the name for online support or instance management.
//
// * Batch operations are supported. Each request can modify up to 100 instances.
//
// * When you modify the security groups associated with an instance is modified, the original security groups are disassociated.
//
// * You can use the API [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1#.E7.A4.BA.E4.BE.8B3-.E6.9F.A5.E8.AF.A2.E5.AE.9E.E4.BE.8B.E7.9A.84.E6.9C.80.E6.96.B0.E6.93.8D.E4.BD.9C.E6.83.85.E5.86.B5) to query the instance operation result. If the 'LatestOperationState' in the response is **SUCCESS**, the operation is successful.
//
// error code that may be returned:
//  AUTHFAILURE_CAMROLENAMEAUTHENTICATEFAILED = "AuthFailure.CamRoleNameAuthenticateFailed"
//  FAILEDOPERATION_SECURITYGROUPACTIONFAILED = "FailedOperation.SecurityGroupActionFailed"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCENAME_TOOLONG = "InvalidInstanceName.TooLong"
//  INVALIDPARAMETER_HOSTNAMEILLEGAL = "InvalidParameter.HostNameIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDUSERDATAFORMAT = "InvalidParameterValue.InvalidUserDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  LIMITEXCEEDED_ASSOCIATEUSGLIMITEXCEEDED = "LimitExceeded.AssociateUSGLimitExceeded"
//  LIMITEXCEEDED_CVMSVIFSPERSECGROUPLIMITEXCEEDED = "LimitExceeded.CvmsVifsPerSecGroupLimitExceeded"
//  LIMITEXCEEDED_SINGLEUSGQUOTA = "LimitExceeded.SingleUSGQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateEnterServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateExitServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesAttributeRequest()
    }
    
    response = NewModifyInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyInstancesAttribute
// The API is used to modify the attributes of an instance. Only the name and the associated security groups can be modified for now.
//
// 
//
// * An attribute must be specified in the request.
//
// * "Instance name" is a custom name for easier management. Tencent Cloud does not use the name for online support or instance management.
//
// * Batch operations are supported. Each request can modify up to 100 instances.
//
// * When you modify the security groups associated with an instance is modified, the original security groups are disassociated.
//
// * You can use the API [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1#.E7.A4.BA.E4.BE.8B3-.E6.9F.A5.E8.AF.A2.E5.AE.9E.E4.BE.8B.E7.9A.84.E6.9C.80.E6.96.B0.E6.93.8D.E4.BD.9C.E6.83.85.E5.86.B5) to query the instance operation result. If the 'LatestOperationState' in the response is **SUCCESS**, the operation is successful.
//
// error code that may be returned:
//  AUTHFAILURE_CAMROLENAMEAUTHENTICATEFAILED = "AuthFailure.CamRoleNameAuthenticateFailed"
//  FAILEDOPERATION_SECURITYGROUPACTIONFAILED = "FailedOperation.SecurityGroupActionFailed"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCENAME_TOOLONG = "InvalidInstanceName.TooLong"
//  INVALIDPARAMETER_HOSTNAMEILLEGAL = "InvalidParameter.HostNameIllegal"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDUSERDATAFORMAT = "InvalidParameterValue.InvalidUserDataFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  LIMITEXCEEDED_ASSOCIATEUSGLIMITEXCEEDED = "LimitExceeded.AssociateUSGLimitExceeded"
//  LIMITEXCEEDED_CVMSVIFSPERSECGROUPLIMITEXCEEDED = "LimitExceeded.CvmsVifsPerSecGroupLimitExceeded"
//  LIMITEXCEEDED_SINGLEUSGQUOTA = "LimitExceeded.SingleUSGQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateEnterServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateExitServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) ModifyInstancesAttributeWithContext(ctx context.Context, request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesProjectRequest() (request *ModifyInstancesProjectRequest) {
    request = &ModifyInstancesProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesProject")
    
    
    return
}

func NewModifyInstancesProjectResponse() (response *ModifyInstancesProjectResponse) {
    response = &ModifyInstancesProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesProject
// This API is used to change the project to which an instance belongs.
//
// 
//
// * Project is a virtual concept. You can create multiple projects under one account, manage different resources in each project, and assign different instances to different projects. You may use the [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API to query instances and use the project ID to filter results.
//
// * You cannot modify the project of an instance that is bound to a load balancer. You need to firstly unbind the load balancer from the instance by using the [`DeregisterInstancesFromLoadBalancer`](https://intl.cloud.tencent.com/document/api/214/1258?from_cn_redirect=1) API.
//
// [^_^]: # (If you modify the project of an instance, security groups associated with the instance will be automatically disassociated. You can use the [`ModifyInstancesAttribute`](https://intl.cloud.tencent.com/document/api/213/15739?from_cn_redirect=1) API to associate the instance with the security groups again.
//
// * Batch operations are supported. You can operate up to 100 instances in each request.
//
// * You can call the [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1#.E7.A4.BA.E4.BE.8B3-.E6.9F.A5.E8.AF.A2.E5.AE.9E.E4.BE.8B.E7.9A.84.E6.9C.80.E6.96.B0.E6.93.8D.E4.BD.9C.E6.83.85.E5.86.B5) API and find the result of the operation in the response parameter `LatestOperationState`. If the value is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
func (c *Client) ModifyInstancesProject(request *ModifyInstancesProjectRequest) (response *ModifyInstancesProjectResponse, err error) {
    if request == nil {
        request = NewModifyInstancesProjectRequest()
    }
    
    response = NewModifyInstancesProjectResponse()
    err = c.Send(request, response)
    return
}

// ModifyInstancesProject
// This API is used to change the project to which an instance belongs.
//
// 
//
// * Project is a virtual concept. You can create multiple projects under one account, manage different resources in each project, and assign different instances to different projects. You may use the [`DescribeInstances`](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) API to query instances and use the project ID to filter results.
//
// * You cannot modify the project of an instance that is bound to a load balancer. You need to firstly unbind the load balancer from the instance by using the [`DeregisterInstancesFromLoadBalancer`](https://intl.cloud.tencent.com/document/api/214/1258?from_cn_redirect=1) API.
//
// [^_^]: # (If you modify the project of an instance, security groups associated with the instance will be automatically disassociated. You can use the [`ModifyInstancesAttribute`](https://intl.cloud.tencent.com/document/api/213/15739?from_cn_redirect=1) API to associate the instance with the security groups again.
//
// * Batch operations are supported. You can operate up to 100 instances in each request.
//
// * You can call the [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1#.E7.A4.BA.E4.BE.8B3-.E6.9F.A5.E8.AF.A2.E5.AE.9E.E4.BE.8B.E7.9A.84.E6.9C.80.E6.96.B0.E6.93.8D.E4.BD.9C.E6.83.85.E5.86.B5) API and find the result of the operation in the response parameter `LatestOperationState`. If the value is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
func (c *Client) ModifyInstancesProjectWithContext(ctx context.Context, request *ModifyInstancesProjectRequest) (response *ModifyInstancesProjectResponse, err error) {
    if request == nil {
        request = NewModifyInstancesProjectRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyInstancesProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesVpcAttributeRequest() (request *ModifyInstancesVpcAttributeRequest) {
    request = &ModifyInstancesVpcAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyInstancesVpcAttribute")
    
    
    return
}

func NewModifyInstancesVpcAttributeResponse() (response *ModifyInstancesVpcAttributeResponse) {
    response = &ModifyInstancesVpcAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesVpcAttribute
// This API is used to modify the VPC attributes of an instance, such as the VPC IP address.
//
// * By default, the instances will shut down when you perform this operation and restart upon completion.
//
// * If the specified VPC ID and subnet ID (the subnet must be in the same availability zone as the instance) are different from the VPC where the specified instance resides, the instance will be migrated to a subnet of the specified VPC. Before performing this operation, make sure that the specified instance is not bound with an [ENI](https://intl.cloud.tencent.com/document/product/576?from_cn_redirect=1) or [CLB](https://intl.cloud.tencent.com/document/product/214?from_cn_redirect=1).
//
// error code that may be returned:
//  ENINOTALLOWEDCHANGESUBNET = "EniNotAllowedChangeSubnet"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCESTATE = "InvalidInstanceState"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IPADDRESSMALFORMED = "InvalidParameterValue.IPAddressMalformed"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MUSTDHCPENABLEDVPC = "InvalidParameterValue.MustDhcpEnabledVpc"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_IPV6NOTSUPPORTVPCMIGRATE = "UnsupportedOperation.IPv6NotSupportVpcMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEBANNING = "UnsupportedOperation.InstanceStateBanning"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateExitServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_MODIFYVPCWITHCLB = "UnsupportedOperation.ModifyVPCWithCLB"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
//  VPCADDRNOTINSUBNET = "VpcAddrNotInSubNet"
//  VPCIPISUSED = "VpcIpIsUsed"
func (c *Client) ModifyInstancesVpcAttribute(request *ModifyInstancesVpcAttributeRequest) (response *ModifyInstancesVpcAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesVpcAttributeRequest()
    }
    
    response = NewModifyInstancesVpcAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyInstancesVpcAttribute
// This API is used to modify the VPC attributes of an instance, such as the VPC IP address.
//
// * By default, the instances will shut down when you perform this operation and restart upon completion.
//
// * If the specified VPC ID and subnet ID (the subnet must be in the same availability zone as the instance) are different from the VPC where the specified instance resides, the instance will be migrated to a subnet of the specified VPC. Before performing this operation, make sure that the specified instance is not bound with an [ENI](https://intl.cloud.tencent.com/document/product/576?from_cn_redirect=1) or [CLB](https://intl.cloud.tencent.com/document/product/214?from_cn_redirect=1).
//
// error code that may be returned:
//  ENINOTALLOWEDCHANGESUBNET = "EniNotAllowedChangeSubnet"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCESTATE = "InvalidInstanceState"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IPADDRESSMALFORMED = "InvalidParameterValue.IPAddressMalformed"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MUSTDHCPENABLEDVPC = "InvalidParameterValue.MustDhcpEnabledVpc"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_IPV6NOTSUPPORTVPCMIGRATE = "UnsupportedOperation.IPv6NotSupportVpcMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEBANNING = "UnsupportedOperation.InstanceStateBanning"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateExitServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_MODIFYVPCWITHCLB = "UnsupportedOperation.ModifyVPCWithCLB"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
//  VPCADDRNOTINSUBNET = "VpcAddrNotInSubNet"
//  VPCIPISUSED = "VpcIpIsUsed"
func (c *Client) ModifyInstancesVpcAttributeWithContext(ctx context.Context, request *ModifyInstancesVpcAttributeRequest) (response *ModifyInstancesVpcAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesVpcAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyInstancesVpcAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyKeyPairAttributeRequest() (request *ModifyKeyPairAttributeRequest) {
    request = &ModifyKeyPairAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyKeyPairAttribute")
    
    
    return
}

func NewModifyKeyPairAttributeResponse() (response *ModifyKeyPairAttributeResponse) {
    response = &ModifyKeyPairAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyKeyPairAttribute
// This API is used to modify the attributes of key pairs.
//
// 
//
// * This API modifies the name and description of the key pair identified by the key pair ID.
//
// * The name of the key pair must be unique.
//
// * Key pair ID is the unique identifier of a key pair and cannot be modified.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDKEYPAIRID_MALFORMED = "InvalidKeyPairId.Malformed"
//  INVALIDKEYPAIRID_NOTFOUND = "InvalidKeyPairId.NotFound"
//  INVALIDKEYPAIRNAME_DUPLICATE = "InvalidKeyPairName.Duplicate"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyKeyPairAttribute(request *ModifyKeyPairAttributeRequest) (response *ModifyKeyPairAttributeResponse, err error) {
    if request == nil {
        request = NewModifyKeyPairAttributeRequest()
    }
    
    response = NewModifyKeyPairAttributeResponse()
    err = c.Send(request, response)
    return
}

// ModifyKeyPairAttribute
// This API is used to modify the attributes of key pairs.
//
// 
//
// * This API modifies the name and description of the key pair identified by the key pair ID.
//
// * The name of the key pair must be unique.
//
// * Key pair ID is the unique identifier of a key pair and cannot be modified.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDKEYPAIRID_MALFORMED = "InvalidKeyPairId.Malformed"
//  INVALIDKEYPAIRID_NOTFOUND = "InvalidKeyPairId.NotFound"
//  INVALIDKEYPAIRNAME_DUPLICATE = "InvalidKeyPairName.Duplicate"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyKeyPairAttributeWithContext(ctx context.Context, request *ModifyKeyPairAttributeRequest) (response *ModifyKeyPairAttributeResponse, err error) {
    if request == nil {
        request = NewModifyKeyPairAttributeRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyKeyPairAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLaunchTemplateDefaultVersionRequest() (request *ModifyLaunchTemplateDefaultVersionRequest) {
    request = &ModifyLaunchTemplateDefaultVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ModifyLaunchTemplateDefaultVersion")
    
    
    return
}

func NewModifyLaunchTemplateDefaultVersionResponse() (response *ModifyLaunchTemplateDefaultVersionResponse) {
    response = &ModifyLaunchTemplateDefaultVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLaunchTemplateDefaultVersion
// This API is used to modify the default version of the instance launch template.
//
// error code that may be returned:
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERSETALREADY = "InvalidParameterValue.LaunchTemplateIdVerSetAlready"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyLaunchTemplateDefaultVersion(request *ModifyLaunchTemplateDefaultVersionRequest) (response *ModifyLaunchTemplateDefaultVersionResponse, err error) {
    if request == nil {
        request = NewModifyLaunchTemplateDefaultVersionRequest()
    }
    
    response = NewModifyLaunchTemplateDefaultVersionResponse()
    err = c.Send(request, response)
    return
}

// ModifyLaunchTemplateDefaultVersion
// This API is used to modify the default version of the instance launch template.
//
// error code that may be returned:
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdVerNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDVERSETALREADY = "InvalidParameterValue.LaunchTemplateIdVerSetAlready"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyLaunchTemplateDefaultVersionWithContext(ctx context.Context, request *ModifyLaunchTemplateDefaultVersionRequest) (response *ModifyLaunchTemplateDefaultVersionResponse, err error) {
    if request == nil {
        request = NewModifyLaunchTemplateDefaultVersionRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyLaunchTemplateDefaultVersionResponse()
    err = c.Send(request, response)
    return
}

func NewPurchaseReservedInstancesOfferingRequest() (request *PurchaseReservedInstancesOfferingRequest) {
    request = &PurchaseReservedInstancesOfferingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "PurchaseReservedInstancesOffering")
    
    
    return
}

func NewPurchaseReservedInstancesOfferingResponse() (response *PurchaseReservedInstancesOfferingResponse) {
    response = &PurchaseReservedInstancesOfferingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PurchaseReservedInstancesOffering
// This API is used to purchase one or more specific Reserved Instances.
//
// error code that may be returned:
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDCLIENTTOKEN_TOOLONG = "InvalidClientToken.TooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_INVALIDPERMISSIONNONINTERNATIONALACCOUNT = "UnsupportedOperation.InvalidPermissionNonInternationalAccount"
//  UNSUPPORTEDOPERATION_RESERVEDINSTANCEINVISIBLEFORUSER = "UnsupportedOperation.ReservedInstanceInvisibleForUser"
//  UNSUPPORTEDOPERATION_RESERVEDINSTANCEOUTOFQUATA = "UnsupportedOperation.ReservedInstanceOutofQuata"
func (c *Client) PurchaseReservedInstancesOffering(request *PurchaseReservedInstancesOfferingRequest) (response *PurchaseReservedInstancesOfferingResponse, err error) {
    if request == nil {
        request = NewPurchaseReservedInstancesOfferingRequest()
    }
    
    response = NewPurchaseReservedInstancesOfferingResponse()
    err = c.Send(request, response)
    return
}

// PurchaseReservedInstancesOffering
// This API is used to purchase one or more specific Reserved Instances.
//
// error code that may be returned:
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDCLIENTTOKEN_TOOLONG = "InvalidClientToken.TooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  UNSUPPORTEDOPERATION_INVALIDPERMISSIONNONINTERNATIONALACCOUNT = "UnsupportedOperation.InvalidPermissionNonInternationalAccount"
//  UNSUPPORTEDOPERATION_RESERVEDINSTANCEINVISIBLEFORUSER = "UnsupportedOperation.ReservedInstanceInvisibleForUser"
//  UNSUPPORTEDOPERATION_RESERVEDINSTANCEOUTOFQUATA = "UnsupportedOperation.ReservedInstanceOutofQuata"
func (c *Client) PurchaseReservedInstancesOfferingWithContext(ctx context.Context, request *PurchaseReservedInstancesOfferingRequest) (response *PurchaseReservedInstancesOfferingResponse, err error) {
    if request == nil {
        request = NewPurchaseReservedInstancesOfferingRequest()
    }
    request.SetContext(ctx)
    
    response = NewPurchaseReservedInstancesOfferingResponse()
    err = c.Send(request, response)
    return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
    request = &RebootInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "RebootInstances")
    
    
    return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
    response = &RebootInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RebootInstances
// This API is used to restart instances.
//
// 
//
// * You can only perform this operation on instances whose status is `RUNNING`.
//
// * If the API is called successfully, the instance status will become `REBOOTING`. After the instance is restarted, its status will become `RUNNING` again.
//
// * Forced restart is supported. A forced restart is similar to switching off the power of a physical computer and starting it again. It may cause data loss or file system corruption. Be sure to only force start a CVM when it cannot be restarted normally.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATECORRUPTED = "UnsupportedOperation.InstanceStateCorrupted"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateEnterServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPED = "UnsupportedOperation.InstanceStateStopped"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

// RebootInstances
// This API is used to restart instances.
//
// 
//
// * You can only perform this operation on instances whose status is `RUNNING`.
//
// * If the API is called successfully, the instance status will become `REBOOTING`. After the instance is restarted, its status will become `RUNNING` again.
//
// * Forced restart is supported. A forced restart is similar to switching off the power of a physical computer and starting it again. It may cause data loss or file system corruption. Be sure to only force start a CVM when it cannot be restarted normally.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATECORRUPTED = "UnsupportedOperation.InstanceStateCorrupted"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateEnterServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPED = "UnsupportedOperation.InstanceStateStopped"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) RebootInstancesWithContext(ctx context.Context, request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstanceRequest() (request *ResetInstanceRequest) {
    request = &ResetInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ResetInstance")
    
    
    return
}

func NewResetInstanceResponse() (response *ResetInstanceResponse) {
    response = &ResetInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstance
// This API is used to reinstall the operating system of the specified instance.
//
// 
//
// * If you specify an `ImageId`, the specified image is used. Otherwise, the image used by the current instance is used.
//
// * The system disk will be formatted and reset. Therefore, make sure that no important files are stored on the system disk.
//
// * If the operating system switches between `Linux` and `Windows`, the system disk `ID` of the instance will change, and the snapshots that are associated with the system disk can no longer be used to roll back and restore data.
//
// * If no password is specified, you will get a random password via internal message.
//
// * You can only use this API to switch the operating system between `Linux` and `Windows` for instances whose [system disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#SystemDisk) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`.
//
// * Currently, this API only supports instances in Mainland China regions.
//
// error code that may be returned:
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER_HOSTNAMEILLEGAL = "InvalidParameter.HostNameIllegal"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALHOSTNAME = "InvalidParameterValue.IllegalHostName"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTFOUND = "InvalidParameterValue.InstanceTypeNotFound"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORGIVENINSTANCETYPE = "InvalidParameterValue.InvalidImageForGivenInstanceType"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEIDFORRETSETINSTANCE = "InvalidParameterValue.InvalidImageIdForRetsetInstance"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_KEYPAIRNOTFOUND = "InvalidParameterValue.KeyPairNotFound"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPASSWORD = "InvalidPassword"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MONITORSERVICE = "MissingParameter.MonitorService"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_CHCINSTALLCLOUDIMAGEWITHOUTDEPLOYNETWORK = "OperationDenied.ChcInstallCloudImageWithoutDeployNetwork"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCECHARGETYPE = "UnsupportedOperation.InstanceChargeType"
//  UNSUPPORTEDOPERATION_INSTANCESTATECORRUPTED = "UnsupportedOperation.InstanceStateCorrupted"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateEnterServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateExitServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEFREEZING = "UnsupportedOperation.InstanceStateFreezing"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_KEYPAIRUNSUPPORTEDWINDOWS = "UnsupportedOperation.KeyPairUnsupportedWindows"
//  UNSUPPORTEDOPERATION_RAWLOCALDISKINSREINSTALLTOQCOW2 = "UnsupportedOperation.RawLocalDiskInsReinstalltoQcow2"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
    if request == nil {
        request = NewResetInstanceRequest()
    }
    
    response = NewResetInstanceResponse()
    err = c.Send(request, response)
    return
}

// ResetInstance
// This API is used to reinstall the operating system of the specified instance.
//
// 
//
// * If you specify an `ImageId`, the specified image is used. Otherwise, the image used by the current instance is used.
//
// * The system disk will be formatted and reset. Therefore, make sure that no important files are stored on the system disk.
//
// * If the operating system switches between `Linux` and `Windows`, the system disk `ID` of the instance will change, and the snapshots that are associated with the system disk can no longer be used to roll back and restore data.
//
// * If no password is specified, you will get a random password via internal message.
//
// * You can only use this API to switch the operating system between `Linux` and `Windows` for instances whose [system disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#SystemDisk) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`.
//
// * Currently, this API only supports instances in Mainland China regions.
//
// error code that may be returned:
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER_HOSTNAMEILLEGAL = "InvalidParameter.HostNameIllegal"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALHOSTNAME = "InvalidParameterValue.IllegalHostName"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTFOUND = "InvalidParameterValue.InstanceTypeNotFound"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORGIVENINSTANCETYPE = "InvalidParameterValue.InvalidImageForGivenInstanceType"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEIDFORRETSETINSTANCE = "InvalidParameterValue.InvalidImageIdForRetsetInstance"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_KEYPAIRNOTFOUND = "InvalidParameterValue.KeyPairNotFound"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPASSWORD = "InvalidPassword"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MONITORSERVICE = "MissingParameter.MonitorService"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_CHCINSTALLCLOUDIMAGEWITHOUTDEPLOYNETWORK = "OperationDenied.ChcInstallCloudImageWithoutDeployNetwork"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCECHARGETYPE = "UnsupportedOperation.InstanceChargeType"
//  UNSUPPORTEDOPERATION_INSTANCESTATECORRUPTED = "UnsupportedOperation.InstanceStateCorrupted"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateEnterServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateExitServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEFREEZING = "UnsupportedOperation.InstanceStateFreezing"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_KEYPAIRUNSUPPORTEDWINDOWS = "UnsupportedOperation.KeyPairUnsupportedWindows"
//  UNSUPPORTEDOPERATION_RAWLOCALDISKINSREINSTALLTOQCOW2 = "UnsupportedOperation.RawLocalDiskInsReinstalltoQcow2"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) ResetInstanceWithContext(ctx context.Context, request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
    if request == nil {
        request = NewResetInstanceRequest()
    }
    request.SetContext(ctx)
    
    response = NewResetInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesInternetMaxBandwidthRequest() (request *ResetInstancesInternetMaxBandwidthRequest) {
    request = &ResetInstancesInternetMaxBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ResetInstancesInternetMaxBandwidth")
    
    
    return
}

func NewResetInstancesInternetMaxBandwidthResponse() (response *ResetInstancesInternetMaxBandwidthResponse) {
    response = &ResetInstancesInternetMaxBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstancesInternetMaxBandwidth
// This API is used to change the public bandwidth cap of an instance.
//
// 
//
// * The allowed bandwidth cap varies for different models. For details, see [Purchasing Network Bandwidth](https://intl.cloud.tencent.com/document/product/213/509?from_cn_redirect=1).
//
// * For bandwidth billed by the `TRAFFIC_POSTPAID_BY_HOUR` method, changing the bandwidth cap through this API takes effect in real time. Users can increase or reduce bandwidth within applicable limits.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUNDEIP = "FailedOperation.NotFoundEIP"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPERMISSION = "InvalidPermission"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateEnterServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) ResetInstancesInternetMaxBandwidth(request *ResetInstancesInternetMaxBandwidthRequest) (response *ResetInstancesInternetMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewResetInstancesInternetMaxBandwidthRequest()
    }
    
    response = NewResetInstancesInternetMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

// ResetInstancesInternetMaxBandwidth
// This API is used to change the public bandwidth cap of an instance.
//
// 
//
// * The allowed bandwidth cap varies for different models. For details, see [Purchasing Network Bandwidth](https://intl.cloud.tencent.com/document/product/213/509?from_cn_redirect=1).
//
// * For bandwidth billed by the `TRAFFIC_POSTPAID_BY_HOUR` method, changing the bandwidth cap through this API takes effect in real time. Users can increase or reduce bandwidth within applicable limits.
//
// error code that may be returned:
//  FAILEDOPERATION_NOTFOUNDEIP = "FailedOperation.NotFoundEIP"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPERMISSION = "InvalidPermission"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateEnterServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) ResetInstancesInternetMaxBandwidthWithContext(ctx context.Context, request *ResetInstancesInternetMaxBandwidthRequest) (response *ResetInstancesInternetMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewResetInstancesInternetMaxBandwidthRequest()
    }
    request.SetContext(ctx)
    
    response = NewResetInstancesInternetMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesPasswordRequest() (request *ResetInstancesPasswordRequest) {
    request = &ResetInstancesPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ResetInstancesPassword")
    
    
    return
}

func NewResetInstancesPasswordResponse() (response *ResetInstancesPasswordResponse) {
    response = &ResetInstancesPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstancesPassword
// This API is used to reset the password of the operating system instances to a user-specified password.
//
// 
//
// * To modify the password of the administrator account: the name of the administrator account varies with the operating system. In Windows, it is `Administrator`; in Ubuntu, it is `ubuntu`; in Linux, it is `root`.
//
// * To reset the password of a running instance, you need to set the parameter `ForceStop` to `True` for a forced shutdown. If not, only passwords of stopped instances can be reset.
//
// * Batch operations are supported. You can reset the passwords of up to 100 instances to the same value once.
//
// * You can call the [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1#.E7.A4.BA.E4.BE.8B3-.E6.9F.A5.E8.AF.A2.E5.AE.9E.E4.BE.8B.E7.9A.84.E6.9C.80.E6.96.B0.E6.93.8D.E4.BD.9C.E6.83.85.E5.86.B5) API and find the result of the operation in the response parameter `LatestOperationState`. If the value is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPASSWORD = "InvalidPassword"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    if request == nil {
        request = NewResetInstancesPasswordRequest()
    }
    
    response = NewResetInstancesPasswordResponse()
    err = c.Send(request, response)
    return
}

// ResetInstancesPassword
// This API is used to reset the password of the operating system instances to a user-specified password.
//
// 
//
// * To modify the password of the administrator account: the name of the administrator account varies with the operating system. In Windows, it is `Administrator`; in Ubuntu, it is `ubuntu`; in Linux, it is `root`.
//
// * To reset the password of a running instance, you need to set the parameter `ForceStop` to `True` for a forced shutdown. If not, only passwords of stopped instances can be reset.
//
// * Batch operations are supported. You can reset the passwords of up to 100 instances to the same value once.
//
// * You can call the [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1#.E7.A4.BA.E4.BE.8B3-.E6.9F.A5.E8.AF.A2.E5.AE.9E.E4.BE.8B.E7.9A.84.E6.9C.80.E6.96.B0.E6.93.8D.E4.BD.9C.E6.83.85.E5.86.B5) API and find the result of the operation in the response parameter `LatestOperationState`. If the value is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPASSWORD = "InvalidPassword"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) ResetInstancesPasswordWithContext(ctx context.Context, request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    if request == nil {
        request = NewResetInstancesPasswordRequest()
    }
    request.SetContext(ctx)
    
    response = NewResetInstancesPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesTypeRequest() (request *ResetInstancesTypeRequest) {
    request = &ResetInstancesTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ResetInstancesType")
    
    
    return
}

func NewResetInstancesTypeResponse() (response *ResetInstancesTypeResponse) {
    response = &ResetInstancesTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstancesType
// This API is used to change the model of an instance.
//
// * You can only use this API to change the models of instances whose [system disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`.
//
// * Currently, you cannot use this API to change the models of [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDINSTANCEAPPLICATIONROLEEMR = "FailedOperation.InvalidInstanceApplicationRoleEmr"
//  FAILEDOPERATION_PROMOTIONALPERIORESTRICTION = "FailedOperation.PromotionalPerioRestriction"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GPUINSTANCEFAMILY = "InvalidParameterValue.GPUInstanceFamily"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDPERMISSION = "InvalidPermission"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_ENINUMLIMIT = "LimitExceeded.EniNumLimit"
//  LIMITEXCEEDED_INSTANCETYPEBANDWIDTH = "LimitExceeded.InstanceTypeBandwidth"
//  LIMITEXCEEDED_SPOTQUOTA = "LimitExceeded.SpotQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  RESOURCEUNAVAILABLE_INSTANCETYPE = "ResourceUnavailable.InstanceType"
//  RESOURCESSOLDOUT_AVAILABLEZONE = "ResourcesSoldOut.AvailableZone"
//  RESOURCESSOLDOUT_SPECIFIEDINSTANCETYPE = "ResourcesSoldOut.SpecifiedInstanceType"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_NOINSTANCETYPESUPPORTSPOT = "UnsupportedOperation.NoInstanceTypeSupportSpot"
//  UNSUPPORTEDOPERATION_SPECIALINSTANCETYPE = "UnsupportedOperation.SpecialInstanceType"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDCHANGEINSTANCEFAMILY = "UnsupportedOperation.UnsupportedChangeInstanceFamily"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDCHANGEINSTANCEFAMILYTOARM = "UnsupportedOperation.UnsupportedChangeInstanceFamilyToARM"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDCHANGEINSTANCEFAMILYTOSA3 = "UnsupportedOperation.UnsupportedChangeInstanceFamilyToSA3"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDCHANGEINSTANCETOTHISINSTANCEFAMILY = "UnsupportedOperation.UnsupportedChangeInstanceToThisInstanceFamily"
func (c *Client) ResetInstancesType(request *ResetInstancesTypeRequest) (response *ResetInstancesTypeResponse, err error) {
    if request == nil {
        request = NewResetInstancesTypeRequest()
    }
    
    response = NewResetInstancesTypeResponse()
    err = c.Send(request, response)
    return
}

// ResetInstancesType
// This API is used to change the model of an instance.
//
// * You can only use this API to change the models of instances whose [system disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`.
//
// * Currently, you cannot use this API to change the models of [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDINSTANCEAPPLICATIONROLEEMR = "FailedOperation.InvalidInstanceApplicationRoleEmr"
//  FAILEDOPERATION_PROMOTIONALPERIORESTRICTION = "FailedOperation.PromotionalPerioRestriction"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GPUINSTANCEFAMILY = "InvalidParameterValue.GPUInstanceFamily"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDPERMISSION = "InvalidPermission"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_ENINUMLIMIT = "LimitExceeded.EniNumLimit"
//  LIMITEXCEEDED_INSTANCETYPEBANDWIDTH = "LimitExceeded.InstanceTypeBandwidth"
//  LIMITEXCEEDED_SPOTQUOTA = "LimitExceeded.SpotQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  RESOURCEUNAVAILABLE_INSTANCETYPE = "ResourceUnavailable.InstanceType"
//  RESOURCESSOLDOUT_AVAILABLEZONE = "ResourcesSoldOut.AvailableZone"
//  RESOURCESSOLDOUT_SPECIFIEDINSTANCETYPE = "ResourcesSoldOut.SpecifiedInstanceType"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_NOINSTANCETYPESUPPORTSPOT = "UnsupportedOperation.NoInstanceTypeSupportSpot"
//  UNSUPPORTEDOPERATION_SPECIALINSTANCETYPE = "UnsupportedOperation.SpecialInstanceType"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDCHANGEINSTANCEFAMILY = "UnsupportedOperation.UnsupportedChangeInstanceFamily"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDCHANGEINSTANCEFAMILYTOARM = "UnsupportedOperation.UnsupportedChangeInstanceFamilyToARM"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDCHANGEINSTANCEFAMILYTOSA3 = "UnsupportedOperation.UnsupportedChangeInstanceFamilyToSA3"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDCHANGEINSTANCETOTHISINSTANCEFAMILY = "UnsupportedOperation.UnsupportedChangeInstanceToThisInstanceFamily"
func (c *Client) ResetInstancesTypeWithContext(ctx context.Context, request *ResetInstancesTypeRequest) (response *ResetInstancesTypeResponse, err error) {
    if request == nil {
        request = NewResetInstancesTypeRequest()
    }
    request.SetContext(ctx)
    
    response = NewResetInstancesTypeResponse()
    err = c.Send(request, response)
    return
}

func NewResizeInstanceDisksRequest() (request *ResizeInstanceDisksRequest) {
    request = &ResizeInstanceDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "ResizeInstanceDisks")
    
    
    return
}

func NewResizeInstanceDisksResponse() (response *ResizeInstanceDisksResponse) {
    response = &ResizeInstanceDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResizeInstanceDisks
// This API (ResizeInstanceDisks) is used to expand the data disks of an instance.
//
// 
//
// * Currently, you can only use the API to expand non-elastic data disks whose [disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`. You can use [`DescribeDisks`](https://intl.cloud.tencent.com/document/api/362/16315?from_cn_redirect=1) to check whether a disk is elastic. If the `Portable` field in the response is `false`, it means that the disk is non-elastic.
//
// * Currently, this API does not support [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances.
//
// * Currently, only one data disk can be expanded at a time.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ATLEASTONE = "MissingParameter.AtLeastOne"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_INSTANCECHARGETYPE = "UnsupportedOperation.InstanceChargeType"
//  UNSUPPORTEDOPERATION_INSTANCESTATECORRUPTED = "UnsupportedOperation.InstanceStateCorrupted"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPED = "UnsupportedOperation.InstanceStateStopped"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_SPECIALINSTANCETYPE = "UnsupportedOperation.SpecialInstanceType"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) ResizeInstanceDisks(request *ResizeInstanceDisksRequest) (response *ResizeInstanceDisksResponse, err error) {
    if request == nil {
        request = NewResizeInstanceDisksRequest()
    }
    
    response = NewResizeInstanceDisksResponse()
    err = c.Send(request, response)
    return
}

// ResizeInstanceDisks
// This API (ResizeInstanceDisks) is used to expand the data disks of an instance.
//
// 
//
// * Currently, you can only use the API to expand non-elastic data disks whose [disk type](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#block_device) is `CLOUD_BASIC`, `CLOUD_PREMIUM`, or `CLOUD_SSD`. You can use [`DescribeDisks`](https://intl.cloud.tencent.com/document/api/362/16315?from_cn_redirect=1) to check whether a disk is elastic. If the `Portable` field in the response is `false`, it means that the disk is non-elastic.
//
// * Currently, this API does not support [CDH](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) instances.
//
// * Currently, only one data disk can be expanded at a time.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ATLEASTONE = "MissingParameter.AtLeastOne"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_INSTANCECHARGETYPE = "UnsupportedOperation.InstanceChargeType"
//  UNSUPPORTEDOPERATION_INSTANCESTATECORRUPTED = "UnsupportedOperation.InstanceStateCorrupted"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPED = "UnsupportedOperation.InstanceStateStopped"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_SPECIALINSTANCETYPE = "UnsupportedOperation.SpecialInstanceType"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) ResizeInstanceDisksWithContext(ctx context.Context, request *ResizeInstanceDisksRequest) (response *ResizeInstanceDisksResponse, err error) {
    if request == nil {
        request = NewResizeInstanceDisksRequest()
    }
    request.SetContext(ctx)
    
    response = NewResizeInstanceDisksResponse()
    err = c.Send(request, response)
    return
}

func NewRunInstancesRequest() (request *RunInstancesRequest) {
    request = &RunInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "RunInstances")
    
    
    return
}

func NewRunInstancesResponse() (response *RunInstancesResponse) {
    response = &RunInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunInstances
// This API is used to create one or more instances with a specified configuration.
//
// 
//
// * After an instance is created successfully, it will start up automatically, and the [instance state](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#instance_state) will become "Running".
//
// * If you create a pay-as-you-go instance billed on an hourly basis, an amount equivalent to the hourly rate will be frozen before the creation. Make sure your account balance is sufficient before calling this API.
//
// * The number of instances you can purchase through this API is subject to the [CVM instance purchase limit](https://intl.cloud.tencent.com/document/product/213/2664?from_cn_redirect=1). Both the instances created through this API and the console will be counted toward the quota.
//
// * This API is an async API. An instance `ID` list will be returned after you successfully make a creation request. However, it does not mean the creation has been completed. The state of the instance will be `Creating` during the creation. You can use [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) to query the status of the instance. If the status changes from `Creating` to `Running`, it means that the instance has been created successfully.
//
// error code that may be returned:
//  ACCOUNTQUALIFICATIONRESTRICTIONS = "AccountQualificationRestrictions"
//  AUTHFAILURE_CAMROLENAMEAUTHENTICATEFAILED = "AuthFailure.CamRoleNameAuthenticateFailed"
//  FAILEDOPERATION_DISASTERRECOVERGROUPNOTFOUND = "FailedOperation.DisasterRecoverGroupNotFound"
//  FAILEDOPERATION_ILLEGALTAGVALUE = "FailedOperation.IllegalTagValue"
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  FAILEDOPERATION_NOAVAILABLEIPADDRESSCOUNTINSUBNET = "FailedOperation.NoAvailableIpAddressCountInSubnet"
//  FAILEDOPERATION_SECURITYGROUPACTIONFAILED = "FailedOperation.SecurityGroupActionFailed"
//  FAILEDOPERATION_SNAPSHOTSIZELARGERTHANDATASIZE = "FailedOperation.SnapshotSizeLargerThanDataSize"
//  FAILEDOPERATION_SNAPSHOTSIZELESSTHANDATASIZE = "FailedOperation.SnapshotSizeLessThanDataSize"
//  FAILEDOPERATION_TAGKEYRESERVED = "FailedOperation.TagKeyReserved"
//  INSTANCESQUOTALIMITEXCEEDED = "InstancesQuotaLimitExceeded"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDCLIENTTOKEN_TOOLONG = "InvalidClientToken.TooLong"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDHOSTID_NOTFOUND = "InvalidHostId.NotFound"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDINSTANCENAME_TOOLONG = "InvalidInstanceName.TooLong"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER_HOSTIDSTATUSNOTSUPPORT = "InvalidParameter.HostIdStatusNotSupport"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETER_INVALIDIPFORMAT = "InvalidParameter.InvalidIpFormat"
//  INVALIDPARAMETER_LACKCORECOUNTORTHREADPERCORE = "InvalidParameter.LackCoreCountOrThreadPerCore"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETER_SNAPSHOTNOTFOUND = "InvalidParameter.SnapshotNotFound"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDNOTFOUND = "InvalidParameterValue.BandwidthPackageIdNotFound"
//  INVALIDPARAMETERVALUE_CHCHOSTSNOTFOUND = "InvalidParameterValue.ChcHostsNotFound"
//  INVALIDPARAMETERVALUE_CLOUDSSDDATADISKSIZETOOSMALL = "InvalidParameterValue.CloudSsdDataDiskSizeTooSmall"
//  INVALIDPARAMETERVALUE_CORECOUNTVALUE = "InvalidParameterValue.CoreCountValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_IPADDRESSMALFORMED = "InvalidParameterValue.IPAddressMalformed"
//  INVALIDPARAMETERVALUE_ILLEGALHOSTNAME = "InvalidParameterValue.IllegalHostName"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTFOUND = "InvalidParameterValue.InstanceTypeNotFound"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTSUPPORTHPCCLUSTER = "InvalidParameterValue.InstanceTypeNotSupportHpcCluster"
//  INVALIDPARAMETERVALUE_INSTANCETYPEREQUIREDHPCCLUSTER = "InvalidParameterValue.InstanceTypeRequiredHpcCluster"
//  INVALIDPARAMETERVALUE_INSUFFICIENTOFFERING = "InvalidParameterValue.InsufficientOffering"
//  INVALIDPARAMETERVALUE_INSUFFICIENTPRICE = "InvalidParameterValue.InsufficientPrice"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORGIVENINSTANCETYPE = "InvalidParameterValue.InvalidImageForGivenInstanceType"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageId"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEOSNAME = "InvalidParameterValue.InvalidImageOsName"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_INVALIDUSERDATAFORMAT = "InvalidParameterValue.InvalidUserDataFormat"
//  INVALIDPARAMETERVALUE_KEYPAIRNOTFOUND = "InvalidParameterValue.KeyPairNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MUSTDHCPENABLEDVPC = "InvalidParameterValue.MustDhcpEnabledVpc"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SUBNETIDMALFORMED = "InvalidParameterValue.SubnetIdMalformed"
//  INVALIDPARAMETERVALUE_SUBNETNOTEXIST = "InvalidParameterValue.SubnetNotExist"
//  INVALIDPARAMETERVALUE_TAGKEYNOTFOUND = "InvalidParameterValue.TagKeyNotFound"
//  INVALIDPARAMETERVALUE_THREADPERCOREVALUE = "InvalidParameterValue.ThreadPerCoreValue"
//  INVALIDPARAMETERVALUE_VPCIDMALFORMED = "InvalidParameterValue.VpcIdMalformed"
//  INVALIDPARAMETERVALUE_VPCIDNOTEXIST = "InvalidParameterValue.VpcIdNotExist"
//  INVALIDPARAMETERVALUE_VPCIDZONEIDNOTMATCH = "InvalidParameterValue.VpcIdZoneIdNotMatch"
//  INVALIDPARAMETERVALUE_VPCNOTSUPPORTIPV6ADDRESS = "InvalidParameterValue.VpcNotSupportIpv6Address"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDPASSWORD = "InvalidPassword"
//  INVALIDPERIOD = "InvalidPeriod"
//  INVALIDPERMISSION = "InvalidPermission"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_CVMSVIFSPERSECGROUPLIMITEXCEEDED = "LimitExceeded.CvmsVifsPerSecGroupLimitExceeded"
//  LIMITEXCEEDED_DISASTERRECOVERGROUP = "LimitExceeded.DisasterRecoverGroup"
//  LIMITEXCEEDED_IPV6ADDRESSNUM = "LimitExceeded.IPv6AddressNum"
//  LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"
//  LIMITEXCEEDED_PREPAYQUOTA = "LimitExceeded.PrepayQuota"
//  LIMITEXCEEDED_SINGLEUSGQUOTA = "LimitExceeded.SingleUSGQuota"
//  LIMITEXCEEDED_SPOTQUOTA = "LimitExceeded.SpotQuota"
//  LIMITEXCEEDED_USERSPOTQUOTA = "LimitExceeded.UserSpotQuota"
//  LIMITEXCEEDED_VPCSUBNETNUM = "LimitExceeded.VpcSubnetNum"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DPDKINSTANCETYPEREQUIREDVPC = "MissingParameter.DPDKInstanceTypeRequiredVPC"
//  MISSINGPARAMETER_MONITORSERVICE = "MissingParameter.MonitorService"
//  OPERATIONDENIED_CHCINSTALLCLOUDIMAGEWITHOUTDEPLOYNETWORK = "OperationDenied.ChcInstallCloudImageWithoutDeployNetwork"
//  RESOURCEINSUFFICIENT_AVAILABILITYZONESOLDOUT = "ResourceInsufficient.AvailabilityZoneSoldOut"
//  RESOURCEINSUFFICIENT_CLOUDDISKSOLDOUT = "ResourceInsufficient.CloudDiskSoldOut"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCEINSUFFICIENT_DISASTERRECOVERGROUPCVMQUOTA = "ResourceInsufficient.DisasterRecoverGroupCvmQuota"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  RESOURCENOTFOUND_HPCCLUSTER = "ResourceNotFound.HpcCluster"
//  RESOURCENOTFOUND_NODEFAULTCBS = "ResourceNotFound.NoDefaultCbs"
//  RESOURCENOTFOUND_NODEFAULTCBSWITHREASON = "ResourceNotFound.NoDefaultCbsWithReason"
//  RESOURCEUNAVAILABLE_INSTANCETYPE = "ResourceUnavailable.InstanceType"
//  RESOURCESSOLDOUT_EIPINSUFFICIENT = "ResourcesSoldOut.EipInsufficient"
//  RESOURCESSOLDOUT_SPECIFIEDINSTANCETYPE = "ResourcesSoldOut.SpecifiedInstanceType"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INVALIDDISK = "UnsupportedOperation.InvalidDisk"
//  UNSUPPORTEDOPERATION_INVALIDREGIONDISKENCRYPT = "UnsupportedOperation.InvalidRegionDiskEncrypt"
//  UNSUPPORTEDOPERATION_KEYPAIRUNSUPPORTEDWINDOWS = "UnsupportedOperation.KeyPairUnsupportedWindows"
//  UNSUPPORTEDOPERATION_NOINSTANCETYPESUPPORTSPOT = "UnsupportedOperation.NoInstanceTypeSupportSpot"
//  UNSUPPORTEDOPERATION_NOTSUPPORTIMPORTINSTANCESACTIONTIMER = "UnsupportedOperation.NotSupportImportInstancesActionTimer"
//  UNSUPPORTEDOPERATION_ONLYFORPREPAIDACCOUNT = "UnsupportedOperation.OnlyForPrepaidAccount"
//  UNSUPPORTEDOPERATION_RAWLOCALDISKINSREINSTALLTOQCOW2 = "UnsupportedOperation.RawLocalDiskInsReinstalltoQcow2"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDINTERNATIONALUSER = "UnsupportedOperation.UnsupportedInternationalUser"
//  VPCADDRNOTINSUBNET = "VpcAddrNotInSubNet"
//  VPCIPISUSED = "VpcIpIsUsed"
func (c *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
    if request == nil {
        request = NewRunInstancesRequest()
    }
    
    response = NewRunInstancesResponse()
    err = c.Send(request, response)
    return
}

// RunInstances
// This API is used to create one or more instances with a specified configuration.
//
// 
//
// * After an instance is created successfully, it will start up automatically, and the [instance state](https://intl.cloud.tencent.com/document/api/213/9452?from_cn_redirect=1#instance_state) will become "Running".
//
// * If you create a pay-as-you-go instance billed on an hourly basis, an amount equivalent to the hourly rate will be frozen before the creation. Make sure your account balance is sufficient before calling this API.
//
// * The number of instances you can purchase through this API is subject to the [CVM instance purchase limit](https://intl.cloud.tencent.com/document/product/213/2664?from_cn_redirect=1). Both the instances created through this API and the console will be counted toward the quota.
//
// * This API is an async API. An instance `ID` list will be returned after you successfully make a creation request. However, it does not mean the creation has been completed. The state of the instance will be `Creating` during the creation. You can use [DescribeInstances](https://intl.cloud.tencent.com/document/api/213/15728?from_cn_redirect=1) to query the status of the instance. If the status changes from `Creating` to `Running`, it means that the instance has been created successfully.
//
// error code that may be returned:
//  ACCOUNTQUALIFICATIONRESTRICTIONS = "AccountQualificationRestrictions"
//  AUTHFAILURE_CAMROLENAMEAUTHENTICATEFAILED = "AuthFailure.CamRoleNameAuthenticateFailed"
//  FAILEDOPERATION_DISASTERRECOVERGROUPNOTFOUND = "FailedOperation.DisasterRecoverGroupNotFound"
//  FAILEDOPERATION_ILLEGALTAGVALUE = "FailedOperation.IllegalTagValue"
//  FAILEDOPERATION_INQUIRYPRICEFAILED = "FailedOperation.InquiryPriceFailed"
//  FAILEDOPERATION_NOAVAILABLEIPADDRESSCOUNTINSUBNET = "FailedOperation.NoAvailableIpAddressCountInSubnet"
//  FAILEDOPERATION_SECURITYGROUPACTIONFAILED = "FailedOperation.SecurityGroupActionFailed"
//  FAILEDOPERATION_SNAPSHOTSIZELARGERTHANDATASIZE = "FailedOperation.SnapshotSizeLargerThanDataSize"
//  FAILEDOPERATION_SNAPSHOTSIZELESSTHANDATASIZE = "FailedOperation.SnapshotSizeLessThanDataSize"
//  FAILEDOPERATION_TAGKEYRESERVED = "FailedOperation.TagKeyReserved"
//  INSTANCESQUOTALIMITEXCEEDED = "InstancesQuotaLimitExceeded"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"
//  INVALIDCLIENTTOKEN_TOOLONG = "InvalidClientToken.TooLong"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDHOSTID_NOTFOUND = "InvalidHostId.NotFound"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDINSTANCENAME_TOOLONG = "InvalidInstanceName.TooLong"
//  INVALIDINSTANCETYPE_MALFORMED = "InvalidInstanceType.Malformed"
//  INVALIDPARAMETER_HOSTIDSTATUSNOTSUPPORT = "InvalidParameter.HostIdStatusNotSupport"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETER_INVALIDIPFORMAT = "InvalidParameter.InvalidIpFormat"
//  INVALIDPARAMETER_LACKCORECOUNTORTHREADPERCORE = "InvalidParameter.LackCoreCountOrThreadPerCore"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETER_SNAPSHOTNOTFOUND = "InvalidParameter.SnapshotNotFound"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDMALFORMED = "InvalidParameterValue.BandwidthPackageIdMalformed"
//  INVALIDPARAMETERVALUE_BANDWIDTHPACKAGEIDNOTFOUND = "InvalidParameterValue.BandwidthPackageIdNotFound"
//  INVALIDPARAMETERVALUE_CHCHOSTSNOTFOUND = "InvalidParameterValue.ChcHostsNotFound"
//  INVALIDPARAMETERVALUE_CLOUDSSDDATADISKSIZETOOSMALL = "InvalidParameterValue.CloudSsdDataDiskSizeTooSmall"
//  INVALIDPARAMETERVALUE_CORECOUNTVALUE = "InvalidParameterValue.CoreCountValue"
//  INVALIDPARAMETERVALUE_DUPLICATE = "InvalidParameterValue.Duplicate"
//  INVALIDPARAMETERVALUE_IPADDRESSMALFORMED = "InvalidParameterValue.IPAddressMalformed"
//  INVALIDPARAMETERVALUE_ILLEGALHOSTNAME = "InvalidParameterValue.IllegalHostName"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTFOUND = "InvalidParameterValue.InstanceTypeNotFound"
//  INVALIDPARAMETERVALUE_INSTANCETYPENOTSUPPORTHPCCLUSTER = "InvalidParameterValue.InstanceTypeNotSupportHpcCluster"
//  INVALIDPARAMETERVALUE_INSTANCETYPEREQUIREDHPCCLUSTER = "InvalidParameterValue.InstanceTypeRequiredHpcCluster"
//  INVALIDPARAMETERVALUE_INSUFFICIENTOFFERING = "InvalidParameterValue.InsufficientOffering"
//  INVALIDPARAMETERVALUE_INSUFFICIENTPRICE = "InvalidParameterValue.InsufficientPrice"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORGIVENINSTANCETYPE = "InvalidParameterValue.InvalidImageForGivenInstanceType"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageId"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEOSNAME = "InvalidParameterValue.InvalidImageOsName"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_INVALIDTIMEFORMAT = "InvalidParameterValue.InvalidTimeFormat"
//  INVALIDPARAMETERVALUE_INVALIDUSERDATAFORMAT = "InvalidParameterValue.InvalidUserDataFormat"
//  INVALIDPARAMETERVALUE_KEYPAIRNOTFOUND = "InvalidParameterValue.KeyPairNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDMALFORMED = "InvalidParameterValue.LaunchTemplateIdMalformed"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEIDNOTEXISTED = "InvalidParameterValue.LaunchTemplateIdNotExisted"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATENOTFOUND = "InvalidParameterValue.LaunchTemplateNotFound"
//  INVALIDPARAMETERVALUE_LAUNCHTEMPLATEVERSION = "InvalidParameterValue.LaunchTemplateVersion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_MUSTDHCPENABLEDVPC = "InvalidParameterValue.MustDhcpEnabledVpc"
//  INVALIDPARAMETERVALUE_RANGE = "InvalidParameterValue.Range"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SUBNETIDMALFORMED = "InvalidParameterValue.SubnetIdMalformed"
//  INVALIDPARAMETERVALUE_SUBNETNOTEXIST = "InvalidParameterValue.SubnetNotExist"
//  INVALIDPARAMETERVALUE_TAGKEYNOTFOUND = "InvalidParameterValue.TagKeyNotFound"
//  INVALIDPARAMETERVALUE_THREADPERCOREVALUE = "InvalidParameterValue.ThreadPerCoreValue"
//  INVALIDPARAMETERVALUE_VPCIDMALFORMED = "InvalidParameterValue.VpcIdMalformed"
//  INVALIDPARAMETERVALUE_VPCIDNOTEXIST = "InvalidParameterValue.VpcIdNotExist"
//  INVALIDPARAMETERVALUE_VPCIDZONEIDNOTMATCH = "InvalidParameterValue.VpcIdZoneIdNotMatch"
//  INVALIDPARAMETERVALUE_VPCNOTSUPPORTIPV6ADDRESS = "InvalidParameterValue.VpcNotSupportIpv6Address"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  INVALIDPASSWORD = "InvalidPassword"
//  INVALIDPERIOD = "InvalidPeriod"
//  INVALIDPERMISSION = "InvalidPermission"
//  INVALIDPROJECTID_NOTFOUND = "InvalidProjectId.NotFound"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_CVMSVIFSPERSECGROUPLIMITEXCEEDED = "LimitExceeded.CvmsVifsPerSecGroupLimitExceeded"
//  LIMITEXCEEDED_DISASTERRECOVERGROUP = "LimitExceeded.DisasterRecoverGroup"
//  LIMITEXCEEDED_IPV6ADDRESSNUM = "LimitExceeded.IPv6AddressNum"
//  LIMITEXCEEDED_INSTANCEQUOTA = "LimitExceeded.InstanceQuota"
//  LIMITEXCEEDED_PREPAYQUOTA = "LimitExceeded.PrepayQuota"
//  LIMITEXCEEDED_SINGLEUSGQUOTA = "LimitExceeded.SingleUSGQuota"
//  LIMITEXCEEDED_SPOTQUOTA = "LimitExceeded.SpotQuota"
//  LIMITEXCEEDED_USERSPOTQUOTA = "LimitExceeded.UserSpotQuota"
//  LIMITEXCEEDED_VPCSUBNETNUM = "LimitExceeded.VpcSubnetNum"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_DPDKINSTANCETYPEREQUIREDVPC = "MissingParameter.DPDKInstanceTypeRequiredVPC"
//  MISSINGPARAMETER_MONITORSERVICE = "MissingParameter.MonitorService"
//  OPERATIONDENIED_CHCINSTALLCLOUDIMAGEWITHOUTDEPLOYNETWORK = "OperationDenied.ChcInstallCloudImageWithoutDeployNetwork"
//  RESOURCEINSUFFICIENT_AVAILABILITYZONESOLDOUT = "ResourceInsufficient.AvailabilityZoneSoldOut"
//  RESOURCEINSUFFICIENT_CLOUDDISKSOLDOUT = "ResourceInsufficient.CloudDiskSoldOut"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCEINSUFFICIENT_DISASTERRECOVERGROUPCVMQUOTA = "ResourceInsufficient.DisasterRecoverGroupCvmQuota"
//  RESOURCEINSUFFICIENT_SPECIFIEDINSTANCETYPE = "ResourceInsufficient.SpecifiedInstanceType"
//  RESOURCENOTFOUND_HPCCLUSTER = "ResourceNotFound.HpcCluster"
//  RESOURCENOTFOUND_NODEFAULTCBS = "ResourceNotFound.NoDefaultCbs"
//  RESOURCENOTFOUND_NODEFAULTCBSWITHREASON = "ResourceNotFound.NoDefaultCbsWithReason"
//  RESOURCEUNAVAILABLE_INSTANCETYPE = "ResourceUnavailable.InstanceType"
//  RESOURCESSOLDOUT_EIPINSUFFICIENT = "ResourcesSoldOut.EipInsufficient"
//  RESOURCESSOLDOUT_SPECIFIEDINSTANCETYPE = "ResourcesSoldOut.SpecifiedInstanceType"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
//  UNSUPPORTEDOPERATION_BANDWIDTHPACKAGEIDNOTSUPPORTED = "UnsupportedOperation.BandwidthPackageIdNotSupported"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INVALIDDISK = "UnsupportedOperation.InvalidDisk"
//  UNSUPPORTEDOPERATION_INVALIDREGIONDISKENCRYPT = "UnsupportedOperation.InvalidRegionDiskEncrypt"
//  UNSUPPORTEDOPERATION_KEYPAIRUNSUPPORTEDWINDOWS = "UnsupportedOperation.KeyPairUnsupportedWindows"
//  UNSUPPORTEDOPERATION_NOINSTANCETYPESUPPORTSPOT = "UnsupportedOperation.NoInstanceTypeSupportSpot"
//  UNSUPPORTEDOPERATION_NOTSUPPORTIMPORTINSTANCESACTIONTIMER = "UnsupportedOperation.NotSupportImportInstancesActionTimer"
//  UNSUPPORTEDOPERATION_ONLYFORPREPAIDACCOUNT = "UnsupportedOperation.OnlyForPrepaidAccount"
//  UNSUPPORTEDOPERATION_RAWLOCALDISKINSREINSTALLTOQCOW2 = "UnsupportedOperation.RawLocalDiskInsReinstalltoQcow2"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDINTERNATIONALUSER = "UnsupportedOperation.UnsupportedInternationalUser"
//  VPCADDRNOTINSUBNET = "VpcAddrNotInSubNet"
//  VPCIPISUSED = "VpcIpIsUsed"
func (c *Client) RunInstancesWithContext(ctx context.Context, request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
    if request == nil {
        request = NewRunInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewRunInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
    request = &StartInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "StartInstances")
    
    
    return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
    response = &StartInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartInstances
// This API is used to start instances.
//
// 
//
// * You can only perform this operation on instances whose status is `STOPPED`.
//
// * The instance status will become `STARTING` when the API is called successfully and `RUNNING` when the instance is successfully started.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_INSTANCESTATECORRUPTED = "UnsupportedOperation.InstanceStateCorrupted"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEFREEZING = "UnsupportedOperation.InstanceStateFreezing"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    if request == nil {
        request = NewStartInstancesRequest()
    }
    
    response = NewStartInstancesResponse()
    err = c.Send(request, response)
    return
}

// StartInstances
// This API is used to start instances.
//
// 
//
// * You can only perform this operation on instances whose status is `STOPPED`.
//
// * The instance status will become `STARTING` when the API is called successfully and `RUNNING` when the instance is successfully started.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNSUPPORTEDOPERATION_INSTANCESTATECORRUPTED = "UnsupportedOperation.InstanceStateCorrupted"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEFREEZING = "UnsupportedOperation.InstanceStateFreezing"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
func (c *Client) StartInstancesWithContext(ctx context.Context, request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    if request == nil {
        request = NewStartInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewStartInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
    request = &StopInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "StopInstances")
    
    
    return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
    response = &StopInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopInstances
// This API is used to shut down instances.
//
// 
//
// * You can only perform this operation on instances whose status is `RUNNING`.
//
// * The instance status will become `STOPPING` when the API is called successfully and `STOPPED` when the instance is successfully shut down.
//
// * Forced shutdown is supported. A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be sht down normally.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATECORRUPTED = "UnsupportedOperation.InstanceStateCorrupted"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateEnterServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPED = "UnsupportedOperation.InstanceStateStopped"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}

// StopInstances
// This API is used to shut down instances.
//
// 
//
// * You can only perform this operation on instances whose status is `RUNNING`.
//
// * The instance status will become `STOPPING` when the API is called successfully and `STOPPED` when the instance is successfully shut down.
//
// * Forced shutdown is supported. A forced shutdown is similar to switching off the power of a physical computer. It may cause data loss or file system corruption. Be sure to only force shut down a CVM when it cannot be sht down normally.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// error code that may be returned:
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDPARAMETERCOMBINATION = "InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCESTATECORRUPTED = "UnsupportedOperation.InstanceStateCorrupted"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateEnterServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATERESCUEMODE = "UnsupportedOperation.InstanceStateRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATESERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATESHUTDOWN = "UnsupportedOperation.InstanceStateShutdown"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPED = "UnsupportedOperation.InstanceStateStopped"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_STOPPEDMODESTOPCHARGING = "UnsupportedOperation.StoppedModeStopCharging"
func (c *Client) StopInstancesWithContext(ctx context.Context, request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewSyncImagesRequest() (request *SyncImagesRequest) {
    request = &SyncImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "SyncImages")
    
    
    return
}

func NewSyncImagesResponse() (response *SyncImagesResponse) {
    response = &SyncImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncImages
// This API is used to sync a custom image to other regions.
//
// 
//
// * Each API call syncs a single image.
//
// * This API supports syncing an image to multiple regions.
//
// * Each account can have up to 10 custom images in each region. 
//
// error code that may be returned:
//  IMAGEQUOTALIMITEXCEEDED = "ImageQuotaLimitExceeded"
//  INVALIDIMAGEID_INCORRECTSTATE = "InvalidImageId.IncorrectState"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDIMAGEID_TOOLARGE = "InvalidImageId.TooLarge"
//  INVALIDIMAGENAME_DUPLICATE = "InvalidImageName.Duplicate"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageId"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDREGION_NOTFOUND = "InvalidRegion.NotFound"
//  INVALIDREGION_UNAVAILABLE = "InvalidRegion.Unavailable"
//  UNSUPPORTEDOPERATION_REGION = "UnsupportedOperation.Region"
func (c *Client) SyncImages(request *SyncImagesRequest) (response *SyncImagesResponse, err error) {
    if request == nil {
        request = NewSyncImagesRequest()
    }
    
    response = NewSyncImagesResponse()
    err = c.Send(request, response)
    return
}

// SyncImages
// This API is used to sync a custom image to other regions.
//
// 
//
// * Each API call syncs a single image.
//
// * This API supports syncing an image to multiple regions.
//
// * Each account can have up to 10 custom images in each region. 
//
// error code that may be returned:
//  IMAGEQUOTALIMITEXCEEDED = "ImageQuotaLimitExceeded"
//  INVALIDIMAGEID_INCORRECTSTATE = "InvalidImageId.IncorrectState"
//  INVALIDIMAGEID_MALFORMED = "InvalidImageId.Malformed"
//  INVALIDIMAGEID_NOTFOUND = "InvalidImageId.NotFound"
//  INVALIDIMAGEID_TOOLARGE = "InvalidImageId.TooLarge"
//  INVALIDIMAGENAME_DUPLICATE = "InvalidImageName.Duplicate"
//  INVALIDPARAMETER_INSTANCEIMAGENOTSUPPORT = "InvalidParameter.InstanceImageNotSupport"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEID = "InvalidParameterValue.InvalidImageId"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESTATE = "InvalidParameterValue.InvalidImageState"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDREGION_NOTFOUND = "InvalidRegion.NotFound"
//  INVALIDREGION_UNAVAILABLE = "InvalidRegion.Unavailable"
//  UNSUPPORTEDOPERATION_REGION = "UnsupportedOperation.Region"
func (c *Client) SyncImagesWithContext(ctx context.Context, request *SyncImagesRequest) (response *SyncImagesResponse, err error) {
    if request == nil {
        request = NewSyncImagesRequest()
    }
    request.SetContext(ctx)
    
    response = NewSyncImagesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cvm", APIVersion, "TerminateInstances")
    
    
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateInstances
// This API is used to return instances.
//
// 
//
// * Use this API to return instances that are no longer required.
//
// * Pay-as-you-go instances can be returned directly through this API.
//
// * When this API is called for the first time, the instance will be moved to the recycle bin. When this API is called for the second time, the instance will be terminated and cannot be recovered.
//
// * Batch operations are supported. The allowed maximum number of instances in each request is 100.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDINSTANCEAPPLICATIONROLEEMR = "FailedOperation.InvalidInstanceApplicationRoleEmr"
//  FAILEDOPERATION_UNRETURNABLE = "FailedOperation.Unreturnable"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCENOTSUPPORTEDPREPAIDINSTANCE = "InvalidInstanceNotSupportedPrepaidInstance"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  LIMITEXCEEDED_USERRETURNQUOTA = "LimitExceeded.UserReturnQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCECHARGETYPE = "UnsupportedOperation.InstanceChargeType"
//  UNSUPPORTEDOPERATION_INSTANCEMIXEDPRICINGMODEL = "UnsupportedOperation.InstanceMixedPricingModel"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateExitServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEFREEZING = "UnsupportedOperation.InstanceStateFreezing"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_INSTANCESPROTECTED = "UnsupportedOperation.InstancesProtected"
//  UNSUPPORTEDOPERATION_REGION = "UnsupportedOperation.Region"
//  UNSUPPORTEDOPERATION_SPECIALINSTANCETYPE = "UnsupportedOperation.SpecialInstanceType"
//  UNSUPPORTEDOPERATION_USERLIMITOPERATIONEXCEEDQUOTA = "UnsupportedOperation.UserLimitOperationExceedQuota"
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}

// TerminateInstances
// This API is used to return instances.
//
// 
//
// * Use this API to return instances that are no longer required.
//
// * Pay-as-you-go instances can be returned directly through this API.
//
// * When this API is called for the first time, the instance will be moved to the recycle bin. When this API is called for the second time, the instance will be terminated and cannot be recovered.
//
// * Batch operations are supported. The allowed maximum number of instances in each request is 100.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDINSTANCEAPPLICATIONROLEEMR = "FailedOperation.InvalidInstanceApplicationRoleEmr"
//  FAILEDOPERATION_UNRETURNABLE = "FailedOperation.Unreturnable"
//  INTERNALERROR_TRADEUNKNOWNERROR = "InternalError.TradeUnknownError"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDACCOUNT_UNPAIDORDER = "InvalidAccount.UnpaidOrder"
//  INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"
//  INVALIDINSTANCENOTSUPPORTEDPREPAIDINSTANCE = "InvalidInstanceNotSupportedPrepaidInstance"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTSUPPORTED = "InvalidParameterValue.NotSupported"
//  LIMITEXCEEDED_USERRETURNQUOTA = "LimitExceeded.UserReturnQuota"
//  MISSINGPARAMETER = "MissingParameter"
//  MUTEXOPERATION_TASKRUNNING = "MutexOperation.TaskRunning"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INSTANCECHARGETYPE = "UnsupportedOperation.InstanceChargeType"
//  UNSUPPORTEDOPERATION_INSTANCEMIXEDPRICINGMODEL = "UnsupportedOperation.InstanceMixedPricingModel"
//  UNSUPPORTEDOPERATION_INSTANCESTATEENTERRESCUEMODE = "UnsupportedOperation.InstanceStateEnterRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITRESCUEMODE = "UnsupportedOperation.InstanceStateExitRescueMode"
//  UNSUPPORTEDOPERATION_INSTANCESTATEEXITSERVICELIVEMIGRATE = "UnsupportedOperation.InstanceStateExitServiceLiveMigrate"
//  UNSUPPORTEDOPERATION_INSTANCESTATEFREEZING = "UnsupportedOperation.InstanceStateFreezing"
//  UNSUPPORTEDOPERATION_INSTANCESTATEISOLATING = "UnsupportedOperation.InstanceStateIsolating"
//  UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"
//  UNSUPPORTEDOPERATION_INSTANCESTATEREBOOTING = "UnsupportedOperation.InstanceStateRebooting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTARTING = "UnsupportedOperation.InstanceStateStarting"
//  UNSUPPORTEDOPERATION_INSTANCESTATESTOPPING = "UnsupportedOperation.InstanceStateStopping"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"
//  UNSUPPORTEDOPERATION_INSTANCESTATETERMINATING = "UnsupportedOperation.InstanceStateTerminating"
//  UNSUPPORTEDOPERATION_INSTANCESPROTECTED = "UnsupportedOperation.InstancesProtected"
//  UNSUPPORTEDOPERATION_REGION = "UnsupportedOperation.Region"
//  UNSUPPORTEDOPERATION_SPECIALINSTANCETYPE = "UnsupportedOperation.SpecialInstanceType"
//  UNSUPPORTEDOPERATION_USERLIMITOPERATIONEXCEEDQUOTA = "UnsupportedOperation.UserLimitOperationExceedQuota"
func (c *Client) TerminateInstancesWithContext(ctx context.Context, request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
