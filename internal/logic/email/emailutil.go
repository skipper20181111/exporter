package email

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	Rand "crypto/rand"
	"exporter/internal/svc"
	"exporter/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"math/rand"
	"net/smtp"
	"strings"
)

type EmailUtilLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailUtilLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailUtilLogic {
	return &EmailUtilLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailUtilLogic) SendMail(EmailInfo *types.EmailInfo, address []string, subject string, body string) (err error) {
	for _, emailUserInfo := range EmailInfo.EmailUser {

		auth := smtp.PlainAuth("", emailUserInfo.User, EasyDecypt(emailUserInfo.Password, svc.Keystr), EmailInfo.Host)
		contentType := "Content-Type: text/html; charset=UTF-8"
		for _, v := range address {
			s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s",
				v, emailUserInfo.NickName, emailUserInfo.User, subject, contentType, body)
			msg := []byte(s)
			addr := fmt.Sprintf("%s:%s", EmailInfo.Host, EmailInfo.Port)
			err = SendMail(addr, auth, emailUserInfo.User, []string{v}, msg)
			if err != nil {
				fmt.Println(err.Error(), "发送邮件产生了异常 ************************************************************")
				return err
			}
		}
	}
	return
}

func (l *EmailUtilLogic) SendMailRandom(EmailInfo *types.EmailInfo, address []string, subject string, body string) (err error) {
	address = EmailInfo.Send2Who
	body = fmt.Sprintf("已同步发送给多人: %s  邮件内容: %s", strings.Join(address, ","), body)
	emailUserInfo := EmailInfo.EmailUser[rand.Intn(len(EmailInfo.EmailUser))]
	auth := smtp.PlainAuth("", emailUserInfo.User, EasyDecypt(emailUserInfo.Password, svc.Keystr), EmailInfo.Host)
	contentType := "Content-Type: text/html; charset=UTF-8"
	for _, v := range address {
		s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s",
			v, emailUserInfo.NickName, emailUserInfo.User, subject, contentType, body)
		msg := []byte(s)
		addr := fmt.Sprintf("%s:%s", EmailInfo.Host, EmailInfo.Port)
		// 使用自制sendMail函数，可以不用tls
		err = SendMail(addr, auth, emailUserInfo.User, []string{v}, msg)
		if err != nil {
			fmt.Println(err.Error(), "发送邮件产生了异常 ************************************************************")
			return err
		}
	}
	return
}
func EasyDecypt(src, keystr string) string {
	defer func() {
		if e := recover(); e != nil {
			return
		}
	}()
	key := []byte(keystr)
	decode_data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	//生成密码数据块cipher.Block
	block, _ := aes.NewCipher(key)
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, decode_data[:aes.BlockSize])
	//输出到[]byte数组
	origin_data := make([]byte, len(decode_data)-aes.BlockSize)
	blockMode.CryptBlocks(origin_data, decode_data[aes.BlockSize:])
	//去除填充,并返回
	return string(unpad(origin_data))
}
func Encrypt(msg string, keystr string) (string, error) {
	key := []byte(keystr)
	src := []byte(msg)
	//生成cipher.Block 数据块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	} else if len(src) == 0 {
		return "", errors.New("src is empty")
	}

	//填充内容，如果不足16位字符
	blockSize := block.BlockSize()
	originData := pad(src, blockSize)

	//加密，输出到[]byte数组
	crypted := make([]byte, aes.BlockSize+len(originData))
	iv := crypted[:aes.BlockSize]
	if _, err := io.ReadFull(Rand.Reader, iv); err != nil {
		return "", nil
	}
	//加密方式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(crypted[aes.BlockSize:], originData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Decrypt(src, keystr string) (string, error) {
	defer func() {
		if e := recover(); e != nil {
			return
		}
	}()
	key := []byte(keystr)
	decode_data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", nil
	}
	//生成密码数据块cipher.Block
	block, _ := aes.NewCipher(key)
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, decode_data[:aes.BlockSize])
	//输出到[]byte数组
	origin_data := make([]byte, len(decode_data)-aes.BlockSize)
	blockMode.CryptBlocks(origin_data, decode_data[aes.BlockSize:])
	//去除填充,并返回
	return string(unpad(origin_data)), nil
}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}
