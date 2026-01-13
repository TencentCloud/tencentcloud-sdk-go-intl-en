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

package v20230308

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2023-03-08"

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


func NewChangeMigratingTopicToNextStageRequest() (request *ChangeMigratingTopicToNextStageRequest) {
    request = &ChangeMigratingTopicToNextStageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ChangeMigratingTopicToNextStage")
    
    
    return
}

func NewChangeMigratingTopicToNextStageResponse() (response *ChangeMigratingTopicToNextStageResponse) {
    response = &ChangeMigratingTopicToNextStageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChangeMigratingTopicToNextStage
// This API is used to modify the Topic status during migration and go to next step.
//
// error code that may be returned:
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
func (c *Client) ChangeMigratingTopicToNextStage(request *ChangeMigratingTopicToNextStageRequest) (response *ChangeMigratingTopicToNextStageResponse, err error) {
    return c.ChangeMigratingTopicToNextStageWithContext(context.Background(), request)
}

// ChangeMigratingTopicToNextStage
// This API is used to modify the Topic status during migration and go to next step.
//
// error code that may be returned:
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
func (c *Client) ChangeMigratingTopicToNextStageWithContext(ctx context.Context, request *ChangeMigratingTopicToNextStageRequest) (response *ChangeMigratingTopicToNextStageResponse, err error) {
    if request == nil {
        request = NewChangeMigratingTopicToNextStageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "ChangeMigratingTopicToNextStage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeMigratingTopicToNextStage require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeMigratingTopicToNextStageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsumerGroupRequest() (request *CreateConsumerGroupRequest) {
    request = &CreateConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "CreateConsumerGroup")
    
    
    return
}

func NewCreateConsumerGroupResponse() (response *CreateConsumerGroupResponse) {
    response = &CreateConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsumerGroup
// Create consumer groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
    return c.CreateConsumerGroupWithContext(context.Background(), request)
}

// CreateConsumerGroup
// Create consumer groups.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateConsumerGroupWithContext(ctx context.Context, request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
    if request == nil {
        request = NewCreateConsumerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "CreateConsumerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "CreateInstance")
    
    
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstance
// This API is used to create a RocketMQ 5.x cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// This API is used to create a RocketMQ 5.x cluster.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "CreateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoleRequest() (request *CreateRoleRequest) {
    request = &CreateRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "CreateRole")
    
    
    return
}

func NewCreateRoleResponse() (response *CreateRoleResponse) {
    response = &CreateRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRole
// This API is used to add a role.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateRole(request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    return c.CreateRoleWithContext(context.Background(), request)
}

// CreateRole
// This API is used to add a role.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
func (c *Client) CreateRoleWithContext(ctx context.Context, request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    if request == nil {
        request = NewCreateRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "CreateRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConsumerGroupRequest() (request *DeleteConsumerGroupRequest) {
    request = &DeleteConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteConsumerGroup")
    
    
    return
}

func NewDeleteConsumerGroupResponse() (response *DeleteConsumerGroupResponse) {
    response = &DeleteConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConsumerGroup
// This API is used to delete a consumer group. After a consumer group is deleted, all configurations and related data of the consumer group are cleared and cannot be restored. After deletion, online consumer clients report errors. It is recommended to take clients offline in advance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_GROUP = "ResourceNotFound.Group"
func (c *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
    return c.DeleteConsumerGroupWithContext(context.Background(), request)
}

// DeleteConsumerGroup
// This API is used to delete a consumer group. After a consumer group is deleted, all configurations and related data of the consumer group are cleared and cannot be restored. After deletion, online consumer clients report errors. It is recommended to take clients offline in advance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_GROUP = "ResourceNotFound.Group"
func (c *Client) DeleteConsumerGroupWithContext(ctx context.Context, request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteConsumerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DeleteConsumerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteInstance")
    
    
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInstance
// This API is used to delete a TDMQ RocketMQ 5.x cluster. Topics, consumer groups, and roles in use should be deleted first.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    return c.DeleteInstanceWithContext(context.Background(), request)
}

// DeleteInstance
// This API is used to delete a TDMQ RocketMQ 5.x cluster. Topics, consumer groups, and roles in use should be deleted first.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DeleteInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoleRequest() (request *DeleteRoleRequest) {
    request = &DeleteRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteRole")
    
    
    return
}

func NewDeleteRoleResponse() (response *DeleteRoleResponse) {
    response = &DeleteRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRole
// This API is used to delete a role. Make sure that the related information on this role is not used in the current code. After the role is deleted, the keys (AccessKey and SecretKey) used to produce or consume messages with this role become invalid immediately.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteRole(request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
    return c.DeleteRoleWithContext(context.Background(), request)
}

// DeleteRole
// This API is used to delete a role. Make sure that the related information on this role is not used in the current code. After the role is deleted, the keys (AccessKey and SecretKey) used to produce or consume messages with this role become invalid immediately.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteRoleWithContext(ctx context.Context, request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
    if request == nil {
        request = NewDeleteRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DeleteRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSmoothMigrationTaskRequest() (request *DeleteSmoothMigrationTaskRequest) {
    request = &DeleteSmoothMigrationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteSmoothMigrationTask")
    
    
    return
}

func NewDeleteSmoothMigrationTaskResponse() (response *DeleteSmoothMigrationTaskResponse) {
    response = &DeleteSmoothMigrationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSmoothMigrationTask
// This API is used to delete a smooth migration task. Only canceled tasks can be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteSmoothMigrationTask(request *DeleteSmoothMigrationTaskRequest) (response *DeleteSmoothMigrationTaskResponse, err error) {
    return c.DeleteSmoothMigrationTaskWithContext(context.Background(), request)
}

// DeleteSmoothMigrationTask
// This API is used to delete a smooth migration task. Only canceled tasks can be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) DeleteSmoothMigrationTaskWithContext(ctx context.Context, request *DeleteSmoothMigrationTaskRequest) (response *DeleteSmoothMigrationTaskResponse, err error) {
    if request == nil {
        request = NewDeleteSmoothMigrationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DeleteSmoothMigrationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSmoothMigrationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSmoothMigrationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
    request = &DeleteTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DeleteTopic")
    
    
    return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
    response = &DeleteTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTopic
// This API is used to delete a topic. After deletion, all configurations and related data of the topic will be cleared and cannot be restored.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    return c.DeleteTopicWithContext(context.Background(), request)
}

// DeleteTopic
// This API is used to delete a topic. After deletion, all configurations and related data of the topic will be cleared and cannot be restored.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DeleteTopicWithContext(ctx context.Context, request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DeleteTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerClientRequest() (request *DescribeConsumerClientRequest) {
    request = &DescribeConsumerClientRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeConsumerClient")
    
    
    return
}

func NewDescribeConsumerClientResponse() (response *DescribeConsumerClientResponse) {
    response = &DescribeConsumerClientResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumerClient
// Query consumer client details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeConsumerClient(request *DescribeConsumerClientRequest) (response *DescribeConsumerClientResponse, err error) {
    return c.DescribeConsumerClientWithContext(context.Background(), request)
}

// DescribeConsumerClient
// Query consumer client details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeConsumerClientWithContext(ctx context.Context, request *DescribeConsumerClientRequest) (response *DescribeConsumerClientResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerClientRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeConsumerClient")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerClient require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerClientResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerClientListRequest() (request *DescribeConsumerClientListRequest) {
    request = &DescribeConsumerClientListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeConsumerClientList")
    
    
    return
}

func NewDescribeConsumerClientListResponse() (response *DescribeConsumerClientListResponse) {
    response = &DescribeConsumerClientListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumerClientList
// This API is used to query the client connection list of a consumer group.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeConsumerClientList(request *DescribeConsumerClientListRequest) (response *DescribeConsumerClientListResponse, err error) {
    return c.DescribeConsumerClientListWithContext(context.Background(), request)
}

// DescribeConsumerClientList
// This API is used to query the client connection list of a consumer group.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeConsumerClientListWithContext(ctx context.Context, request *DescribeConsumerClientListRequest) (response *DescribeConsumerClientListResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerClientListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeConsumerClientList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerClientList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerClientListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerGroupRequest() (request *DescribeConsumerGroupRequest) {
    request = &DescribeConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeConsumerGroup")
    
    
    return
}

func NewDescribeConsumerGroupResponse() (response *DescribeConsumerGroupResponse) {
    response = &DescribeConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumerGroup
// Query consumer group details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_GROUP = "ResourceNotFound.Group"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeConsumerGroup(request *DescribeConsumerGroupRequest) (response *DescribeConsumerGroupResponse, err error) {
    return c.DescribeConsumerGroupWithContext(context.Background(), request)
}

// DescribeConsumerGroup
// Query consumer group details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_GROUP = "ResourceNotFound.Group"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeConsumerGroupWithContext(ctx context.Context, request *DescribeConsumerGroupRequest) (response *DescribeConsumerGroupResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeConsumerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerLagRequest() (request *DescribeConsumerLagRequest) {
    request = &DescribeConsumerLagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeConsumerLag")
    
    
    return
}

func NewDescribeConsumerLagResponse() (response *DescribeConsumerLagResponse) {
    response = &DescribeConsumerLagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumerLag
// This API is used to query the number of heaped messages in a specified consumer group.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeConsumerLag(request *DescribeConsumerLagRequest) (response *DescribeConsumerLagResponse, err error) {
    return c.DescribeConsumerLagWithContext(context.Background(), request)
}

// DescribeConsumerLag
// This API is used to query the number of heaped messages in a specified consumer group.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeConsumerLagWithContext(ctx context.Context, request *DescribeConsumerLagRequest) (response *DescribeConsumerLagResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerLagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeConsumerLag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerLag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerLagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFusionInstanceListRequest() (request *DescribeFusionInstanceListRequest) {
    request = &DescribeFusionInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeFusionInstanceList")
    
    
    return
}

func NewDescribeFusionInstanceListResponse() (response *DescribeFusionInstanceListResponse) {
    response = &DescribeFusionInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFusionInstanceList
// This API is used to describe clusters, supporting both 4.x and 5.x clusters. Among them, parameter usage of Filters is as follows:.
//
// 
//
// -InstanceName, the cluster name, supports fuzzy query and can be obtained from the API return value or console.
//
// -InstanceId Cluster ID, exact query, obtain from the current API or console.
//
// - InstanceType cluster type, see [InstanceItem](https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#InstanceItem) data structure, supports multiple selections.
//
// - Version: cluster edition, enumeration values as follows:.
//
// -4 RocketMQ 4.x clusters.
//
// -Deploy a RocketMQ 5.x cluster.
//
// 
//
// This API is used to demonstrate Filters.
//
//  [{ "Name": "InstanceId", "Values": ["rmq-72mo3a9o"] }]
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFusionInstanceList(request *DescribeFusionInstanceListRequest) (response *DescribeFusionInstanceListResponse, err error) {
    return c.DescribeFusionInstanceListWithContext(context.Background(), request)
}

// DescribeFusionInstanceList
// This API is used to describe clusters, supporting both 4.x and 5.x clusters. Among them, parameter usage of Filters is as follows:.
//
// 
//
// -InstanceName, the cluster name, supports fuzzy query and can be obtained from the API return value or console.
//
// -InstanceId Cluster ID, exact query, obtain from the current API or console.
//
// - InstanceType cluster type, see [InstanceItem](https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#InstanceItem) data structure, supports multiple selections.
//
// - Version: cluster edition, enumeration values as follows:.
//
// -4 RocketMQ 4.x clusters.
//
// -Deploy a RocketMQ 5.x cluster.
//
// 
//
// This API is used to demonstrate Filters.
//
//  [{ "Name": "InstanceId", "Values": ["rmq-72mo3a9o"] }]
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFusionInstanceListWithContext(ctx context.Context, request *DescribeFusionInstanceListRequest) (response *DescribeFusionInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeFusionInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeFusionInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFusionInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFusionInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
    request = &DescribeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeInstance")
    
    
    return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
    response = &DescribeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstance
// This API is used to query RocketMQ 5.x cluster information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    return c.DescribeInstanceWithContext(context.Background(), request)
}

// DescribeInstance
// This API is used to query RocketMQ 5.x cluster information.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceListRequest() (request *DescribeInstanceListRequest) {
    request = &DescribeInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeInstanceList")
    
    
    return
}

func NewDescribeInstanceListResponse() (response *DescribeInstanceListResponse) {
    response = &DescribeInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceList
// This API is used to describe clusters. It only supports 5.x clusters. Description of the Filters parameter use is as follows:.
//
// 
//
// -InstanceName Cluster name, supports fuzzy search.
//
// - InstanceId The Tencent Cloud RocketMQ instance ID, obtained from the [DescribeFusionInstanceList](https://www.tencentcloud.com/document/api/1493/106745?from_cn_redirect=1) API or console.
//
// - InstanceType cluster type, see [InstanceItem](https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#InstanceItem) data structure, supports multiple selections.
//
// -InstanceStatus cluster status, see [InstanceItem](https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#InstanceItem) data structure, supports multiple selections.
//
// 
//
// This API is used to demonstrate Filters.
//
// [{
//
//     "Name": "InstanceId",
//
//     "Values": ["rmq-72mo3a9o"]
//
// }]
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    return c.DescribeInstanceListWithContext(context.Background(), request)
}

// DescribeInstanceList
// This API is used to describe clusters. It only supports 5.x clusters. Description of the Filters parameter use is as follows:.
//
// 
//
// -InstanceName Cluster name, supports fuzzy search.
//
// - InstanceId The Tencent Cloud RocketMQ instance ID, obtained from the [DescribeFusionInstanceList](https://www.tencentcloud.com/document/api/1493/106745?from_cn_redirect=1) API or console.
//
// - InstanceType cluster type, see [InstanceItem](https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#InstanceItem) data structure, supports multiple selections.
//
// -InstanceStatus cluster status, see [InstanceItem](https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#InstanceItem) data structure, supports multiple selections.
//
// 
//
// This API is used to demonstrate Filters.
//
// [{
//
//     "Name": "InstanceId",
//
//     "Values": ["rmq-72mo3a9o"]
//
// }]
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstanceListWithContext(ctx context.Context, request *DescribeInstanceListRequest) (response *DescribeInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMessageRequest() (request *DescribeMessageRequest) {
    request = &DescribeMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMessage")
    
    
    return
}

func NewDescribeMessageResponse() (response *DescribeMessageResponse) {
    response = &DescribeMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMessage
// Query message details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MESSAGE = "ResourceNotFound.Message"
func (c *Client) DescribeMessage(request *DescribeMessageRequest) (response *DescribeMessageResponse, err error) {
    return c.DescribeMessageWithContext(context.Background(), request)
}

// DescribeMessage
// Query message details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MESSAGE = "ResourceNotFound.Message"
func (c *Client) DescribeMessageWithContext(ctx context.Context, request *DescribeMessageRequest) (response *DescribeMessageResponse, err error) {
    if request == nil {
        request = NewDescribeMessageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeMessage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMessageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMessageListRequest() (request *DescribeMessageListRequest) {
    request = &DescribeMessageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMessageList")
    
    
    return
}

func NewDescribeMessageListResponse() (response *DescribeMessageListResponse) {
    response = &DescribeMessageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMessageList
// Query the message list. If querying dead letter messages, set the ConsumerGroup parameter.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMessageList(request *DescribeMessageListRequest) (response *DescribeMessageListResponse, err error) {
    return c.DescribeMessageListWithContext(context.Background(), request)
}

// DescribeMessageList
// Query the message list. If querying dead letter messages, set the ConsumerGroup parameter.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeMessageListWithContext(ctx context.Context, request *DescribeMessageListRequest) (response *DescribeMessageListResponse, err error) {
    if request == nil {
        request = NewDescribeMessageListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeMessageList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMessageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMessageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMessageTraceRequest() (request *DescribeMessageTraceRequest) {
    request = &DescribeMessageTraceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMessageTrace")
    
    
    return
}

func NewDescribeMessageTraceResponse() (response *DescribeMessageTraceResponse) {
    response = &DescribeMessageTraceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMessageTrace
// This API is used to query message trace by message ID.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MESSAGE = "ResourceNotFound.Message"
func (c *Client) DescribeMessageTrace(request *DescribeMessageTraceRequest) (response *DescribeMessageTraceResponse, err error) {
    return c.DescribeMessageTraceWithContext(context.Background(), request)
}

// DescribeMessageTrace
// This API is used to query message trace by message ID.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MESSAGE = "ResourceNotFound.Message"
func (c *Client) DescribeMessageTraceWithContext(ctx context.Context, request *DescribeMessageTraceRequest) (response *DescribeMessageTraceResponse, err error) {
    if request == nil {
        request = NewDescribeMessageTraceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeMessageTrace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMessageTrace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMessageTraceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigratingGroupStatsRequest() (request *DescribeMigratingGroupStatsRequest) {
    request = &DescribeMigratingGroupStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMigratingGroupStats")
    
    
    return
}

func NewDescribeMigratingGroupStatsResponse() (response *DescribeMigratingGroupStatsResponse) {
    response = &DescribeMigratingGroupStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigratingGroupStats
// This API is used to view real-time information of migration consumption groups.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MESSAGE = "ResourceNotFound.Message"
func (c *Client) DescribeMigratingGroupStats(request *DescribeMigratingGroupStatsRequest) (response *DescribeMigratingGroupStatsResponse, err error) {
    return c.DescribeMigratingGroupStatsWithContext(context.Background(), request)
}

// DescribeMigratingGroupStats
// This API is used to view real-time information of migration consumption groups.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MESSAGE = "ResourceNotFound.Message"
func (c *Client) DescribeMigratingGroupStatsWithContext(ctx context.Context, request *DescribeMigratingGroupStatsRequest) (response *DescribeMigratingGroupStatsResponse, err error) {
    if request == nil {
        request = NewDescribeMigratingGroupStatsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeMigratingGroupStats")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigratingGroupStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigratingGroupStatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigratingTopicListRequest() (request *DescribeMigratingTopicListRequest) {
    request = &DescribeMigratingTopicListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMigratingTopicList")
    
    
    return
}

func NewDescribeMigratingTopicListResponse() (response *DescribeMigratingTopicListResponse) {
    response = &DescribeMigratingTopicListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigratingTopicList
// This API is used to query the Topic migration status list.
//
// 
//
// The Filters field is a query filter that supports the following conditions:.
//
// This API is used to get topic names with fuzzy query support.
//
// This api is used to query the migration status. See the data structure in MigratingTopic (https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#MigratingTopic).
//
// This API is used to manage namespaces, valid only for 4.x clusters.
//
// 
//
// This API is used to demonstrate Filters.
//
// [{
//
//     "Name": "TopicName",
//
//     "Values": ["topic-a"]
//
// }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
func (c *Client) DescribeMigratingTopicList(request *DescribeMigratingTopicListRequest) (response *DescribeMigratingTopicListResponse, err error) {
    return c.DescribeMigratingTopicListWithContext(context.Background(), request)
}

// DescribeMigratingTopicList
// This API is used to query the Topic migration status list.
//
// 
//
// The Filters field is a query filter that supports the following conditions:.
//
// This API is used to get topic names with fuzzy query support.
//
// This api is used to query the migration status. See the data structure in MigratingTopic (https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#MigratingTopic).
//
// This API is used to manage namespaces, valid only for 4.x clusters.
//
// 
//
// This API is used to demonstrate Filters.
//
// [{
//
//     "Name": "TopicName",
//
//     "Values": ["topic-a"]
//
// }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
func (c *Client) DescribeMigratingTopicListWithContext(ctx context.Context, request *DescribeMigratingTopicListRequest) (response *DescribeMigratingTopicListResponse, err error) {
    if request == nil {
        request = NewDescribeMigratingTopicListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeMigratingTopicList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigratingTopicList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigratingTopicListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigratingTopicStatsRequest() (request *DescribeMigratingTopicStatsRequest) {
    request = &DescribeMigratingTopicStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMigratingTopicStats")
    
    
    return
}

func NewDescribeMigratingTopicStatsResponse() (response *DescribeMigratingTopicStatsResponse) {
    response = &DescribeMigratingTopicStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigratingTopicStats
// This API is used to query real-time data of migration topics.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeMigratingTopicStats(request *DescribeMigratingTopicStatsRequest) (response *DescribeMigratingTopicStatsResponse, err error) {
    return c.DescribeMigratingTopicStatsWithContext(context.Background(), request)
}

// DescribeMigratingTopicStats
// This API is used to query real-time data of migration topics.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeMigratingTopicStatsWithContext(ctx context.Context, request *DescribeMigratingTopicStatsRequest) (response *DescribeMigratingTopicStatsResponse, err error) {
    if request == nil {
        request = NewDescribeMigratingTopicStatsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeMigratingTopicStats")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigratingTopicStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigratingTopicStatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationTaskListRequest() (request *DescribeMigrationTaskListRequest) {
    request = &DescribeMigrationTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeMigrationTaskList")
    
    
    return
}

func NewDescribeMigrationTaskListResponse() (response *DescribeMigrationTaskListResponse) {
    response = &DescribeMigrationTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigrationTaskList
// This API is used to search the data migration task list. Filter parameter usage instructions are as follows:.
//
// 
//
// This API is used to search precisely according to task ID.
//
// InstanceId. This API is used to precisely search based on instance ID.
//
// This API is used to search by task Type.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeMigrationTaskList(request *DescribeMigrationTaskListRequest) (response *DescribeMigrationTaskListResponse, err error) {
    return c.DescribeMigrationTaskListWithContext(context.Background(), request)
}

// DescribeMigrationTaskList
// This API is used to search the data migration task list. Filter parameter usage instructions are as follows:.
//
// 
//
// This API is used to search precisely according to task ID.
//
// InstanceId. This API is used to precisely search based on instance ID.
//
// This API is used to search by task Type.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeMigrationTaskListWithContext(ctx context.Context, request *DescribeMigrationTaskListRequest) (response *DescribeMigrationTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationTaskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeMigrationTaskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrationTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrationTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProducerListRequest() (request *DescribeProducerListRequest) {
    request = &DescribeProducerListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeProducerList")
    
    
    return
}

func NewDescribeProducerListResponse() (response *DescribeProducerListResponse) {
    response = &DescribeProducerListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProducerList
// This API is used to query producer list information associated with a topic. Filters support the following criteria:.
//
// -client IP.
//
// -client ID.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeProducerList(request *DescribeProducerListRequest) (response *DescribeProducerListResponse, err error) {
    return c.DescribeProducerListWithContext(context.Background(), request)
}

// DescribeProducerList
// This API is used to query producer list information associated with a topic. Filters support the following criteria:.
//
// -client IP.
//
// -client ID.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeProducerListWithContext(ctx context.Context, request *DescribeProducerListRequest) (response *DescribeProducerListResponse, err error) {
    if request == nil {
        request = NewDescribeProducerListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeProducerList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProducerList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProducerListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductSKUsRequest() (request *DescribeProductSKUsRequest) {
    request = &DescribeProductSKUsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeProductSKUs")
    
    
    return
}

func NewDescribeProductSKUsResponse() (response *DescribeProductSKUsResponse) {
    response = &DescribeProductSKUsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProductSKUs
// This API is used to query product sales specifications against RocketMQ 5.x clusters.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeProductSKUs(request *DescribeProductSKUsRequest) (response *DescribeProductSKUsResponse, err error) {
    return c.DescribeProductSKUsWithContext(context.Background(), request)
}

// DescribeProductSKUs
// This API is used to query product sales specifications against RocketMQ 5.x clusters.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeProductSKUsWithContext(ctx context.Context, request *DescribeProductSKUsRequest) (response *DescribeProductSKUsResponse, err error) {
    if request == nil {
        request = NewDescribeProductSKUsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeProductSKUs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductSKUs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductSKUsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoleListRequest() (request *DescribeRoleListRequest) {
    request = &DescribeRoleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeRoleList")
    
    
    return
}

func NewDescribeRoleListResponse() (response *DescribeRoleListResponse) {
    response = &DescribeRoleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoleList
// This API is used to query the list of roles. Filter parameter usage instructions are as follows:.
//
// 
//
// -Role name supports fuzzy search and can be obtained from the API return value or console.
//
// -AccessKey, support fuzzy search, obtain from API return value or console.
//
// 
//
// This API is used to demonstrate Filters. 
//
// [{ "Name": "RoleName", "Values": ["test_role"] }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRoleList(request *DescribeRoleListRequest) (response *DescribeRoleListResponse, err error) {
    return c.DescribeRoleListWithContext(context.Background(), request)
}

// DescribeRoleList
// This API is used to query the list of roles. Filter parameter usage instructions are as follows:.
//
// 
//
// -Role name supports fuzzy search and can be obtained from the API return value or console.
//
// -AccessKey, support fuzzy search, obtain from API return value or console.
//
// 
//
// This API is used to demonstrate Filters. 
//
// [{ "Name": "RoleName", "Values": ["test_role"] }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeRoleListWithContext(ctx context.Context, request *DescribeRoleListRequest) (response *DescribeRoleListResponse, err error) {
    if request == nil {
        request = NewDescribeRoleListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeRoleList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoleList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmoothMigrationTaskListRequest() (request *DescribeSmoothMigrationTaskListRequest) {
    request = &DescribeSmoothMigrationTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeSmoothMigrationTaskList")
    
    
    return
}

func NewDescribeSmoothMigrationTaskListResponse() (response *DescribeSmoothMigrationTaskListResponse) {
    response = &DescribeSmoothMigrationTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSmoothMigrationTaskList
// This API is used to query the migration task list smoothly.
//
// 
//
// This API is used to query the supported fields of the query parameter Filters as follows:.
//
// Task status, supports multiple selections. 
//
// ConnectionType, network connection type, supports multiple selections. See the description of SmoothMigrationTaskItem (https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#SmoothMigrationTaskItem).
//
// This API is used to search for an instance by instance ID with precise matching. 
//
// This API is used to query task names with fuzzy search support.
//
// 
//
// This API is used to demonstrate Filters.
//
// [{
//
//     "Name": "InstanceId",
//
//     "Values": ["rmq-1gzecldfg"]
//
// }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSmoothMigrationTaskList(request *DescribeSmoothMigrationTaskListRequest) (response *DescribeSmoothMigrationTaskListResponse, err error) {
    return c.DescribeSmoothMigrationTaskListWithContext(context.Background(), request)
}

// DescribeSmoothMigrationTaskList
// This API is used to query the migration task list smoothly.
//
// 
//
// This API is used to query the supported fields of the query parameter Filters as follows:.
//
// Task status, supports multiple selections. 
//
// ConnectionType, network connection type, supports multiple selections. See the description of SmoothMigrationTaskItem (https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#SmoothMigrationTaskItem).
//
// This API is used to search for an instance by instance ID with precise matching. 
//
// This API is used to query task names with fuzzy search support.
//
// 
//
// This API is used to demonstrate Filters.
//
// [{
//
//     "Name": "InstanceId",
//
//     "Values": ["rmq-1gzecldfg"]
//
// }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSmoothMigrationTaskListWithContext(ctx context.Context, request *DescribeSmoothMigrationTaskListRequest) (response *DescribeSmoothMigrationTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeSmoothMigrationTaskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeSmoothMigrationTaskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmoothMigrationTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSmoothMigrationTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceClusterGroupListRequest() (request *DescribeSourceClusterGroupListRequest) {
    request = &DescribeSourceClusterGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeSourceClusterGroupList")
    
    
    return
}

func NewDescribeSourceClusterGroupListResponse() (response *DescribeSourceClusterGroupListResponse) {
    response = &DescribeSourceClusterGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSourceClusterGroupList
// This API is used to obtain the group list of the source cluster during the smooth migration process.
//
// 
//
// The Filters field is a query filter that supports the following fields:.
//
// This API is used to query consumer group names with fuzzy search support.
//
// This API is used to check whether the data is Imported.
//
// This api is used to check the import status. See SourceClusterGroupConfig (https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#SourceClusterGroupConfig) for details.
//
// This API is used to manage namespaces, valid only for 4.x clusters.
//
// 
//
// This API is used to demonstrate Filters.
//
// [{
//
//     "Name": "GroupName",
//
//     "Values": ["group-a"]
//
// }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSourceClusterGroupList(request *DescribeSourceClusterGroupListRequest) (response *DescribeSourceClusterGroupListResponse, err error) {
    return c.DescribeSourceClusterGroupListWithContext(context.Background(), request)
}

// DescribeSourceClusterGroupList
// This API is used to obtain the group list of the source cluster during the smooth migration process.
//
// 
//
// The Filters field is a query filter that supports the following fields:.
//
// This API is used to query consumer group names with fuzzy search support.
//
// This API is used to check whether the data is Imported.
//
// This api is used to check the import status. See SourceClusterGroupConfig (https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#SourceClusterGroupConfig) for details.
//
// This API is used to manage namespaces, valid only for 4.x clusters.
//
// 
//
// This API is used to demonstrate Filters.
//
// [{
//
//     "Name": "GroupName",
//
//     "Values": ["group-a"]
//
// }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_MIGRATIONTASK = "ResourceNotFound.MigrationTask"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeSourceClusterGroupListWithContext(ctx context.Context, request *DescribeSourceClusterGroupListRequest) (response *DescribeSourceClusterGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeSourceClusterGroupListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeSourceClusterGroupList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSourceClusterGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSourceClusterGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicRequest() (request *DescribeTopicRequest) {
    request = &DescribeTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeTopic")
    
    
    return
}

func NewDescribeTopicResponse() (response *DescribeTopicResponse) {
    response = &DescribeTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopic
// This API is used to query topic details. The Offset and Limit parameters are pagination parameters for consumer groups subscribing to this topic. The Filter parameter usage instructions are as follows:.
//
// 
//
// -The ConsumerGroup name can be obtained from the [ConsumeGroupItem](https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#ConsumeGroupItem) in the API response of [DescribeConsumerGroupList](https://www.tencentcloud.com/document/api/1493/101535?from_cn_redirect=1) or from the console.
//
// 
//
// This API is used to demonstrate Filters. 
//
// [{ "Name": "ConsumerGroup", "Values": ["test_group"] }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeTopic(request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
    return c.DescribeTopicWithContext(context.Background(), request)
}

// DescribeTopic
// This API is used to query topic details. The Offset and Limit parameters are pagination parameters for consumer groups subscribing to this topic. The Filter parameter usage instructions are as follows:.
//
// 
//
// -The ConsumerGroup name can be obtained from the [ConsumeGroupItem](https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#ConsumeGroupItem) in the API response of [DescribeConsumerGroupList](https://www.tencentcloud.com/document/api/1493/101535?from_cn_redirect=1) or from the console.
//
// 
//
// This API is used to demonstrate Filters. 
//
// [{ "Name": "ConsumerGroup", "Values": ["test_group"] }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"
func (c *Client) DescribeTopicWithContext(ctx context.Context, request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
    if request == nil {
        request = NewDescribeTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicListRequest() (request *DescribeTopicListRequest) {
    request = &DescribeTopicListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeTopicList")
    
    
    return
}

func NewDescribeTopicListResponse() (response *DescribeTopicListResponse) {
    response = &DescribeTopicListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopicList
// This API is used to search the topic list. Filter parameter usage instructions are as follows:.
//
// 
//
// -TopicName supports fuzzy search. Obtain it from the [TopicItem](https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#TopicItem) in the [DescribeTopicList](https://www.tencentcloud.com/document/api/1493/96030?from_cn_redirect=1) API response or the console.
//
// -Search by TopicType, support multiple selections. See the TopicType field in the [DescribeTopic](https://www.tencentcloud.com/document/api/1493/97945?from_cn_redirect=1) API.
//
// 
//
// This API is used to demonstrate Filters.
//
//  [{ "Name": "TopicName", "Values": ["test_topic"] }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeTopicList(request *DescribeTopicListRequest) (response *DescribeTopicListResponse, err error) {
    return c.DescribeTopicListWithContext(context.Background(), request)
}

// DescribeTopicList
// This API is used to search the topic list. Filter parameter usage instructions are as follows:.
//
// 
//
// -TopicName supports fuzzy search. Obtain it from the [TopicItem](https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#TopicItem) in the [DescribeTopicList](https://www.tencentcloud.com/document/api/1493/96030?from_cn_redirect=1) API response or the console.
//
// -Search by TopicType, support multiple selections. See the TopicType field in the [DescribeTopic](https://www.tencentcloud.com/document/api/1493/97945?from_cn_redirect=1) API.
//
// 
//
// This API is used to demonstrate Filters.
//
//  [{ "Name": "TopicName", "Values": ["test_topic"] }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeTopicListWithContext(ctx context.Context, request *DescribeTopicListRequest) (response *DescribeTopicListResponse, err error) {
    if request == nil {
        request = NewDescribeTopicListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeTopicList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicListByGroupRequest() (request *DescribeTopicListByGroupRequest) {
    request = &DescribeTopicListByGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DescribeTopicListByGroup")
    
    
    return
}

func NewDescribeTopicListByGroupResponse() (response *DescribeTopicListByGroupResponse) {
    response = &DescribeTopicListByGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopicListByGroup
// This API is used to get topic list by consumer group. Filter parameter usage instructions are as follows:.
//
// 
//
// -TopicName. It can be obtained from [TopicItem](https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#TopicItem) returned by the API [DescribeTopicList](https://www.tencentcloud.com/document/api/1493/96030?from_cn_redirect=1) or from the console.
//
// 
//
// This API is used to demonstrate Filters. 
//
// [{ "Name": "TopicName", "Values": ["test_topic"] }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeTopicListByGroup(request *DescribeTopicListByGroupRequest) (response *DescribeTopicListByGroupResponse, err error) {
    return c.DescribeTopicListByGroupWithContext(context.Background(), request)
}

// DescribeTopicListByGroup
// This API is used to get topic list by consumer group. Filter parameter usage instructions are as follows:.
//
// 
//
// -TopicName. It can be obtained from [TopicItem](https://www.tencentcloud.com/document/api/1493/96031?from_cn_redirect=1#TopicItem) returned by the API [DescribeTopicList](https://www.tencentcloud.com/document/api/1493/96030?from_cn_redirect=1) or from the console.
//
// 
//
// This API is used to demonstrate Filters. 
//
// [{ "Name": "TopicName", "Values": ["test_topic"] }]
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DescribeTopicListByGroupWithContext(ctx context.Context, request *DescribeTopicListByGroupRequest) (response *DescribeTopicListByGroupResponse, err error) {
    if request == nil {
        request = NewDescribeTopicListByGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DescribeTopicListByGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicListByGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicListByGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDoHealthCheckOnMigratingTopicRequest() (request *DoHealthCheckOnMigratingTopicRequest) {
    request = &DoHealthCheckOnMigratingTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "DoHealthCheckOnMigratingTopic")
    
    
    return
}

func NewDoHealthCheckOnMigratingTopicResponse() (response *DoHealthCheckOnMigratingTopicResponse) {
    response = &DoHealthCheckOnMigratingTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DoHealthCheckOnMigratingTopic
// This API is used to check whether the topics being migrated are in normal status. Only topics in normal status can enter the next migration stage.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DoHealthCheckOnMigratingTopic(request *DoHealthCheckOnMigratingTopicRequest) (response *DoHealthCheckOnMigratingTopicResponse, err error) {
    return c.DoHealthCheckOnMigratingTopicWithContext(context.Background(), request)
}

// DoHealthCheckOnMigratingTopic
// This API is used to check whether the topics being migrated are in normal status. Only topics in normal status can enter the next migration stage.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) DoHealthCheckOnMigratingTopicWithContext(ctx context.Context, request *DoHealthCheckOnMigratingTopicRequest) (response *DoHealthCheckOnMigratingTopicResponse, err error) {
    if request == nil {
        request = NewDoHealthCheckOnMigratingTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "DoHealthCheckOnMigratingTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DoHealthCheckOnMigratingTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDoHealthCheckOnMigratingTopicResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsumerGroupRequest() (request *ModifyConsumerGroupRequest) {
    request = &ModifyConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyConsumerGroup")
    
    
    return
}

func NewModifyConsumerGroupResponse() (response *ModifyConsumerGroupResponse) {
    response = &ModifyConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConsumerGroup
// Modify consumer group attributes.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyConsumerGroup(request *ModifyConsumerGroupRequest) (response *ModifyConsumerGroupResponse, err error) {
    return c.ModifyConsumerGroupWithContext(context.Background(), request)
}

// ModifyConsumerGroup
// Modify consumer group attributes.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
func (c *Client) ModifyConsumerGroupWithContext(ctx context.Context, request *ModifyConsumerGroupRequest) (response *ModifyConsumerGroupResponse, err error) {
    if request == nil {
        request = NewModifyConsumerGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "ModifyConsumerGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyInstance")
    
    
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstance
// This API is used to modify attributes of a TDMQ RocketMQ 5.x cluster. Only running clusters support this operation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION_INSTANCETOPICNUMDOWNGRADE = "UnsupportedOperation.InstanceTopicNumDowngrade"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    return c.ModifyInstanceWithContext(context.Background(), request)
}

// ModifyInstance
// This API is used to modify attributes of a TDMQ RocketMQ 5.x cluster. Only running clusters support this operation.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION_INSTANCETOPICNUMDOWNGRADE = "UnsupportedOperation.InstanceTopicNumDowngrade"
func (c *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "ModifyInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceEndpointRequest() (request *ModifyInstanceEndpointRequest) {
    request = &ModifyInstanceEndpointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyInstanceEndpoint")
    
    
    return
}

func NewModifyInstanceEndpointResponse() (response *ModifyInstanceEndpointResponse) {
    response = &ModifyInstanceEndpointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceEndpoint
// This API is used to modify access points of a TDMQ RocketMQ 5.x cluster. Make sure that the access points exist before the operation.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_ENDPOINT = "ResourceNotFound.Endpoint"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstanceEndpoint(request *ModifyInstanceEndpointRequest) (response *ModifyInstanceEndpointResponse, err error) {
    return c.ModifyInstanceEndpointWithContext(context.Background(), request)
}

// ModifyInstanceEndpoint
// This API is used to modify access points of a TDMQ RocketMQ 5.x cluster. Make sure that the access points exist before the operation.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_ENDPOINT = "ResourceNotFound.Endpoint"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstanceEndpointWithContext(ctx context.Context, request *ModifyInstanceEndpointRequest) (response *ModifyInstanceEndpointResponse, err error) {
    if request == nil {
        request = NewModifyInstanceEndpointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "ModifyInstanceEndpoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceEndpoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceEndpointResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoleRequest() (request *ModifyRoleRequest) {
    request = &ModifyRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyRole")
    
    
    return
}

func NewModifyRoleResponse() (response *ModifyRoleResponse) {
    response = &ModifyRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRole
// This API is used to modify a role.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) ModifyRole(request *ModifyRoleRequest) (response *ModifyRoleResponse, err error) {
    return c.ModifyRoleWithContext(context.Background(), request)
}

// ModifyRole
// This API is used to modify a role.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"
//  RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"
func (c *Client) ModifyRoleWithContext(ctx context.Context, request *ModifyRoleRequest) (response *ModifyRoleResponse, err error) {
    if request == nil {
        request = NewModifyRoleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "ModifyRole")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRoleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicRequest() (request *ModifyTopicRequest) {
    request = &ModifyTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ModifyTopic")
    
    
    return
}

func NewModifyTopicResponse() (response *ModifyTopicResponse) {
    response = &ModifyTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTopic
// Modify topic attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTopic(request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    return c.ModifyTopicWithContext(context.Background(), request)
}

// ModifyTopic
// Modify topic attributes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyTopicWithContext(ctx context.Context, request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
    if request == nil {
        request = NewModifyTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "ModifyTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTopicResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveMigratingTopicRequest() (request *RemoveMigratingTopicRequest) {
    request = &RemoveMigratingTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "RemoveMigratingTopic")
    
    
    return
}

func NewRemoveMigratingTopicResponse() (response *RemoveMigratingTopicResponse) {
    response = &RemoveMigratingTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveMigratingTopic
// This API is used to remove a topic from the migration list. It is valid only when the topic is in the initial state.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RemoveMigratingTopic(request *RemoveMigratingTopicRequest) (response *RemoveMigratingTopicResponse, err error) {
    return c.RemoveMigratingTopicWithContext(context.Background(), request)
}

// RemoveMigratingTopic
// This API is used to remove a topic from the migration list. It is valid only when the topic is in the initial state.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RemoveMigratingTopicWithContext(ctx context.Context, request *RemoveMigratingTopicRequest) (response *RemoveMigratingTopicResponse, err error) {
    if request == nil {
        request = NewRemoveMigratingTopicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "RemoveMigratingTopic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveMigratingTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveMigratingTopicResponse()
    err = c.Send(request, response)
    return
}

func NewResendDeadLetterMessageRequest() (request *ResendDeadLetterMessageRequest) {
    request = &ResendDeadLetterMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ResendDeadLetterMessage")
    
    
    return
}

func NewResendDeadLetterMessageResponse() (response *ResendDeadLetterMessageResponse) {
    response = &ResendDeadLetterMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResendDeadLetterMessage
// This API is used to resend dead letter messages.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResendDeadLetterMessage(request *ResendDeadLetterMessageRequest) (response *ResendDeadLetterMessageResponse, err error) {
    return c.ResendDeadLetterMessageWithContext(context.Background(), request)
}

// ResendDeadLetterMessage
// This API is used to resend dead letter messages.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResendDeadLetterMessageWithContext(ctx context.Context, request *ResendDeadLetterMessageRequest) (response *ResendDeadLetterMessageResponse, err error) {
    if request == nil {
        request = NewResendDeadLetterMessageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "ResendDeadLetterMessage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResendDeadLetterMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewResendDeadLetterMessageResponse()
    err = c.Send(request, response)
    return
}

func NewResetConsumerGroupOffsetRequest() (request *ResetConsumerGroupOffsetRequest) {
    request = &ResetConsumerGroupOffsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "ResetConsumerGroupOffset")
    
    
    return
}

func NewResetConsumerGroupOffsetResponse() (response *ResetConsumerGroupOffsetResponse) {
    response = &ResetConsumerGroupOffsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetConsumerGroupOffset
// Reset the consumption position.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResetConsumerGroupOffset(request *ResetConsumerGroupOffsetRequest) (response *ResetConsumerGroupOffsetResponse, err error) {
    return c.ResetConsumerGroupOffsetWithContext(context.Background(), request)
}

// ResetConsumerGroupOffset
// Reset the consumption position.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ResetConsumerGroupOffsetWithContext(ctx context.Context, request *ResetConsumerGroupOffsetRequest) (response *ResetConsumerGroupOffsetResponse, err error) {
    if request == nil {
        request = NewResetConsumerGroupOffsetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "ResetConsumerGroupOffset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetConsumerGroupOffset require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetConsumerGroupOffsetResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackMigratingTopicStageRequest() (request *RollbackMigratingTopicStageRequest) {
    request = &RollbackMigratingTopicStageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trocket", APIVersion, "RollbackMigratingTopicStage")
    
    
    return
}

func NewRollbackMigratingTopicStageResponse() (response *RollbackMigratingTopicStageResponse) {
    response = &RollbackMigratingTopicStageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RollbackMigratingTopicStage
// This API is used to roll back the topic that is undergoing migration to the previous stage.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RollbackMigratingTopicStage(request *RollbackMigratingTopicStageRequest) (response *RollbackMigratingTopicStageResponse, err error) {
    return c.RollbackMigratingTopicStageWithContext(context.Background(), request)
}

// RollbackMigratingTopicStage
// This API is used to roll back the topic that is undergoing migration to the previous stage.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RollbackMigratingTopicStageWithContext(ctx context.Context, request *RollbackMigratingTopicStageRequest) (response *RollbackMigratingTopicStageResponse, err error) {
    if request == nil {
        request = NewRollbackMigratingTopicStageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trocket", APIVersion, "RollbackMigratingTopicStage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackMigratingTopicStage require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackMigratingTopicStageResponse()
    err = c.Send(request, response)
    return
}
