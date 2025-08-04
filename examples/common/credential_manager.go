package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/cvm/v20170312"
)

func main() {
	// Obtain credentials through the default provider chain
	pc := common.DefaultProviderChain()
	cre, err := pc.GetCredential()
	if err != nil {
		fmt.Println("get cre failed!", err)
		return
	}

	client, _ := cvm.NewClient(cre, regions.Guangzhou, profile.NewClientProfile())
	request := cvm.NewDescribeInstancesRequest()
	response, err := client.DescribeInstances(request)
	if err != nil {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	fmt.Printf("%s\n", response.ToJsonString())
}
