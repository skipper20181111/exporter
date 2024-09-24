package util

import (
	"bufio"
	"context"
	"exporter/internal/logic/email"
	"exporter/internal/svc"
	"exporter/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"strings"
	"syscall"
)

type MountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	eul    *email.EmailUtilLogic
}

func NewMountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MountLogic {
	return &MountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		eul:    email.NewEmailUtilLogic(ctx, svcCtx),
	}
}
func (l *MountLogic) MountReport() {
	code, message := CheckMountPoints()
	if code == 1 {
		return
	}
	EmailInfoList := make([]*types.EmailInfo, 0)
	emaillist, ok := l.svcCtx.LocalCache.Get(svc.EmailListKey)
	if ok {
		EmailInfoList = emaillist.([]*types.EmailInfo)
	}
	for _, emailInfo := range EmailInfoList {
		err := l.eul.SendMailRandom(emailInfo, emailInfo.Send2Who, "CDP集群告警Subject", message)
		if err != nil {
			fmt.Println("################发送邮件失败")
		}
	}
}

func CheckMountPoints() (code int, report_message string) {
	code = 1
	// 获取主机名
	hostname := "CDP告警"
	hostname, _ = os.Hostname()
	// 打开 /proc/mounts 文件
	file, err := os.Open("/proc/mounts")
	if err != nil {
		fmt.Println("Error reading /proc/mounts:", err)
		code = 0
	}
	defer file.Close()

	// 逐行读取挂载点信息
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 3 {
			continue
		}

		mountPoint := fields[1] // 第二个字段是挂载点路径
		fsType := fields[2]     // 第三个字段是文件系统类型

		// 排除非硬盘挂载类文件系统
		if fsType != "xfs" {
			continue
		}
		testFile := fmt.Sprintf("%s/.testfile", mountPoint)

		// 尝试在挂载点上创建一个临时文件
		f, err := os.Create(testFile)
		if err != nil {
			code = 0
			if pathError, ok := err.(*os.PathError); ok && pathError.Err == syscall.EIO {
				report_message += "\n"
				report_message += fmt.Sprintf("I/O error detected on mount point: %s\n", mountPoint)
				fmt.Printf("I/O error detected on mount point: %s\n", mountPoint)
			} else {
				report_message += "\n"
				report_message += fmt.Sprintf("Error creating file on mount point %s: %v\n", mountPoint, err)
				fmt.Printf("Error creating file on mount point %s: %v\n", mountPoint, err)
			}
			continue
		}

		// 尝试写入文件
		_, err = f.WriteString("test")
		if err != nil {
			code = 0
			if pathError, ok := err.(*os.PathError); ok && pathError.Err == syscall.EIO {
				report_message += "\n"
				report_message += fmt.Sprintf("I/O error detected on mount point: %s\n", mountPoint)
				fmt.Printf("I/O error detected on mount point: %s\n", mountPoint)
			} else {
				report_message += "\n"
				report_message += fmt.Sprintf("Error writing file on mount point %s: %v\n", mountPoint, err)
				fmt.Printf("Error writing file on mount point %s: %v\n", mountPoint, err)
			}
			f.Close()
			os.Remove(testFile)
			continue
		}

		// 关闭并删除临时文件
		f.Close()
		os.Remove(testFile)

		report_message += "\n"
		report_message += fmt.Sprintf("Mount point %s is functioning normally.\n", mountPoint)
		fmt.Printf("Mount point %s is functioning normally.\n", mountPoint)
	}

	if err := scanner.Err(); err != nil {
		report_message += "\n"
		report_message += fmt.Sprintln("Error scanning /proc/mounts:", err)
		fmt.Println("Error scanning /proc/mounts:", err)
	}
	report_message = hostname + "告警\n" + report_message

	return code, report_message
}
