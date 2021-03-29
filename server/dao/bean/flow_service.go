/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: flow_service
 * @Date: 2021/3/12 10:24
 */

package bean

import "time"

type FlowService struct {
	FId                     uint64
	FServiceId              string
	FMachineQrn             string
	FRoleQrn                string
	FTemplateId             string
	FServiceName            string
	FServiceChineseName     string
	FMachineType            string
	FStatus                 string
	FCreateTime             time.Time
	FModifyTime             time.Time
	FCreator                string
	FDescription            string
}
