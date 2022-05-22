package test

import (
	"fmt"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
	"testing"
)

// Test_md5 测试MD5
func Test_md5(t *testing.T) {

	md5 := utils.MD5("123456")
	fmt.Println(md5)
}
