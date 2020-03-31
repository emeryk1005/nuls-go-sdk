/*
 *  MIT License
 *  Copyright (c) 2019-2020 niels.wang
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 *
 */

// @Title
// @Description
// @Author  Niels  2020/3/31
package ps

import (
	"errors"
	"github.com/niels1286/nuls-go-sdk/client/jsonrpc"
	"math/rand"
	"time"
)

type AgentInfo struct {
	txHash            string  `json:"txHash"`            //string	创建共识节点的交易hash
	agentId           string  `json:"agentId"`           //string	节点id
	agentAddress      string  `json:"agentAddress"`      //string	创建节点的代理账户地址
	packingAddress    string  `json:"packingAddress"`    //string	节点负责打包区块的账户地址
	rewardAddress     string  `json:"rewardAddress"`     //string	节点获取共识奖励的账户地址
	agentAlias        string  `json:"agentAlias"`        //string	节点的代理地址别名
	deposit           string  `json:"deposit"`           //bigInt	创建节点时代理节点的保证金
	commissionRat     int     `json:"commissionRat"`     //int		节点收取的佣金比例，单位%
	createTime        uint64  `json:"createTime"`        //long		节点的创建时间
	status            int     `json:"status"`            //int		节点状态::待共识, 1:共识中, 2:已注销
	totalDeposit      string  `json:"totalDeposit"`      //bigInt	委托参与共识总金额
	depositCount      int     `json:"depositCount"`      //int		委托次数
	creditValue       float32 `json:"creditValue"`       //double	信用值 取值[-1,1]
	totalPackingCount int     `json:"totalPackingCount"` //int		节点已打包出块总数
	lostRate          float32 `json:"lostRate"`          //double	丢块率
	lastRewardHeight  uint64  `json:"lastRewardHeight"`  //long		最后一次出块获取奖励的区块高度
	deleteHash        string  `json:"deleteHash"`        //string	注销节点的交易hash
	blockHeight       uint64  `json:"blockHeight"`       //long		创建节点时的区块高度
	deleteHeight      uint64  `json:"deleteHeight"`      //long		注销节点时的区块高度
	totalReward       string  `json:"totalReward"`       //bigInt	总共识奖励 totalReward=commissionReward+agentReward
	commissionReward  string  `json:"commissionReward"`  //bigInt	委托共识奖励
	agentReward       string  `json:"agentReward"`       //bigInt	节点获取奖励
	roundPackingTime  uint64  `json:"roundPackingTime"`  //long		当前轮次节点负责打包区块的时间
	version           int     `json:"version"`           //int		节点的协议版本号
	agentType         int     `json:"agentType"`         //int		1:普通节点,2:开
}

//将根据账户地址，查询对应的节点详情
func GetAgentByAddress(client *jsonrpc.NulsPSClient, chainId uint16, address string) (*AgentInfo, error) {
	if client == nil || address == "" {
		return nil, errors.New("parameter wrong.")
	}
	rand.Seed(time.Now().Unix())
	param := jsonrpc.NewRequestParam(rand.Intn(10000), "getAccountConsensusNode", []interface{}{chainId, address})
	result, err := client.PSRequest(param)
	if err != nil {
		return nil, err
	}
	resultMap := result.Result.(*AgentInfo)
	return resultMap, nil
}
