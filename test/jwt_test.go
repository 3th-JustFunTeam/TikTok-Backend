package test

import (
	"fmt"
	"github.com/3th-JustFunTeam/Tiktok-Backend/utils"
	"testing"
)

// Test_createJWT 测试创建token
func Test_createJWT(t *testing.T) {
	userId := 10001
	token, err := utils.GenerateToken(userId)
	if err != nil {
		fmt.Println("token 创建失败")
		return
	}
	fmt.Println(token)
}

// Test_ParseToken 测试解析token
func Test_ParseToken(t *testing.T) {
	// 用上面创建的token值进行测试 默认创建的是1天过期, 这个时间可以在 tokenUtil 里设置
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMDAwMSwiZXhwIjoxNjUzMjI1OTkzLCJpc3MiOiJUaWtUb2sifQ.JGsfa6f0YhV6NfNBUIx47PH7ZVswft-_dta_C3GkEhU"
	claims, err := utils.ParseToken(token)
	if err != nil {
		fmt.Println("token 解密失败, 或者 token 过期")
		return
	}
	// claims 有创建token时放里面的信息 (userId, ExpiresAt, Issuer) 还需要什么数据可以在自己加
	fmt.Println(claims.UserId, claims.ExpiresAt, claims.Issuer)
}
