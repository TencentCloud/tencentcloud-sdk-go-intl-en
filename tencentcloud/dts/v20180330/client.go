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

package v20180330

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-03-30"

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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewActivateSubscribeRequest() (request *ActivateSubscribeRequest) {
    request = &ActivateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ActivateSubscribe")
    return
}

func NewActivateSubscribeResponse() (response *ActivateSubscribeResponse) {
    response = &ActivateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to configure a data subscription, which can be called only for subscription instances in unconfigured status.
func (c *Client) ActivateSubscribe(request *ActivateSubscribeRequest) (response *ActivateSubscribeResponse, err error) {
    if request == nil {
        request = NewActivateSubscribeRequest()
    }
    response = NewActivateSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCompleteMigrateJobRequest() (request *CompleteMigrateJobRequest) {
    request = &CompleteMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CompleteMigrateJob")
    return
}

func NewCompleteMigrateJobResponse() (response *CompleteMigrateJobResponse) {
    response = &CompleteMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (CompleteMigrateJob) is used to complete a data migration task.
// For tasks in incremental migration mode, you need to call this API before migration gets ready, so as to stop migrating incremental data.
// If the task status queried through the (DescribeMigrateJobs) API is ready (status=8), you can call this API to complete the migration task.
func (c *Client) CompleteMigrateJob(request *CompleteMigrateJobRequest) (response *CompleteMigrateJobResponse, err error) {
    if request == nil {
        request = NewCompleteMigrateJobRequest()
    }
    response = NewCompleteMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMigrateCheckJobRequest() (request *CreateMigrateCheckJobRequest) {
    request = &CreateMigrateCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CreateMigrateCheckJob")
    return
}

func NewCreateMigrateCheckJobResponse() (response *CreateMigrateCheckJobResponse) {
    response = &CreateMigrateCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a migration check task.
// Before migration, you should call this API to create a check. Migration will start only if the check succeeds. You can view the check result through the DescribeMigrateCheckJob API.
// After successful check, if the migration task needs to be modified, a new check task should be created and migration will begin only after the new check succeeds.
func (c *Client) CreateMigrateCheckJob(request *CreateMigrateCheckJobRequest) (response *CreateMigrateCheckJobResponse, err error) {
    if request == nil {
        request = NewCreateMigrateCheckJobRequest()
    }
    response = NewCreateMigrateCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMigrateJobRequest() (request *CreateMigrateJobRequest) {
    request = &CreateMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CreateMigrateJob")
    return
}

func NewCreateMigrateJobResponse() (response *CreateMigrateJobResponse) {
    response = &CreateMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (CreateMigrateJob) is used to create a data migration task.
// 
// For a finance zone linkage, please use the domain name dts.ap-shenzhen-fsi.tencentcloudapi.com.
func (c *Client) CreateMigrateJob(request *CreateMigrateJobRequest) (response *CreateMigrateJobResponse, err error) {
    if request == nil {
        request = NewCreateMigrateJobRequest()
    }
    response = NewCreateMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubscribeRequest() (request *CreateSubscribeRequest) {
    request = &CreateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CreateSubscribe")
    return
}

func NewCreateSubscribeResponse() (response *CreateSubscribeResponse) {
    response = &CreateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a data subscription instance.
func (c *Client) CreateSubscribe(request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
    if request == nil {
        request = NewCreateSubscribeRequest()
    }
    response = NewCreateSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSyncCheckJobRequest() (request *CreateSyncCheckJobRequest) {
    request = &CreateSyncCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CreateSyncCheckJob")
    return
}

func NewCreateSyncCheckJobResponse() (response *CreateSyncCheckJobResponse) {
    response = &CreateSyncCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Before the StartSyncJob API is called to start disaster recovery sync, this API should be called first to create a check. Data sync can start only if the check succeeds. You can view the check result through the DescribeSyncCheckJob API.
// Sync can begin only if the check succeeds.
func (c *Client) CreateSyncCheckJob(request *CreateSyncCheckJobRequest) (response *CreateSyncCheckJobResponse, err error) {
    if request == nil {
        request = NewCreateSyncCheckJobRequest()
    }
    response = NewCreateSyncCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSyncJobRequest() (request *CreateSyncJobRequest) {
    request = &CreateSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CreateSyncJob")
    return
}

func NewCreateSyncJobResponse() (response *CreateSyncJobResponse) {
    response = &CreateSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (CreateSyncJob) is used to create a disaster recovery sync task.
// After successful creation, check can be initiated through the CreateSyncCheckJob API. The sync task can be started through the StartSyncJob API only if the check succeeds.
func (c *Client) CreateSyncJob(request *CreateSyncJobRequest) (response *CreateSyncJobResponse, err error) {
    if request == nil {
        request = NewCreateSyncJobRequest()
    }
    response = NewCreateSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMigrateJobRequest() (request *DeleteMigrateJobRequest) {
    request = &DeleteMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DeleteMigrateJob")
    return
}

func NewDeleteMigrateJobResponse() (response *DeleteMigrateJobResponse) {
    response = &DeleteMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (DeleteMigrationJob) is used to delete a data migration task. If the task status queried through the DescribeMigrateJobs API is checking (status=3), running (status=7), ready (status=8), canceling (status=11), or completing (status=12), the task cannot be deleted.
func (c *Client) DeleteMigrateJob(request *DeleteMigrateJobRequest) (response *DeleteMigrateJobResponse, err error) {
    if request == nil {
        request = NewDeleteMigrateJobRequest()
    }
    response = NewDeleteMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSyncJobRequest() (request *DeleteSyncJobRequest) {
    request = &DeleteSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DeleteSyncJob")
    return
}

func NewDeleteSyncJobResponse() (response *DeleteSyncJobResponse) {
    response = &DeleteSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a disaster recovery sync task. Sync tasks that are running cannot be deleted.
func (c *Client) DeleteSyncJob(request *DeleteSyncJobRequest) (response *DeleteSyncJobResponse, err error) {
    if request == nil {
        request = NewDeleteSyncJobRequest()
    }
    response = NewDeleteSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsyncRequestInfoRequest() (request *DescribeAsyncRequestInfoRequest) {
    request = &DescribeAsyncRequestInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeAsyncRequestInfo")
    return
}

func NewDescribeAsyncRequestInfoResponse() (response *DescribeAsyncRequestInfoResponse) {
    response = &DescribeAsyncRequestInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the execution result of a task.
func (c *Client) DescribeAsyncRequestInfo(request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncRequestInfoRequest()
    }
    response = NewDescribeAsyncRequestInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrateCheckJobRequest() (request *DescribeMigrateCheckJobRequest) {
    request = &DescribeMigrateCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateCheckJob")
    return
}

func NewDescribeMigrateCheckJobResponse() (response *DescribeMigrateCheckJobResponse) {
    response = &DescribeMigrateCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the check result and query check status and progress after a check is created. 
// If the check succeeds, you can call the StartMigrateJob API to start migration.
// If the check fails, the reason can be queried. Please modify the migration configuration or adjust relevant parameters of the source/target instances through the ModifyMigrateJob API based on the error message.
func (c *Client) DescribeMigrateCheckJob(request *DescribeMigrateCheckJobRequest) (response *DescribeMigrateCheckJobResponse, err error) {
    if request == nil {
        request = NewDescribeMigrateCheckJobRequest()
    }
    response = NewDescribeMigrateCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrateJobsRequest() (request *DescribeMigrateJobsRequest) {
    request = &DescribeMigrateJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateJobs")
    return
}

func NewDescribeMigrateJobsResponse() (response *DescribeMigrateJobsResponse) {
    response = &DescribeMigrateJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query data migration tasks.
// For a finance zone linkage, please use the domain name https://dts.ap-shenzhen-fsi.tencentcloudapi.com.
func (c *Client) DescribeMigrateJobs(request *DescribeMigrateJobsRequest) (response *DescribeMigrateJobsResponse, err error) {
    if request == nil {
        request = NewDescribeMigrateJobsRequest()
    }
    response = NewDescribeMigrateJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionConfRequest() (request *DescribeRegionConfRequest) {
    request = &DescribeRegionConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeRegionConf")
    return
}

func NewDescribeRegionConfResponse() (response *DescribeRegionConfResponse) {
    response = &DescribeRegionConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the purchasable subscription instance regions.
func (c *Client) DescribeRegionConf(request *DescribeRegionConfRequest) (response *DescribeRegionConfResponse, err error) {
    if request == nil {
        request = NewDescribeRegionConfRequest()
    }
    response = NewDescribeRegionConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeConfRequest() (request *DescribeSubscribeConfRequest) {
    request = &DescribeSubscribeConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeConf")
    return
}

func NewDescribeSubscribeConfResponse() (response *DescribeSubscribeConfResponse) {
    response = &DescribeSubscribeConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the subscription instance configuration.
func (c *Client) DescribeSubscribeConf(request *DescribeSubscribeConfRequest) (response *DescribeSubscribeConfResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeConfRequest()
    }
    response = NewDescribeSubscribeConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribesRequest() (request *DescribeSubscribesRequest) {
    request = &DescribeSubscribesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribes")
    return
}

func NewDescribeSubscribesResponse() (response *DescribeSubscribesResponse) {
    response = &DescribeSubscribesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the information list of data subscription instances. Pagination is enabled by default with 20 results returned each time.
func (c *Client) DescribeSubscribes(request *DescribeSubscribesRequest) (response *DescribeSubscribesResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribesRequest()
    }
    response = NewDescribeSubscribesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSyncCheckJobRequest() (request *DescribeSyncCheckJobRequest) {
    request = &DescribeSyncCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSyncCheckJob")
    return
}

func NewDescribeSyncCheckJobResponse() (response *DescribeSyncCheckJobResponse) {
    response = &DescribeSyncCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the check result after a disaster recovery sync check task is created through the CreateSyncCheckJob API. Check status and progress can be queried.
// If the check succeeds, you can call the StartSyncJob API to start the sync task.
// If the check fails, the reason will be returned. You can modify the configuration through the ModifySyncJob API and initiate check again.
// It takes about 30 seconds to complete the check task. If the returned status is not "finished", the check has not been completed, and this API needs to be polled.
// If Status=finished and CheckFlag=1, the check succeeds.
// If Status=finished and CheckFlag !=1, the check fails.
func (c *Client) DescribeSyncCheckJob(request *DescribeSyncCheckJobRequest) (response *DescribeSyncCheckJobResponse, err error) {
    if request == nil {
        request = NewDescribeSyncCheckJobRequest()
    }
    response = NewDescribeSyncCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSyncJobsRequest() (request *DescribeSyncJobsRequest) {
    request = &DescribeSyncJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSyncJobs")
    return
}

func NewDescribeSyncJobsResponse() (response *DescribeSyncJobsResponse) {
    response = &DescribeSyncJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query disaster recovery sync tasks initiated on the DTS platform.
func (c *Client) DescribeSyncJobs(request *DescribeSyncJobsRequest) (response *DescribeSyncJobsResponse, err error) {
    if request == nil {
        request = NewDescribeSyncJobsRequest()
    }
    response = NewDescribeSyncJobsResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateSubscribeRequest() (request *IsolateSubscribeRequest) {
    request = &IsolateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "IsolateSubscribe")
    return
}

func NewIsolateSubscribeResponse() (response *IsolateSubscribeResponse) {
    response = &IsolateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to isolate an hourly billed subscription instance. After this API is called, the instance will become unavailable and billing will stop for it.
func (c *Client) IsolateSubscribe(request *IsolateSubscribeRequest) (response *IsolateSubscribeResponse, err error) {
    if request == nil {
        request = NewIsolateSubscribeRequest()
    }
    response = NewIsolateSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrateJobRequest() (request *ModifyMigrateJobRequest) {
    request = &ModifyMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateJob")
    return
}

func NewModifyMigrateJobResponse() (response *ModifyMigrateJobResponse) {
    response = &ModifyMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (ModifyMigrateJob) is used to modify a data migration task.
// If the status of a migration task is creating (status=1), check succeeded (status=4), check failed (status=5), or migration failed (status=10), this API can be called to modify the task, but the type of the source and target instances and the region of the target instance cannot be modified.
// 
// For a finance zone linkage, please use the domain name dts.ap-shenzhen-fsi.tencentcloudapi.com.
func (c *Client) ModifyMigrateJob(request *ModifyMigrateJobRequest) (response *ModifyMigrateJobResponse, err error) {
    if request == nil {
        request = NewModifyMigrateJobRequest()
    }
    response = NewModifyMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeConsumeTimeRequest() (request *ModifySubscribeConsumeTimeRequest) {
    request = &ModifySubscribeConsumeTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeConsumeTime")
    return
}

func NewModifySubscribeConsumeTimeResponse() (response *ModifySubscribeConsumeTimeResponse) {
    response = &ModifySubscribeConsumeTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the consumption time point of a data subscription channel.
func (c *Client) ModifySubscribeConsumeTime(request *ModifySubscribeConsumeTimeRequest) (response *ModifySubscribeConsumeTimeResponse, err error) {
    if request == nil {
        request = NewModifySubscribeConsumeTimeRequest()
    }
    response = NewModifySubscribeConsumeTimeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeNameRequest() (request *ModifySubscribeNameRequest) {
    request = &ModifySubscribeNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeName")
    return
}

func NewModifySubscribeNameResponse() (response *ModifySubscribeNameResponse) {
    response = &ModifySubscribeNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to rename a data subscription instance.
func (c *Client) ModifySubscribeName(request *ModifySubscribeNameRequest) (response *ModifySubscribeNameResponse, err error) {
    if request == nil {
        request = NewModifySubscribeNameRequest()
    }
    response = NewModifySubscribeNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeObjectsRequest() (request *ModifySubscribeObjectsRequest) {
    request = &ModifySubscribeObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeObjects")
    return
}

func NewModifySubscribeObjectsResponse() (response *ModifySubscribeObjectsResponse) {
    response = &ModifySubscribeObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the subscription rule of a data subscription channel.
func (c *Client) ModifySubscribeObjects(request *ModifySubscribeObjectsRequest) (response *ModifySubscribeObjectsResponse, err error) {
    if request == nil {
        request = NewModifySubscribeObjectsRequest()
    }
    response = NewModifySubscribeObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeVipVportRequest() (request *ModifySubscribeVipVportRequest) {
    request = &ModifySubscribeVipVportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeVipVport")
    return
}

func NewModifySubscribeVipVportResponse() (response *ModifySubscribeVipVportResponse) {
    response = &ModifySubscribeVipVportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the IP and port number of a data subscription instance.
func (c *Client) ModifySubscribeVipVport(request *ModifySubscribeVipVportRequest) (response *ModifySubscribeVipVportResponse, err error) {
    if request == nil {
        request = NewModifySubscribeVipVportRequest()
    }
    response = NewModifySubscribeVipVportResponse()
    err = c.Send(request, response)
    return
}

func NewModifySyncJobRequest() (request *ModifySyncJobRequest) {
    request = &ModifySyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifySyncJob")
    return
}

func NewModifySyncJobResponse() (response *ModifySyncJobResponse) {
    response = &ModifySyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify a disaster recovery sync task. 
// If the status of a sync task is creating, created, check succeeded, or check failed, this API can be called to modify the task. 
// The information of the source and target instances cannot be modified, but the task name and the tables to be synced can.
func (c *Client) ModifySyncJob(request *ModifySyncJobRequest) (response *ModifySyncJobResponse, err error) {
    if request == nil {
        request = NewModifySyncJobRequest()
    }
    response = NewModifySyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineIsolatedSubscribeRequest() (request *OfflineIsolatedSubscribeRequest) {
    request = &OfflineIsolatedSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "OfflineIsolatedSubscribe")
    return
}

func NewOfflineIsolatedSubscribeResponse() (response *OfflineIsolatedSubscribeResponse) {
    response = &OfflineIsolatedSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to deactivate an isolated data subscription instance.
func (c *Client) OfflineIsolatedSubscribe(request *OfflineIsolatedSubscribeRequest) (response *OfflineIsolatedSubscribeResponse, err error) {
    if request == nil {
        request = NewOfflineIsolatedSubscribeRequest()
    }
    response = NewOfflineIsolatedSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewResetSubscribeRequest() (request *ResetSubscribeRequest) {
    request = &ResetSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ResetSubscribe")
    return
}

func NewResetSubscribeResponse() (response *ResetSubscribeResponse) {
    response = &ResetSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to reset a data subscription instance. Once reset, an activated instance can be bound to other database instances through the `ActivateSubscribe` API.
func (c *Client) ResetSubscribe(request *ResetSubscribeRequest) (response *ResetSubscribeResponse, err error) {
    if request == nil {
        request = NewResetSubscribeRequest()
    }
    response = NewResetSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewStartMigrateJobRequest() (request *StartMigrateJobRequest) {
    request = &StartMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "StartMigrateJob")
    return
}

func NewStartMigrateJobResponse() (response *StartMigrateJobResponse) {
    response = &StartMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (StartMigrationJob) is used to start a migration task. After the API is called, non-scheduled migration tasks will start the migration immediately, while scheduled tasks will start the countdown.
// Before calling this API, be sure to use the CreateMigrateCheckJob API to check the data migration task, which can be started only if its status queried through the DescribeMigrateJobs API is check succeeded (status=4).
func (c *Client) StartMigrateJob(request *StartMigrateJobRequest) (response *StartMigrateJobResponse, err error) {
    if request == nil {
        request = NewStartMigrateJobRequest()
    }
    response = NewStartMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewStartSyncJobRequest() (request *StartSyncJobRequest) {
    request = &StartSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "StartSyncJob")
    return
}

func NewStartSyncJobResponse() (response *StartSyncJobResponse) {
    response = &StartSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to start a disaster recovery sync task after it is successfully checked through the CreateSyncCheckJob and DescribeSyncCheckJob APIs.
func (c *Client) StartSyncJob(request *StartSyncJobRequest) (response *StartSyncJobResponse, err error) {
    if request == nil {
        request = NewStartSyncJobRequest()
    }
    response = NewStartSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewStopMigrateJobRequest() (request *StopMigrateJobRequest) {
    request = &StopMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "StopMigrateJob")
    return
}

func NewStopMigrateJobResponse() (response *StopMigrateJobResponse) {
    response = &StopMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API (StopMigrateJob) is used to cancel a data migration task.
// During migration, this API can be used to cancel migration if the task status queried through the DescribeMigrateJobs API is running (status=7) or ready (status=8), and the migration task will fail.
func (c *Client) StopMigrateJob(request *StopMigrateJobRequest) (response *StopMigrateJobResponse, err error) {
    if request == nil {
        request = NewStopMigrateJobRequest()
    }
    response = NewStopMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchDrToMasterRequest() (request *SwitchDrToMasterRequest) {
    request = &SwitchDrToMasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "SwitchDrToMaster")
    return
}

func NewSwitchDrToMasterResponse() (response *SwitchDrToMasterResponse) {
    response = &SwitchDrToMasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to promote a disaster recovery instance to a master instance, which will stop sync from the original master instance and end the master/slave relationship.
func (c *Client) SwitchDrToMaster(request *SwitchDrToMasterRequest) (response *SwitchDrToMasterResponse, err error) {
    if request == nil {
        request = NewSwitchDrToMasterRequest()
    }
    response = NewSwitchDrToMasterResponse()
    err = c.Send(request, response)
    return
}
