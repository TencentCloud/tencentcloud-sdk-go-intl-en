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

package v20200303

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-03-03"

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


func NewAnalyzeFaceRequest() (request *AnalyzeFaceRequest) {
    request = &AnalyzeFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "AnalyzeFace")
    return
}

func NewAnalyzeFaceResponse() (response *AnalyzeFaceResponse) {
    response = &AnalyzeFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to perform facial feature localization (aka facial keypoint localization) on a given image and calculate 90 facial keypoints that make up the contour of the face, including eyebrows (8 points on the left and 8 on the right), eyes (8 points on the left and 8 on the right), nose (13 points), mouth (22 points), face contour (21 points), and eyeballs or pupils (2 points).
// 
// >     
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
func (c *Client) AnalyzeFace(request *AnalyzeFaceRequest) (response *AnalyzeFaceResponse, err error) {
    if request == nil {
        request = NewAnalyzeFaceRequest()
    }
    response = NewAnalyzeFaceResponse()
    err = c.Send(request, response)
    return
}

func NewCheckSimilarPersonRequest() (request *CheckSimilarPersonRequest) {
    request = &CheckSimilarPersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "CheckSimilarPerson")
    return
}

func NewCheckSimilarPersonResponse() (response *CheckSimilarPersonResponse) {
    response = &CheckSimilarPersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to check a specified group for suspected duplicate persons and list their information.
// 
// You can use this API to check for duplicate persons in one group so as to avoid situations where the same person has multiple roles in the group. You can also use it to check for duplicate persons across multiple groups to see whether the same person exists in multiple groups at the same time.
// 
// Duplicate check across algorithm model versions is not supported. Currently, this feature is available only to groups with algorithm model v3.0.
// 
// >     
// - If you perform a duplicate check on the same group again, you need to wait for the last operation to complete, that is, when the `GroupIds` entered in the two requests are the same, if the first request is not completed, the second request will fail.
// 
// >     
// - The status of the group on which the duplicate check is to be performed is that when the duplicate check task really starts, that is, after you initiate the duplicate check request; if your duplicate check task needs to queue up, any addition or deletion operation performed on the group during the queuing will affect the duplicate check result. Tencent Cloud will use the group status when the duplicate check task actually starts. After the task starts, any operation on the group will not affect the task execution; however, you are still recommended not to add/delete persons or faces to/from the group after the task starts.
func (c *Client) CheckSimilarPerson(request *CheckSimilarPersonRequest) (response *CheckSimilarPersonResponse, err error) {
    if request == nil {
        request = NewCheckSimilarPersonRequest()
    }
    response = NewCheckSimilarPersonResponse()
    err = c.Send(request, response)
    return
}

func NewCompareFaceRequest() (request *CompareFaceRequest) {
    request = &CompareFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "CompareFace")
    return
}

func NewCompareFaceResponse() (response *CompareFaceResponse) {
    response = &CompareFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to calculate the similarity of faces in two images and return the face similarity score.
// 
// If you need to judge "whether the person in the image is someone specified" in scenarios such as face login, i.e., checking whether the person in a given image is someone with a known identity, you are recommended to use the [VerifyFace](https://cloud.tencent.com/document/product/867/32806) or [VerifyPerson](https://cloud.tencent.com/document/product/867/38879) API.
// 
// >     
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
func (c *Client) CompareFace(request *CompareFaceRequest) (response *CompareFaceResponse, err error) {
    if request == nil {
        request = NewCompareFaceRequest()
    }
    response = NewCompareFaceResponse()
    err = c.Send(request, response)
    return
}

func NewCopyPersonRequest() (request *CopyPersonRequest) {
    request = &CopyPersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "CopyPerson")
    return
}

func NewCopyPersonResponse() (response *CopyPersonResponse) {
    response = &CopyPersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to copy a person in a group to another group (without copying the description). One person can exist in up to 100 groups at the same time.
// >     
// - Note: if the version of the algorithm model was 2.0 when the person was created, the copy operation will fail if it is to copy to a group not on algorithm model v2.0.
func (c *Client) CopyPerson(request *CopyPersonRequest) (response *CopyPersonResponse, err error) {
    if request == nil {
        request = NewCopyPersonRequest()
    }
    response = NewCopyPersonResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFaceRequest() (request *CreateFaceRequest) {
    request = &CreateFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "CreateFace")
    return
}

func NewCreateFaceResponse() (response *CreateFaceResponse) {
    response = &CreateFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add a set of face images to a person. One person can have up to 5 images. If a person exists in multiple groups, the images will be added to all those groups for the person.
// 
// >     
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
func (c *Client) CreateFace(request *CreateFaceRequest) (response *CreateFaceResponse, err error) {
    if request == nil {
        request = NewCreateFaceRequest()
    }
    response = NewCreateFaceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "CreateGroup")
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create an empty group. If the group already exists, an error will be returned.
// Custom description fields can be created as needed to describe persons in the group.
// 
// A maximum of 100,000 groups or 50 million faces can be created under one `APPID`.
// 
// The maximum number of faces that can be included in one group varies by algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePersonRequest() (request *CreatePersonRequest) {
    request = &CreatePersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "CreatePerson")
    return
}

func NewCreatePersonResponse() (response *CreatePersonResponse) {
    response = &CreatePersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a person and add face, name, gender, and other related information.
// 
// >     
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
func (c *Client) CreatePerson(request *CreatePersonRequest) (response *CreatePersonResponse, err error) {
    if request == nil {
        request = NewCreatePersonRequest()
    }
    response = NewCreatePersonResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFaceRequest() (request *DeleteFaceRequest) {
    request = &DeleteFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DeleteFace")
    return
}

func NewDeleteFaceResponse() (response *DeleteFaceResponse) {
    response = &DeleteFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete the face images of a person. If the person has only one face image, an error will be returned.
func (c *Client) DeleteFace(request *DeleteFaceRequest) (response *DeleteFaceResponse, err error) {
    if request == nil {
        request = NewDeleteFaceRequest()
    }
    response = NewDeleteFaceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DeleteGroup")
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a group and all persons in it. Meanwhile, all face information corresponding to the persons will be deleted. If a person exists in multiple groups at the same time, deleting a group will not delete the person, but the custom description field information in the group will be deleted. Custom description field information in other groups will not be affected.
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePersonRequest() (request *DeletePersonRequest) {
    request = &DeletePersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DeletePerson")
    return
}

func NewDeletePersonResponse() (response *DeletePersonResponse) {
    response = &DeletePersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a person from all groups. Meanwhile, all face information of the person will be deleted.
func (c *Client) DeletePerson(request *DeletePersonRequest) (response *DeletePersonResponse, err error) {
    if request == nil {
        request = NewDeletePersonRequest()
    }
    response = NewDeletePersonResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePersonFromGroupRequest() (request *DeletePersonFromGroupRequest) {
    request = &DeletePersonFromGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DeletePersonFromGroup")
    return
}

func NewDeletePersonFromGroupResponse() (response *DeletePersonFromGroupResponse) {
    response = &DeletePersonFromGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to remove a person from a specified group. This operation only affects the group. If the person exists only in the group, the person will be deleted, and all face information of the person will also be deleted.
func (c *Client) DeletePersonFromGroup(request *DeletePersonFromGroupRequest) (response *DeletePersonFromGroupResponse, err error) {
    if request == nil {
        request = NewDeletePersonFromGroupRequest()
    }
    response = NewDeletePersonFromGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDetectFaceRequest() (request *DetectFaceRequest) {
    request = &DetectFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DetectFace")
    return
}

func NewDetectFaceResponse() (response *DetectFaceResponse) {
    response = &DetectFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to detect the position, attributes, and quality information of a face in the given image. The position information includes (x, y, w, h); the face attributes include gender, age, expression, beauty, glass, hair, mask, and pose (pitch, roll, yaw); and the face quality information includes the overall quality score, sharpness, brightness, and completeness.
// 
//  
// The face quality information is mainly used to evaluate the quality of the input face image. When using the Face Recognition service, you are recommended to evaluate the quality of the input face image first to improve the effects of subsequent processing. Application scenarios of this feature include:
// 
// 1). [Creating](https://cloud.tencent.com/document/product/867/32793)/[Adding](https://cloud.tencent.com/document/product/867/32795) a person in a group: this is to ensure the quality of the face information to facilitate subsequent processing.
// 
// 2). [Face search](https://cloud.tencent.com/document/product/867/32798): this is to ensure the quality of the input image to quickly find the corresponding person.
// 
// 3). [Face verification](https://cloud.tencent.com/document/product/867/32806): this is to ensure the quality of the face information to avoid cases where the verification incorrectly fails.
// 
// 4). [Face fusion](https://cloud.tencent.com/product/facefusion): this is to ensure the quality of the uploaded face images to improve the fusion effect.
// 
// >     
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
// 
func (c *Client) DetectFace(request *DetectFaceRequest) (response *DetectFaceResponse, err error) {
    if request == nil {
        request = NewDetectFaceRequest()
    }
    response = NewDetectFaceResponse()
    err = c.Send(request, response)
    return
}

func NewDetectLiveFaceRequest() (request *DetectLiveFaceRequest) {
    request = &DetectLiveFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "DetectLiveFace")
    return
}

func NewDetectLiveFaceResponse() (response *DetectLiveFaceResponse) {
    response = &DetectLiveFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to detect the liveness of a user with a user-uploaded image. Its difference from video-based liveness detection lies in that the user does not need to speak, shake their head, or wink for detection.
// 
// Image-based liveness detection is suitable for scenarios where the image is a selfie or the requirement for attack defense is not high. If you have a higher security requirement for liveness detection, please use [Faceid](https://cloud.tencent.com/product/faceid).
// 
// >     
// - The aspect ratio of the image should be close to 3:4 (width:height); otherwise, the score returned for the image will be meaningless. This API is suitable for selfie scenarios, and the score returned in other scenarios will be meaningless.
// 
// >     
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
func (c *Client) DetectLiveFace(request *DetectLiveFaceRequest) (response *DetectLiveFaceResponse, err error) {
    if request == nil {
        request = NewDetectLiveFaceRequest()
    }
    response = NewDetectLiveFaceResponse()
    err = c.Send(request, response)
    return
}

func NewEstimateCheckSimilarPersonCostTimeRequest() (request *EstimateCheckSimilarPersonCostTimeRequest) {
    request = &EstimateCheckSimilarPersonCostTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "EstimateCheckSimilarPersonCostTime")
    return
}

func NewEstimateCheckSimilarPersonCostTimeResponse() (response *EstimateCheckSimilarPersonCostTimeResponse) {
    response = &EstimateCheckSimilarPersonCostTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the estimated duration of a duplicate person check task.
// 
// If the `EndTimestamp` meets your expectations, please initiate the duplicate person check request as soon as possible; otherwise, the task may take more time.
// 
// If the estimated duration is more than 5 hours, the duplicate person check feature cannot be used.
func (c *Client) EstimateCheckSimilarPersonCostTime(request *EstimateCheckSimilarPersonCostTimeRequest) (response *EstimateCheckSimilarPersonCostTimeResponse, err error) {
    if request == nil {
        request = NewEstimateCheckSimilarPersonCostTimeRequest()
    }
    response = NewEstimateCheckSimilarPersonCostTimeResponse()
    err = c.Send(request, response)
    return
}

func NewGetCheckSimilarPersonJobIdListRequest() (request *GetCheckSimilarPersonJobIdListRequest) {
    request = &GetCheckSimilarPersonJobIdListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetCheckSimilarPersonJobIdList")
    return
}

func NewGetCheckSimilarPersonJobIdListResponse() (response *GetCheckSimilarPersonJobIdListResponse) {
    response = &GetCheckSimilarPersonJobIdListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of duplicate person check tasks and sort them in reverse order by task creation time (i.e., the newest one is at the top)
// 
// Only data in the past year is retained.
func (c *Client) GetCheckSimilarPersonJobIdList(request *GetCheckSimilarPersonJobIdListRequest) (response *GetCheckSimilarPersonJobIdListResponse, err error) {
    if request == nil {
        request = NewGetCheckSimilarPersonJobIdListRequest()
    }
    response = NewGetCheckSimilarPersonJobIdListResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupInfoRequest() (request *GetGroupInfoRequest) {
    request = &GetGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetGroupInfo")
    return
}

func NewGetGroupInfoResponse() (response *GetGroupInfoResponse) {
    response = &GetGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the group information.
func (c *Client) GetGroupInfo(request *GetGroupInfoRequest) (response *GetGroupInfoResponse, err error) {
    if request == nil {
        request = NewGetGroupInfoRequest()
    }
    response = NewGetGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupListRequest() (request *GetGroupListRequest) {
    request = &GetGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetGroupList")
    return
}

func NewGetGroupListResponse() (response *GetGroupListResponse) {
    response = &GetGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of groups.
func (c *Client) GetGroupList(request *GetGroupListRequest) (response *GetGroupListResponse, err error) {
    if request == nil {
        request = NewGetGroupListRequest()
    }
    response = NewGetGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPersonBaseInfoRequest() (request *GetPersonBaseInfoRequest) {
    request = &GetPersonBaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetPersonBaseInfo")
    return
}

func NewGetPersonBaseInfoResponse() (response *GetPersonBaseInfoResponse) {
    response = &GetPersonBaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the information of a specified person, including name, gender, face, etc.
func (c *Client) GetPersonBaseInfo(request *GetPersonBaseInfoRequest) (response *GetPersonBaseInfoResponse, err error) {
    if request == nil {
        request = NewGetPersonBaseInfoRequest()
    }
    response = NewGetPersonBaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetPersonGroupInfoRequest() (request *GetPersonGroupInfoRequest) {
    request = &GetPersonGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetPersonGroupInfo")
    return
}

func NewGetPersonGroupInfoResponse() (response *GetPersonGroupInfoResponse) {
    response = &GetPersonGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the information of a specified person, including group, description, etc.
func (c *Client) GetPersonGroupInfo(request *GetPersonGroupInfoRequest) (response *GetPersonGroupInfoResponse, err error) {
    if request == nil {
        request = NewGetPersonGroupInfoRequest()
    }
    response = NewGetPersonGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetPersonListRequest() (request *GetPersonListRequest) {
    request = &GetPersonListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetPersonList")
    return
}

func NewGetPersonListResponse() (response *GetPersonListResponse) {
    response = &GetPersonListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of persons in a specified group.
func (c *Client) GetPersonList(request *GetPersonListRequest) (response *GetPersonListResponse, err error) {
    if request == nil {
        request = NewGetPersonListRequest()
    }
    response = NewGetPersonListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPersonListNumRequest() (request *GetPersonListNumRequest) {
    request = &GetPersonListNumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetPersonListNum")
    return
}

func NewGetPersonListNumResponse() (response *GetPersonListNumResponse) {
    response = &GetPersonListNumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the number of persons in a specified group.
func (c *Client) GetPersonListNum(request *GetPersonListNumRequest) (response *GetPersonListNumResponse, err error) {
    if request == nil {
        request = NewGetPersonListNumRequest()
    }
    response = NewGetPersonListNumResponse()
    err = c.Send(request, response)
    return
}

func NewGetSimilarPersonResultRequest() (request *GetSimilarPersonResultRequest) {
    request = &GetSimilarPersonResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "GetSimilarPersonResult")
    return
}

func NewGetSimilarPersonResultResponse() (response *GetSimilarPersonResultResponse) {
    response = &GetSimilarPersonResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the result of the `CheckSimilarPerson` API.
func (c *Client) GetSimilarPersonResult(request *GetSimilarPersonResultRequest) (response *GetSimilarPersonResultResponse, err error) {
    if request == nil {
        request = NewGetSimilarPersonResultRequest()
    }
    response = NewGetSimilarPersonResultResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupRequest() (request *ModifyGroupRequest) {
    request = &ModifyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "ModifyGroup")
    return
}

func NewModifyGroupResponse() (response *ModifyGroupResponse) {
    response = &ModifyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the name, tag, and custom description field of a group.
func (c *Client) ModifyGroup(request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    if request == nil {
        request = NewModifyGroupRequest()
    }
    response = NewModifyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonBaseInfoRequest() (request *ModifyPersonBaseInfoRequest) {
    request = &ModifyPersonBaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "ModifyPersonBaseInfo")
    return
}

func NewModifyPersonBaseInfoResponse() (response *ModifyPersonBaseInfoResponse) {
    response = &ModifyPersonBaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the information of a person, including name, gender, etc. The changes of person name and gender will be synced to all the groups that contain the person.
func (c *Client) ModifyPersonBaseInfo(request *ModifyPersonBaseInfoRequest) (response *ModifyPersonBaseInfoResponse, err error) {
    if request == nil {
        request = NewModifyPersonBaseInfoRequest()
    }
    response = NewModifyPersonBaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonGroupInfoRequest() (request *ModifyPersonGroupInfoRequest) {
    request = &ModifyPersonGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "ModifyPersonGroupInfo")
    return
}

func NewModifyPersonGroupInfoResponse() (response *ModifyPersonGroupInfoResponse) {
    response = &ModifyPersonGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the description of a specified person in a group.
func (c *Client) ModifyPersonGroupInfo(request *ModifyPersonGroupInfoRequest) (response *ModifyPersonGroupInfoResponse, err error) {
    if request == nil {
        request = NewModifyPersonGroupInfoRequest()
    }
    response = NewModifyPersonGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewSearchFacesRequest() (request *SearchFacesRequest) {
    request = &SearchFacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "SearchFaces")
    return
}

func NewSearchFacesResponse() (response *SearchFacesResponse) {
    response = &SearchFacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to recognize top K persons in one or more groups who are similar to the person in a given image and rank the similarity in a descending order.
// 
// Up to 10 faces in an image can be recognized at a time, and up to 100 groups can be searched in at a time.
// 
// The maximum number of faces in a group that can be searched for at a time is subject to the group's algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
// 
// This API recognizes each face of a person as an independent one. By contrast, the [SearchPersons](https://cloud.tencent.com/document/product/867/38881) and [SearchPersonsReturnsByGroup](https://cloud.tencent.com/document/product/867/38880) APIs fuse the features of all faces of a person; for example, if a person has 4 faces, they will fuse the features of the 4 faces and generate the summarized facial features of the person to make the search more accurate.
// 
// 
// This API should be used together with the [CreateGroup API](https://cloud.tencent.com/document/product/867/32794).
// 
// >     
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
func (c *Client) SearchFaces(request *SearchFacesRequest) (response *SearchFacesResponse, err error) {
    if request == nil {
        request = NewSearchFacesRequest()
    }
    response = NewSearchFacesResponse()
    err = c.Send(request, response)
    return
}

func NewSearchFacesReturnsByGroupRequest() (request *SearchFacesReturnsByGroupRequest) {
    request = &SearchFacesReturnsByGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "SearchFacesReturnsByGroup")
    return
}

func NewSearchFacesReturnsByGroupResponse() (response *SearchFacesReturnsByGroupResponse) {
    response = &SearchFacesReturnsByGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to recognize top K persons in one or more groups who are similar to the person in a given image, display the results **by group**, and rank the similarity within each group in a descending order.
// 
// Up to 10 faces in the image can be recognized at a time, and cross-group search is supported.
// 
// The maximum number of faces in a group that can be searched for at a time is subject to the group's algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
// 
// This API recognizes each face of a person as an independent one. By contrast, the [SearchPersons](https://cloud.tencent.com/document/product/867/38881) and [SearchPersonsReturnsByGroup](https://cloud.tencent.com/document/product/867/38880) APIs fuse the features of all faces of a person; for example, if a person has 4 faces, they will fuse the features of the 4 faces and generate the summarized facial features of the person to make the search more accurate.
// 
// This API should be used together with the [CreateGroup API](https://cloud.tencent.com/document/product/867/32794).
// 
// >     
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
// 
func (c *Client) SearchFacesReturnsByGroup(request *SearchFacesReturnsByGroupRequest) (response *SearchFacesReturnsByGroupResponse, err error) {
    if request == nil {
        request = NewSearchFacesReturnsByGroupRequest()
    }
    response = NewSearchFacesReturnsByGroupResponse()
    err = c.Send(request, response)
    return
}

func NewSearchPersonsRequest() (request *SearchPersonsRequest) {
    request = &SearchPersonsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "SearchPersons")
    return
}

func NewSearchPersonsResponse() (response *SearchPersonsResponse) {
    response = &SearchPersonsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to recognize top K persons in one or more groups who are similar to the person in a given image and rank the similarity in a descending order.
// 
// Up to 10 faces in an image can be recognized at a time, and up to 100 groups can be searched in at a time.
// 
// The maximum number of faces in a group that can be searched for at a time is subject to the group's algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
// 
// This API fuses the features of all faces of a person; for example, if a person has 4 faces, it will fuse the features of the 4 faces and generate the summarized facial features of the person to make the person search (i.e., judging whether the face image to be recognized is of a specified person) more accurate. By contrast, the [SearchFaces](https://cloud.tencent.com/document/product/867/32798) and [SearchFacesReturnsByGroup](https://cloud.tencent.com/document/product/867/38882) APIs recognize each face of a person as an independent one for search.
// 
// >     
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
// - This feature is available only to groups whose algorithm model version (`FaceModelVersion`) is 3.0.
func (c *Client) SearchPersons(request *SearchPersonsRequest) (response *SearchPersonsResponse, err error) {
    if request == nil {
        request = NewSearchPersonsRequest()
    }
    response = NewSearchPersonsResponse()
    err = c.Send(request, response)
    return
}

func NewSearchPersonsReturnsByGroupRequest() (request *SearchPersonsReturnsByGroupRequest) {
    request = &SearchPersonsReturnsByGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "SearchPersonsReturnsByGroup")
    return
}

func NewSearchPersonsReturnsByGroupResponse() (response *SearchPersonsReturnsByGroupResponse) {
    response = &SearchPersonsReturnsByGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to recognize top K persons in one or more groups who are similar to the person in a given image, display the results **by group**, and rank the similarity within each group in a descending order.
// 
// Up to 10 faces in the image can be recognized at a time, and cross-group search is supported.
// 
// The maximum number of faces in a group that can be searched for at a time is subject to the group's algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
// 
// This API fuses the features of all faces of a person; for example, if a person has 4 faces, it will fuse the features of the 4 faces and generate the summarized facial features of the person to make the person search (i.e., judging whether the face image to be recognized is of a specified person) more accurate. By contrast, the [SearchFaces](https://cloud.tencent.com/document/product/867/32798) and [SearchFacesReturnsByGroup](https://cloud.tencent.com/document/product/867/38882) APIs recognize each face of a person as an independent one for search.
// >     
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
// - This feature is available only to groups whose algorithm model version (`FaceModelVersion`) is 3.0.
func (c *Client) SearchPersonsReturnsByGroup(request *SearchPersonsReturnsByGroupRequest) (response *SearchPersonsReturnsByGroupResponse, err error) {
    if request == nil {
        request = NewSearchPersonsReturnsByGroupRequest()
    }
    response = NewSearchPersonsReturnsByGroupResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyFaceRequest() (request *VerifyFaceRequest) {
    request = &VerifyFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "VerifyFace")
    return
}

func NewVerifyFaceResponse() (response *VerifyFaceResponse) {
    response = &VerifyFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to judge whether a person in an image corresponds to a given `PersonId`. For more information on `PersonId`, please see [CreateGroup](https://cloud.tencent.com/document/product/867/32794). 
// 
// Unlike the [CompareFace](https://cloud.tencent.com/document/product/867/32802) API that is used to judge the similarity between two faces, this API is used to judge "whether the person in the image is someone specified" whose information is stored in a group. This "someone" may have multiple face images.
// 
// This API recognizes each face of a person as an independent one. By contrast, the [VerifyPerson](https://cloud.tencent.com/document/product/867/38879) API fuses the features of all faces of a person; for example, if a person has 4 faces, it will fuse the features of the 4 faces and generate the summarized facial features of the person to make the person verification (i.e., judging whether the face image to be recognized is of a specified person) more accurate.
// 
// >     
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
func (c *Client) VerifyFace(request *VerifyFaceRequest) (response *VerifyFaceResponse, err error) {
    if request == nil {
        request = NewVerifyFaceRequest()
    }
    response = NewVerifyFaceResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyPersonRequest() (request *VerifyPersonRequest) {
    request = &VerifyPersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iai", APIVersion, "VerifyPerson")
    return
}

func NewVerifyPersonResponse() (response *VerifyPersonResponse) {
    response = &VerifyPersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to judge whether a person in an image corresponds to a given `PersonId`. For more information on `PersonId`, please see [CreateGroup](https://cloud.tencent.com/document/product/867/32794).
// This API fuses the features of all faces of a person; for example, if a person has 4 faces, it will fuse the features of the 4 faces and generate the summarized facial features of the person to make the person verification (i.e., judging whether the face image to be recognized is of a specified person) more accurate.
// 
//  Unlike the `CompareFace` API that is used to judge the similarity between two faces, this API is used to judge "whether the person in the image is someone specified" whose information is stored in a group. This "someone" may have multiple face images.
// 
// 
// >     
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
// - This feature is available only to groups whose algorithm model version (`FaceModelVersion`) is 3.0.
func (c *Client) VerifyPerson(request *VerifyPersonRequest) (response *VerifyPersonResponse, err error) {
    if request == nil {
        request = NewVerifyPersonRequest()
    }
    response = NewVerifyPersonResponse()
    err = c.Send(request, response)
    return
}
