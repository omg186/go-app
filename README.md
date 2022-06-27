# go version
    go version go1.18.2
# 开启GO111MODULE
    go env -w GO111MODULE=on
# 使用中国代理
    go env -w GOPROXY=https://goproxy.cn,direct
# 安装第三方依赖
    go get -u github.com/gin-gonic/gin
    go get -u github.com/gin-contrib/cors
    go get -u gorm.io/driver/mysql
    go get -u github.com/jinzhu/copier
    go get -u github.com/go-redis/redis/v9
    go get -u github.com/swaggo/swag/cmd/swag
    go get -u github.com/swaggo/files
    go get -u github.com/swaggo/gin-swagger
    go get -u github.com/sirupsen/logrus
    go get -u golang.org/x/sys
    go get github.com/satori/go.uuid
    go get gorm.io/gorm // **注意** GORM `v2.0.0` 发布的 git tag 是 `v1.20.0`
# 初始化项目
    go mod init yungengny.com/system
# swagger
## 检查swagger是否安装成功
    swag -version
## 生成接口文档
    在main.go平级下执行swag init命令生成swagger的docs文件夹，每次接口有变化都需要执行swag init

