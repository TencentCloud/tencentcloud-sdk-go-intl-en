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

package v20220817

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-08-17"

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


func NewAddGroupMemberRequest() (request *AddGroupMemberRequest) {
    request = &AddGroupMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "AddGroupMember")
    
    
    return
}

func NewAddGroupMemberResponse() (response *AddGroupMemberResponse) {
    response = &AddGroupMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddGroupMember
// This API is used to add users to a group.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddGroupMember(request *AddGroupMemberRequest) (response *AddGroupMemberResponse, err error) {
    return c.AddGroupMemberWithContext(context.Background(), request)
}

// AddGroupMember
// This API is used to add users to a group.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddGroupMemberWithContext(ctx context.Context, request *AddGroupMemberRequest) (response *AddGroupMemberResponse, err error) {
    if request == nil {
        request = NewAddGroupMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddGroupMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddGroupMemberResponse()
    err = c.Send(request, response)
    return
}

func NewBatchAddGroupMemberRequest() (request *BatchAddGroupMemberRequest) {
    request = &BatchAddGroupMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchAddGroupMember")
    
    
    return
}

func NewBatchAddGroupMemberResponse() (response *BatchAddGroupMemberResponse) {
    response = &BatchAddGroupMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchAddGroupMember
// This API is used to add users to multiple groups at a time.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchAddGroupMember(request *BatchAddGroupMemberRequest) (response *BatchAddGroupMemberResponse, err error) {
    return c.BatchAddGroupMemberWithContext(context.Background(), request)
}

// BatchAddGroupMember
// This API is used to add users to multiple groups at a time.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchAddGroupMemberWithContext(ctx context.Context, request *BatchAddGroupMemberRequest) (response *BatchAddGroupMemberResponse, err error) {
    if request == nil {
        request = NewBatchAddGroupMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchAddGroupMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchAddGroupMemberResponse()
    err = c.Send(request, response)
    return
}

func NewBatchCreateGroupWithMembersRequest() (request *BatchCreateGroupWithMembersRequest) {
    request = &BatchCreateGroupWithMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchCreateGroupWithMembers")
    
    
    return
}

func NewBatchCreateGroupWithMembersResponse() (response *BatchCreateGroupWithMembersResponse) {
    response = &BatchCreateGroupWithMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchCreateGroupWithMembers
// This API is used to create multiple groups at a time.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchCreateGroupWithMembers(request *BatchCreateGroupWithMembersRequest) (response *BatchCreateGroupWithMembersResponse, err error) {
    return c.BatchCreateGroupWithMembersWithContext(context.Background(), request)
}

// BatchCreateGroupWithMembers
// This API is used to create multiple groups at a time.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchCreateGroupWithMembersWithContext(ctx context.Context, request *BatchCreateGroupWithMembersRequest) (response *BatchCreateGroupWithMembersResponse, err error) {
    if request == nil {
        request = NewBatchCreateGroupWithMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchCreateGroupWithMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchCreateGroupWithMembersResponse()
    err = c.Send(request, response)
    return
}

func NewBatchCreateRoomRequest() (request *BatchCreateRoomRequest) {
    request = &BatchCreateRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchCreateRoom")
    
    
    return
}

func NewBatchCreateRoomResponse() (response *BatchCreateRoomResponse) {
    response = &BatchCreateRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchCreateRoom
// This API is used to create multiple rooms at a time.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchCreateRoom(request *BatchCreateRoomRequest) (response *BatchCreateRoomResponse, err error) {
    return c.BatchCreateRoomWithContext(context.Background(), request)
}

// BatchCreateRoom
// This API is used to create multiple rooms at a time.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BatchCreateRoomWithContext(ctx context.Context, request *BatchCreateRoomRequest) (response *BatchCreateRoomResponse, err error) {
    if request == nil {
        request = NewBatchCreateRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchCreateRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchCreateRoomResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteGroupMemberRequest() (request *BatchDeleteGroupMemberRequest) {
    request = &BatchDeleteGroupMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchDeleteGroupMember")
    
    
    return
}

func NewBatchDeleteGroupMemberResponse() (response *BatchDeleteGroupMemberResponse) {
    response = &BatchDeleteGroupMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeleteGroupMember
// This API is used to remove users from multiple groups at a time.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) BatchDeleteGroupMember(request *BatchDeleteGroupMemberRequest) (response *BatchDeleteGroupMemberResponse, err error) {
    return c.BatchDeleteGroupMemberWithContext(context.Background(), request)
}

// BatchDeleteGroupMember
// This API is used to remove users from multiple groups at a time.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) BatchDeleteGroupMemberWithContext(ctx context.Context, request *BatchDeleteGroupMemberRequest) (response *BatchDeleteGroupMemberResponse, err error) {
    if request == nil {
        request = NewBatchDeleteGroupMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteGroupMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteGroupMemberResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteRecordRequest() (request *BatchDeleteRecordRequest) {
    request = &BatchDeleteRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchDeleteRecord")
    
    
    return
}

func NewBatchDeleteRecordResponse() (response *BatchDeleteRecordResponse) {
    response = &BatchDeleteRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeleteRecord
// This API is used to delete the recordings of multiple rooms.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) BatchDeleteRecord(request *BatchDeleteRecordRequest) (response *BatchDeleteRecordResponse, err error) {
    return c.BatchDeleteRecordWithContext(context.Background(), request)
}

// BatchDeleteRecord
// This API is used to delete the recordings of multiple rooms.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) BatchDeleteRecordWithContext(ctx context.Context, request *BatchDeleteRecordRequest) (response *BatchDeleteRecordResponse, err error) {
    if request == nil {
        request = NewBatchDeleteRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteRecordResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDescribeDocumentRequest() (request *BatchDescribeDocumentRequest) {
    request = &BatchDescribeDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchDescribeDocument")
    
    
    return
}

func NewBatchDescribeDocumentResponse() (response *BatchDescribeDocumentResponse) {
    response = &BatchDescribeDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDescribeDocument
// This API is used to get courseware information.
//
// error code that may be returned:
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) BatchDescribeDocument(request *BatchDescribeDocumentRequest) (response *BatchDescribeDocumentResponse, err error) {
    return c.BatchDescribeDocumentWithContext(context.Background(), request)
}

// BatchDescribeDocument
// This API is used to get courseware information.
//
// error code that may be returned:
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) BatchDescribeDocumentWithContext(ctx context.Context, request *BatchDescribeDocumentRequest) (response *BatchDescribeDocumentResponse, err error) {
    if request == nil {
        request = NewBatchDescribeDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDescribeDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDescribeDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewBatchRegisterRequest() (request *BatchRegisterRequest) {
    request = &BatchRegisterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BatchRegister")
    
    
    return
}

func NewBatchRegisterResponse() (response *BatchRegisterResponse) {
    response = &BatchRegisterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchRegister
// This API is used to register multiple users (up to 1,000) at a time. If a user ID already exists, the existing one will be overwritten.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchRegister(request *BatchRegisterRequest) (response *BatchRegisterResponse, err error) {
    return c.BatchRegisterWithContext(context.Background(), request)
}

// BatchRegister
// This API is used to register multiple users (up to 1,000) at a time. If a user ID already exists, the existing one will be overwritten.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BatchRegisterWithContext(ctx context.Context, request *BatchRegisterRequest) (response *BatchRegisterResponse, err error) {
    if request == nil {
        request = NewBatchRegisterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchRegister require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchRegisterResponse()
    err = c.Send(request, response)
    return
}

func NewBindDocumentToRoomRequest() (request *BindDocumentToRoomRequest) {
    request = &BindDocumentToRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "BindDocumentToRoom")
    
    
    return
}

func NewBindDocumentToRoomResponse() (response *BindDocumentToRoomResponse) {
    response = &BindDocumentToRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindDocumentToRoom
// This API is used to bind a document to a room.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) BindDocumentToRoom(request *BindDocumentToRoomRequest) (response *BindDocumentToRoomResponse, err error) {
    return c.BindDocumentToRoomWithContext(context.Background(), request)
}

// BindDocumentToRoom
// This API is used to bind a document to a room.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) BindDocumentToRoomWithContext(ctx context.Context, request *BindDocumentToRoomRequest) (response *BindDocumentToRoomResponse, err error) {
    if request == nil {
        request = NewBindDocumentToRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindDocumentToRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindDocumentToRoomResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDocumentRequest() (request *CreateDocumentRequest) {
    request = &CreateDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateDocument")
    
    
    return
}

func NewCreateDocumentResponse() (response *CreateDocumentResponse) {
    response = &CreateDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDocument
// This API is used to create a document to be used in a room.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateDocument(request *CreateDocumentRequest) (response *CreateDocumentResponse, err error) {
    return c.CreateDocumentWithContext(context.Background(), request)
}

// CreateDocument
// This API is used to create a document to be used in a room.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateDocumentWithContext(ctx context.Context, request *CreateDocumentRequest) (response *CreateDocumentResponse, err error) {
    if request == nil {
        request = NewCreateDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupWithMembersRequest() (request *CreateGroupWithMembersRequest) {
    request = &CreateGroupWithMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateGroupWithMembers")
    
    
    return
}

func NewCreateGroupWithMembersResponse() (response *CreateGroupWithMembersResponse) {
    response = &CreateGroupWithMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGroupWithMembers
// his API is used to create a group and specify its members.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGroupWithMembers(request *CreateGroupWithMembersRequest) (response *CreateGroupWithMembersResponse, err error) {
    return c.CreateGroupWithMembersWithContext(context.Background(), request)
}

// CreateGroupWithMembers
// his API is used to create a group and specify its members.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGroupWithMembersWithContext(ctx context.Context, request *CreateGroupWithMembersRequest) (response *CreateGroupWithMembersResponse, err error) {
    if request == nil {
        request = NewCreateGroupWithMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGroupWithMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGroupWithMembersResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupWithSubGroupRequest() (request *CreateGroupWithSubGroupRequest) {
    request = &CreateGroupWithSubGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateGroupWithSubGroup")
    
    
    return
}

func NewCreateGroupWithSubGroupResponse() (response *CreateGroupWithSubGroupResponse) {
    response = &CreateGroupWithSubGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGroupWithSubGroup
// This API is used to merge groups.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGroupWithSubGroup(request *CreateGroupWithSubGroupRequest) (response *CreateGroupWithSubGroupResponse, err error) {
    return c.CreateGroupWithSubGroupWithContext(context.Background(), request)
}

// CreateGroupWithSubGroup
// This API is used to merge groups.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateGroupWithSubGroupWithContext(ctx context.Context, request *CreateGroupWithSubGroupRequest) (response *CreateGroupWithSubGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupWithSubGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGroupWithSubGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGroupWithSubGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoomRequest() (request *CreateRoomRequest) {
    request = &CreateRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateRoom")
    
    
    return
}

func NewCreateRoomResponse() (response *CreateRoomResponse) {
    response = &CreateRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRoom
// This API is used to create a room.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ROOMTYPEINVALID = "InvalidParameter.RoomTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateRoom(request *CreateRoomRequest) (response *CreateRoomResponse, err error) {
    return c.CreateRoomWithContext(context.Background(), request)
}

// CreateRoom
// This API is used to create a room.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_ROOMTYPEINVALID = "InvalidParameter.RoomTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateRoomWithContext(ctx context.Context, request *CreateRoomRequest) (response *CreateRoomResponse, err error) {
    if request == nil {
        request = NewCreateRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoomResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSupervisorRequest() (request *CreateSupervisorRequest) {
    request = &CreateSupervisorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "CreateSupervisor")
    
    
    return
}

func NewCreateSupervisorResponse() (response *CreateSupervisorResponse) {
    response = &CreateSupervisorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSupervisor
// This API is used to create a spectator.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateSupervisor(request *CreateSupervisorRequest) (response *CreateSupervisorResponse, err error) {
    return c.CreateSupervisorWithContext(context.Background(), request)
}

// CreateSupervisor
// This API is used to create a spectator.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) CreateSupervisorWithContext(ctx context.Context, request *CreateSupervisorRequest) (response *CreateSupervisorResponse, err error) {
    if request == nil {
        request = NewCreateSupervisorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSupervisor require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSupervisorResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppCustomContentRequest() (request *DeleteAppCustomContentRequest) {
    request = &DeleteAppCustomContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteAppCustomContent")
    
    
    return
}

func NewDeleteAppCustomContentResponse() (response *DeleteAppCustomContentResponse) {
    response = &DeleteAppCustomContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAppCustomContent
// This API is used to delete the custom elements. The `Scenes` parameter specifies the custom elements to delete. If `Scenes` is empty, all custom elements will be deleted.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DeleteAppCustomContent(request *DeleteAppCustomContentRequest) (response *DeleteAppCustomContentResponse, err error) {
    return c.DeleteAppCustomContentWithContext(context.Background(), request)
}

// DeleteAppCustomContent
// This API is used to delete the custom elements. The `Scenes` parameter specifies the custom elements to delete. If `Scenes` is empty, all custom elements will be deleted.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DeleteAppCustomContentWithContext(ctx context.Context, request *DeleteAppCustomContentRequest) (response *DeleteAppCustomContentResponse, err error) {
    if request == nil {
        request = NewDeleteAppCustomContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAppCustomContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppCustomContentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDocumentRequest() (request *DeleteDocumentRequest) {
    request = &DeleteDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteDocument")
    
    
    return
}

func NewDeleteDocumentResponse() (response *DeleteDocumentResponse) {
    response = &DeleteDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDocument
// This API is used to delete a document.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
func (c *Client) DeleteDocument(request *DeleteDocumentRequest) (response *DeleteDocumentResponse, err error) {
    return c.DeleteDocumentWithContext(context.Background(), request)
}

// DeleteDocument
// This API is used to delete a document.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
func (c *Client) DeleteDocumentWithContext(ctx context.Context, request *DeleteDocumentRequest) (response *DeleteDocumentResponse, err error) {
    if request == nil {
        request = NewDeleteDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteGroup")
    
    
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGroup
// This API is used to delete one or multiple groups.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

// DeleteGroup
// This API is used to delete one or multiple groups.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupMemberRequest() (request *DeleteGroupMemberRequest) {
    request = &DeleteGroupMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteGroupMember")
    
    
    return
}

func NewDeleteGroupMemberResponse() (response *DeleteGroupMemberResponse) {
    response = &DeleteGroupMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGroupMember
// This API is used to remove users from a group.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteGroupMember(request *DeleteGroupMemberRequest) (response *DeleteGroupMemberResponse, err error) {
    return c.DeleteGroupMemberWithContext(context.Background(), request)
}

// DeleteGroupMember
// This API is used to remove users from a group.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTEACHERSNOTEXIST = "InvalidParameter.GroupTeachersNotExist"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteGroupMemberWithContext(ctx context.Context, request *DeleteGroupMemberRequest) (response *DeleteGroupMemberResponse, err error) {
    if request == nil {
        request = NewDeleteGroupMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroupMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupMemberResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordRequest() (request *DeleteRecordRequest) {
    request = &DeleteRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteRecord")
    
    
    return
}

func NewDeleteRecordResponse() (response *DeleteRecordResponse) {
    response = &DeleteRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecord
// This example shows you how to delete the recording files of a specific room.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRecord(request *DeleteRecordRequest) (response *DeleteRecordResponse, err error) {
    return c.DeleteRecordWithContext(context.Background(), request)
}

// DeleteRecord
// This example shows you how to delete the recording files of a specific room.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRecordWithContext(ctx context.Context, request *DeleteRecordRequest) (response *DeleteRecordResponse, err error) {
    if request == nil {
        request = NewDeleteRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoomRequest() (request *DeleteRoomRequest) {
    request = &DeleteRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteRoom")
    
    
    return
}

func NewDeleteRoomResponse() (response *DeleteRoomResponse) {
    response = &DeleteRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoom
// This API is used to delete a room.
//
// error code that may be returned:
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DeleteRoom(request *DeleteRoomRequest) (response *DeleteRoomResponse, err error) {
    return c.DeleteRoomWithContext(context.Background(), request)
}

// DeleteRoom
// This API is used to delete a room.
//
// error code that may be returned:
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DeleteRoomWithContext(ctx context.Context, request *DeleteRoomRequest) (response *DeleteRoomResponse, err error) {
    if request == nil {
        request = NewDeleteRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSupervisorRequest() (request *DeleteSupervisorRequest) {
    request = &DeleteSupervisorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DeleteSupervisor")
    
    
    return
}

func NewDeleteSupervisorResponse() (response *DeleteSupervisorResponse) {
    response = &DeleteSupervisorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSupervisor
// This API is used to delete spectators.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DeleteSupervisor(request *DeleteSupervisorRequest) (response *DeleteSupervisorResponse, err error) {
    return c.DeleteSupervisorWithContext(context.Background(), request)
}

// DeleteSupervisor
// This API is used to delete spectators.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DeleteSupervisorWithContext(ctx context.Context, request *DeleteSupervisorRequest) (response *DeleteSupervisorResponse, err error) {
    if request == nil {
        request = NewDeleteSupervisorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSupervisor require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSupervisorResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAnswerListRequest() (request *DescribeAnswerListRequest) {
    request = &DescribeAnswerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeAnswerList")
    
    
    return
}

func NewDescribeAnswerListResponse() (response *DescribeAnswerListResponse) {
    response = &DescribeAnswerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAnswerList
// This API is used to get the answers to a quiz question in a room.
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAnswerList(request *DescribeAnswerListRequest) (response *DescribeAnswerListResponse, err error) {
    return c.DescribeAnswerListWithContext(context.Background(), request)
}

// DescribeAnswerList
// This API is used to get the answers to a quiz question in a room.
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAnswerListWithContext(ctx context.Context, request *DescribeAnswerListRequest) (response *DescribeAnswerListResponse, err error) {
    if request == nil {
        request = NewDescribeAnswerListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAnswerList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAnswerListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCurrentMemberListRequest() (request *DescribeCurrentMemberListRequest) {
    request = &DescribeCurrentMemberListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeCurrentMemberList")
    
    
    return
}

func NewDescribeCurrentMemberListResponse() (response *DescribeCurrentMemberListResponse) {
    response = &DescribeCurrentMemberListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCurrentMemberList
// This API is used to get the user list of a room. This API will not work if a room has ended or expired.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeCurrentMemberList(request *DescribeCurrentMemberListRequest) (response *DescribeCurrentMemberListResponse, err error) {
    return c.DescribeCurrentMemberListWithContext(context.Background(), request)
}

// DescribeCurrentMemberList
// This API is used to get the user list of a room. This API will not work if a room has ended or expired.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeCurrentMemberListWithContext(ctx context.Context, request *DescribeCurrentMemberListRequest) (response *DescribeCurrentMemberListResponse, err error) {
    if request == nil {
        request = NewDescribeCurrentMemberListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCurrentMemberList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCurrentMemberListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeveloperRequest() (request *DescribeDeveloperRequest) {
    request = &DescribeDeveloperRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeDeveloper")
    
    
    return
}

func NewDescribeDeveloperResponse() (response *DescribeDeveloperResponse) {
    response = &DescribeDeveloperResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeveloper
// This API is used to get developer information.
//
// error code that may be returned:
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeveloper(request *DescribeDeveloperRequest) (response *DescribeDeveloperResponse, err error) {
    return c.DescribeDeveloperWithContext(context.Background(), request)
}

// DescribeDeveloper
// This API is used to get developer information.
//
// error code that may be returned:
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeveloperWithContext(ctx context.Context, request *DescribeDeveloperRequest) (response *DescribeDeveloperResponse, err error) {
    if request == nil {
        request = NewDescribeDeveloperRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeveloper require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeveloperResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDocumentRequest() (request *DescribeDocumentRequest) {
    request = &DescribeDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeDocument")
    
    
    return
}

func NewDescribeDocumentResponse() (response *DescribeDocumentResponse) {
    response = &DescribeDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDocument
// This API is used to get the information of a specific document.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
func (c *Client) DescribeDocument(request *DescribeDocumentRequest) (response *DescribeDocumentResponse, err error) {
    return c.DescribeDocumentWithContext(context.Background(), request)
}

// DescribeDocument
// This API is used to get the information of a specific document.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
func (c *Client) DescribeDocumentWithContext(ctx context.Context, request *DescribeDocumentRequest) (response *DescribeDocumentResponse, err error) {
    if request == nil {
        request = NewDescribeDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDocumentsRequest() (request *DescribeDocumentsRequest) {
    request = &DescribeDocumentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeDocuments")
    
    
    return
}

func NewDescribeDocumentsResponse() (response *DescribeDocumentsResponse) {
    response = &DescribeDocumentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDocuments
// A new API is offered for this action now.
//
// 
//
// This API is used to query courseware. It has been deprecated. Please use `BatchDescribeDocument` instead.
//
// error code that may be returned:
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeDocuments(request *DescribeDocumentsRequest) (response *DescribeDocumentsResponse, err error) {
    return c.DescribeDocumentsWithContext(context.Background(), request)
}

// DescribeDocuments
// A new API is offered for this action now.
//
// 
//
// This API is used to query courseware. It has been deprecated. Please use `BatchDescribeDocument` instead.
//
// error code that may be returned:
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeDocumentsWithContext(ctx context.Context, request *DescribeDocumentsRequest) (response *DescribeDocumentsResponse, err error) {
    if request == nil {
        request = NewDescribeDocumentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDocuments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocumentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDocumentsByRoomRequest() (request *DescribeDocumentsByRoomRequest) {
    request = &DescribeDocumentsByRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeDocumentsByRoom")
    
    
    return
}

func NewDescribeDocumentsByRoomResponse() (response *DescribeDocumentsByRoomResponse) {
    response = &DescribeDocumentsByRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDocumentsByRoom
// This API is used to get the document list of a specific room.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDocumentsByRoom(request *DescribeDocumentsByRoomRequest) (response *DescribeDocumentsByRoomResponse, err error) {
    return c.DescribeDocumentsByRoomWithContext(context.Background(), request)
}

// DescribeDocumentsByRoom
// This API is used to get the document list of a specific room.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDocumentsByRoomWithContext(ctx context.Context, request *DescribeDocumentsByRoomRequest) (response *DescribeDocumentsByRoomResponse, err error) {
    if request == nil {
        request = NewDescribeDocumentsByRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDocumentsByRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocumentsByRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupRequest() (request *DescribeGroupRequest) {
    request = &DescribeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeGroup")
    
    
    return
}

func NewDescribeGroupResponse() (response *DescribeGroupResponse) {
    response = &DescribeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroup
// This API is used to get the details of a group.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroup(request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    return c.DescribeGroupWithContext(context.Background(), request)
}

// DescribeGroup
// This API is used to get the details of a group.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupWithContext(ctx context.Context, request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupListRequest() (request *DescribeGroupListRequest) {
    request = &DescribeGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeGroupList")
    
    
    return
}

func NewDescribeGroupListResponse() (response *DescribeGroupListResponse) {
    response = &DescribeGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupList
// This API is used to query groups.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupList(request *DescribeGroupListRequest) (response *DescribeGroupListResponse, err error) {
    return c.DescribeGroupListWithContext(context.Background(), request)
}

// DescribeGroupList
// This API is used to query groups.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPMEMBEROVERLIMIT = "InvalidParameter.GroupMemberOverLimit"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupListWithContext(ctx context.Context, request *DescribeGroupListRequest) (response *DescribeGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupMemberListRequest() (request *DescribeGroupMemberListRequest) {
    request = &DescribeGroupMemberListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeGroupMemberList")
    
    
    return
}

func NewDescribeGroupMemberListResponse() (response *DescribeGroupMemberListResponse) {
    response = &DescribeGroupMemberListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGroupMemberList
// This API is used to get the members of a group.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupMemberList(request *DescribeGroupMemberListRequest) (response *DescribeGroupMemberListResponse, err error) {
    return c.DescribeGroupMemberListWithContext(context.Background(), request)
}

// DescribeGroupMemberList
// This API is used to get the members of a group.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeGroupMemberListWithContext(ctx context.Context, request *DescribeGroupMemberListRequest) (response *DescribeGroupMemberListResponse, err error) {
    if request == nil {
        request = NewDescribeGroupMemberListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupMemberList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupMemberListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuestionListRequest() (request *DescribeQuestionListRequest) {
    request = &DescribeQuestionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeQuestionList")
    
    
    return
}

func NewDescribeQuestionListResponse() (response *DescribeQuestionListResponse) {
    response = &DescribeQuestionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQuestionList
// This API is used to get the quiz details of a room.
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuestionList(request *DescribeQuestionListRequest) (response *DescribeQuestionListResponse, err error) {
    return c.DescribeQuestionListWithContext(context.Background(), request)
}

// DescribeQuestionList
// This API is used to get the quiz details of a room.
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeQuestionListWithContext(ctx context.Context, request *DescribeQuestionListRequest) (response *DescribeQuestionListResponse, err error) {
    if request == nil {
        request = NewDescribeQuestionListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQuestionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQuestionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomRequest() (request *DescribeRoomRequest) {
    request = &DescribeRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeRoom")
    
    
    return
}

func NewDescribeRoomResponse() (response *DescribeRoomResponse) {
    response = &DescribeRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoom
// This API is used to get room information.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeRoom(request *DescribeRoomRequest) (response *DescribeRoomResponse, err error) {
    return c.DescribeRoomWithContext(context.Background(), request)
}

// DescribeRoom
// This API is used to get room information.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) DescribeRoomWithContext(ctx context.Context, request *DescribeRoomRequest) (response *DescribeRoomResponse, err error) {
    if request == nil {
        request = NewDescribeRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomStatisticsRequest() (request *DescribeRoomStatisticsRequest) {
    request = &DescribeRoomStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeRoomStatistics")
    
    
    return
}

func NewDescribeRoomStatisticsResponse() (response *DescribeRoomStatisticsResponse) {
    response = &DescribeRoomStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoomStatistics
// This API is used to obtain the statistics of a room. It can be called only after the room is ended.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
func (c *Client) DescribeRoomStatistics(request *DescribeRoomStatisticsRequest) (response *DescribeRoomStatisticsResponse, err error) {
    return c.DescribeRoomStatisticsWithContext(context.Background(), request)
}

// DescribeRoomStatistics
// This API is used to obtain the statistics of a room. It can be called only after the room is ended.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
func (c *Client) DescribeRoomStatisticsWithContext(ctx context.Context, request *DescribeRoomStatisticsRequest) (response *DescribeRoomStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeRoomStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoomStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSdkAppIdUsersRequest() (request *DescribeSdkAppIdUsersRequest) {
    request = &DescribeSdkAppIdUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeSdkAppIdUsers")
    
    
    return
}

func NewDescribeSdkAppIdUsersResponse() (response *DescribeSdkAppIdUsersResponse) {
    response = &DescribeSdkAppIdUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSdkAppIdUsers
// This API is used to get the user list of a specific application.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSdkAppIdUsers(request *DescribeSdkAppIdUsersRequest) (response *DescribeSdkAppIdUsersResponse, err error) {
    return c.DescribeSdkAppIdUsersWithContext(context.Background(), request)
}

// DescribeSdkAppIdUsers
// This API is used to get the user list of a specific application.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSdkAppIdUsersWithContext(ctx context.Context, request *DescribeSdkAppIdUsersRequest) (response *DescribeSdkAppIdUsersResponse, err error) {
    if request == nil {
        request = NewDescribeSdkAppIdUsersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSdkAppIdUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSdkAppIdUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSupervisorsRequest() (request *DescribeSupervisorsRequest) {
    request = &DescribeSupervisorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeSupervisors")
    
    
    return
}

func NewDescribeSupervisorsResponse() (response *DescribeSupervisorsResponse) {
    response = &DescribeSupervisorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSupervisors
// This API is used to get the spectators of a room.
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSupervisors(request *DescribeSupervisorsRequest) (response *DescribeSupervisorsResponse, err error) {
    return c.DescribeSupervisorsWithContext(context.Background(), request)
}

// DescribeSupervisors
// This API is used to get the spectators of a room.
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSupervisorsWithContext(ctx context.Context, request *DescribeSupervisorsRequest) (response *DescribeSupervisorsResponse, err error) {
    if request == nil {
        request = NewDescribeSupervisorsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSupervisors require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSupervisorsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserRequest() (request *DescribeUserRequest) {
    request = &DescribeUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "DescribeUser")
    
    
    return
}

func NewDescribeUserResponse() (response *DescribeUserResponse) {
    response = &DescribeUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUser
// This API is used to obtain user profile.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DescribeUser(request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    return c.DescribeUserWithContext(context.Background(), request)
}

// DescribeUser
// This API is used to obtain user profile.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) DescribeUserWithContext(ctx context.Context, request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    if request == nil {
        request = NewDescribeUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResponse()
    err = c.Send(request, response)
    return
}

func NewEndRoomRequest() (request *EndRoomRequest) {
    request = &EndRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "EndRoom")
    
    
    return
}

func NewEndRoomResponse() (response *EndRoomResponse) {
    response = &EndRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EndRoom
// This API is used to end a room.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EndRoom(request *EndRoomRequest) (response *EndRoomResponse, err error) {
    return c.EndRoomWithContext(context.Background(), request)
}

// EndRoom
// This API is used to end a room.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EndRoomWithContext(ctx context.Context, request *EndRoomRequest) (response *EndRoomResponse, err error) {
    if request == nil {
        request = NewEndRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EndRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewEndRoomResponse()
    err = c.Send(request, response)
    return
}

func NewGetRoomEventRequest() (request *GetRoomEventRequest) {
    request = &GetRoomEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "GetRoomEvent")
    
    
    return
}

func NewGetRoomEventResponse() (response *GetRoomEventResponse) {
    response = &GetRoomEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetRoomEvent
// This API is used to get the events of a room. It only works within one hour after a class ends.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) GetRoomEvent(request *GetRoomEventRequest) (response *GetRoomEventResponse, err error) {
    return c.GetRoomEventWithContext(context.Background(), request)
}

// GetRoomEvent
// This API is used to get the events of a room. It only works within one hour after a class ends.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) GetRoomEventWithContext(ctx context.Context, request *GetRoomEventRequest) (response *GetRoomEventResponse, err error) {
    if request == nil {
        request = NewGetRoomEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRoomEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRoomEventResponse()
    err = c.Send(request, response)
    return
}

func NewGetRoomMessageRequest() (request *GetRoomMessageRequest) {
    request = &GetRoomMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "GetRoomMessage")
    
    
    return
}

func NewGetRoomMessageResponse() (response *GetRoomMessageResponse) {
    response = &GetRoomMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetRoomMessage
// This API is used to get the message history of a room (room messages are retained for seven days).
//
// error code that may be returned:
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRoomMessage(request *GetRoomMessageRequest) (response *GetRoomMessageResponse, err error) {
    return c.GetRoomMessageWithContext(context.Background(), request)
}

// GetRoomMessage
// This API is used to get the message history of a room (room messages are retained for seven days).
//
// error code that may be returned:
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRoomMessageWithContext(ctx context.Context, request *GetRoomMessageRequest) (response *GetRoomMessageResponse, err error) {
    if request == nil {
        request = NewGetRoomMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRoomMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRoomMessageResponse()
    err = c.Send(request, response)
    return
}

func NewGetRoomsRequest() (request *GetRoomsRequest) {
    request = &GetRoomsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "GetRooms")
    
    
    return
}

func NewGetRoomsResponse() (response *GetRoomsResponse) {
    response = &GetRoomsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetRooms
// This API is used to get the room list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRooms(request *GetRoomsRequest) (response *GetRoomsResponse, err error) {
    return c.GetRoomsWithContext(context.Background(), request)
}

// GetRooms
// This API is used to get the room list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRoomsWithContext(ctx context.Context, request *GetRoomsRequest) (response *GetRoomsResponse, err error) {
    if request == nil {
        request = NewGetRoomsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRooms require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRoomsResponse()
    err = c.Send(request, response)
    return
}

func NewGetWatermarkRequest() (request *GetWatermarkRequest) {
    request = &GetWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "GetWatermark")
    
    
    return
}

func NewGetWatermarkResponse() (response *GetWatermarkResponse) {
    response = &GetWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWatermark
// This API is used to get watermark settings.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) GetWatermark(request *GetWatermarkRequest) (response *GetWatermarkResponse, err error) {
    return c.GetWatermarkWithContext(context.Background(), request)
}

// GetWatermark
// This API is used to get watermark settings.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) GetWatermarkWithContext(ctx context.Context, request *GetWatermarkRequest) (response *GetWatermarkResponse, err error) {
    if request == nil {
        request = NewGetWatermarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewKickUserFromRoomRequest() (request *KickUserFromRoomRequest) {
    request = &KickUserFromRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "KickUserFromRoom")
    
    
    return
}

func NewKickUserFromRoomResponse() (response *KickUserFromRoomResponse) {
    response = &KickUserFromRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KickUserFromRoom
// This API is used to remove a user from the room.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) KickUserFromRoom(request *KickUserFromRoomRequest) (response *KickUserFromRoomResponse, err error) {
    return c.KickUserFromRoomWithContext(context.Background(), request)
}

// KickUserFromRoom
// This API is used to remove a user from the room.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) KickUserFromRoomWithContext(ctx context.Context, request *KickUserFromRoomRequest) (response *KickUserFromRoomResponse, err error) {
    if request == nil {
        request = NewKickUserFromRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("KickUserFromRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewKickUserFromRoomResponse()
    err = c.Send(request, response)
    return
}

func NewLoginOriginIdRequest() (request *LoginOriginIdRequest) {
    request = &LoginOriginIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "LoginOriginId")
    
    
    return
}

func NewLoginOriginIdResponse() (response *LoginOriginIdResponse) {
    response = &LoginOriginIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// LoginOriginId
// This API is used to log in with an origin account, which is the `originId` entered during registration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) LoginOriginId(request *LoginOriginIdRequest) (response *LoginOriginIdResponse, err error) {
    return c.LoginOriginIdWithContext(context.Background(), request)
}

// LoginOriginId
// This API is used to log in with an origin account, which is the `originId` entered during registration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) LoginOriginIdWithContext(ctx context.Context, request *LoginOriginIdRequest) (response *LoginOriginIdResponse, err error) {
    if request == nil {
        request = NewLoginOriginIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LoginOriginId require credential")
    }

    request.SetContext(ctx)
    
    response = NewLoginOriginIdResponse()
    err = c.Send(request, response)
    return
}

func NewLoginUserRequest() (request *LoginUserRequest) {
    request = &LoginUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "LoginUser")
    
    
    return
}

func NewLoginUserResponse() (response *LoginUserResponse) {
    response = &LoginUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// LoginUser
// This API is used to log in.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) LoginUser(request *LoginUserRequest) (response *LoginUserResponse, err error) {
    return c.LoginUserWithContext(context.Background(), request)
}

// LoginUser
// This API is used to log in.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
func (c *Client) LoginUserWithContext(ctx context.Context, request *LoginUserRequest) (response *LoginUserResponse, err error) {
    if request == nil {
        request = NewLoginUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LoginUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewLoginUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppRequest() (request *ModifyAppRequest) {
    request = &ModifyAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "ModifyApp")
    
    
    return
}

func NewModifyAppResponse() (response *ModifyAppResponse) {
    response = &ModifyAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApp
// This API is used to modify an application.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) ModifyApp(request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    return c.ModifyAppWithContext(context.Background(), request)
}

// ModifyApp
// This API is used to modify an application.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) ModifyAppWithContext(ctx context.Context, request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    if request == nil {
        request = NewModifyAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupRequest() (request *ModifyGroupRequest) {
    request = &ModifyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "ModifyGroup")
    
    
    return
}

func NewModifyGroupResponse() (response *ModifyGroupResponse) {
    response = &ModifyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGroup
// This API is used to modify a group.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyGroup(request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    return c.ModifyGroupWithContext(context.Background(), request)
}

// ModifyGroup
// This API is used to modify a group.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_GROUPPARAMINVALID = "InvalidParameter.GroupParamInvalid"
//  INVALIDPARAMETER_GROUPTEACHERNOTMEMBER = "InvalidParameter.GroupTeacherNotMember"
//  INVALIDPARAMETER_GROUPTYPEINVALID = "InvalidParameter.GroupTypeInvalid"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
//  RESOURCENOTFOUND_GROUPPARTUSERSNOTEXIST = "ResourceNotFound.GroupPartUsersNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyGroupWithContext(ctx context.Context, request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    if request == nil {
        request = NewModifyGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoomRequest() (request *ModifyRoomRequest) {
    request = &ModifyRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "ModifyRoom")
    
    
    return
}

func NewModifyRoomResponse() (response *ModifyRoomResponse) {
    response = &ModifyRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRoom
// This API is used to modify a room.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyRoom(request *ModifyRoomRequest) (response *ModifyRoomResponse, err error) {
    return c.ModifyRoomWithContext(context.Background(), request)
}

// ModifyRoom
// This API is used to modify a room.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyRoomWithContext(ctx context.Context, request *ModifyRoomRequest) (response *ModifyRoomResponse, err error) {
    if request == nil {
        request = NewModifyRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRoomResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserProfileRequest() (request *ModifyUserProfileRequest) {
    request = &ModifyUserProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "ModifyUserProfile")
    
    
    return
}

func NewModifyUserProfileResponse() (response *ModifyUserProfileResponse) {
    response = &ModifyUserProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserProfile
// This API is used to modify a user profile such as the nickname and profile photo.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserProfile(request *ModifyUserProfileRequest) (response *ModifyUserProfileResponse, err error) {
    return c.ModifyUserProfileWithContext(context.Background(), request)
}

// ModifyUserProfile
// This API is used to modify a user profile such as the nickname and profile photo.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_CLASSTOOLONG = "FailedOperation.ClassTooLong"
//  FAILEDOPERATION_IMAGEARGINVALID = "FailedOperation.ImageArgInvalid"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_RECORD = "ResourceInsufficient.Record"
//  RESOURCEINSUFFICIENT_ROOM = "ResourceInsufficient.Room"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_ROOMSTATISTICS = "ResourceUnavailable.RoomStatistics"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserProfileWithContext(ctx context.Context, request *ModifyUserProfileRequest) (response *ModifyUserProfileResponse, err error) {
    if request == nil {
        request = NewModifyUserProfileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserProfileResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterUserRequest() (request *RegisterUserRequest) {
    request = &RegisterUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "RegisterUser")
    
    
    return
}

func NewRegisterUserResponse() (response *RegisterUserResponse) {
    response = &RegisterUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterUser
// This API is used to register a user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) RegisterUser(request *RegisterUserRequest) (response *RegisterUserResponse, err error) {
    return c.RegisterUserWithContext(context.Background(), request)
}

// RegisterUser
// This API is used to register a user.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ORIGINIDEXISTS = "FailedOperation.OriginIdExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) RegisterUserWithContext(ctx context.Context, request *RegisterUserRequest) (response *RegisterUserResponse, err error) {
    if request == nil {
        request = NewRegisterUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterUserResponse()
    err = c.Send(request, response)
    return
}

func NewSetAppCustomContentRequest() (request *SetAppCustomContentRequest) {
    request = &SetAppCustomContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "SetAppCustomContent")
    
    
    return
}

func NewSetAppCustomContentResponse() (response *SetAppCustomContentResponse) {
    response = &SetAppCustomContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetAppCustomContent
// This API is used to set or update the custom content of an application, including the application icon and custom code. After you update JS and CSS content, you also need to call this API for the updates to take effect.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CSSORJS = "InvalidParameter.CssOrJs"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) SetAppCustomContent(request *SetAppCustomContentRequest) (response *SetAppCustomContentResponse, err error) {
    return c.SetAppCustomContentWithContext(context.Background(), request)
}

// SetAppCustomContent
// This API is used to set or update the custom content of an application, including the application icon and custom code. After you update JS and CSS content, you also need to call this API for the updates to take effect.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CSSORJS = "InvalidParameter.CssOrJs"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) SetAppCustomContentWithContext(ctx context.Context, request *SetAppCustomContentRequest) (response *SetAppCustomContentResponse, err error) {
    if request == nil {
        request = NewSetAppCustomContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetAppCustomContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetAppCustomContentResponse()
    err = c.Send(request, response)
    return
}

func NewSetWatermarkRequest() (request *SetWatermarkRequest) {
    request = &SetWatermarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "SetWatermark")
    
    
    return
}

func NewSetWatermarkResponse() (response *SetWatermarkResponse) {
    response = &SetWatermarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetWatermark
// This API is used to configure watermarks.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetWatermark(request *SetWatermarkRequest) (response *SetWatermarkResponse, err error) {
    return c.SetWatermarkWithContext(context.Background(), request)
}

// SetWatermark
// This API is used to configure watermarks.
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SetWatermarkWithContext(ctx context.Context, request *SetWatermarkRequest) (response *SetWatermarkResponse, err error) {
    if request == nil {
        request = NewSetWatermarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetWatermark require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetWatermarkResponse()
    err = c.Send(request, response)
    return
}

func NewStartRoomRequest() (request *StartRoomRequest) {
    request = &StartRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "StartRoom")
    
    
    return
}

func NewStartRoomResponse() (response *StartRoomResponse) {
    response = &StartRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartRoom
// This API is used to start a room. Before you call this API, make sure a user has entered the room so that the class is initialized.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartRoom(request *StartRoomRequest) (response *StartRoomResponse, err error) {
    return c.StartRoomWithContext(context.Background(), request)
}

// StartRoom
// This API is used to start a room. Before you call this API, make sure a user has entered the room so that the class is initialized.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLASSENDED = "FailedOperation.ClassEnded"
//  FAILEDOPERATION_CLASSEXPIRED = "FailedOperation.ClassExpired"
//  FAILEDOPERATION_CLASSSTARTED = "FailedOperation.ClassStarted"
//  FAILEDOPERATION_REQUESTTIMEDOUT = "FailedOperation.RequestTimedOut"
//  FAILEDOPERATION_ROOMNOTEND = "FailedOperation.RoomNotEnd"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartRoomWithContext(ctx context.Context, request *StartRoomRequest) (response *StartRoomResponse, err error) {
    if request == nil {
        request = NewStartRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartRoomResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindDocumentFromRoomRequest() (request *UnbindDocumentFromRoomRequest) {
    request = &UnbindDocumentFromRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lcic", APIVersion, "UnbindDocumentFromRoom")
    
    
    return
}

func NewUnbindDocumentFromRoomResponse() (response *UnbindDocumentFromRoomResponse) {
    response = &UnbindDocumentFromRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindDocumentFromRoom
// This API is used to unbind a document from a room.
//
// error code that may be returned:
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) UnbindDocumentFromRoom(request *UnbindDocumentFromRoomRequest) (response *UnbindDocumentFromRoomResponse, err error) {
    return c.UnbindDocumentFromRoomWithContext(context.Background(), request)
}

// UnbindDocumentFromRoom
// This API is used to unbind a document from a room.
//
// error code that may be returned:
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ROOM = "ResourceNotFound.Room"
func (c *Client) UnbindDocumentFromRoomWithContext(ctx context.Context, request *UnbindDocumentFromRoomRequest) (response *UnbindDocumentFromRoomResponse, err error) {
    if request == nil {
        request = NewUnbindDocumentFromRoomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindDocumentFromRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindDocumentFromRoomResponse()
    err = c.Send(request, response)
    return
}
