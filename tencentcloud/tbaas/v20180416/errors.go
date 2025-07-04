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

package v20180416

const (
	// error codes for specific actions

	// The block does not exist.
	FAILEDOPERATION_FABRICBLOCKNOEXIST = "FailedOperation.FabricBlockNoExist"

	// Contract call failed.
	FAILEDOPERATION_FABRICCHAINCODEINVOKEFAILED = "FailedOperation.FabricChaincodeInvokeFailed"

	// The contract does not exist.
	FAILEDOPERATION_FABRICCHAINCODENOEXIST = "FailedOperation.FabricChaincodeNoExist"

	// Contract query failed.
	FAILEDOPERATION_FABRICCHAINCODEQUERYFAILED = "FailedOperation.FabricChaincodeQueryFailed"

	// The transaction does not exist.
	FAILEDOPERATION_FABRICTRANSACTIONNOEXIST = "FailedOperation.FabricTransactionNoExist"

	// User has no access permission.
	FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"

	// Internal error. please try again later or contact technical personnel.
	INTERNALERROR_SERVERERROR = "InternalError.ServerError"
)
