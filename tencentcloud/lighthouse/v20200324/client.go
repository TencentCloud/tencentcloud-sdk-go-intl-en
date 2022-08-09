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

package v20200324

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-03-24"

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


func NewApplyInstanceSnapshotRequest() (request *ApplyInstanceSnapshotRequest) {
    request = &ApplyInstanceSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ApplyInstanceSnapshot")
    
    
    return
}

func NewApplyInstanceSnapshotResponse() (response *ApplyInstanceSnapshotResponse) {
    response = &ApplyInstanceSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyInstanceSnapshot
// This API is used to roll back the system disk snapshot of the specified instance.
//
// <li>Only rollback to the original system disk is supported.</li>
//
// <li>Only snapshots in `NORMAL` status can be used for rollback. To query the status of a snapshot, you can call the `DescribeSnapshots` API and see the `SnapshotState` field in the response.</li>
//
// <li>When a snapshot is rolled back, the status of the instance must be `STOPPED` or `RUNNING`. You can call the `DescribeInstances` API to query the instance status. Instances in `RUNNING` status will be forcibly shut down before snapshot rollback.</li>
//
// error code that may be returned:
//  FAILEDOPERATION_SNAPSHOTOPERATIONFAILED = "FailedOperation.SnapshotOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DISKSIZENOTMATCH = "InvalidParameterValue.DiskSizeNotMatch"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_DISKBUSY = "UnsupportedOperation.DiskBusy"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_INVALIDSNAPSHOTSTATE = "UnsupportedOperation.InvalidSnapshotState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_SNAPSHOTBUSY = "UnsupportedOperation.SnapshotBusy"
//  UNSUPPORTEDOPERATION_SYSTEMBUSY = "UnsupportedOperation.SystemBusy"
func (c *Client) ApplyInstanceSnapshot(request *ApplyInstanceSnapshotRequest) (response *ApplyInstanceSnapshotResponse, err error) {
    return c.ApplyInstanceSnapshotWithContext(context.Background(), request)
}

// ApplyInstanceSnapshot
// This API is used to roll back the system disk snapshot of the specified instance.
//
// <li>Only rollback to the original system disk is supported.</li>
//
// <li>Only snapshots in `NORMAL` status can be used for rollback. To query the status of a snapshot, you can call the `DescribeSnapshots` API and see the `SnapshotState` field in the response.</li>
//
// <li>When a snapshot is rolled back, the status of the instance must be `STOPPED` or `RUNNING`. You can call the `DescribeInstances` API to query the instance status. Instances in `RUNNING` status will be forcibly shut down before snapshot rollback.</li>
//
// error code that may be returned:
//  FAILEDOPERATION_SNAPSHOTOPERATIONFAILED = "FailedOperation.SnapshotOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DISKSIZENOTMATCH = "InvalidParameterValue.DiskSizeNotMatch"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_DISKBUSY = "UnsupportedOperation.DiskBusy"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_INVALIDSNAPSHOTSTATE = "UnsupportedOperation.InvalidSnapshotState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_SNAPSHOTBUSY = "UnsupportedOperation.SnapshotBusy"
//  UNSUPPORTEDOPERATION_SYSTEMBUSY = "UnsupportedOperation.SystemBusy"
func (c *Client) ApplyInstanceSnapshotWithContext(ctx context.Context, request *ApplyInstanceSnapshotRequest) (response *ApplyInstanceSnapshotResponse, err error) {
    if request == nil {
        request = NewApplyInstanceSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyInstanceSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyInstanceSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateInstancesKeyPairsRequest() (request *AssociateInstancesKeyPairsRequest) {
    request = &AssociateInstancesKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "AssociateInstancesKeyPairs")
    
    
    return
}

func NewAssociateInstancesKeyPairsResponse() (response *AssociateInstancesKeyPairsResponse) {
    response = &AssociateInstancesKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateInstancesKeyPairs
// This API is used to bind a user-specified key pair to an instance.
//
// * Only instances on LINUX_UNIX in [RUNNING, STOPPED] status are supported. Instances in `RUNNING` status will be forcibly shut down before binding.
//
// * If the public key of a key pair is written to the SSH configuration of the instance, you will be able to log in to the instance with the private key of the key pair.
//
// * If the instance is already associated with a key, the old key will become invalid.
//
// * If you currently use a password to log in, you will no longer be able to do so after you associate the instance with a key.
//
// * Batch operations are supported. The maximum number of instances in each request is 100. If instances not available for the operation are selected, you will get an error code.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_KEYPAIRBINDDUPLICATE = "UnsupportedOperation.KeyPairBindDuplicate"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_WINDOWSNOTALLOWTOASSOCIATEKEYPAIR = "UnsupportedOperation.WindowsNotAllowToAssociateKeyPair"
func (c *Client) AssociateInstancesKeyPairs(request *AssociateInstancesKeyPairsRequest) (response *AssociateInstancesKeyPairsResponse, err error) {
    return c.AssociateInstancesKeyPairsWithContext(context.Background(), request)
}

// AssociateInstancesKeyPairs
// This API is used to bind a user-specified key pair to an instance.
//
// * Only instances on LINUX_UNIX in [RUNNING, STOPPED] status are supported. Instances in `RUNNING` status will be forcibly shut down before binding.
//
// * If the public key of a key pair is written to the SSH configuration of the instance, you will be able to log in to the instance with the private key of the key pair.
//
// * If the instance is already associated with a key, the old key will become invalid.
//
// * If you currently use a password to log in, you will no longer be able to do so after you associate the instance with a key.
//
// * Batch operations are supported. The maximum number of instances in each request is 100. If instances not available for the operation are selected, you will get an error code.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_KEYPAIRBINDDUPLICATE = "UnsupportedOperation.KeyPairBindDuplicate"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_WINDOWSNOTALLOWTOASSOCIATEKEYPAIR = "UnsupportedOperation.WindowsNotAllowToAssociateKeyPair"
func (c *Client) AssociateInstancesKeyPairsWithContext(ctx context.Context, request *AssociateInstancesKeyPairsRequest) (response *AssociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewAssociateInstancesKeyPairsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateInstancesKeyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewAttachCcnRequest() (request *AttachCcnRequest) {
    request = &AttachCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "AttachCcn")
    
    
    return
}

func NewAttachCcnResponse() (response *AttachCcnResponse) {
    response = &AttachCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachCcn
// This API is used to associate a CCN instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHCCNCONDITIONUNSATISFIED = "UnsupportedOperation.AttachCcnConditionUnsatisfied"
//  UNSUPPORTEDOPERATION_ATTACHCCNFAILED = "UnsupportedOperation.AttachCcnFailed"
//  UNSUPPORTEDOPERATION_CCNALREADYATTACHED = "UnsupportedOperation.CcnAlreadyAttached"
func (c *Client) AttachCcn(request *AttachCcnRequest) (response *AttachCcnResponse, err error) {
    return c.AttachCcnWithContext(context.Background(), request)
}

// AttachCcn
// This API is used to associate a CCN instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHCCNCONDITIONUNSATISFIED = "UnsupportedOperation.AttachCcnConditionUnsatisfied"
//  UNSUPPORTEDOPERATION_ATTACHCCNFAILED = "UnsupportedOperation.AttachCcnFailed"
//  UNSUPPORTEDOPERATION_CCNALREADYATTACHED = "UnsupportedOperation.CcnAlreadyAttached"
func (c *Client) AttachCcnWithContext(ctx context.Context, request *AttachCcnRequest) (response *AttachCcnResponse, err error) {
    if request == nil {
        request = NewAttachCcnRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachCcn require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachCcnResponse()
    err = c.Send(request, response)
    return
}

func NewAttachDisksRequest() (request *AttachDisksRequest) {
    request = &AttachDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "AttachDisks")
    
    
    return
}

func NewAttachDisksResponse() (response *AttachDisksResponse) {
    response = &AttachDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachDisks
// This API is used to attach one or more cloud disks.
//
// error code that may be returned:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  LIMITEXCEEDED_ATTACHDATADISKQUOTALIMITEXCEEDED = "LimitExceeded.AttachDataDiskQuotaLimitExceeded"
//  OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) AttachDisks(request *AttachDisksRequest) (response *AttachDisksResponse, err error) {
    return c.AttachDisksWithContext(context.Background(), request)
}

// AttachDisks
// This API is used to attach one or more cloud disks.
//
// error code that may be returned:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  LIMITEXCEEDED_ATTACHDATADISKQUOTALIMITEXCEEDED = "LimitExceeded.AttachDataDiskQuotaLimitExceeded"
//  OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
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

func NewCreateBlueprintRequest() (request *CreateBlueprintRequest) {
    request = &CreateBlueprintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateBlueprint")
    
    
    return
}

func NewCreateBlueprintResponse() (response *CreateBlueprintResponse) {
    response = &CreateBlueprintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBlueprint
// This API is used to create an image.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEBLUEPRINTFAILED = "FailedOperation.CreateBlueprintFailed"
//  FAILEDOPERATION_UNABLETOCREATEBLUEPRINT = "FailedOperation.UnableToCreateBlueprint"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) CreateBlueprint(request *CreateBlueprintRequest) (response *CreateBlueprintResponse, err error) {
    return c.CreateBlueprintWithContext(context.Background(), request)
}

// CreateBlueprint
// This API is used to create an image.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEBLUEPRINTFAILED = "FailedOperation.CreateBlueprintFailed"
//  FAILEDOPERATION_UNABLETOCREATEBLUEPRINT = "FailedOperation.UnableToCreateBlueprint"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) CreateBlueprintWithContext(ctx context.Context, request *CreateBlueprintRequest) (response *CreateBlueprintResponse, err error) {
    if request == nil {
        request = NewCreateBlueprintRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBlueprint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBlueprintResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFirewallRulesRequest() (request *CreateFirewallRulesRequest) {
    request = &CreateFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateFirewallRules")
    
    
    return
}

func NewCreateFirewallRulesResponse() (response *CreateFirewallRulesResponse) {
    response = &CreateFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFirewallRules
// This API is used to add a firewall rule on an instance.
//
// 
//
// 
//
// * `FirewallVersion` is the firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the updated rule from expiring. If it is left empty, conflicts will not be considered.
//
// 
//
// In the `FirewallRules` parameter:
//
// * Valid values of the `Protocol` field include `TCP`, `UDP`, `ICMP`, and `ALL`.
//
// * For the `Port` field, you can enter only `ALL`, a single port number, several port numbers separated by commas, or a port range indicated by two port numbers separated by a minus sign. If `Port` is a range, the port number on the left of the minus sign must be smaller than the one on the right. If `Protocol` is not `TCP` or `UDP`, `Port` can only be empty or `ALL`. The length of the `Port` field cannot exceed 64 characters.
//
// * For the `CidrBlock` field, you can enter any string that conforms to the CIDR format standard. Multi-Tenant network isolation rules take precedence over private network rules in the firewall.
//
// * For the `Action` field, you can enter only `ACCEPT` or `DROP`.
//
// * The length of the `FirewallRuleDescription` field cannot exceed 64 characters.
//
// error code that may be returned:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETER_FIREWALLRULESEXIST = "InvalidParameter.FirewallRulesExist"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  LIMITEXCEEDED_FIREWALLRULESLIMITEXCEEDED = "LimitExceeded.FirewallRulesLimitExceeded"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) CreateFirewallRules(request *CreateFirewallRulesRequest) (response *CreateFirewallRulesResponse, err error) {
    return c.CreateFirewallRulesWithContext(context.Background(), request)
}

// CreateFirewallRules
// This API is used to add a firewall rule on an instance.
//
// 
//
// 
//
// * `FirewallVersion` is the firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the updated rule from expiring. If it is left empty, conflicts will not be considered.
//
// 
//
// In the `FirewallRules` parameter:
//
// * Valid values of the `Protocol` field include `TCP`, `UDP`, `ICMP`, and `ALL`.
//
// * For the `Port` field, you can enter only `ALL`, a single port number, several port numbers separated by commas, or a port range indicated by two port numbers separated by a minus sign. If `Port` is a range, the port number on the left of the minus sign must be smaller than the one on the right. If `Protocol` is not `TCP` or `UDP`, `Port` can only be empty or `ALL`. The length of the `Port` field cannot exceed 64 characters.
//
// * For the `CidrBlock` field, you can enter any string that conforms to the CIDR format standard. Multi-Tenant network isolation rules take precedence over private network rules in the firewall.
//
// * For the `Action` field, you can enter only `ACCEPT` or `DROP`.
//
// * The length of the `FirewallRuleDescription` field cannot exceed 64 characters.
//
// error code that may be returned:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETER_FIREWALLRULESEXIST = "InvalidParameter.FirewallRulesExist"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  LIMITEXCEEDED_FIREWALLRULESLIMITEXCEEDED = "LimitExceeded.FirewallRulesLimitExceeded"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) CreateFirewallRulesWithContext(ctx context.Context, request *CreateFirewallRulesRequest) (response *CreateFirewallRulesResponse, err error) {
    if request == nil {
        request = NewCreateFirewallRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFirewallRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceSnapshotRequest() (request *CreateInstanceSnapshotRequest) {
    request = &CreateInstanceSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateInstanceSnapshot")
    
    
    return
}

func NewCreateInstanceSnapshotResponse() (response *CreateInstanceSnapshotResponse) {
    response = &CreateInstanceSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstanceSnapshot
// This API is used to create a system disk snapshot of the specified instance.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_GETSNAPSHOTALLOCQUOTALOCKERROR = "InternalError.GetSnapshotAllocQuotaLockError"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  LIMITEXCEEDED_SNAPSHOTQUOTALIMITEXCEEDED = "LimitExceeded.SnapshotQuotaLimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIEDCREATESNAPSHOTFORSTORAGEBUNDLE = "OperationDenied.OperationDeniedCreateSnapshotForStorageBundle"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_DISKBUSY = "UnsupportedOperation.DiskBusy"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) CreateInstanceSnapshot(request *CreateInstanceSnapshotRequest) (response *CreateInstanceSnapshotResponse, err error) {
    return c.CreateInstanceSnapshotWithContext(context.Background(), request)
}

// CreateInstanceSnapshot
// This API is used to create a system disk snapshot of the specified instance.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_GETSNAPSHOTALLOCQUOTALOCKERROR = "InternalError.GetSnapshotAllocQuotaLockError"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  LIMITEXCEEDED_SNAPSHOTQUOTALIMITEXCEEDED = "LimitExceeded.SnapshotQuotaLimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIEDCREATESNAPSHOTFORSTORAGEBUNDLE = "OperationDenied.OperationDeniedCreateSnapshotForStorageBundle"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_DISKBUSY = "UnsupportedOperation.DiskBusy"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) CreateInstanceSnapshotWithContext(ctx context.Context, request *CreateInstanceSnapshotRequest) (response *CreateInstanceSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateInstanceSnapshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceSnapshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
    request = &CreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateInstances")
    
    
    return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
    response = &CreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstances
// This API is used to create one or more Lighthouse instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNABLETOCREATEINSTANCES = "FailedOperation.UnableToCreateInstances"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDACTIONNOTFOUND = "InternalError.InvalidActionNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUNDLEANDBLUEPRINTNOTMATCH = "InvalidParameter.BundleAndBlueprintNotMatch"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTID = "InvalidParameterValue.BlueprintId"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTID = "InvalidParameterValue.InvalidBlueprintId"
//  INVALIDPARAMETERVALUE_INVALIDBUNDLE = "InvalidParameterValue.InvalidBundle"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERCOMBINATION = "InvalidParameterValue.InvalidParameterCombination"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INSTANCEQUOTALIMITEXCEEDED = "LimitExceeded.InstanceQuotaLimitExceeded"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_BLUEPRINTNOTFOUND = "ResourceNotFound.BlueprintNotFound"
//  RESOURCEUNAVAILABLE_BUNDLEUNAVAILABLE = "ResourceUnavailable.BundleUnavailable"
//  RESOURCESSOLDOUT_PURCHASESOURCEHASNOBUNDLECONFIGS = "ResourcesSoldOut.PurchaseSourceHasNoBundleConfigs"
//  RESOURCESSOLDOUT_ZONESHASNOBUNDLECONFIGS = "ResourcesSoldOut.ZonesHasNoBundleConfigs"
//  UNSUPPORTEDOPERATION_INSTANCELINUXUNIXCREATINGNOTSUPPORTPASSWORD = "UnsupportedOperation.InstanceLinuxUnixCreatingNotSupportPassword"
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    return c.CreateInstancesWithContext(context.Background(), request)
}

// CreateInstances
// This API is used to create one or more Lighthouse instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_UNABLETOCREATEINSTANCES = "FailedOperation.UnableToCreateInstances"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDACTIONNOTFOUND = "InternalError.InvalidActionNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUNDLEANDBLUEPRINTNOTMATCH = "InvalidParameter.BundleAndBlueprintNotMatch"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTID = "InvalidParameterValue.BlueprintId"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTID = "InvalidParameterValue.InvalidBlueprintId"
//  INVALIDPARAMETERVALUE_INVALIDBUNDLE = "InvalidParameterValue.InvalidBundle"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERCOMBINATION = "InvalidParameterValue.InvalidParameterCombination"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INSTANCEQUOTALIMITEXCEEDED = "LimitExceeded.InstanceQuotaLimitExceeded"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_BLUEPRINTNOTFOUND = "ResourceNotFound.BlueprintNotFound"
//  RESOURCEUNAVAILABLE_BUNDLEUNAVAILABLE = "ResourceUnavailable.BundleUnavailable"
//  RESOURCESSOLDOUT_PURCHASESOURCEHASNOBUNDLECONFIGS = "ResourcesSoldOut.PurchaseSourceHasNoBundleConfigs"
//  RESOURCESSOLDOUT_ZONESHASNOBUNDLECONFIGS = "ResourcesSoldOut.ZonesHasNoBundleConfigs"
//  UNSUPPORTEDOPERATION_INSTANCELINUXUNIXCREATINGNOTSUPPORTPASSWORD = "UnsupportedOperation.InstanceLinuxUnixCreatingNotSupportPassword"
func (c *Client) CreateInstancesWithContext(ctx context.Context, request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    if request == nil {
        request = NewCreateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKeyPairRequest() (request *CreateKeyPairRequest) {
    request = &CreateKeyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateKeyPair")
    
    
    return
}

func NewCreateKeyPairResponse() (response *CreateKeyPairResponse) {
    response = &CreateKeyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateKeyPair
// This API is used to create a key pair.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEKEYPAIRFAILED = "FailedOperation.CreateKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEPARAMETERVALUE = "InvalidParameterValue.DuplicateParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEEMPTY = "InvalidParameterValue.InvalidKeyPairNameEmpty"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEINCLUDEILLEGALCHAR = "InvalidParameterValue.InvalidKeyPairNameIncludeIllegalChar"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMETOOLONG = "InvalidParameterValue.InvalidKeyPairNameTooLong"
//  LIMITEXCEEDED_KEYPAIRLIMITEXCEEDED = "LimitExceeded.KeyPairLimitExceeded"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateKeyPair(request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
    return c.CreateKeyPairWithContext(context.Background(), request)
}

// CreateKeyPair
// This API is used to create a key pair.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEKEYPAIRFAILED = "FailedOperation.CreateKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEPARAMETERVALUE = "InvalidParameterValue.DuplicateParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEEMPTY = "InvalidParameterValue.InvalidKeyPairNameEmpty"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEINCLUDEILLEGALCHAR = "InvalidParameterValue.InvalidKeyPairNameIncludeIllegalChar"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMETOOLONG = "InvalidParameterValue.InvalidKeyPairNameTooLong"
//  LIMITEXCEEDED_KEYPAIRLIMITEXCEEDED = "LimitExceeded.KeyPairLimitExceeded"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateKeyPairWithContext(ctx context.Context, request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
    if request == nil {
        request = NewCreateKeyPairRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKeyPair require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKeyPairResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlueprintsRequest() (request *DeleteBlueprintsRequest) {
    request = &DeleteBlueprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteBlueprints")
    
    
    return
}

func NewDeleteBlueprintsResponse() (response *DeleteBlueprintsResponse) {
    response = &DeleteBlueprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBlueprints
// This API is used to delete an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_BLUEPRINTNOTFOUND = "ResourceNotFound.BlueprintNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION_BLUEPRINTCURSTATEINVALID = "UnsupportedOperation.BlueprintCurStateInvalid"
//  UNSUPPORTEDOPERATION_BLUEPRINTOCCUPIED = "UnsupportedOperation.BlueprintOccupied"
//  UNSUPPORTEDOPERATION_NOTSUPPORTSHAREDBLUEPRINT = "UnsupportedOperation.NotSupportSharedBlueprint"
//  UNSUPPORTEDOPERATION_POSTDESTROYRESOURCEFAILED = "UnsupportedOperation.PostDestroyResourceFailed"
func (c *Client) DeleteBlueprints(request *DeleteBlueprintsRequest) (response *DeleteBlueprintsResponse, err error) {
    return c.DeleteBlueprintsWithContext(context.Background(), request)
}

// DeleteBlueprints
// This API is used to delete an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_BLUEPRINTNOTFOUND = "ResourceNotFound.BlueprintNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION_BLUEPRINTCURSTATEINVALID = "UnsupportedOperation.BlueprintCurStateInvalid"
//  UNSUPPORTEDOPERATION_BLUEPRINTOCCUPIED = "UnsupportedOperation.BlueprintOccupied"
//  UNSUPPORTEDOPERATION_NOTSUPPORTSHAREDBLUEPRINT = "UnsupportedOperation.NotSupportSharedBlueprint"
//  UNSUPPORTEDOPERATION_POSTDESTROYRESOURCEFAILED = "UnsupportedOperation.PostDestroyResourceFailed"
func (c *Client) DeleteBlueprintsWithContext(ctx context.Context, request *DeleteBlueprintsRequest) (response *DeleteBlueprintsResponse, err error) {
    if request == nil {
        request = NewDeleteBlueprintsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBlueprints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBlueprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFirewallRulesRequest() (request *DeleteFirewallRulesRequest) {
    request = &DeleteFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteFirewallRules")
    
    
    return
}

func NewDeleteFirewallRulesResponse() (response *DeleteFirewallRulesResponse) {
    response = &DeleteFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteFirewallRules
// This API is used to delete a firewall rule of an instance.
//
// 
//
// * `FirewallVersion` is used to specify the version of the firewall to be manipulated. If the `FirewallVersion` value passed in is not equal to the current latest version of the firewall, a failure will be returned. If `FirewallVersion` is not passed in, the specified rule will be deleted directly.
//
// 
//
// In the `FirewallRules` parameter:
//
// * Valid values of the `Protocol` field include `TCP`, `UDP`, `ICMP`, and `ALL`.
//
// * For the `Port` field, you can enter only `ALL`, a single port number, several port numbers separated by commas, or a port range indicated by two port numbers separated by a minus sign. If `Port` is a range, the port number on the left of the minus sign must be smaller than the one on the right. If `Protocol` is not `TCP` or `UDP`, `Port` can only be empty or `ALL`. The length of the `Port` field cannot exceed 64 characters.
//
// * For the `CidrBlock` field, you can enter any string that conforms to the CIDR format standard. Multi-Tenant network isolation rules take precedence over private network rules in the firewall.
//
// * For the `Action` field, you can enter only `ACCEPT` or `DROP`.
//
// * The length of the `FirewallRuleDescription` field cannot exceed 64 characters.
//
// error code that may be returned:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_FIREWALLRULESNOTFOUND = "ResourceNotFound.FirewallRulesNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) DeleteFirewallRules(request *DeleteFirewallRulesRequest) (response *DeleteFirewallRulesResponse, err error) {
    return c.DeleteFirewallRulesWithContext(context.Background(), request)
}

// DeleteFirewallRules
// This API is used to delete a firewall rule of an instance.
//
// 
//
// * `FirewallVersion` is used to specify the version of the firewall to be manipulated. If the `FirewallVersion` value passed in is not equal to the current latest version of the firewall, a failure will be returned. If `FirewallVersion` is not passed in, the specified rule will be deleted directly.
//
// 
//
// In the `FirewallRules` parameter:
//
// * Valid values of the `Protocol` field include `TCP`, `UDP`, `ICMP`, and `ALL`.
//
// * For the `Port` field, you can enter only `ALL`, a single port number, several port numbers separated by commas, or a port range indicated by two port numbers separated by a minus sign. If `Port` is a range, the port number on the left of the minus sign must be smaller than the one on the right. If `Protocol` is not `TCP` or `UDP`, `Port` can only be empty or `ALL`. The length of the `Port` field cannot exceed 64 characters.
//
// * For the `CidrBlock` field, you can enter any string that conforms to the CIDR format standard. Multi-Tenant network isolation rules take precedence over private network rules in the firewall.
//
// * For the `Action` field, you can enter only `ACCEPT` or `DROP`.
//
// * The length of the `FirewallRuleDescription` field cannot exceed 64 characters.
//
// error code that may be returned:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_FIREWALLRULESNOTFOUND = "ResourceNotFound.FirewallRulesNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) DeleteFirewallRulesWithContext(ctx context.Context, request *DeleteFirewallRulesRequest) (response *DeleteFirewallRulesResponse, err error) {
    if request == nil {
        request = NewDeleteFirewallRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFirewallRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKeyPairsRequest() (request *DeleteKeyPairsRequest) {
    request = &DeleteKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteKeyPairs")
    
    
    return
}

func NewDeleteKeyPairsResponse() (response *DeleteKeyPairsResponse) {
    response = &DeleteKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteKeyPairs
// This API is used to delete a key pair.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEKEYPAIRFAILED = "FailedOperation.DeleteKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  RESOURCEINUSE_KEYPAIRINUSE = "ResourceInUse.KeyPairInUse"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNSUPPORTEDOPERATION_KEYPAIRBINDTOBLUEPRINTS = "UnsupportedOperation.KeyPairBindToBlueprints"
func (c *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (response *DeleteKeyPairsResponse, err error) {
    return c.DeleteKeyPairsWithContext(context.Background(), request)
}

// DeleteKeyPairs
// This API is used to delete a key pair.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEKEYPAIRFAILED = "FailedOperation.DeleteKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  RESOURCEINUSE_KEYPAIRINUSE = "ResourceInUse.KeyPairInUse"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNSUPPORTEDOPERATION_KEYPAIRBINDTOBLUEPRINTS = "UnsupportedOperation.KeyPairBindToBlueprints"
func (c *Client) DeleteKeyPairsWithContext(ctx context.Context, request *DeleteKeyPairsRequest) (response *DeleteKeyPairsResponse, err error) {
    if request == nil {
        request = NewDeleteKeyPairsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteKeyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
    request = &DeleteSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteSnapshots")
    
    
    return
}

func NewDeleteSnapshotsResponse() (response *DeleteSnapshotsResponse) {
    response = &DeleteSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSnapshots
// This API is used to delete a snapshot.
//
// The snapshot must be in `NORMAL` status. To query the status of a snapshot, you can call the `DescribeSnapshots` API and see the `SnapshotState` field in the response.
//
// error code that may be returned:
//  FAILEDOPERATION_SNAPSHOTOPERATIONFAILED = "FailedOperation.SnapshotOperationFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSNAPSHOTSTATE = "UnsupportedOperation.InvalidSnapshotState"
//  UNSUPPORTEDOPERATION_SNAPSHOTBUSY = "UnsupportedOperation.SnapshotBusy"
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    return c.DeleteSnapshotsWithContext(context.Background(), request)
}

// DeleteSnapshots
// This API is used to delete a snapshot.
//
// The snapshot must be in `NORMAL` status. To query the status of a snapshot, you can call the `DescribeSnapshots` API and see the `SnapshotState` field in the response.
//
// error code that may be returned:
//  FAILEDOPERATION_SNAPSHOTOPERATIONFAILED = "FailedOperation.SnapshotOperationFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSNAPSHOTSTATE = "UnsupportedOperation.InvalidSnapshotState"
//  UNSUPPORTEDOPERATION_SNAPSHOTBUSY = "UnsupportedOperation.SnapshotBusy"
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

func NewDescribeBlueprintInstancesRequest() (request *DescribeBlueprintInstancesRequest) {
    request = &DescribeBlueprintInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBlueprintInstances")
    
    
    return
}

func NewDescribeBlueprintInstancesResponse() (response *DescribeBlueprintInstancesResponse) {
    response = &DescribeBlueprintInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlueprintInstances
// This API is used to query the information of an image instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeBlueprintInstances(request *DescribeBlueprintInstancesRequest) (response *DescribeBlueprintInstancesResponse, err error) {
    return c.DescribeBlueprintInstancesWithContext(context.Background(), request)
}

// DescribeBlueprintInstances
// This API is used to query the information of an image instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeBlueprintInstancesWithContext(ctx context.Context, request *DescribeBlueprintInstancesRequest) (response *DescribeBlueprintInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeBlueprintInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlueprintInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlueprintInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlueprintsRequest() (request *DescribeBlueprintsRequest) {
    request = &DescribeBlueprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBlueprints")
    
    
    return
}

func NewDescribeBlueprintsResponse() (response *DescribeBlueprintsResponse) {
    response = &DescribeBlueprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlueprints
// This API is used to query the information of an image.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTSTATE = "InvalidParameterValue.InvalidBlueprintState"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTTYPE = "InvalidParameterValue.InvalidBlueprintType"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeBlueprints(request *DescribeBlueprintsRequest) (response *DescribeBlueprintsResponse, err error) {
    return c.DescribeBlueprintsWithContext(context.Background(), request)
}

// DescribeBlueprints
// This API is used to query the information of an image.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTSTATE = "InvalidParameterValue.InvalidBlueprintState"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTTYPE = "InvalidParameterValue.InvalidBlueprintType"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeBlueprintsWithContext(ctx context.Context, request *DescribeBlueprintsRequest) (response *DescribeBlueprintsResponse, err error) {
    if request == nil {
        request = NewDescribeBlueprintsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBlueprints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBlueprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBundleDiscountRequest() (request *DescribeBundleDiscountRequest) {
    request = &DescribeBundleDiscountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBundleDiscount")
    
    
    return
}

func NewDescribeBundleDiscountResponse() (response *DescribeBundleDiscountResponse) {
    response = &DescribeBundleDiscountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBundleDiscount
// This API is used to query the discount information of a package.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
func (c *Client) DescribeBundleDiscount(request *DescribeBundleDiscountRequest) (response *DescribeBundleDiscountResponse, err error) {
    return c.DescribeBundleDiscountWithContext(context.Background(), request)
}

// DescribeBundleDiscount
// This API is used to query the discount information of a package.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
func (c *Client) DescribeBundleDiscountWithContext(ctx context.Context, request *DescribeBundleDiscountRequest) (response *DescribeBundleDiscountResponse, err error) {
    if request == nil {
        request = NewDescribeBundleDiscountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBundleDiscount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBundleDiscountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBundlesRequest() (request *DescribeBundlesRequest) {
    request = &DescribeBundlesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBundles")
    
    
    return
}

func NewDescribeBundlesResponse() (response *DescribeBundlesResponse) {
    response = &DescribeBundlesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBundles
// This API is used to query the information of a package.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
func (c *Client) DescribeBundles(request *DescribeBundlesRequest) (response *DescribeBundlesResponse, err error) {
    return c.DescribeBundlesWithContext(context.Background(), request)
}

// DescribeBundles
// This API is used to query the information of a package.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
func (c *Client) DescribeBundlesWithContext(ctx context.Context, request *DescribeBundlesRequest) (response *DescribeBundlesResponse, err error) {
    if request == nil {
        request = NewDescribeBundlesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBundles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBundlesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcnAttachedInstancesRequest() (request *DescribeCcnAttachedInstancesRequest) {
    request = &DescribeCcnAttachedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeCcnAttachedInstances")
    
    
    return
}

func NewDescribeCcnAttachedInstancesResponse() (response *DescribeCcnAttachedInstancesResponse) {
    response = &DescribeCcnAttachedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCcnAttachedInstances
// This API is used to query the information of instances associated with CCN.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DESCRIBECCNATTACHEDINSTANCESFAILED = "UnsupportedOperation.DescribeCcnAttachedInstancesFailed"
func (c *Client) DescribeCcnAttachedInstances(request *DescribeCcnAttachedInstancesRequest) (response *DescribeCcnAttachedInstancesResponse, err error) {
    return c.DescribeCcnAttachedInstancesWithContext(context.Background(), request)
}

// DescribeCcnAttachedInstances
// This API is used to query the information of instances associated with CCN.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DESCRIBECCNATTACHEDINSTANCESFAILED = "UnsupportedOperation.DescribeCcnAttachedInstancesFailed"
func (c *Client) DescribeCcnAttachedInstancesWithContext(ctx context.Context, request *DescribeCcnAttachedInstancesRequest) (response *DescribeCcnAttachedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeCcnAttachedInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCcnAttachedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCcnAttachedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiskConfigsRequest() (request *DescribeDiskConfigsRequest) {
    request = &DescribeDiskConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeDiskConfigs")
    
    
    return
}

func NewDescribeDiskConfigsResponse() (response *DescribeDiskConfigsResponse) {
    response = &DescribeDiskConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskConfigs
// This API is used to query the cloud disk configuration.
//
// error code that may be returned:
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDiskConfigs(request *DescribeDiskConfigsRequest) (response *DescribeDiskConfigsResponse, err error) {
    return c.DescribeDiskConfigsWithContext(context.Background(), request)
}

// DescribeDiskConfigs
// This API is used to query the cloud disk configuration.
//
// error code that may be returned:
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDiskConfigsWithContext(ctx context.Context, request *DescribeDiskConfigsRequest) (response *DescribeDiskConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeDiskConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiskDiscountRequest() (request *DescribeDiskDiscountRequest) {
    request = &DescribeDiskDiscountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeDiskDiscount")
    
    
    return
}

func NewDescribeDiskDiscountResponse() (response *DescribeDiskDiscountResponse) {
    response = &DescribeDiskDiscountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiskDiscount
// This API is used to query the discount information of a cloud disk.
//
// error code that may be returned:
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDiskDiscount(request *DescribeDiskDiscountRequest) (response *DescribeDiskDiscountResponse, err error) {
    return c.DescribeDiskDiscountWithContext(context.Background(), request)
}

// DescribeDiskDiscount
// This API is used to query the discount information of a cloud disk.
//
// error code that may be returned:
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDiskDiscountWithContext(ctx context.Context, request *DescribeDiskDiscountRequest) (response *DescribeDiskDiscountResponse, err error) {
    if request == nil {
        request = NewDescribeDiskDiscountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiskDiscount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiskDiscountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
    request = &DescribeDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeDisks")
    
    
    return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
    response = &DescribeDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDisks
// This API is used to query the information of one or more cloud disks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DISKNAMETOOLONG = "InvalidParameterValue.DiskNameTooLong"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
    return c.DescribeDisksWithContext(context.Background(), request)
}

// DescribeDisks
// This API is used to query the information of one or more cloud disks.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DISKNAMETOOLONG = "InvalidParameterValue.DiskNameTooLong"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
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

func NewDescribeDisksDeniedActionsRequest() (request *DescribeDisksDeniedActionsRequest) {
    request = &DescribeDisksDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeDisksDeniedActions")
    
    
    return
}

func NewDescribeDisksDeniedActionsResponse() (response *DescribeDisksDeniedActionsResponse) {
    response = &DescribeDisksDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDisksDeniedActions
// This API is used to query the list of operation limits of one or more cloud disks.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
func (c *Client) DescribeDisksDeniedActions(request *DescribeDisksDeniedActionsRequest) (response *DescribeDisksDeniedActionsResponse, err error) {
    return c.DescribeDisksDeniedActionsWithContext(context.Background(), request)
}

// DescribeDisksDeniedActions
// This API is used to query the list of operation limits of one or more cloud disks.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
func (c *Client) DescribeDisksDeniedActionsWithContext(ctx context.Context, request *DescribeDisksDeniedActionsRequest) (response *DescribeDisksDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeDisksDeniedActionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisksDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisksDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDisksReturnableRequest() (request *DescribeDisksReturnableRequest) {
    request = &DescribeDisksReturnableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeDisksReturnable")
    
    
    return
}

func NewDescribeDisksReturnableResponse() (response *DescribeDisksReturnableResponse) {
    response = &DescribeDisksReturnableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDisksReturnable
// This API is used to query whether the specified cloud disk can be returned.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEDISKSRETURNABLEERROR = "InternalError.DescribeDisksReturnableError"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
func (c *Client) DescribeDisksReturnable(request *DescribeDisksReturnableRequest) (response *DescribeDisksReturnableResponse, err error) {
    return c.DescribeDisksReturnableWithContext(context.Background(), request)
}

// DescribeDisksReturnable
// This API is used to query whether the specified cloud disk can be returned.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEDISKSRETURNABLEERROR = "InternalError.DescribeDisksReturnableError"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
func (c *Client) DescribeDisksReturnableWithContext(ctx context.Context, request *DescribeDisksReturnableRequest) (response *DescribeDisksReturnableResponse, err error) {
    if request == nil {
        request = NewDescribeDisksReturnableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDisksReturnable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDisksReturnableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirewallRulesRequest() (request *DescribeFirewallRulesRequest) {
    request = &DescribeFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeFirewallRules")
    
    
    return
}

func NewDescribeFirewallRulesResponse() (response *DescribeFirewallRulesResponse) {
    response = &DescribeFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirewallRules
// This API is used to query the firewall rules of an instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeFirewallRules(request *DescribeFirewallRulesRequest) (response *DescribeFirewallRulesResponse, err error) {
    return c.DescribeFirewallRulesWithContext(context.Background(), request)
}

// DescribeFirewallRules
// This API is used to query the firewall rules of an instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeFirewallRulesWithContext(ctx context.Context, request *DescribeFirewallRulesRequest) (response *DescribeFirewallRulesResponse, err error) {
    if request == nil {
        request = NewDescribeFirewallRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFirewallRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirewallRulesTemplateRequest() (request *DescribeFirewallRulesTemplateRequest) {
    request = &DescribeFirewallRulesTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeFirewallRulesTemplate")
    
    
    return
}

func NewDescribeFirewallRulesTemplateResponse() (response *DescribeFirewallRulesTemplateResponse) {
    response = &DescribeFirewallRulesTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirewallRulesTemplate
// This API is used to query a firewall rule template.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeFirewallRulesTemplate(request *DescribeFirewallRulesTemplateRequest) (response *DescribeFirewallRulesTemplateResponse, err error) {
    return c.DescribeFirewallRulesTemplateWithContext(context.Background(), request)
}

// DescribeFirewallRulesTemplate
// This API is used to query a firewall rule template.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeFirewallRulesTemplateWithContext(ctx context.Context, request *DescribeFirewallRulesTemplateRequest) (response *DescribeFirewallRulesTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeFirewallRulesTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFirewallRulesTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFirewallRulesTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralResourceQuotasRequest() (request *DescribeGeneralResourceQuotasRequest) {
    request = &DescribeGeneralResourceQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeGeneralResourceQuotas")
    
    
    return
}

func NewDescribeGeneralResourceQuotasResponse() (response *DescribeGeneralResourceQuotasResponse) {
    response = &DescribeGeneralResourceQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGeneralResourceQuotas
// This API is used to query the quota information of general resources.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDRESOURCEQUOTARESOURCENAME = "InvalidParameterValue.InvalidResourceQuotaResourceName"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeGeneralResourceQuotas(request *DescribeGeneralResourceQuotasRequest) (response *DescribeGeneralResourceQuotasResponse, err error) {
    return c.DescribeGeneralResourceQuotasWithContext(context.Background(), request)
}

// DescribeGeneralResourceQuotas
// This API is used to query the quota information of general resources.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDRESOURCEQUOTARESOURCENAME = "InvalidParameterValue.InvalidResourceQuotaResourceName"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeGeneralResourceQuotasWithContext(ctx context.Context, request *DescribeGeneralResourceQuotasRequest) (response *DescribeGeneralResourceQuotasResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralResourceQuotasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralResourceQuotas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralResourceQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLoginKeyPairAttributeRequest() (request *DescribeInstanceLoginKeyPairAttributeRequest) {
    request = &DescribeInstanceLoginKeyPairAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstanceLoginKeyPairAttribute")
    
    
    return
}

func NewDescribeInstanceLoginKeyPairAttributeResponse() (response *DescribeInstanceLoginKeyPairAttributeResponse) {
    response = &DescribeInstanceLoginKeyPairAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceLoginKeyPairAttribute
// This API is used to query the attributes of the default login key of an instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeInstanceLoginKeyPairAttribute(request *DescribeInstanceLoginKeyPairAttributeRequest) (response *DescribeInstanceLoginKeyPairAttributeResponse, err error) {
    return c.DescribeInstanceLoginKeyPairAttributeWithContext(context.Background(), request)
}

// DescribeInstanceLoginKeyPairAttribute
// This API is used to query the attributes of the default login key of an instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeInstanceLoginKeyPairAttributeWithContext(ctx context.Context, request *DescribeInstanceLoginKeyPairAttributeRequest) (response *DescribeInstanceLoginKeyPairAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLoginKeyPairAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLoginKeyPairAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLoginKeyPairAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceVncUrlRequest() (request *DescribeInstanceVncUrlRequest) {
    request = &DescribeInstanceVncUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstanceVncUrl")
    
    
    return
}

func NewDescribeInstanceVncUrlResponse() (response *DescribeInstanceVncUrlResponse) {
    response = &DescribeInstanceVncUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceVncUrl
// This API is used to query the URL for VNC login.
//
// 
//
// * It does not support `STOPPED` CVMs.
//
// * A VNC URL is only valid for 15 seconds. If you do not access the URL within 15 seconds, you will need to query another one.
//
// * The VNC URL can be used once only. You need to query a new one if you want to log in again.
//
// * Up to 30 re-connection attempts allowed in one minute.
//
// After you get the value of `InstanceVncUrl`, you need to append `InstanceVncUrl=xxxx` to the end of the link `https://img.qcloud.com/qcloud/app/active_vnc/index.html?`.
//
// 
//
//   - `InstanceVncUrl`: Its value will be returned after the API is successfully called.
//
// 
//
//     The final URL can be in the following formats:
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
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
    return c.DescribeInstanceVncUrlWithContext(context.Background(), request)
}

// DescribeInstanceVncUrl
// This API is used to query the URL for VNC login.
//
// 
//
// * It does not support `STOPPED` CVMs.
//
// * A VNC URL is only valid for 15 seconds. If you do not access the URL within 15 seconds, you will need to query another one.
//
// * The VNC URL can be used once only. You need to query a new one if you want to log in again.
//
// * Up to 30 re-connection attempts allowed in one minute.
//
// After you get the value of `InstanceVncUrl`, you need to append `InstanceVncUrl=xxxx` to the end of the link `https://img.qcloud.com/qcloud/app/active_vnc/index.html?`.
//
// 
//
//   - `InstanceVncUrl`: Its value will be returned after the API is successfully called.
//
// 
//
//     The final URL can be in the following formats:
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
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) DescribeInstanceVncUrlWithContext(ctx context.Context, request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceVncUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceVncUrl require credential")
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
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// This API is used to query the details of one or multiple instances.
//
// 
//
// * You can query the details of an instance according to its ID, name, or private IP.
//
// * For more information on filters, please see [Filters](https://intl.cloud.tencent.com/document/product/1207/47576?from_cn_redirect=1#Filter).
//
// * If no parameter is defined, the status of a certain number of instances under the current account will be returned. The number is specified by `Limit` and is 20 by default.
//
// * The latest operation (LatestOperation) and the latest operation status (LatestOperationState) of the instance can be queried.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// This API is used to query the details of one or multiple instances.
//
// 
//
// * You can query the details of an instance according to its ID, name, or private IP.
//
// * For more information on filters, please see [Filters](https://intl.cloud.tencent.com/document/product/1207/47576?from_cn_redirect=1#Filter).
//
// * If no parameter is defined, the status of a certain number of instances under the current account will be returned. The number is specified by `Limit` and is 20 by default.
//
// * The latest operation (LatestOperation) and the latest operation status (LatestOperationState) of the instance can be queried.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesDeniedActionsRequest() (request *DescribeInstancesDeniedActionsRequest) {
    request = &DescribeInstancesDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstancesDeniedActions")
    
    
    return
}

func NewDescribeInstancesDeniedActionsResponse() (response *DescribeInstancesDeniedActionsResponse) {
    response = &DescribeInstancesDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesDeniedActions
// This API is used to query the list of operation limits of one or more instances.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesDeniedActions(request *DescribeInstancesDeniedActionsRequest) (response *DescribeInstancesDeniedActionsResponse, err error) {
    return c.DescribeInstancesDeniedActionsWithContext(context.Background(), request)
}

// DescribeInstancesDeniedActions
// This API is used to query the list of operation limits of one or more instances.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesDeniedActionsWithContext(ctx context.Context, request *DescribeInstancesDeniedActionsRequest) (response *DescribeInstancesDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesDeniedActionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesDiskNumRequest() (request *DescribeInstancesDiskNumRequest) {
    request = &DescribeInstancesDiskNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstancesDiskNum")
    
    
    return
}

func NewDescribeInstancesDiskNumResponse() (response *DescribeInstancesDiskNumResponse) {
    response = &DescribeInstancesDiskNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesDiskNum
// This API is used to query the number of cloud disks attached to instances.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeInstancesDiskNum(request *DescribeInstancesDiskNumRequest) (response *DescribeInstancesDiskNumResponse, err error) {
    return c.DescribeInstancesDiskNumWithContext(context.Background(), request)
}

// DescribeInstancesDiskNum
// This API is used to query the number of cloud disks attached to instances.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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

func NewDescribeInstancesReturnableRequest() (request *DescribeInstancesReturnableRequest) {
    request = &DescribeInstancesReturnableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstancesReturnable")
    
    
    return
}

func NewDescribeInstancesReturnableResponse() (response *DescribeInstancesReturnableResponse) {
    response = &DescribeInstancesReturnableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesReturnable
// This API is used to query whether the specified instance can be returned.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESRETURNABLEERROR = "InternalError.DescribeInstancesReturnableError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesReturnable(request *DescribeInstancesReturnableRequest) (response *DescribeInstancesReturnableResponse, err error) {
    return c.DescribeInstancesReturnableWithContext(context.Background(), request)
}

// DescribeInstancesReturnable
// This API is used to query whether the specified instance can be returned.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESRETURNABLEERROR = "InternalError.DescribeInstancesReturnableError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesReturnableWithContext(ctx context.Context, request *DescribeInstancesReturnableRequest) (response *DescribeInstancesReturnableResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesReturnableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesReturnable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesReturnableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesTrafficPackagesRequest() (request *DescribeInstancesTrafficPackagesRequest) {
    request = &DescribeInstancesTrafficPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstancesTrafficPackages")
    
    
    return
}

func NewDescribeInstancesTrafficPackagesResponse() (response *DescribeInstancesTrafficPackagesResponse) {
    response = &DescribeInstancesTrafficPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesTrafficPackages
// This API is used to query the traffic package details of one or more instances.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTRAFFICPACKAGESFAILED = "InternalError.DescribeInstancesTrafficPackagesFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesTrafficPackages(request *DescribeInstancesTrafficPackagesRequest) (response *DescribeInstancesTrafficPackagesResponse, err error) {
    return c.DescribeInstancesTrafficPackagesWithContext(context.Background(), request)
}

// DescribeInstancesTrafficPackages
// This API is used to query the traffic package details of one or more instances.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTRAFFICPACKAGESFAILED = "InternalError.DescribeInstancesTrafficPackagesFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeInstancesTrafficPackagesWithContext(ctx context.Context, request *DescribeInstancesTrafficPackagesRequest) (response *DescribeInstancesTrafficPackagesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesTrafficPackagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesTrafficPackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesTrafficPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeyPairsRequest() (request *DescribeKeyPairsRequest) {
    request = &DescribeKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeKeyPairs")
    
    
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
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
    return c.DescribeKeyPairsWithContext(context.Background(), request)
}

// DescribeKeyPairs
// This API is used to query key pairs.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeKeyPairsWithContext(ctx context.Context, request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
    if request == nil {
        request = NewDescribeKeyPairsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKeyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModifyInstanceBundlesRequest() (request *DescribeModifyInstanceBundlesRequest) {
    request = &DescribeModifyInstanceBundlesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeModifyInstanceBundles")
    
    
    return
}

func NewDescribeModifyInstanceBundlesResponse() (response *DescribeModifyInstanceBundlesResponse) {
    response = &DescribeModifyInstanceBundlesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModifyInstanceBundles
// This API is used to query the list of package options of an instance.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATION = "InternalError.DescribeInstancesModification"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATIONERROR = "InternalError.DescribeInstancesModificationError"
//  INTERNALERROR_INVALIDBUNDLEPRICE = "InternalError.InvalidBundlePrice"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_BUNDLENOTSUPPORTMODIFY = "OperationDenied.BundleNotSupportModify"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEEXPIRED = "UnsupportedOperation.InstanceExpired"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeModifyInstanceBundles(request *DescribeModifyInstanceBundlesRequest) (response *DescribeModifyInstanceBundlesResponse, err error) {
    return c.DescribeModifyInstanceBundlesWithContext(context.Background(), request)
}

// DescribeModifyInstanceBundles
// This API is used to query the list of package options of an instance.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATION = "InternalError.DescribeInstancesModification"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATIONERROR = "InternalError.DescribeInstancesModificationError"
//  INTERNALERROR_INVALIDBUNDLEPRICE = "InternalError.InvalidBundlePrice"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_BUNDLENOTSUPPORTMODIFY = "OperationDenied.BundleNotSupportModify"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEEXPIRED = "UnsupportedOperation.InstanceExpired"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeModifyInstanceBundlesWithContext(ctx context.Context, request *DescribeModifyInstanceBundlesRequest) (response *DescribeModifyInstanceBundlesResponse, err error) {
    if request == nil {
        request = NewDescribeModifyInstanceBundlesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModifyInstanceBundles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModifyInstanceBundlesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegions
// This API is used to query the information of regions.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATION = "InternalError.DescribeInstancesModification"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATIONERROR = "InternalError.DescribeInstancesModificationError"
//  INTERNALERROR_INVALIDBUNDLEPRICE = "InternalError.InvalidBundlePrice"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_BUNDLENOTSUPPORTMODIFY = "OperationDenied.BundleNotSupportModify"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEEXPIRED = "UnsupportedOperation.InstanceExpired"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// This API is used to query the information of regions.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATION = "InternalError.DescribeInstancesModification"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATIONERROR = "InternalError.DescribeInstancesModificationError"
//  INTERNALERROR_INVALIDBUNDLEPRICE = "InternalError.InvalidBundlePrice"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_BUNDLENOTSUPPORTMODIFY = "OperationDenied.BundleNotSupportModify"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEEXPIRED = "UnsupportedOperation.InstanceExpired"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResetInstanceBlueprintsRequest() (request *DescribeResetInstanceBlueprintsRequest) {
    request = &DescribeResetInstanceBlueprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeResetInstanceBlueprints")
    
    
    return
}

func NewDescribeResetInstanceBlueprintsResponse() (response *DescribeResetInstanceBlueprintsResponse) {
    response = &DescribeResetInstanceBlueprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResetInstanceBlueprints
// This API is used to query the image information of a reset instance.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTSTATE = "InvalidParameterValue.InvalidBlueprintState"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTTYPE = "InvalidParameterValue.InvalidBlueprintType"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeResetInstanceBlueprints(request *DescribeResetInstanceBlueprintsRequest) (response *DescribeResetInstanceBlueprintsResponse, err error) {
    return c.DescribeResetInstanceBlueprintsWithContext(context.Background(), request)
}

// DescribeResetInstanceBlueprints
// This API is used to query the image information of a reset instance.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTSTATE = "InvalidParameterValue.InvalidBlueprintState"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTTYPE = "InvalidParameterValue.InvalidBlueprintType"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeResetInstanceBlueprintsWithContext(ctx context.Context, request *DescribeResetInstanceBlueprintsRequest) (response *DescribeResetInstanceBlueprintsResponse, err error) {
    if request == nil {
        request = NewDescribeResetInstanceBlueprintsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResetInstanceBlueprints require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResetInstanceBlueprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeSnapshots")
    
    
    return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
    response = &DescribeSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshots
// This API is used to query the list of snapshots.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    return c.DescribeSnapshotsWithContext(context.Background(), request)
}

// DescribeSnapshots
// This API is used to query the list of snapshots.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
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

func NewDescribeSnapshotsDeniedActionsRequest() (request *DescribeSnapshotsDeniedActionsRequest) {
    request = &DescribeSnapshotsDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeSnapshotsDeniedActions")
    
    
    return
}

func NewDescribeSnapshotsDeniedActionsResponse() (response *DescribeSnapshotsDeniedActionsResponse) {
    response = &DescribeSnapshotsDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshotsDeniedActions
// This API is used to query the list of operation limits of one or more snapshots.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
func (c *Client) DescribeSnapshotsDeniedActions(request *DescribeSnapshotsDeniedActionsRequest) (response *DescribeSnapshotsDeniedActionsResponse, err error) {
    return c.DescribeSnapshotsDeniedActionsWithContext(context.Background(), request)
}

// DescribeSnapshotsDeniedActions
// This API is used to query the list of operation limits of one or more snapshots.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
func (c *Client) DescribeSnapshotsDeniedActionsWithContext(ctx context.Context, request *DescribeSnapshotsDeniedActionsRequest) (response *DescribeSnapshotsDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsDeniedActionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotsDeniedActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotsDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZones
// This API is used to query the list of AZs in a region.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// This API is used to query the list of AZs in a region.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDetachCcnRequest() (request *DetachCcnRequest) {
    request = &DetachCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DetachCcn")
    
    
    return
}

func NewDetachCcnResponse() (response *DetachCcnResponse) {
    response = &DetachCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachCcn
// This API is used to unassociate a CCN instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_DETACHCCNFAILED = "UnsupportedOperation.DetachCcnFailed"
func (c *Client) DetachCcn(request *DetachCcnRequest) (response *DetachCcnResponse, err error) {
    return c.DetachCcnWithContext(context.Background(), request)
}

// DetachCcn
// This API is used to unassociate a CCN instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_DETACHCCNFAILED = "UnsupportedOperation.DetachCcnFailed"
func (c *Client) DetachCcnWithContext(ctx context.Context, request *DetachCcnRequest) (response *DetachCcnResponse, err error) {
    if request == nil {
        request = NewDetachCcnRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachCcn require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachCcnResponse()
    err = c.Send(request, response)
    return
}

func NewDetachDisksRequest() (request *DetachDisksRequest) {
    request = &DetachDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DetachDisks")
    
    
    return
}

func NewDetachDisksResponse() (response *DetachDisksResponse) {
    response = &DetachDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachDisks
// This API is used to detach one or more cloud disks.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) DetachDisks(request *DetachDisksRequest) (response *DetachDisksResponse, err error) {
    return c.DetachDisksWithContext(context.Background(), request)
}

// DetachDisks
// This API is used to detach one or more cloud disks.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
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

func NewDisassociateInstancesKeyPairsRequest() (request *DisassociateInstancesKeyPairsRequest) {
    request = &DisassociateInstancesKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DisassociateInstancesKeyPairs")
    
    
    return
}

func NewDisassociateInstancesKeyPairsResponse() (response *DisassociateInstancesKeyPairsResponse) {
    response = &DisassociateInstancesKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateInstancesKeyPairs
// This API is used to unbind an instance from the specified key pair.
//
// 
//
// * Only instances on LINUX_UNIX in [RUNNING, STOPPED] status are supported. Instances in `RUNNING` status will be forcibly shut down before unbinding.
//
// * After a key pair is unassociated from an instance, you can log in to the instance with password.
//
// * If no password was set, you cannot log in to the instance with SSH after unbinding. You can call the ResetInstancesPassword API to set a login password.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_KEYPAIRNOTBOUNDTOINSTANCE = "UnsupportedOperation.KeyPairNotBoundToInstance"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) DisassociateInstancesKeyPairs(request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
    return c.DisassociateInstancesKeyPairsWithContext(context.Background(), request)
}

// DisassociateInstancesKeyPairs
// This API is used to unbind an instance from the specified key pair.
//
// 
//
// * Only instances on LINUX_UNIX in [RUNNING, STOPPED] status are supported. Instances in `RUNNING` status will be forcibly shut down before unbinding.
//
// * After a key pair is unassociated from an instance, you can log in to the instance with password.
//
// * If no password was set, you cannot log in to the instance with SSH after unbinding. You can call the ResetInstancesPassword API to set a login password.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_KEYPAIRNOTBOUNDTOINSTANCE = "UnsupportedOperation.KeyPairNotBoundToInstance"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) DisassociateInstancesKeyPairsWithContext(ctx context.Context, request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewDisassociateInstancesKeyPairsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateInstancesKeyPairs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewImportKeyPairRequest() (request *ImportKeyPairRequest) {
    request = &ImportKeyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ImportKeyPair")
    
    
    return
}

func NewImportKeyPairResponse() (response *ImportKeyPairResponse) {
    response = &ImportKeyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportKeyPair
// This API is used to import the specified key pair.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEKEYPAIRFAILED = "FailedOperation.CreateKeyPairFailed"
//  FAILEDOPERATION_DELETEKEYPAIRFAILED = "FailedOperation.DeleteKeyPairFailed"
//  FAILEDOPERATION_IMPORTKEYPAIRFAILED = "FailedOperation.ImportKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEPARAMETERVALUE = "InvalidParameterValue.DuplicateParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEEMPTY = "InvalidParameterValue.InvalidKeyPairNameEmpty"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMETOOLONG = "InvalidParameterValue.InvalidKeyPairNameTooLong"
//  INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYDUPLICATED = "InvalidParameterValue.KeyPairPublicKeyDuplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYMALFORMED = "InvalidParameterValue.KeyPairPublicKeyMalformed"
//  LIMITEXCEEDED_KEYPAIRLIMITEXCEEDED = "LimitExceeded.KeyPairLimitExceeded"
func (c *Client) ImportKeyPair(request *ImportKeyPairRequest) (response *ImportKeyPairResponse, err error) {
    return c.ImportKeyPairWithContext(context.Background(), request)
}

// ImportKeyPair
// This API is used to import the specified key pair.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEKEYPAIRFAILED = "FailedOperation.CreateKeyPairFailed"
//  FAILEDOPERATION_DELETEKEYPAIRFAILED = "FailedOperation.DeleteKeyPairFailed"
//  FAILEDOPERATION_IMPORTKEYPAIRFAILED = "FailedOperation.ImportKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEPARAMETERVALUE = "InvalidParameterValue.DuplicateParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEEMPTY = "InvalidParameterValue.InvalidKeyPairNameEmpty"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMETOOLONG = "InvalidParameterValue.InvalidKeyPairNameTooLong"
//  INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYDUPLICATED = "InvalidParameterValue.KeyPairPublicKeyDuplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYMALFORMED = "InvalidParameterValue.KeyPairPublicKeyMalformed"
//  LIMITEXCEEDED_KEYPAIRLIMITEXCEEDED = "LimitExceeded.KeyPairLimitExceeded"
func (c *Client) ImportKeyPairWithContext(ctx context.Context, request *ImportKeyPairRequest) (response *ImportKeyPairResponse, err error) {
    if request == nil {
        request = NewImportKeyPairRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportKeyPair require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportKeyPairResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateBlueprintRequest() (request *InquirePriceCreateBlueprintRequest) {
    request = &InquirePriceCreateBlueprintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceCreateBlueprint")
    
    
    return
}

func NewInquirePriceCreateBlueprintResponse() (response *InquirePriceCreateBlueprintResponse) {
    response = &InquirePriceCreateBlueprintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceCreateBlueprint
// This API is used to query the price of a created image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) InquirePriceCreateBlueprint(request *InquirePriceCreateBlueprintRequest) (response *InquirePriceCreateBlueprintResponse, err error) {
    return c.InquirePriceCreateBlueprintWithContext(context.Background(), request)
}

// InquirePriceCreateBlueprint
// This API is used to query the price of a created image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) InquirePriceCreateBlueprintWithContext(ctx context.Context, request *InquirePriceCreateBlueprintRequest) (response *InquirePriceCreateBlueprintResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateBlueprintRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateBlueprint require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateBlueprintResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateDisksRequest() (request *InquirePriceCreateDisksRequest) {
    request = &InquirePriceCreateDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceCreateDisks")
    
    
    return
}

func NewInquirePriceCreateDisksResponse() (response *InquirePriceCreateDisksResponse) {
    response = &InquirePriceCreateDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceCreateDisks
// This API is used to query the price of purchasing cloud disks.
//
// error code that may be returned:
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
func (c *Client) InquirePriceCreateDisks(request *InquirePriceCreateDisksRequest) (response *InquirePriceCreateDisksResponse, err error) {
    return c.InquirePriceCreateDisksWithContext(context.Background(), request)
}

// InquirePriceCreateDisks
// This API is used to query the price of purchasing cloud disks.
//
// error code that may be returned:
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
func (c *Client) InquirePriceCreateDisksWithContext(ctx context.Context, request *InquirePriceCreateDisksRequest) (response *InquirePriceCreateDisksResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateDisksResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateInstancesRequest() (request *InquirePriceCreateInstancesRequest) {
    request = &InquirePriceCreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceCreateInstances")
    
    
    return
}

func NewInquirePriceCreateInstancesResponse() (response *InquirePriceCreateInstancesResponse) {
    response = &InquirePriceCreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceCreateInstances
// This API is used to query the price of a created instance.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTID = "InvalidParameterValue.InvalidBlueprintId"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) InquirePriceCreateInstances(request *InquirePriceCreateInstancesRequest) (response *InquirePriceCreateInstancesResponse, err error) {
    return c.InquirePriceCreateInstancesWithContext(context.Background(), request)
}

// InquirePriceCreateInstances
// This API is used to query the price of a created instance.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTID = "InvalidParameterValue.InvalidBlueprintId"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) InquirePriceCreateInstancesWithContext(ctx context.Context, request *InquirePriceCreateInstancesRequest) (response *InquirePriceCreateInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRenewDisksRequest() (request *InquirePriceRenewDisksRequest) {
    request = &InquirePriceRenewDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceRenewDisks")
    
    
    return
}

func NewInquirePriceRenewDisksResponse() (response *InquirePriceRenewDisksResponse) {
    response = &InquirePriceRenewDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceRenewDisks
// This API is used to query the price of renewing cloud disks.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
func (c *Client) InquirePriceRenewDisks(request *InquirePriceRenewDisksRequest) (response *InquirePriceRenewDisksResponse, err error) {
    return c.InquirePriceRenewDisksWithContext(context.Background(), request)
}

// InquirePriceRenewDisks
// This API is used to query the price of renewing cloud disks.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
func (c *Client) InquirePriceRenewDisksWithContext(ctx context.Context, request *InquirePriceRenewDisksRequest) (response *InquirePriceRenewDisksResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewDisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceRenewDisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceRenewDisksResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRenewInstancesRequest() (request *InquirePriceRenewInstancesRequest) {
    request = &InquirePriceRenewInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceRenewInstances")
    
    
    return
}

func NewInquirePriceRenewInstancesResponse() (response *InquirePriceRenewInstancesResponse) {
    response = &InquirePriceRenewInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceRenewInstances
// This API is used to query the price of renewed instance.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEDATADISKNOTFOUND = "ResourceNotFound.InstanceDataDiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) InquirePriceRenewInstances(request *InquirePriceRenewInstancesRequest) (response *InquirePriceRenewInstancesResponse, err error) {
    return c.InquirePriceRenewInstancesWithContext(context.Background(), request)
}

// InquirePriceRenewInstances
// This API is used to query the price of renewed instance.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEDATADISKNOTFOUND = "ResourceNotFound.InstanceDataDiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) InquirePriceRenewInstancesWithContext(ctx context.Context, request *InquirePriceRenewInstancesRequest) (response *InquirePriceRenewInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceRenewInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceRenewInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlueprintAttributeRequest() (request *ModifyBlueprintAttributeRequest) {
    request = &ModifyBlueprintAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyBlueprintAttribute")
    
    
    return
}

func NewModifyBlueprintAttributeResponse() (response *ModifyBlueprintAttributeResponse) {
    response = &ModifyBlueprintAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBlueprintAttribute
// This API is used to modify an image attribute.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  UNSUPPORTEDOPERATION_BLUEPRINTCURSTATEINVALID = "UnsupportedOperation.BlueprintCurStateInvalid"
//  UNSUPPORTEDOPERATION_NOTSUPPORTSHAREDBLUEPRINT = "UnsupportedOperation.NotSupportSharedBlueprint"
func (c *Client) ModifyBlueprintAttribute(request *ModifyBlueprintAttributeRequest) (response *ModifyBlueprintAttributeResponse, err error) {
    return c.ModifyBlueprintAttributeWithContext(context.Background(), request)
}

// ModifyBlueprintAttribute
// This API is used to modify an image attribute.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  UNSUPPORTEDOPERATION_BLUEPRINTCURSTATEINVALID = "UnsupportedOperation.BlueprintCurStateInvalid"
//  UNSUPPORTEDOPERATION_NOTSUPPORTSHAREDBLUEPRINT = "UnsupportedOperation.NotSupportSharedBlueprint"
func (c *Client) ModifyBlueprintAttributeWithContext(ctx context.Context, request *ModifyBlueprintAttributeRequest) (response *ModifyBlueprintAttributeResponse, err error) {
    if request == nil {
        request = NewModifyBlueprintAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBlueprintAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBlueprintAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDisksAttributeRequest() (request *ModifyDisksAttributeRequest) {
    request = &ModifyDisksAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyDisksAttribute")
    
    
    return
}

func NewModifyDisksAttributeResponse() (response *ModifyDisksAttributeResponse) {
    response = &ModifyDisksAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDisksAttribute
// This API is used to modify cloud disk attributes.
//
// error code that may be returned:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
func (c *Client) ModifyDisksAttribute(request *ModifyDisksAttributeRequest) (response *ModifyDisksAttributeResponse, err error) {
    return c.ModifyDisksAttributeWithContext(context.Background(), request)
}

// ModifyDisksAttribute
// This API is used to modify cloud disk attributes.
//
// error code that may be returned:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
func (c *Client) ModifyDisksAttributeWithContext(ctx context.Context, request *ModifyDisksAttributeRequest) (response *ModifyDisksAttributeResponse, err error) {
    if request == nil {
        request = NewModifyDisksAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDisksAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDisksAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDisksRenewFlagRequest() (request *ModifyDisksRenewFlagRequest) {
    request = &ModifyDisksRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyDisksRenewFlag")
    
    
    return
}

func NewModifyDisksRenewFlagResponse() (response *ModifyDisksRenewFlagResponse) {
    response = &ModifyDisksRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDisksRenewFlag
// This API is used to modify the configuration of auto-renewal of cloud disks.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
func (c *Client) ModifyDisksRenewFlag(request *ModifyDisksRenewFlagRequest) (response *ModifyDisksRenewFlagResponse, err error) {
    return c.ModifyDisksRenewFlagWithContext(context.Background(), request)
}

// ModifyDisksRenewFlag
// This API is used to modify the configuration of auto-renewal of cloud disks.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
func (c *Client) ModifyDisksRenewFlagWithContext(ctx context.Context, request *ModifyDisksRenewFlagRequest) (response *ModifyDisksRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyDisksRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDisksRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDisksRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFirewallRuleDescriptionRequest() (request *ModifyFirewallRuleDescriptionRequest) {
    request = &ModifyFirewallRuleDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyFirewallRuleDescription")
    
    
    return
}

func NewModifyFirewallRuleDescriptionResponse() (response *ModifyFirewallRuleDescriptionResponse) {
    response = &ModifyFirewallRuleDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFirewallRuleDescription
// This API is used to modify the description of a single firewall rule.
//
// 
//
// * `FirewallVersion` is used to specify the version of the firewall to be manipulated. If the `FirewallVersion` value passed in is not equal to the current latest version of the firewall, a failure will be returned. If `FirewallVersion` is not passed in, the description of the specified rule will be modified directly.
//
// 
//
// In the `FirewallRule` parameter:
//
// * Valid values of the `Protocol` field include `TCP`, `UDP`, `ICMP`, and `ALL`.
//
// * For the `Port` field, you can enter only `ALL`, a single port number, several port numbers separated by commas, or a port range indicated by two port numbers separated by a minus sign. If `Port` is a range, the port number on the left of the minus sign must be smaller than the one on the right. If `Protocol` is not `TCP` or `UDP`, `Port` can only be empty or `ALL`. The length of the `Port` field cannot exceed 64 characters.
//
// * For the `CidrBlock` field, you can enter any string that conforms to the CIDR format standard. Multi-Tenant network isolation rules take precedence over private network rules in the firewall.
//
// * For the `Action` field, you can enter only `ACCEPT` or `DROP`.
//
// * The length of the `FirewallRuleDescription` field cannot exceed 64 characters.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_FIREWALLRULESNOTFOUND = "ResourceNotFound.FirewallRulesNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) ModifyFirewallRuleDescription(request *ModifyFirewallRuleDescriptionRequest) (response *ModifyFirewallRuleDescriptionResponse, err error) {
    return c.ModifyFirewallRuleDescriptionWithContext(context.Background(), request)
}

// ModifyFirewallRuleDescription
// This API is used to modify the description of a single firewall rule.
//
// 
//
// * `FirewallVersion` is used to specify the version of the firewall to be manipulated. If the `FirewallVersion` value passed in is not equal to the current latest version of the firewall, a failure will be returned. If `FirewallVersion` is not passed in, the description of the specified rule will be modified directly.
//
// 
//
// In the `FirewallRule` parameter:
//
// * Valid values of the `Protocol` field include `TCP`, `UDP`, `ICMP`, and `ALL`.
//
// * For the `Port` field, you can enter only `ALL`, a single port number, several port numbers separated by commas, or a port range indicated by two port numbers separated by a minus sign. If `Port` is a range, the port number on the left of the minus sign must be smaller than the one on the right. If `Protocol` is not `TCP` or `UDP`, `Port` can only be empty or `ALL`. The length of the `Port` field cannot exceed 64 characters.
//
// * For the `CidrBlock` field, you can enter any string that conforms to the CIDR format standard. Multi-Tenant network isolation rules take precedence over private network rules in the firewall.
//
// * For the `Action` field, you can enter only `ACCEPT` or `DROP`.
//
// * The length of the `FirewallRuleDescription` field cannot exceed 64 characters.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_FIREWALLRULESNOTFOUND = "ResourceNotFound.FirewallRulesNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) ModifyFirewallRuleDescriptionWithContext(ctx context.Context, request *ModifyFirewallRuleDescriptionRequest) (response *ModifyFirewallRuleDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyFirewallRuleDescriptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFirewallRuleDescription require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFirewallRuleDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFirewallRulesRequest() (request *ModifyFirewallRulesRequest) {
    request = &ModifyFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyFirewallRules")
    
    
    return
}

func NewModifyFirewallRulesResponse() (response *ModifyFirewallRulesResponse) {
    response = &ModifyFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFirewallRules
// This API is used to reset the firewall rules of an instance.
//
// 
//
// This API deletes all firewall rules of the current instance first and then adds new rules.
//
// 
//
// * `FirewallVersion` is used to specify the version of the firewall to be manipulated. If the `FirewallVersion` value passed in is not equal to the current latest version of the firewall, a failure will be returned. If `FirewallVersion` is not passed in, the specified rule will be reset directly.
//
// 
//
// In the `FirewallRules` parameter:
//
// * Valid values of the `Protocol` field include `TCP`, `UDP`, `ICMP`, and `ALL`.
//
// * For the `Port` field, you can enter only `ALL`, a single port number, several port numbers separated by commas, or a port range indicated by two port numbers separated by a minus sign. If `Port` is a range, the port number on the left of the minus sign must be smaller than the one on the right. If `Protocol` is not `TCP` or `UDP`, `Port` can only be empty or `ALL`. The length of the `Port` field cannot exceed 64 characters.
//
// * For the `CidrBlock` field, you can enter any string that conforms to the CIDR format standard. Multi-Tenant network isolation rules take precedence over private network rules in the firewall.
//
// * For the `Action` field, you can enter only `ACCEPT` or `DROP`.
//
// * The length of the `FirewallRuleDescription` field cannot exceed 64 characters.
//
// error code that may be returned:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) ModifyFirewallRules(request *ModifyFirewallRulesRequest) (response *ModifyFirewallRulesResponse, err error) {
    return c.ModifyFirewallRulesWithContext(context.Background(), request)
}

// ModifyFirewallRules
// This API is used to reset the firewall rules of an instance.
//
// 
//
// This API deletes all firewall rules of the current instance first and then adds new rules.
//
// 
//
// * `FirewallVersion` is used to specify the version of the firewall to be manipulated. If the `FirewallVersion` value passed in is not equal to the current latest version of the firewall, a failure will be returned. If `FirewallVersion` is not passed in, the specified rule will be reset directly.
//
// 
//
// In the `FirewallRules` parameter:
//
// * Valid values of the `Protocol` field include `TCP`, `UDP`, `ICMP`, and `ALL`.
//
// * For the `Port` field, you can enter only `ALL`, a single port number, several port numbers separated by commas, or a port range indicated by two port numbers separated by a minus sign. If `Port` is a range, the port number on the left of the minus sign must be smaller than the one on the right. If `Protocol` is not `TCP` or `UDP`, `Port` can only be empty or `ALL`. The length of the `Port` field cannot exceed 64 characters.
//
// * For the `CidrBlock` field, you can enter any string that conforms to the CIDR format standard. Multi-Tenant network isolation rules take precedence over private network rules in the firewall.
//
// * For the `Action` field, you can enter only `ACCEPT` or `DROP`.
//
// * The length of the `FirewallRuleDescription` field cannot exceed 64 characters.
//
// error code that may be returned:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) ModifyFirewallRulesWithContext(ctx context.Context, request *ModifyFirewallRulesRequest) (response *ModifyFirewallRulesResponse, err error) {
    if request == nil {
        request = NewModifyFirewallRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFirewallRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
    request = &ModifyInstancesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyInstancesAttribute")
    
    
    return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
    response = &ModifyInstancesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesAttribute
// This API is used to modify the attributes of instances.
//
// * The instance name is used only for users’ convenience.
//
// * Batch operations are supported. Each request can contain up to 100 instances at a time.
//
// * This API is async. A successful request will return a `RequestId`, it does not mean the operation is completed. You can call the `DescribeInstances` API to query the operation result. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETER_ONLYALLOWMODIFYONEATTRIBUTE = "InvalidParameter.OnlyAllowModifyOneAttribute"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    return c.ModifyInstancesAttributeWithContext(context.Background(), request)
}

// ModifyInstancesAttribute
// This API is used to modify the attributes of instances.
//
// * The instance name is used only for users’ convenience.
//
// * Batch operations are supported. Each request can contain up to 100 instances at a time.
//
// * This API is async. A successful request will return a `RequestId`, it does not mean the operation is completed. You can call the `DescribeInstances` API to query the operation result. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETER_ONLYALLOWMODIFYONEATTRIBUTE = "InvalidParameter.OnlyAllowModifyOneAttribute"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesAttributeWithContext(ctx context.Context, request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancesAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesLoginKeyPairAttributeRequest() (request *ModifyInstancesLoginKeyPairAttributeRequest) {
    request = &ModifyInstancesLoginKeyPairAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyInstancesLoginKeyPairAttribute")
    
    
    return
}

func NewModifyInstancesLoginKeyPairAttributeResponse() (response *ModifyInstancesLoginKeyPairAttributeResponse) {
    response = &ModifyInstancesLoginKeyPairAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesLoginKeyPairAttribute
// This API is used to set the attributes of the default login key pair of an instance.
//
// 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCELOGINKEYPAIRPERMITLOGIN = "InvalidParameterValue.InvalidInstanceLoginKeyPairPermitLogin"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesLoginKeyPairAttribute(request *ModifyInstancesLoginKeyPairAttributeRequest) (response *ModifyInstancesLoginKeyPairAttributeResponse, err error) {
    return c.ModifyInstancesLoginKeyPairAttributeWithContext(context.Background(), request)
}

// ModifyInstancesLoginKeyPairAttribute
// This API is used to set the attributes of the default login key pair of an instance.
//
// 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCELOGINKEYPAIRPERMITLOGIN = "InvalidParameterValue.InvalidInstanceLoginKeyPairPermitLogin"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesLoginKeyPairAttributeWithContext(ctx context.Context, request *ModifyInstancesLoginKeyPairAttributeRequest) (response *ModifyInstancesLoginKeyPairAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesLoginKeyPairAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancesLoginKeyPairAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancesLoginKeyPairAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesRenewFlagRequest() (request *ModifyInstancesRenewFlagRequest) {
    request = &ModifyInstancesRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyInstancesRenewFlag")
    
    
    return
}

func NewModifyInstancesRenewFlagResponse() (response *ModifyInstancesRenewFlagResponse) {
    response = &ModifyInstancesRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesRenewFlag
// This API is used to modify the renewal flags of monthly subscribed instances.
//
// 
//
// * Instances marked with "auto-renewal" will be automatically renewed for one month when they expire.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesRenewFlag(request *ModifyInstancesRenewFlagRequest) (response *ModifyInstancesRenewFlagResponse, err error) {
    return c.ModifyInstancesRenewFlagWithContext(context.Background(), request)
}

// ModifyInstancesRenewFlag
// This API is used to modify the renewal flags of monthly subscribed instances.
//
// 
//
// * Instances marked with "auto-renewal" will be automatically renewed for one month when they expire.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesRenewFlagWithContext(ctx context.Context, request *ModifyInstancesRenewFlagRequest) (response *ModifyInstancesRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyInstancesRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancesRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancesRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotAttributeRequest() (request *ModifySnapshotAttributeRequest) {
    request = &ModifySnapshotAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifySnapshotAttribute")
    
    
    return
}

func NewModifySnapshotAttributeResponse() (response *ModifySnapshotAttributeResponse) {
    response = &ModifySnapshotAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySnapshotAttribute
// This API is used to modify the attributes of a snapshot.
//
// <li>The snapshot name is used only for users’ convenience.</li>
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
func (c *Client) ModifySnapshotAttribute(request *ModifySnapshotAttributeRequest) (response *ModifySnapshotAttributeResponse, err error) {
    return c.ModifySnapshotAttributeWithContext(context.Background(), request)
}

// ModifySnapshotAttribute
// This API is used to modify the attributes of a snapshot.
//
// <li>The snapshot name is used only for users’ convenience.</li>
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
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

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
    request = &RebootInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "RebootInstances")
    
    
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
// * The instance status will become `REBOOTING` when the API is called successfully and will become `RUNNING` when the instance is successfully restarted.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    return c.RebootInstancesWithContext(context.Background(), request)
}

// RebootInstances
// This API is used to restart instances.
//
// 
//
// * You can only perform this operation on instances whose status is `RUNNING`.
//
// * The instance status will become `REBOOTING` when the API is called successfully and will become `RUNNING` when the instance is successfully restarted.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) RebootInstancesWithContext(ctx context.Context, request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RebootInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetAttachCcnRequest() (request *ResetAttachCcnRequest) {
    request = &ResetAttachCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ResetAttachCcn")
    
    
    return
}

func NewResetAttachCcnResponse() (response *ResetAttachCcnResponse) {
    response = &ResetAttachCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetAttachCcn
// This API is used to apply for association again after a CCN instance association application expires.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_RESETATTACHCCNFAILED = "UnsupportedOperation.ResetAttachCcnFailed"
func (c *Client) ResetAttachCcn(request *ResetAttachCcnRequest) (response *ResetAttachCcnResponse, err error) {
    return c.ResetAttachCcnWithContext(context.Background(), request)
}

// ResetAttachCcn
// This API is used to apply for association again after a CCN instance association application expires.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_RESETATTACHCCNFAILED = "UnsupportedOperation.ResetAttachCcnFailed"
func (c *Client) ResetAttachCcnWithContext(ctx context.Context, request *ResetAttachCcnRequest) (response *ResetAttachCcnResponse, err error) {
    if request == nil {
        request = NewResetAttachCcnRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetAttachCcn require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetAttachCcnResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstanceRequest() (request *ResetInstanceRequest) {
    request = &ResetInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ResetInstance")
    
    
    return
}

func NewResetInstanceResponse() (response *ResetInstanceResponse) {
    response = &ResetInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstance
// This API is used to reinstall the image on the specified instance.
//
// 
//
// * If you specify a `BlueprintId`, the specified image is used; otherwise, the image used by the current instance is used.
//
// * The system disk will be formatted and reset. Therefore, make sure that no important files are stored on the system disk.
//
// * Currently, this API does not support switching the operating system between `LINUX_UNIX` and `WINDOWS` for instances.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTCONFIGNOTMATCH = "InvalidParameterValue.BlueprintConfigNotMatch"
//  INVALIDPARAMETERVALUE_BLUEPRINTID = "InvalidParameterValue.BlueprintId"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTALLOWTOCHANGEPLATFORMTYPE = "InvalidParameterValue.NotAllowToChangePlatformType"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_BLUEPRINTUNAVAILABLE = "ResourceUnavailable.BlueprintUnavailable"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_SYSTEMBUSY = "UnsupportedOperation.SystemBusy"
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
    return c.ResetInstanceWithContext(context.Background(), request)
}

// ResetInstance
// This API is used to reinstall the image on the specified instance.
//
// 
//
// * If you specify a `BlueprintId`, the specified image is used; otherwise, the image used by the current instance is used.
//
// * The system disk will be formatted and reset. Therefore, make sure that no important files are stored on the system disk.
//
// * Currently, this API does not support switching the operating system between `LINUX_UNIX` and `WINDOWS` for instances.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTCONFIGNOTMATCH = "InvalidParameterValue.BlueprintConfigNotMatch"
//  INVALIDPARAMETERVALUE_BLUEPRINTID = "InvalidParameterValue.BlueprintId"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTALLOWTOCHANGEPLATFORMTYPE = "InvalidParameterValue.NotAllowToChangePlatformType"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_BLUEPRINTUNAVAILABLE = "ResourceUnavailable.BlueprintUnavailable"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_SYSTEMBUSY = "UnsupportedOperation.SystemBusy"
func (c *Client) ResetInstanceWithContext(ctx context.Context, request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
    if request == nil {
        request = NewResetInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesPasswordRequest() (request *ResetInstancesPasswordRequest) {
    request = &ResetInstancesPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ResetInstancesPassword")
    
    
    return
}

func NewResetInstancesPasswordResponse() (response *ResetInstancesPasswordResponse) {
    response = &ResetInstancesPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstancesPassword
// This API is used to reset the password of the instance OS to a user-specified password.
//
// * You can only use this API to modify the password of the admin account. The name of the admin account varies by OS (on Windows, it is `Administrator`; on Ubuntu, it is `ubuntu`; on other systems, it is `root`).
//
// * Batch operations are supported. You can reset the passwords of multiple instances to the same one. The maximum number of instances in each request is 100.
//
// * It’s recommended to shut down the instance first and then reset the password. If the instance is running, this API will try to shut it down normally. If the attempt fails, it will force to instance to shut down.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// Note: Just like powering off a physical PC, a forced shutdown may cause data loss or the corruption of file system.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTALLOWTOCHANGEPLATFORMTYPE = "InvalidParameterValue.NotAllowToChangePlatformType"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    return c.ResetInstancesPasswordWithContext(context.Background(), request)
}

// ResetInstancesPassword
// This API is used to reset the password of the instance OS to a user-specified password.
//
// * You can only use this API to modify the password of the admin account. The name of the admin account varies by OS (on Windows, it is `Administrator`; on Ubuntu, it is `ubuntu`; on other systems, it is `root`).
//
// * Batch operations are supported. You can reset the passwords of multiple instances to the same one. The maximum number of instances in each request is 100.
//
// * It’s recommended to shut down the instance first and then reset the password. If the instance is running, this API will try to shut it down normally. If the attempt fails, it will force to instance to shut down.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// Note: Just like powering off a physical PC, a forced shutdown may cause data loss or the corruption of file system.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTALLOWTOCHANGEPLATFORMTYPE = "InvalidParameterValue.NotAllowToChangePlatformType"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ResetInstancesPasswordWithContext(ctx context.Context, request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    if request == nil {
        request = NewResetInstancesPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetInstancesPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetInstancesPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
    request = &StartInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "StartInstances")
    
    
    return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
    response = &StartInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartInstances
// This API is used to start one or more instances.
//
// 
//
// * You can only perform this operation on instances whose status is `STOPPED`.
//
// * The instance status will become `STARTING` when the API is called successfully and will become `RUNNING` when the instance is successfully started.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    return c.StartInstancesWithContext(context.Background(), request)
}

// StartInstances
// This API is used to start one or more instances.
//
// 
//
// * You can only perform this operation on instances whose status is `STOPPED`.
//
// * The instance status will become `STARTING` when the API is called successfully and will become `RUNNING` when the instance is successfully started.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) StartInstancesWithContext(ctx context.Context, request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    if request == nil {
        request = NewStartInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartInstances require credential")
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
    request.Init().WithApiInfo("lighthouse", APIVersion, "StopInstances")
    
    
    return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
    response = &StopInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopInstances
// This API is used to shut down one or more instances.
//
// * You can only perform this operation on instances whose status is `RUNNING`.
//
// * The instance status will become `STOPPING` when the API is called successfully and will become `STOPPED` when the instance is successfully shut down.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    return c.StopInstancesWithContext(context.Background(), request)
}

// StopInstances
// This API is used to shut down one or more instances.
//
// * You can only perform this operation on instances whose status is `RUNNING`.
//
// * The instance status will become `STOPPING` when the API is called successfully and will become `STOPPED` when the instance is successfully shut down.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) StopInstancesWithContext(ctx context.Context, request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDisksRequest() (request *TerminateDisksRequest) {
    request = &TerminateDisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "TerminateDisks")
    
    
    return
}

func NewTerminateDisksResponse() (response *TerminateDisksResponse) {
    response = &TerminateDisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateDisks
// This API is used to terminate one or more cloud disk.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
func (c *Client) TerminateDisks(request *TerminateDisksRequest) (response *TerminateDisksResponse, err error) {
    return c.TerminateDisksWithContext(context.Background(), request)
}

// TerminateDisks
// This API is used to terminate one or more cloud disk.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"
//  OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"
//  RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"
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

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "TerminateInstances")
    
    
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateInstances
// This API is used to terminate one or more instances.
//
// 
//
// * Instances in `SHUTDOWN` status can be terminated through this API and cannot be recovered.
//
// * Batch operations are supported. The allowed maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    return c.TerminateInstancesWithContext(context.Background(), request)
}

// TerminateInstances
// This API is used to terminate one or more instances.
//
// 
//
// * Instances in `SHUTDOWN` status can be terminated through this API and cannot be recovered.
//
// * Batch operations are supported. The allowed maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) TerminateInstancesWithContext(ctx context.Context, request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
