# issue-2163
https://github.com/kevwan/go-zero/tree/fix/cors

debug the issue-2163

steps:

1、 Reference the branche fix/cors to my local repostory 
[link: github.com/Holyson/test-go-zero-cors v0.0.1]

![image](https://user-images.githubusercontent.com/5540291/185849151-1f10fa76-c3b4-459d-b90e-f5a5852b2e3c.png)


2、 Start the main program [api/weixinapi.go]
if your port conflict,please change your port at etc/xx.yaml

![image](https://user-images.githubusercontent.com/5540291/185848996-a7b5a4f5-0092-4e33-bbca-85fcb5dd3f29.png)

3、Curl the url
```bash
curl -i -X POST http://127.0.0.1:1112/test
```
the response header include the CORS info

<img width="566" alt="image" src="https://user-images.githubusercontent.com/5540291/185849646-ecae5e39-8928-47a0-9658-8fb3ce3bad2c.png">

4、Test pass 
