# middleware-paddle
基于paddleOcr的web组件 [PaddleOCR](https://github.com/PaddlePaddle/PaddleOCR)




## Usage

### 下载项目

``` bash
git clone https://github.com/Anderyly/middleware-paddle.git
```

### 编译并后台运行

```bash 
cd middleware-paddle
go mod init && go mod tidy
go build
chmod +x middlewarePaddle
nohup ./middlewarePaddle &
```

>url地址

```text
http://ip:8080/api/picture
```
>post请求参数
>
1. 参数1 type 1 参数2 file 文件
2. 参数1 type 2 参数2 base 图片base64

>返回

![截图1](https://raw.githubusercontent.com/Anderyly/middleware-paddle/master/1.png)
