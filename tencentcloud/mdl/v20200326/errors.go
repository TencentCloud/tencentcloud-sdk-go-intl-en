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

package v20200326

const (
	// error codes for specific actions

	// Internal error.
	INTERNALERROR = "InternalError"

	// Invalid parameter.
	INVALIDPARAMETER = "InvalidParameter"

	// `Channel` already associated.
	INVALIDPARAMETER_ALREADYASSOCIATEDCHANNEL = "InvalidParameter.AlreadyAssociatedChannel"

	// `Input` already associated.
	INVALIDPARAMETER_ALREADYASSOCIATEDINPUT = "InvalidParameter.AlreadyAssociatedInput"

	// Invalid `AttachedInputs`.
	INVALIDPARAMETER_ATTACHEDINPUTS = "InvalidParameter.AttachedInputs"

	// Incorrect audio transcoding template.
	INVALIDPARAMETER_AUDIOTEMPLATES = "InvalidParameter.AudioTemplates"

	// Invalid `EndTime`.
	INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"

	// The quantity exceeds the limit.
	INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"

	// Invalid `Id`.
	INVALIDPARAMETER_ID = "InvalidParameter.Id"

	// Invalid `InputSettings`.
	INVALIDPARAMETER_INPUTSETTINGS = "InvalidParameter.InputSettings"

	// Invalid `Name`.
	INVALIDPARAMETER_NAME = "InvalidParameter.Name"

	// Not found.
	INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"

	// Invalid `OutputGroups`.
	INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"

	// Invalid `SecurityGroups`.
	INVALIDPARAMETER_SECURITYGROUPS = "InvalidParameter.SecurityGroups"

	// Invalid `StartTime`.
	INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"

	// Exceptional status.
	INVALIDPARAMETER_STATE = "InvalidParameter.State"

	// Incorrect status.
	INVALIDPARAMETER_STATEERROR = "InvalidParameter.StateError"

	// Invalid `Type`.
	INVALIDPARAMETER_TYPE = "InvalidParameter.Type"

	// Incorrect video transcoding template.
	INVALIDPARAMETER_VIDEOTEMPLATES = "InvalidParameter.VideoTemplates"

	// Invalid `Whitelist`.
	INVALIDPARAMETER_WHITELIST = "InvalidParameter.Whitelist"
)
