package main

import (
	"flag"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"
	"mxshop-api/user-web/validators"
	"strconv"
)

func main() {
	host := flag.String("host", "localhost", "Host address")
	port := flag.Int("port", 8021, "Port")
	flag.Parse()
	addr := *host + ":" + strconv.Itoa(*port)

	//初始化配置文件
	initialize.InitConfig()

	//初始化logger
	initialize.InitLogger()

	//初始化routers
	engine := initialize.InitRouter()

	// 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", validators.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	//初始化srv的连接
	zap.S().Infof("服务启动 %s\n", addr)
	if err := engine.Run(addr); err != nil {
		zap.S().Panic("服务启动失败 err: ", err.Error())
	}

	//服务注册

	//接收终止信号
}
