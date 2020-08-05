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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAttachInstancesRequest() (request *AttachInstancesRequest) {
    request = &AttachInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "AttachInstances")
    return
}

func NewAttachInstancesResponse() (response *AttachInstancesResponse) {
    response = &AttachInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add existing instances to the compute environment.
// Considerations: <br/>
// 1. The instance should not be in the batch compute system.<br/>
// 2. The instance status should be 'running'.<br/>
// 3. It supports dedicated CVMs and pay-as-you-go instances that billed on an hourly basis. Spot instances are not supported.<b/>
// 
// For instances added to the compute environment, their UserData will be reset and the operating systems will be reinstalled.
func (c *Client) AttachInstances(request *AttachInstancesRequest) (response *AttachInstancesResponse, err error) {
    if request == nil {
        request = NewAttachInstancesRequest()
    }
    response = NewAttachInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateComputeEnvRequest() (request *CreateComputeEnvRequest) {
    request = &CreateComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "CreateComputeEnv")
    return
}

func NewCreateComputeEnvResponse() (response *CreateComputeEnvResponse) {
    response = &CreateComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a compute environment.
func (c *Client) CreateComputeEnv(request *CreateComputeEnvRequest) (response *CreateComputeEnvResponse, err error) {
    if request == nil {
        request = NewCreateComputeEnvRequest()
    }
    response = NewCreateComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskTemplateRequest() (request *CreateTaskTemplateRequest) {
    request = &CreateTaskTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "CreateTaskTemplate")
    return
}

func NewCreateTaskTemplateResponse() (response *CreateTaskTemplateResponse) {
    response = &CreateTaskTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a task template.
func (c *Client) CreateTaskTemplate(request *CreateTaskTemplateRequest) (response *CreateTaskTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTaskTemplateRequest()
    }
    response = NewCreateTaskTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteComputeEnvRequest() (request *DeleteComputeEnvRequest) {
    request = &DeleteComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DeleteComputeEnv")
    return
}

func NewDeleteComputeEnvResponse() (response *DeleteComputeEnvResponse) {
    response = &DeleteComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a compute environment.
func (c *Client) DeleteComputeEnv(request *DeleteComputeEnvRequest) (response *DeleteComputeEnvResponse, err error) {
    if request == nil {
        request = NewDeleteComputeEnvRequest()
    }
    response = NewDeleteComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJobRequest() (request *DeleteJobRequest) {
    request = &DeleteJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DeleteJob")
    return
}

func NewDeleteJobResponse() (response *DeleteJobResponse) {
    response = &DeleteJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a job.
// Deleting a job is equivalent to deleting all the information related to the job. After successful deletion, all information related to the job cannot be queried.
// The job to be deleted must be in a completed state, and all task instances contained in it must also be in a completed state; otherwise, the operation will be prohibited. The completed state can be either SUCCEED or FAILED.
func (c *Client) DeleteJob(request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    if request == nil {
        request = NewDeleteJobRequest()
    }
    response = NewDeleteJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskTemplatesRequest() (request *DeleteTaskTemplatesRequest) {
    request = &DeleteTaskTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DeleteTaskTemplates")
    return
}

func NewDeleteTaskTemplatesResponse() (response *DeleteTaskTemplatesResponse) {
    response = &DeleteTaskTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete task template information.
func (c *Client) DeleteTaskTemplates(request *DeleteTaskTemplatesRequest) (response *DeleteTaskTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteTaskTemplatesRequest()
    }
    response = NewDeleteTaskTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableCvmInstanceTypesRequest() (request *DescribeAvailableCvmInstanceTypesRequest) {
    request = &DescribeAvailableCvmInstanceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeAvailableCvmInstanceTypes")
    return
}

func NewDescribeAvailableCvmInstanceTypesResponse() (response *DescribeAvailableCvmInstanceTypesResponse) {
    response = &DescribeAvailableCvmInstanceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to view the information of available CVM model configurations.
func (c *Client) DescribeAvailableCvmInstanceTypes(request *DescribeAvailableCvmInstanceTypesRequest) (response *DescribeAvailableCvmInstanceTypesResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableCvmInstanceTypesRequest()
    }
    response = NewDescribeAvailableCvmInstanceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvRequest() (request *DescribeComputeEnvRequest) {
    request = &DescribeComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnv")
    return
}

func NewDescribeComputeEnvResponse() (response *DescribeComputeEnvResponse) {
    response = &DescribeComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query compute environment details.
func (c *Client) DescribeComputeEnv(request *DescribeComputeEnvRequest) (response *DescribeComputeEnvResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvRequest()
    }
    response = NewDescribeComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvActivitiesRequest() (request *DescribeComputeEnvActivitiesRequest) {
    request = &DescribeComputeEnvActivitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnvActivities")
    return
}

func NewDescribeComputeEnvActivitiesResponse() (response *DescribeComputeEnvActivitiesResponse) {
    response = &DescribeComputeEnvActivitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the information of activities in the compute environment.
func (c *Client) DescribeComputeEnvActivities(request *DescribeComputeEnvActivitiesRequest) (response *DescribeComputeEnvActivitiesResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvActivitiesRequest()
    }
    response = NewDescribeComputeEnvActivitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvCreateInfoRequest() (request *DescribeComputeEnvCreateInfoRequest) {
    request = &DescribeComputeEnvCreateInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnvCreateInfo")
    return
}

func NewDescribeComputeEnvCreateInfoResponse() (response *DescribeComputeEnvCreateInfoResponse) {
    response = &DescribeComputeEnvCreateInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Views compute environment creation information.
func (c *Client) DescribeComputeEnvCreateInfo(request *DescribeComputeEnvCreateInfoRequest) (response *DescribeComputeEnvCreateInfoResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvCreateInfoRequest()
    }
    response = NewDescribeComputeEnvCreateInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvCreateInfosRequest() (request *DescribeComputeEnvCreateInfosRequest) {
    request = &DescribeComputeEnvCreateInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnvCreateInfos")
    return
}

func NewDescribeComputeEnvCreateInfosResponse() (response *DescribeComputeEnvCreateInfosResponse) {
    response = &DescribeComputeEnvCreateInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to view the list of information of compute environment creation, including name, description, type, environment parameters, notifications, and number of desired nodes.
func (c *Client) DescribeComputeEnvCreateInfos(request *DescribeComputeEnvCreateInfosRequest) (response *DescribeComputeEnvCreateInfosResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvCreateInfosRequest()
    }
    response = NewDescribeComputeEnvCreateInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvsRequest() (request *DescribeComputeEnvsRequest) {
    request = &DescribeComputeEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnvs")
    return
}

func NewDescribeComputeEnvsResponse() (response *DescribeComputeEnvsResponse) {
    response = &DescribeComputeEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to view the list of compute environments.
func (c *Client) DescribeComputeEnvs(request *DescribeComputeEnvsRequest) (response *DescribeComputeEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvsRequest()
    }
    response = NewDescribeComputeEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCvmZoneInstanceConfigInfosRequest() (request *DescribeCvmZoneInstanceConfigInfosRequest) {
    request = &DescribeCvmZoneInstanceConfigInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeCvmZoneInstanceConfigInfos")
    return
}

func NewDescribeCvmZoneInstanceConfigInfosResponse() (response *DescribeCvmZoneInstanceConfigInfosResponse) {
    response = &DescribeCvmZoneInstanceConfigInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the model configuration information of the availability zone of BatchCompute.
func (c *Client) DescribeCvmZoneInstanceConfigInfos(request *DescribeCvmZoneInstanceConfigInfosRequest) (response *DescribeCvmZoneInstanceConfigInfosResponse, err error) {
    if request == nil {
        request = NewDescribeCvmZoneInstanceConfigInfosRequest()
    }
    response = NewDescribeCvmZoneInstanceConfigInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceCategoriesRequest() (request *DescribeInstanceCategoriesRequest) {
    request = &DescribeInstanceCategoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeInstanceCategories")
    return
}

func NewDescribeInstanceCategoriesResponse() (response *DescribeInstanceCategoriesResponse) {
    response = &DescribeInstanceCategoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Currently, CVM instance families are classified into different category, and each category contains several instance families. This API is used to query the instance category information.
func (c *Client) DescribeInstanceCategories(request *DescribeInstanceCategoriesRequest) (response *DescribeInstanceCategoriesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceCategoriesRequest()
    }
    response = NewDescribeInstanceCategoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobRequest() (request *DescribeJobRequest) {
    request = &DescribeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeJob")
    return
}

func NewDescribeJobResponse() (response *DescribeJobResponse) {
    response = &DescribeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to view the details of a job, including internal task (Task) and dependency (Dependence) information.
func (c *Client) DescribeJob(request *DescribeJobRequest) (response *DescribeJobResponse, err error) {
    if request == nil {
        request = NewDescribeJobRequest()
    }
    response = NewDescribeJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobSubmitInfoRequest() (request *DescribeJobSubmitInfoRequest) {
    request = &DescribeJobSubmitInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeJobSubmitInfo")
    return
}

func NewDescribeJobSubmitInfoResponse() (response *DescribeJobSubmitInfoResponse) {
    response = &DescribeJobSubmitInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the submission information of the specified job, with the return including the job submission information used as input parameters in the JobId and SubmitJob APIs.
func (c *Client) DescribeJobSubmitInfo(request *DescribeJobSubmitInfoRequest) (response *DescribeJobSubmitInfoResponse, err error) {
    if request == nil {
        request = NewDescribeJobSubmitInfoRequest()
    }
    response = NewDescribeJobSubmitInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobsRequest() (request *DescribeJobsRequest) {
    request = &DescribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeJobs")
    return
}

func NewDescribeJobsResponse() (response *DescribeJobsResponse) {
    response = &DescribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the overview information of several instances.
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeJobsRequest()
    }
    response = NewDescribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
    request = &DescribeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeTask")
    return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
    response = &DescribeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the details of a specified task, including information of the task instances inside the task.
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskLogsRequest() (request *DescribeTaskLogsRequest) {
    request = &DescribeTaskLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeTaskLogs")
    return
}

func NewDescribeTaskLogsResponse() (response *DescribeTaskLogsResponse) {
    response = &DescribeTaskLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the standard outputs and standard error logs of multiple task instances.
func (c *Client) DescribeTaskLogs(request *DescribeTaskLogsRequest) (response *DescribeTaskLogsResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLogsRequest()
    }
    response = NewDescribeTaskLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskTemplatesRequest() (request *DescribeTaskTemplatesRequest) {
    request = &DescribeTaskTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DescribeTaskTemplates")
    return
}

func NewDescribeTaskTemplatesResponse() (response *DescribeTaskTemplatesResponse) {
    response = &DescribeTaskTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to query the information of task templates.
func (c *Client) DescribeTaskTemplates(request *DescribeTaskTemplatesRequest) (response *DescribeTaskTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTaskTemplatesRequest()
    }
    response = NewDescribeTaskTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDetachInstancesRequest() (request *DetachInstancesRequest) {
    request = &DetachInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "DetachInstances")
    return
}

func NewDetachInstancesResponse() (response *DetachInstancesResponse) {
    response = &DetachInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to remove instances that from compute environment. 
func (c *Client) DetachInstances(request *DetachInstancesRequest) (response *DetachInstancesResponse, err error) {
    if request == nil {
        request = NewDetachInstancesRequest()
    }
    response = NewDetachInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyComputeEnvRequest() (request *ModifyComputeEnvRequest) {
    request = &ModifyComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "ModifyComputeEnv")
    return
}

func NewModifyComputeEnvResponse() (response *ModifyComputeEnvResponse) {
    response = &ModifyComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the attributes of a compute environment.
func (c *Client) ModifyComputeEnv(request *ModifyComputeEnvRequest) (response *ModifyComputeEnvResponse, err error) {
    if request == nil {
        request = NewModifyComputeEnvRequest()
    }
    response = NewModifyComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskTemplateRequest() (request *ModifyTaskTemplateRequest) {
    request = &ModifyTaskTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "ModifyTaskTemplate")
    return
}

func NewModifyTaskTemplateResponse() (response *ModifyTaskTemplateResponse) {
    response = &ModifyTaskTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify a task template.
func (c *Client) ModifyTaskTemplate(request *ModifyTaskTemplateRequest) (response *ModifyTaskTemplateResponse, err error) {
    if request == nil {
        request = NewModifyTaskTemplateRequest()
    }
    response = NewModifyTaskTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewRetryJobsRequest() (request *RetryJobsRequest) {
    request = &RetryJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "RetryJobs")
    return
}

func NewRetryJobsResponse() (response *RetryJobsResponse) {
    response = &RetryJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to retry the failing task instance in instances.
// Job retry is supported only if a job is in the "FAILED" state. After the retry operation succeeds, the instance will retry the failing task instances in each task in turn according to the task dependencies specified in the "DAG". The history information of the task instances will be reset, the instances will participate in subsequent scheduling and execution as if they are run for the first time.
func (c *Client) RetryJobs(request *RetryJobsRequest) (response *RetryJobsResponse, err error) {
    if request == nil {
        request = NewRetryJobsRequest()
    }
    response = NewRetryJobsResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitJobRequest() (request *SubmitJobRequest) {
    request = &SubmitJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "SubmitJob")
    return
}

func NewSubmitJobResponse() (response *SubmitJobResponse) {
    response = &SubmitJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to submit a instance.
func (c *Client) SubmitJob(request *SubmitJobRequest) (response *SubmitJobResponse, err error) {
    if request == nil {
        request = NewSubmitJobRequest()
    }
    response = NewSubmitJobResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateComputeNodeRequest() (request *TerminateComputeNodeRequest) {
    request = &TerminateComputeNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "TerminateComputeNode")
    return
}

func NewTerminateComputeNodeResponse() (response *TerminateComputeNodeResponse) {
    response = &TerminateComputeNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to terminate a compute node.
// Termination is allowed for nodes in the CREATED, CREATION_FAILED, RUNNING or ABNORMAL state.
func (c *Client) TerminateComputeNode(request *TerminateComputeNodeRequest) (response *TerminateComputeNodeResponse, err error) {
    if request == nil {
        request = NewTerminateComputeNodeRequest()
    }
    response = NewTerminateComputeNodeResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateComputeNodesRequest() (request *TerminateComputeNodesRequest) {
    request = &TerminateComputeNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "TerminateComputeNodes")
    return
}

func NewTerminateComputeNodesResponse() (response *TerminateComputeNodesResponse) {
    response = &TerminateComputeNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to terminate compute nodes in batches. It is not allowed to repeatedly terminate the same node.
func (c *Client) TerminateComputeNodes(request *TerminateComputeNodesRequest) (response *TerminateComputeNodesResponse, err error) {
    if request == nil {
        request = NewTerminateComputeNodesRequest()
    }
    response = NewTerminateComputeNodesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateJobRequest() (request *TerminateJobRequest) {
    request = &TerminateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "TerminateJob")
    return
}

func NewTerminateJobResponse() (response *TerminateJobResponse) {
    response = &TerminateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to terminate an instance.
// Termination is prohibited if an instance is in the "SUBMITTED" state and does not take effect if it is in the "SUCCEED" state.
// Instance termination is an asynchronous process, and the time it takes to complete the entire process is directly proportional to the total number of tasks. The effect of terminating an instance is equivalent to performing the TerminateTaskInstance operation on all the task instances contained in the job. For more information on the specific effect and usage, see TerminateTaskInstance.
func (c *Client) TerminateJob(request *TerminateJobRequest) (response *TerminateJobResponse, err error) {
    if request == nil {
        request = NewTerminateJobRequest()
    }
    response = NewTerminateJobResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateTaskInstanceRequest() (request *TerminateTaskInstanceRequest) {
    request = &TerminateTaskInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("batch", APIVersion, "TerminateTaskInstance")
    return
}

func NewTerminateTaskInstanceResponse() (response *TerminateTaskInstanceResponse) {
    response = &TerminateTaskInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to terminate a task instance.
// Task instances in the "SUCCEED" or "FAILED" state will not be processed.
// Task instances in the "SUBMITTED", "PENDING", or "RUNNABLE" state will be set to the "FAILED" state.
// For task instances in the "STARTING", "RUNNING", or "FAILED_INTERRUPTED" state, there will be two cases: if no compute environment is specified, the CVM instance will be terminated first, and then the state will be set to "FAILED"; if a compute environment's EnvId is specified, the state of the task instances will be set to "FAILED" first, and then the CVM instance that performs the task will be restarted. Both cases takes a certain amount of time to be completed.
// For task instances in the "FAILED_INTERRUPTED" state, the related resources and quotas will be released only after the termination actually succeeds.
func (c *Client) TerminateTaskInstance(request *TerminateTaskInstanceRequest) (response *TerminateTaskInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateTaskInstanceRequest()
    }
    response = NewTerminateTaskInstanceResponse()
    err = c.Send(request, response)
    return
}
