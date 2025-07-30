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

package v20200303

import (
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// AnalyzeFace
// This API is used to perform facial feature localization (aka facial keypoint localization) on a given image and calculate 90 facial keypoints that make up the contour of the face, including eyebrows (8 points on the left and 8 on the right), eyes (8 points on the left and 8 on the right), nose (13 points), mouth (22 points), face contour (21 points), and eyeballs or pupils (2 points).
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) AnalyzeFace(request *AnalyzeFaceRequest) (response *AnalyzeFaceResponse, err error) {
    return c.AnalyzeFaceWithContext(context.Background(), request)
}

// AnalyzeFace
// This API is used to perform facial feature localization (aka facial keypoint localization) on a given image and calculate 90 facial keypoints that make up the contour of the face, including eyebrows (8 points on the left and 8 on the right), eyes (8 points on the left and 8 on the right), nose (13 points), mouth (22 points), face contour (21 points), and eyeballs or pupils (2 points).
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) AnalyzeFaceWithContext(ctx context.Context, request *AnalyzeFaceRequest) (response *AnalyzeFaceResponse, err error) {
    if request == nil {
        request = NewAnalyzeFaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "AnalyzeFace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AnalyzeFace require credential")
    }

    request.SetContext(ctx)
    
    response = NewAnalyzeFaceResponse()
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

// CompareFace
// This API is used to calculate the similarity of faces in two images and return the face similarity score.
//
// 
//
// If you need to judge "whether the person in the image is someone specified" in scenarios such as face login, i.e., checking whether the person in a given image is someone with a known identity, we recommend using the [VerifyFace](https://intl.cloud.tencent.com/document/product/867/44983?from_cn_redirect=1) or [VerifyPerson](https://intl.cloud.tencent.com/document/product/867/44982?from_cn_redirect=1) API.
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CompareFace(request *CompareFaceRequest) (response *CompareFaceResponse, err error) {
    return c.CompareFaceWithContext(context.Background(), request)
}

// CompareFace
// This API is used to calculate the similarity of faces in two images and return the face similarity score.
//
// 
//
// If you need to judge "whether the person in the image is someone specified" in scenarios such as face login, i.e., checking whether the person in a given image is someone with a known identity, we recommend using the [VerifyFace](https://intl.cloud.tencent.com/document/product/867/44983?from_cn_redirect=1) or [VerifyPerson](https://intl.cloud.tencent.com/document/product/867/44982?from_cn_redirect=1) API.
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CompareFaceWithContext(ctx context.Context, request *CompareFaceRequest) (response *CompareFaceResponse, err error) {
    if request == nil {
        request = NewCompareFaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "CompareFace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CompareFace require credential")
    }

    request.SetContext(ctx)
    
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

// CopyPerson
// This API is used to copy a person in a group to another group (without copying the description). One person can exist in up to 100 groups at the same time.
//
// >- Note: in the case that the version of the algorithm model was 2.0 when the person was created, the copy operation will fail if the target group is not of algorithm model 2.0.
//
// error code that may be returned:
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CopyPerson(request *CopyPersonRequest) (response *CopyPersonResponse, err error) {
    return c.CopyPersonWithContext(context.Background(), request)
}

// CopyPerson
// This API is used to copy a person in a group to another group (without copying the description). One person can exist in up to 100 groups at the same time.
//
// >- Note: in the case that the version of the algorithm model was 2.0 when the person was created, the copy operation will fail if the target group is not of algorithm model 2.0.
//
// error code that may be returned:
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CopyPersonWithContext(ctx context.Context, request *CopyPersonRequest) (response *CopyPersonResponse, err error) {
    if request == nil {
        request = NewCopyPersonRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "CopyPerson")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyPerson require credential")
    }

    request.SetContext(ctx)
    
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

// CreateFace
// This API is used to add a set of face images to a person. One person can have up to 5 images. If a person exists in multiple groups, the images will be added to all those groups for the person.
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CreateFace(request *CreateFaceRequest) (response *CreateFaceResponse, err error) {
    return c.CreateFaceWithContext(context.Background(), request)
}

// CreateFace
// This API is used to add a set of face images to a person. One person can have up to 5 images. If a person exists in multiple groups, the images will be added to all those groups for the person.
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CreateFaceWithContext(ctx context.Context, request *CreateFaceRequest) (response *CreateFaceResponse, err error) {
    if request == nil {
        request = NewCreateFaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "CreateFace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFace require credential")
    }

    request.SetContext(ctx)
    
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

// CreateGroup
// This API is used to create an empty group. If the group already exists, an error will be returned.
//
// Custom description fields can be created as needed to describe persons in the group.
//
// 
//
// A maximum of 100,000 groups or 50 million faces can be created under one `APPID`.
//
// 
//
// The maximum number of faces that can be included in one group varies by algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    return c.CreateGroupWithContext(context.Background(), request)
}

// CreateGroup
// This API is used to create an empty group. If the group already exists, an error will be returned.
//
// Custom description fields can be created as needed to describe persons in the group.
//
// 
//
// A maximum of 100,000 groups or 50 million faces can be created under one `APPID`.
//
// 
//
// The maximum number of faces that can be included in one group varies by algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "CreateGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGroup require credential")
    }

    request.SetContext(ctx)
    
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

// CreatePerson
// This API is used to create a person and add face, name, gender, and other related information.
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UNIQUEPERSONCONTROLILLEGAL = "InvalidParameterValue.UniquePersonControlIllegal"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CreatePerson(request *CreatePersonRequest) (response *CreatePersonResponse, err error) {
    return c.CreatePersonWithContext(context.Background(), request)
}

// CreatePerson
// This API is used to create a person and add face, name, gender, and other related information.
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UNIQUEPERSONCONTROLILLEGAL = "InvalidParameterValue.UniquePersonControlIllegal"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CreatePersonWithContext(ctx context.Context, request *CreatePersonRequest) (response *CreatePersonResponse, err error) {
    if request == nil {
        request = NewCreatePersonRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "CreatePerson")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePerson require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteFace
// This API is used to delete the face images of a person. If the person has only one face image, an error will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeleteFace(request *DeleteFaceRequest) (response *DeleteFaceResponse, err error) {
    return c.DeleteFaceWithContext(context.Background(), request)
}

// DeleteFace
// This API is used to delete the face images of a person. If the person has only one face image, an error will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeleteFaceWithContext(ctx context.Context, request *DeleteFaceRequest) (response *DeleteFaceResponse, err error) {
    if request == nil {
        request = NewDeleteFaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "DeleteFace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFace require credential")
    }

    request.SetContext(ctx)
    
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

// DeleteGroup
// This API is used to delete a group and all persons in it. Meanwhile, all face information corresponding to the persons will be deleted. If a person exists in multiple groups at the same time, deleting a group will not delete the person, but the custom description field information in the group will be deleted. Custom description field information in other groups will not be affected.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

// DeleteGroup
// This API is used to delete a group and all persons in it. Meanwhile, all face information corresponding to the persons will be deleted. If a person exists in multiple groups at the same time, deleting a group will not delete the person, but the custom description field information in the group will be deleted. Custom description field information in other groups will not be affected.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "DeleteGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroup require credential")
    }

    request.SetContext(ctx)
    
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

// DeletePerson
// This API is used to delete a person from all groups. Meanwhile, all face information of the person will be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeletePerson(request *DeletePersonRequest) (response *DeletePersonResponse, err error) {
    return c.DeletePersonWithContext(context.Background(), request)
}

// DeletePerson
// This API is used to delete a person from all groups. Meanwhile, all face information of the person will be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeletePersonWithContext(ctx context.Context, request *DeletePersonRequest) (response *DeletePersonResponse, err error) {
    if request == nil {
        request = NewDeletePersonRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "DeletePerson")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePerson require credential")
    }

    request.SetContext(ctx)
    
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

// DeletePersonFromGroup
// This API is used to remove a person from a specified group. This operation only affects the group. If the person exists only in the group, the person will be deleted, and all face information of the person will also be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeletePersonFromGroup(request *DeletePersonFromGroupRequest) (response *DeletePersonFromGroupResponse, err error) {
    return c.DeletePersonFromGroupWithContext(context.Background(), request)
}

// DeletePersonFromGroup
// This API is used to remove a person from a specified group. This operation only affects the group. If the person exists only in the group, the person will be deleted, and all face information of the person will also be deleted.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_CREATEFACECONCURRENT = "FailedOperation.CreateFaceConcurrent"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeletePersonFromGroupWithContext(ctx context.Context, request *DeletePersonFromGroupRequest) (response *DeletePersonFromGroupResponse, err error) {
    if request == nil {
        request = NewDeletePersonFromGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "DeletePersonFromGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePersonFromGroup require credential")
    }

    request.SetContext(ctx)
    
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

// DetectFace
// This API is used to detect the position, attributes, and quality information of a face in the given image. The position information includes (x, y, w, h); the face attributes include gender, age, expression, beauty, glass, hair, mask, and pose (pitch, roll, yaw); and the face quality information includes the overall quality score, sharpness, brightness, and completeness.
//
// 
//
//  
//
// The face quality information is mainly used to evaluate the quality of the input face image. When using the Face Recognition service, we recommend evaluating the quality of the input face image first to improve the effects of subsequent processing. Application scenarios of this feature include:
//
// 
//
// 1. [Creating](https://intl.cloud.tencent.com/document/product/867/45014?from_cn_redirect=1)/[Adding](https://intl.cloud.tencent.com/document/product/867/45016?from_cn_redirect=1) a person in a group: this is to ensure the quality of the face information to facilitate subsequent processing.
//
// 
//
// 2. [Face search](https://intl.cloud.tencent.com/document/product/867/44994?from_cn_redirect=1): this is to ensure the quality of the input image to quickly find the corresponding person.
//
// 
//
// 3. [Face verification](https://intl.cloud.tencent.com/document/product/867/44983?from_cn_redirect=1): this is to ensure the quality of the face information to avoid cases where the verification incorrectly fails.
//
// 
//
// 4. Face fusion: this is to ensure the quality of the uploaded face images to improve the fusion effect.
//
// 
//
// >- Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectFace(request *DetectFaceRequest) (response *DetectFaceResponse, err error) {
    return c.DetectFaceWithContext(context.Background(), request)
}

// DetectFace
// This API is used to detect the position, attributes, and quality information of a face in the given image. The position information includes (x, y, w, h); the face attributes include gender, age, expression, beauty, glass, hair, mask, and pose (pitch, roll, yaw); and the face quality information includes the overall quality score, sharpness, brightness, and completeness.
//
// 
//
//  
//
// The face quality information is mainly used to evaluate the quality of the input face image. When using the Face Recognition service, we recommend evaluating the quality of the input face image first to improve the effects of subsequent processing. Application scenarios of this feature include:
//
// 
//
// 1. [Creating](https://intl.cloud.tencent.com/document/product/867/45014?from_cn_redirect=1)/[Adding](https://intl.cloud.tencent.com/document/product/867/45016?from_cn_redirect=1) a person in a group: this is to ensure the quality of the face information to facilitate subsequent processing.
//
// 
//
// 2. [Face search](https://intl.cloud.tencent.com/document/product/867/44994?from_cn_redirect=1): this is to ensure the quality of the input image to quickly find the corresponding person.
//
// 
//
// 3. [Face verification](https://intl.cloud.tencent.com/document/product/867/44983?from_cn_redirect=1): this is to ensure the quality of the face information to avoid cases where the verification incorrectly fails.
//
// 
//
// 4. Face fusion: this is to ensure the quality of the uploaded face images to improve the fusion effect.
//
// 
//
// >- Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectFaceWithContext(ctx context.Context, request *DetectFaceRequest) (response *DetectFaceResponse, err error) {
    if request == nil {
        request = NewDetectFaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "DetectFace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectFace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectFaceResponse()
    err = c.Send(request, response)
    return
}

func NewDetectFaceAttributesRequest() (request *DetectFaceAttributesRequest) {
    request = &DetectFaceAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iai", APIVersion, "DetectFaceAttributes")
    
    
    return
}

func NewDetectFaceAttributesResponse() (response *DetectFaceAttributesResponse) {
    response = &DetectFaceAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetectFaceAttributes
// This API is used to detect the position, attributes, and quality information of a face in the given image. The position information includes (x, y, w, h); the face attributes include gender, age, expression, beauty, glass, hair, mask, and pose (pitch, roll, yaw); and the face quality information includes the overall quality score, sharpness, brightness, and completeness.
//
// 
//
//  
//
// The face quality information is mainly used to evaluate the quality of the input face image. When using the Face Recognition service, we recommend evaluating the quality of the input face image first to improve the effects of subsequent processing. Application scenarios of this feature include:
//
// 
//
// 1. [Creating](https://intl.cloud.tencent.com/document/api/1059/36964)/[Adding](https://intl.cloud.tencent.com/document/api/1059/36966) a person in a group: This is to ensure the quality of the face information to facilitate subsequent processing.
//
// 
//
// 2. [Face search](https://intl.cloud.tencent.com/document/api/1059/36977): This is to ensure the quality of the input image to quickly find the corresponding person.
//
// 
//
// 3. [Face verification](https://intl.cloud.tencent.com/document/api/1059/36972): This is to ensure the quality of the face information to avoid cases where the verification fails unexpectedly.
//
// 
//
// 4. Face fusion: This is to ensure the quality of the uploaded face images to improve the fusion effect.
//
// 
//
// >     
//
// - This API is an upgrade of [DetectFace](https://intl.cloud.tencent.com/document/api/1059/36979); specifically:
//
// 1. This API can be used to specify the face attributes that need to be computed and returned, which avoids ineffective computation and reduces time consumption.
//
// 2. This API supports more detailed attribute items and will continue providing new features in the future.
//
// Use this API for corresponding face detection and attribute analysis.
//
// 
//
// >     
//
// - Use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the parameter `SignatureMethod` to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectFaceAttributes(request *DetectFaceAttributesRequest) (response *DetectFaceAttributesResponse, err error) {
    return c.DetectFaceAttributesWithContext(context.Background(), request)
}

// DetectFaceAttributes
// This API is used to detect the position, attributes, and quality information of a face in the given image. The position information includes (x, y, w, h); the face attributes include gender, age, expression, beauty, glass, hair, mask, and pose (pitch, roll, yaw); and the face quality information includes the overall quality score, sharpness, brightness, and completeness.
//
// 
//
//  
//
// The face quality information is mainly used to evaluate the quality of the input face image. When using the Face Recognition service, we recommend evaluating the quality of the input face image first to improve the effects of subsequent processing. Application scenarios of this feature include:
//
// 
//
// 1. [Creating](https://intl.cloud.tencent.com/document/api/1059/36964)/[Adding](https://intl.cloud.tencent.com/document/api/1059/36966) a person in a group: This is to ensure the quality of the face information to facilitate subsequent processing.
//
// 
//
// 2. [Face search](https://intl.cloud.tencent.com/document/api/1059/36977): This is to ensure the quality of the input image to quickly find the corresponding person.
//
// 
//
// 3. [Face verification](https://intl.cloud.tencent.com/document/api/1059/36972): This is to ensure the quality of the face information to avoid cases where the verification fails unexpectedly.
//
// 
//
// 4. Face fusion: This is to ensure the quality of the uploaded face images to improve the fusion effect.
//
// 
//
// >     
//
// - This API is an upgrade of [DetectFace](https://intl.cloud.tencent.com/document/api/1059/36979); specifically:
//
// 1. This API can be used to specify the face attributes that need to be computed and returned, which avoids ineffective computation and reduces time consumption.
//
// 2. This API supports more detailed attribute items and will continue providing new features in the future.
//
// Use this API for corresponding face detection and attribute analysis.
//
// 
//
// >     
//
// - Use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the parameter `SignatureMethod` to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectFaceAttributesWithContext(ctx context.Context, request *DetectFaceAttributesRequest) (response *DetectFaceAttributesResponse, err error) {
    if request == nil {
        request = NewDetectFaceAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "DetectFaceAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectFaceAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectFaceAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDetectFaceSimilarityRequest() (request *DetectFaceSimilarityRequest) {
    request = &DetectFaceSimilarityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iai", APIVersion, "DetectFaceSimilarity")
    
    
    return
}

func NewDetectFaceSimilarityResponse() (response *DetectFaceSimilarityResponse) {
    response = &DetectFaceSimilarityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetectFaceSimilarity
// Compare the faces in the two pictures for similarity and return the face similarity score. If you need to determine "whether this person is someone", that is, to verify whether the person in a picture is someone with a known identity, such as a common face login scenario, it is recommended to use [VerifyFace] (https://www.tencentcloud.com/document/product/1059/36972) or [VerifyPerson] (https://www.tencentcloud.com/document/product/1059/36971) inferface. 
//
// Please use the V3 version for the signature method in the public parameters, that is, configure the SignatureMethod parameter to TC3-HMAC-SHA256
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectFaceSimilarity(request *DetectFaceSimilarityRequest) (response *DetectFaceSimilarityResponse, err error) {
    return c.DetectFaceSimilarityWithContext(context.Background(), request)
}

// DetectFaceSimilarity
// Compare the faces in the two pictures for similarity and return the face similarity score. If you need to determine "whether this person is someone", that is, to verify whether the person in a picture is someone with a known identity, such as a common face login scenario, it is recommended to use [VerifyFace] (https://www.tencentcloud.com/document/product/1059/36972) or [VerifyPerson] (https://www.tencentcloud.com/document/product/1059/36971) inferface. 
//
// Please use the V3 version for the signature method in the public parameters, that is, configure the SignatureMethod parameter to TC3-HMAC-SHA256
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectFaceSimilarityWithContext(ctx context.Context, request *DetectFaceSimilarityRequest) (response *DetectFaceSimilarityResponse, err error) {
    if request == nil {
        request = NewDetectFaceSimilarityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "DetectFaceSimilarity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectFaceSimilarity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectFaceSimilarityResponse()
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

// DetectLiveFace
// This API is used to detect the liveness of a face in a static image uploaded by a user. Compared with dynamic liveness detection, static liveness detection does not require moving lips, shaking head, or blinking for recognition.
//
// 
//
// Image-based liveness detection is suitable for scenarios where the image is a selfie or the requirement for attack defense is not high. If you have a higher security requirement for liveness detection, please use [FaceID](https://intl.cloud.tencent.com/product/faceid?from_cn_redirect=1).
//
// 
//
// >     
//
// - The aspect ratio of the image should be close to 3:4 (width:height); otherwise, the score returned for the image will be meaningless. This API is suitable for selfie scenarios, and the score returned in other scenarios will be meaningless.
//
// 
//
// >
//
// - During the process, please directly face the camera and keep a suitable distance to completely display your face in the recognition frame. During the recognition, keep your device still and fully show your face. You are advised to perform the detection in an environment with appropriate light and without filters.
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the parameter `SignatureMethod` to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectLiveFace(request *DetectLiveFaceRequest) (response *DetectLiveFaceResponse, err error) {
    return c.DetectLiveFaceWithContext(context.Background(), request)
}

// DetectLiveFace
// This API is used to detect the liveness of a face in a static image uploaded by a user. Compared with dynamic liveness detection, static liveness detection does not require moving lips, shaking head, or blinking for recognition.
//
// 
//
// Image-based liveness detection is suitable for scenarios where the image is a selfie or the requirement for attack defense is not high. If you have a higher security requirement for liveness detection, please use [FaceID](https://intl.cloud.tencent.com/product/faceid?from_cn_redirect=1).
//
// 
//
// >     
//
// - The aspect ratio of the image should be close to 3:4 (width:height); otherwise, the score returned for the image will be meaningless. This API is suitable for selfie scenarios, and the score returned in other scenarios will be meaningless.
//
// 
//
// >
//
// - During the process, please directly face the camera and keep a suitable distance to completely display your face in the recognition frame. During the recognition, keep your device still and fully show your face. You are advised to perform the detection in an environment with appropriate light and without filters.
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the parameter `SignatureMethod` to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectLiveFaceWithContext(ctx context.Context, request *DetectLiveFaceRequest) (response *DetectLiveFaceResponse, err error) {
    if request == nil {
        request = NewDetectLiveFaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "DetectLiveFace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectLiveFace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectLiveFaceResponse()
    err = c.Send(request, response)
    return
}

func NewDetectLiveFaceAccurateRequest() (request *DetectLiveFaceAccurateRequest) {
    request = &DetectLiveFaceAccurateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("iai", APIVersion, "DetectLiveFaceAccurate")
    
    
    return
}

func NewDetectLiveFaceAccurateResponse() (response *DetectLiveFaceAccurateResponse) {
    response = &DetectLiveFaceAccurateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetectLiveFaceAccurate
// This API is used to detect the liveness of faces in images uploaded by users and determine whether these images are photographed.
//
// 
//
// Compared with normal Image-based Liveness Detection services, this API enhances the defense capability against attacks from HD screens, printed photos, and 3D masks, as well as improves attack blocking four to five times the competing products, while maintaining high accuracy. It also supports face verification in different use cases, and satisfies the image-based liveness detection needs on mobile or PCs, making it ideal for liveness detection applications in various industries.
//
// 
//
// Pay-as-you-go billing officially started for this API at 00:00, August 1, 2022. For more information, see [Billing Overview](https://intl.cloud.tencent.com/document/product/867/17640?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectLiveFaceAccurate(request *DetectLiveFaceAccurateRequest) (response *DetectLiveFaceAccurateResponse, err error) {
    return c.DetectLiveFaceAccurateWithContext(context.Background(), request)
}

// DetectLiveFaceAccurate
// This API is used to detect the liveness of faces in images uploaded by users and determine whether these images are photographed.
//
// 
//
// Compared with normal Image-based Liveness Detection services, this API enhances the defense capability against attacks from HD screens, printed photos, and 3D masks, as well as improves attack blocking four to five times the competing products, while maintaining high accuracy. It also supports face verification in different use cases, and satisfies the image-based liveness detection needs on mobile or PCs, making it ideal for liveness detection applications in various industries.
//
// 
//
// Pay-as-you-go billing officially started for this API at 00:00, August 1, 2022. For more information, see [Billing Overview](https://intl.cloud.tencent.com/document/product/867/17640?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMODELVERSIONILLEGAL = "InvalidParameterValue.FaceModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectLiveFaceAccurateWithContext(ctx context.Context, request *DetectLiveFaceAccurateRequest) (response *DetectLiveFaceAccurateResponse, err error) {
    if request == nil {
        request = NewDetectLiveFaceAccurateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "DetectLiveFaceAccurate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectLiveFaceAccurate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectLiveFaceAccurateResponse()
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

// GetGroupInfo
// This API is used to get the group information.
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetGroupInfo(request *GetGroupInfoRequest) (response *GetGroupInfoResponse, err error) {
    return c.GetGroupInfoWithContext(context.Background(), request)
}

// GetGroupInfo
// This API is used to get the group information.
//
// error code that may be returned:
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetGroupInfoWithContext(ctx context.Context, request *GetGroupInfoRequest) (response *GetGroupInfoResponse, err error) {
    if request == nil {
        request = NewGetGroupInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "GetGroupInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetGroupInfo require credential")
    }

    request.SetContext(ctx)
    
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

// GetGroupList
// This API is used to get the list of groups.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetGroupList(request *GetGroupListRequest) (response *GetGroupListResponse, err error) {
    return c.GetGroupListWithContext(context.Background(), request)
}

// GetGroupList
// This API is used to get the list of groups.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetGroupListWithContext(ctx context.Context, request *GetGroupListRequest) (response *GetGroupListResponse, err error) {
    if request == nil {
        request = NewGetGroupListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "GetGroupList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetGroupList require credential")
    }

    request.SetContext(ctx)
    
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

// GetPersonBaseInfo
// This API is used to get the information of a specified person, including name, gender, face, etc.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonBaseInfo(request *GetPersonBaseInfoRequest) (response *GetPersonBaseInfoResponse, err error) {
    return c.GetPersonBaseInfoWithContext(context.Background(), request)
}

// GetPersonBaseInfo
// This API is used to get the information of a specified person, including name, gender, face, etc.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonBaseInfoWithContext(ctx context.Context, request *GetPersonBaseInfoRequest) (response *GetPersonBaseInfoResponse, err error) {
    if request == nil {
        request = NewGetPersonBaseInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "GetPersonBaseInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPersonBaseInfo require credential")
    }

    request.SetContext(ctx)
    
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

// GetPersonGroupInfo
// This API is used to get the information of a specified person, including group, description, etc.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonGroupInfo(request *GetPersonGroupInfoRequest) (response *GetPersonGroupInfoResponse, err error) {
    return c.GetPersonGroupInfoWithContext(context.Background(), request)
}

// GetPersonGroupInfo
// This API is used to get the information of a specified person, including group, description, etc.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonGroupInfoWithContext(ctx context.Context, request *GetPersonGroupInfoRequest) (response *GetPersonGroupInfoResponse, err error) {
    if request == nil {
        request = NewGetPersonGroupInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "GetPersonGroupInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPersonGroupInfo require credential")
    }

    request.SetContext(ctx)
    
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

// GetPersonList
// This API is used to get the list of persons in a specified group.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonList(request *GetPersonListRequest) (response *GetPersonListResponse, err error) {
    return c.GetPersonListWithContext(context.Background(), request)
}

// GetPersonList
// This API is used to get the list of persons in a specified group.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonListWithContext(ctx context.Context, request *GetPersonListRequest) (response *GetPersonListResponse, err error) {
    if request == nil {
        request = NewGetPersonListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "GetPersonList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPersonList require credential")
    }

    request.SetContext(ctx)
    
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

// GetPersonListNum
// This API is used to get the number of persons in a specified group.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonListNum(request *GetPersonListNumRequest) (response *GetPersonListNumResponse, err error) {
    return c.GetPersonListNumWithContext(context.Background(), request)
}

// GetPersonListNum
// This API is used to get the number of persons in a specified group.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonListNumWithContext(ctx context.Context, request *GetPersonListNumRequest) (response *GetPersonListNumResponse, err error) {
    if request == nil {
        request = NewGetPersonListNumRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "GetPersonListNum")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPersonListNum require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPersonListNumResponse()
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

// ModifyGroup
// This API is used to modify the name, tag, and custom description field of a group.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) ModifyGroup(request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    return c.ModifyGroupWithContext(context.Background(), request)
}

// ModifyGroup
// This API is used to modify the name, tag, and custom description field of a group.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) ModifyGroupWithContext(ctx context.Context, request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    if request == nil {
        request = NewModifyGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "ModifyGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGroupResponse()
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

// ModifyPersonGroupInfo
// This API is used to modify the description of a specified person in a group.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) ModifyPersonGroupInfo(request *ModifyPersonGroupInfoRequest) (response *ModifyPersonGroupInfoResponse, err error) {
    return c.ModifyPersonGroupInfoWithContext(context.Background(), request)
}

// ModifyPersonGroupInfo
// This API is used to modify the description of a specified person in a group.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) ModifyPersonGroupInfoWithContext(ctx context.Context, request *ModifyPersonGroupInfoRequest) (response *ModifyPersonGroupInfoResponse, err error) {
    if request == nil {
        request = NewModifyPersonGroupInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "ModifyPersonGroupInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPersonGroupInfo require credential")
    }

    request.SetContext(ctx)
    
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

// SearchFaces
// This API is used to recognize top K persons in one or more groups who are similar to the person in a given image and rank the similarity in descending order.
//
// 
//
// Up to 10 faces in an image can be recognized at a time, and up to 100 groups can be searched in at a time.
//
// 
//
// The maximum number of faces in groups that can be searched for at a time is subject to the algorithm model version (`FaceModelVersion`) of groups, which is 1 million for v2.0 or 3 million for v3.0.
//
// 
//
// This API recognizes each face image of a person as an independent one. By contrast, the [SearchPersons](https://intl.cloud.tencent.com/document/product/867/44992?from_cn_redirect=1) and [SearchPersonsReturnsByGroup](https://intl.cloud.tencent.com/document/product/867/44991?from_cn_redirect=1) APIs fuse the features of all face images of a person; for example, if a person has 4 face images, they will fuse the features of the 4 face images and generate the summarized facial features of the person to make the search more accurate.
//
// 
//
// 
//
// This API should be used together with [Group Management APIs](https://intl.cloud.tencent.com/document/product/867/45015?from_cn_redirect=1).
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the parameter `SignatureMethod` to `TC3-HMAC-SHA256`.
//
// 
//
// >     
//
// - You cannot search for groups using different algorithm model versions (`FaceModelVersion`) at a time.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchFaces(request *SearchFacesRequest) (response *SearchFacesResponse, err error) {
    return c.SearchFacesWithContext(context.Background(), request)
}

// SearchFaces
// This API is used to recognize top K persons in one or more groups who are similar to the person in a given image and rank the similarity in descending order.
//
// 
//
// Up to 10 faces in an image can be recognized at a time, and up to 100 groups can be searched in at a time.
//
// 
//
// The maximum number of faces in groups that can be searched for at a time is subject to the algorithm model version (`FaceModelVersion`) of groups, which is 1 million for v2.0 or 3 million for v3.0.
//
// 
//
// This API recognizes each face image of a person as an independent one. By contrast, the [SearchPersons](https://intl.cloud.tencent.com/document/product/867/44992?from_cn_redirect=1) and [SearchPersonsReturnsByGroup](https://intl.cloud.tencent.com/document/product/867/44991?from_cn_redirect=1) APIs fuse the features of all face images of a person; for example, if a person has 4 face images, they will fuse the features of the 4 face images and generate the summarized facial features of the person to make the search more accurate.
//
// 
//
// 
//
// This API should be used together with [Group Management APIs](https://intl.cloud.tencent.com/document/product/867/45015?from_cn_redirect=1).
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the parameter `SignatureMethod` to `TC3-HMAC-SHA256`.
//
// 
//
// >     
//
// - You cannot search for groups using different algorithm model versions (`FaceModelVersion`) at a time.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchFacesWithContext(ctx context.Context, request *SearchFacesRequest) (response *SearchFacesResponse, err error) {
    if request == nil {
        request = NewSearchFacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "SearchFaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchFaces require credential")
    }

    request.SetContext(ctx)
    
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

// SearchFacesReturnsByGroup
// This API is used to recognize top K persons in one or more groups who are similar to the person in a given image, display the results **by group**, and rank the similarity within each group in descending order.
//
// 
//
// Up to 10 faces in the image can be recognized at a time, and cross-group search is supported.
//
// 
//
// The maximum number of faces in groups that can be searched for at a time is subject to the group's algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
//
// 
//
// This API recognizes each face image of a person as an independent one. By contrast, the [SearchPersons](https://intl.cloud.tencent.com/document/product/867/44992?from_cn_redirect=1) and [SearchPersonsReturnsByGroup](https://intl.cloud.tencent.com/document/product/867/44991?from_cn_redirect=1) APIs fuse the features of all face images of a person; for example, if a person has 4 face images, they will fuse the features of the 4 face images and generate the summarized facial features of the person to make the search more accurate.
//
// 
//
// This API should be used together with [Group Management APIs](https://intl.cloud.tencent.com/document/product/867/45015?from_cn_redirect=1).
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// 
//
// >     
//
// - You cannot search for groups using different algorithm model versions (`FaceModelVersion`) at a time.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchFacesReturnsByGroup(request *SearchFacesReturnsByGroupRequest) (response *SearchFacesReturnsByGroupResponse, err error) {
    return c.SearchFacesReturnsByGroupWithContext(context.Background(), request)
}

// SearchFacesReturnsByGroup
// This API is used to recognize top K persons in one or more groups who are similar to the person in a given image, display the results **by group**, and rank the similarity within each group in descending order.
//
// 
//
// Up to 10 faces in the image can be recognized at a time, and cross-group search is supported.
//
// 
//
// The maximum number of faces in groups that can be searched for at a time is subject to the group's algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
//
// 
//
// This API recognizes each face image of a person as an independent one. By contrast, the [SearchPersons](https://intl.cloud.tencent.com/document/product/867/44992?from_cn_redirect=1) and [SearchPersonsReturnsByGroup](https://intl.cloud.tencent.com/document/product/867/44991?from_cn_redirect=1) APIs fuse the features of all face images of a person; for example, if a person has 4 face images, they will fuse the features of the 4 face images and generate the summarized facial features of the person to make the search more accurate.
//
// 
//
// This API should be used together with [Group Management APIs](https://intl.cloud.tencent.com/document/product/867/45015?from_cn_redirect=1).
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// 
//
// >     
//
// - You cannot search for groups using different algorithm model versions (`FaceModelVersion`) at a time.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchFacesReturnsByGroupWithContext(ctx context.Context, request *SearchFacesReturnsByGroupRequest) (response *SearchFacesReturnsByGroupResponse, err error) {
    if request == nil {
        request = NewSearchFacesReturnsByGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "SearchFacesReturnsByGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchFacesReturnsByGroup require credential")
    }

    request.SetContext(ctx)
    
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

// SearchPersons
// This API is used to recognize top K persons in one or more groups who are similar to the person in a given image and rank the similarity in a descending order.
//
// 
//
// Up to 10 faces in an image can be recognized at a time, and up to 100 groups can be searched in at a time.
//
// 
//
// The maximum number of faces in groups that can be searched for at a time is subject to the group's algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
//
// 
//
// This API fuses the features of all face images of a person; for example, if a person has 4 face images, it will fuse the features of the 4 face images and generate the summarized facial features of the person to make the person search (i.e., judging whether the face image to be recognized is of a specified person) more accurate. By contrast, the [SearchFaces](https://intl.cloud.tencent.com/document/product/867/44994?from_cn_redirect=1) and [SearchFacesReturnsByGroup](https://intl.cloud.tencent.com/document/product/867/44993?from_cn_redirect=1) APIs recognize each face image of a person as an independent one for search.
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// - This feature is available only to groups whose algorithm model version (`FaceModelVersion`) is 3.0.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_IMAGEEMPTYERROR = "InvalidParameterValue.ImageEmptyError"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGROUPFACEMODELVERSION = "InvalidParameterValue.UnsupportedGroupFaceModelVersion"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_CHARGESTATUSEXCEPTION = "ResourceUnavailable.ChargeStatusException"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchPersons(request *SearchPersonsRequest) (response *SearchPersonsResponse, err error) {
    return c.SearchPersonsWithContext(context.Background(), request)
}

// SearchPersons
// This API is used to recognize top K persons in one or more groups who are similar to the person in a given image and rank the similarity in a descending order.
//
// 
//
// Up to 10 faces in an image can be recognized at a time, and up to 100 groups can be searched in at a time.
//
// 
//
// The maximum number of faces in groups that can be searched for at a time is subject to the group's algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
//
// 
//
// This API fuses the features of all face images of a person; for example, if a person has 4 face images, it will fuse the features of the 4 face images and generate the summarized facial features of the person to make the person search (i.e., judging whether the face image to be recognized is of a specified person) more accurate. By contrast, the [SearchFaces](https://intl.cloud.tencent.com/document/product/867/44994?from_cn_redirect=1) and [SearchFacesReturnsByGroup](https://intl.cloud.tencent.com/document/product/867/44993?from_cn_redirect=1) APIs recognize each face image of a person as an independent one for search.
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// - This feature is available only to groups whose algorithm model version (`FaceModelVersion`) is 3.0.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_IMAGEEMPTYERROR = "InvalidParameterValue.ImageEmptyError"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGROUPFACEMODELVERSION = "InvalidParameterValue.UnsupportedGroupFaceModelVersion"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_CHARGESTATUSEXCEPTION = "ResourceUnavailable.ChargeStatusException"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchPersonsWithContext(ctx context.Context, request *SearchPersonsRequest) (response *SearchPersonsResponse, err error) {
    if request == nil {
        request = NewSearchPersonsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "SearchPersons")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchPersons require credential")
    }

    request.SetContext(ctx)
    
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

// SearchPersonsReturnsByGroup
// This API is used to recognize top K persons in one or more groups who are similar to the person in a given image, display the results **by group**, and rank the similarity within each group in a descending order.
//
// 
//
// Up to 10 faces in the image can be recognized at a time, and cross-group search is supported.
//
// 
//
// The maximum number of faces in groups that can be searched for at a time is subject to the group's algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
//
// 
//
// This API fuses the features of all face images of a person; for example, if a person has 4 face images, it will fuse the features of the 4 face images and generate the summarized facial features of the person to make the person search (i.e., judging whether the face image to be recognized is of a specified person) more accurate. By contrast, the [SearchFaces](https://intl.cloud.tencent.com/document/product/867/44994?from_cn_redirect=1) and [SearchFacesReturnsByGroup](https://intl.cloud.tencent.com/document/product/867/44993?from_cn_redirect=1) APIs recognize each face image of a person as an independent one for search.
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// - This feature is available only to groups whose algorithm model version (`FaceModelVersion`) is 3.0.
//
// error code that may be returned:
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXISTS = "InvalidParameterValue.GroupIdNotExists"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGROUPFACEMODELVERSION = "InvalidParameterValue.UnsupportedGroupFaceModelVersion"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_CHARGESTATUSEXCEPTION = "ResourceUnavailable.ChargeStatusException"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchPersonsReturnsByGroup(request *SearchPersonsReturnsByGroupRequest) (response *SearchPersonsReturnsByGroupResponse, err error) {
    return c.SearchPersonsReturnsByGroupWithContext(context.Background(), request)
}

// SearchPersonsReturnsByGroup
// This API is used to recognize top K persons in one or more groups who are similar to the person in a given image, display the results **by group**, and rank the similarity within each group in a descending order.
//
// 
//
// Up to 10 faces in the image can be recognized at a time, and cross-group search is supported.
//
// 
//
// The maximum number of faces in groups that can be searched for at a time is subject to the group's algorithm model version (`FaceModelVersion`), which is 1 million for v2.0 or 3 million for v3.0.
//
// 
//
// This API fuses the features of all face images of a person; for example, if a person has 4 face images, it will fuse the features of the 4 face images and generate the summarized facial features of the person to make the person search (i.e., judging whether the face image to be recognized is of a specified person) more accurate. By contrast, the [SearchFaces](https://intl.cloud.tencent.com/document/product/867/44994?from_cn_redirect=1) and [SearchFacesReturnsByGroup](https://intl.cloud.tencent.com/document/product/867/44993?from_cn_redirect=1) APIs recognize each face image of a person as an independent one for search.
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// - This feature is available only to groups whose algorithm model version (`FaceModelVersion`) is 3.0.
//
// error code that may be returned:
//  FAILEDOPERATION_ACROSSVERSIONSERROR = "FailedOperation.AcrossVersionsError"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_FACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.FaceMatchThresholdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXISTS = "InvalidParameterValue.GroupIdNotExists"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGROUPFACEMODELVERSION = "InvalidParameterValue.UnsupportedGroupFaceModelVersion"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_CHARGESTATUSEXCEPTION = "ResourceUnavailable.ChargeStatusException"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchPersonsReturnsByGroupWithContext(ctx context.Context, request *SearchPersonsReturnsByGroupRequest) (response *SearchPersonsReturnsByGroupResponse, err error) {
    if request == nil {
        request = NewSearchPersonsReturnsByGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "SearchPersonsReturnsByGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchPersonsReturnsByGroup require credential")
    }

    request.SetContext(ctx)
    
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

// VerifyFace
// This API is used to judge whether a person in an image corresponds to a given `PersonId`. For more information on `PersonId`, please see [Group Management APIs](https://intl.cloud.tencent.com/document/product/867/45015?from_cn_redirect=1). 
//
// 
//
// The `VerifyFace` API judges whether a person is someone specified whose information is stored in a group, and there may be multiple face images of "someone". By contrast, the [CompareFace](https://intl.cloud.tencent.com/document/product/867/44987?from_cn_redirect=1) API judges the similarity between two faces.
//
// 
//
// This API recognizes each face image of a person as an independent one. By contrast, the [VerifyPerson](https://intl.cloud.tencent.com/document/product/867/44982?from_cn_redirect=1) API fuses the features of all face images of a person; for example, if a person has 4 face images, the VerifyPerson API will fuse the features of the 4 face images and generate the summarized facial features of the person to make the person verification (i.e., judging whether the face image to be recognized is of a specified person) more accurate.
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the parameter `SignatureMethod` to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) VerifyFace(request *VerifyFaceRequest) (response *VerifyFaceResponse, err error) {
    return c.VerifyFaceWithContext(context.Background(), request)
}

// VerifyFace
// This API is used to judge whether a person in an image corresponds to a given `PersonId`. For more information on `PersonId`, please see [Group Management APIs](https://intl.cloud.tencent.com/document/product/867/45015?from_cn_redirect=1). 
//
// 
//
// The `VerifyFace` API judges whether a person is someone specified whose information is stored in a group, and there may be multiple face images of "someone". By contrast, the [CompareFace](https://intl.cloud.tencent.com/document/product/867/44987?from_cn_redirect=1) API judges the similarity between two faces.
//
// 
//
// This API recognizes each face image of a person as an independent one. By contrast, the [VerifyPerson](https://intl.cloud.tencent.com/document/product/867/44982?from_cn_redirect=1) API fuses the features of all face images of a person; for example, if a person has 4 face images, the VerifyPerson API will fuse the features of the 4 face images and generate the summarized facial features of the person to make the person verification (i.e., judging whether the face image to be recognized is of a specified person) more accurate.
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the parameter `SignatureMethod` to `TC3-HMAC-SHA256`.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CONFLICTOPERATION = "FailedOperation.ConflictOperation"
//  FAILEDOPERATION_DUPLICATEDGROUPDESCRIPTION = "FailedOperation.DuplicatedGroupDescription"
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_GROUPINDELETEDSTATE = "FailedOperation.GroupInDeletedState"
//  FAILEDOPERATION_GROUPPERSONMAPEXIST = "FailedOperation.GroupPersonMapExist"
//  FAILEDOPERATION_GROUPPERSONMAPNOTEXIST = "FailedOperation.GroupPersonMapNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEARCHFACESEXCEED = "FailedOperation.SearchFacesExceed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTFACENUMEXCEED = "InvalidParameterValue.AccountFaceNumExceed"
//  INVALIDPARAMETERVALUE_DELETEFACENUMEXCEED = "InvalidParameterValue.DeleteFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSEXCEED = "InvalidParameterValue.GroupExDescriptionsExceed"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.GroupExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.GroupExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.GroupExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPFACENUMEXCEED = "InvalidParameterValue.GroupFaceNumExceed"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPIDSEXCEED = "InvalidParameterValue.GroupIdsExceed"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMEILLEGAL = "InvalidParameterValue.GroupNameIllegal"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPNUMPERPERSONEXCEED = "InvalidParameterValue.GroupNumPerPersonExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGILLEGAL = "InvalidParameterValue.GroupTagIllegal"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONINFOSEXCEED = "InvalidParameterValue.PersonExDescriptionInfosExceed"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEIDENTICAL = "InvalidParameterValue.PersonExDescriptionsNameIdentical"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMEILLEGAL = "InvalidParameterValue.PersonExDescriptionsNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONEXDESCRIPTIONSNAMETOOLONG = "InvalidParameterValue.PersonExDescriptionsNameTooLong"
//  INVALIDPARAMETERVALUE_PERSONEXISTINGROUP = "InvalidParameterValue.PersonExistInGroup"
//  INVALIDPARAMETERVALUE_PERSONFACENUMEXCEED = "InvalidParameterValue.PersonFaceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONGENDERILLEGAL = "InvalidParameterValue.PersonGenderIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_UPLOADFACENUMEXCEED = "InvalidParameterValue.UploadFaceNumExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_ERRORFACENUMEXCEED = "LimitExceeded.ErrorFaceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) VerifyFaceWithContext(ctx context.Context, request *VerifyFaceRequest) (response *VerifyFaceResponse, err error) {
    if request == nil {
        request = NewVerifyFaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "VerifyFace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyFace require credential")
    }

    request.SetContext(ctx)
    
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

// VerifyPerson
// This API is used to judge whether a person in an image corresponds to a given `PersonId`. For more information on `PersonId`, please see [Group Management APIs](https://intl.cloud.tencent.com/document/product/867/45015?from_cn_redirect=1).
//
// This API fuses the features of all face images of a person; for example, if a person has 4 face images, it will fuse the features of the 4 face images and generate the summarized facial features of the person to make the person verification (i.e., judging whether the face image to be recognized is of a specified person) more accurate.
//
// 
//
//  The face verification APIs judge whether a person is someone specified whose information is stored in a group, and the "someone" may have multiple face images. By contrast, the face comparison APIs judge the similarity between two faces.
//
// 
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// - This feature is available only to groups whose algorithm model version (`FaceModelVersion`) is 3.0.
//
// error code that may be returned:
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGROUPFACEMODELVERSION = "InvalidParameterValue.UnsupportedGroupFaceModelVersion"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE_CHARGESTATUSEXCEPTION = "ResourceUnavailable.ChargeStatusException"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) VerifyPerson(request *VerifyPersonRequest) (response *VerifyPersonResponse, err error) {
    return c.VerifyPersonWithContext(context.Background(), request)
}

// VerifyPerson
// This API is used to judge whether a person in an image corresponds to a given `PersonId`. For more information on `PersonId`, please see [Group Management APIs](https://intl.cloud.tencent.com/document/product/867/45015?from_cn_redirect=1).
//
// This API fuses the features of all face images of a person; for example, if a person has 4 face images, it will fuse the features of the 4 face images and generate the summarized facial features of the person to make the person verification (i.e., judging whether the face image to be recognized is of a specified person) more accurate.
//
// 
//
//  The face verification APIs judge whether a person is someone specified whose information is stored in a group, and the "someone" may have multiple face images. By contrast, the face comparison APIs judge the similarity between two faces.
//
// 
//
// 
//
// >     
//
// - Please use the signature algorithm v3 to calculate the signature in the common parameters, that is, set the `SignatureMethod` parameter to `TC3-HMAC-SHA256`.
//
// - This feature is available only to groups whose algorithm model version (`FaceModelVersion`) is 3.0.
//
// error code that may be returned:
//  FAILEDOPERATION_FACEQUALITYNOTQUALIFIED = "FailedOperation.FaceQualityNotQualified"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_NOFACEINGROUPS = "InvalidParameterValue.NoFaceInGroups"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_QUALITYCONTROLILLEGAL = "InvalidParameterValue.QualityControlIllegal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGROUPFACEMODELVERSION = "InvalidParameterValue.UnsupportedGroupFaceModelVersion"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE_CHARGESTATUSEXCEPTION = "ResourceUnavailable.ChargeStatusException"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) VerifyPersonWithContext(ctx context.Context, request *VerifyPersonRequest) (response *VerifyPersonResponse, err error) {
    if request == nil {
        request = NewVerifyPersonRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "iai", APIVersion, "VerifyPerson")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyPerson require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyPersonResponse()
    err = c.Send(request, response)
    return
}
