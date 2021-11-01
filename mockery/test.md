### 运行  
* `go test user_test.go user.go`
* `go test -v -run="TestGetOrderList|TestNewUserInfo"`  
    // 测试多个方法使用|分隔
* `go test -v -run="(?i)double" ./`		
    //执行当前目录下包含double的测试用例(不区分大小写)
* `go test -v  ./...`		
    //递归执行当前目录及下级目录所有测试文件
* `go test -v -coverprofile cover.out user_test.go user.go` 	
    //输出测试覆盖率文件
* `go tool cover -html=cover.out -o cover.html`		
    //根据测试覆盖文件转换为html文件
* `go test -bench="." -run="none"`   
    //压测所有压测用例(-run="none" 表示执行不存在的测试用例,就是不执行测试用例)
* `go test -v -bench="." -run="none" -benchmem -benchtime=3s`
    // 执行基准测试, 5秒,打印内存分配次数和占用字节数
* 其他测试选项
  
### tdd方式
* 自己作为第一个使用者,构建更好的api包
* 提高代码可测试性