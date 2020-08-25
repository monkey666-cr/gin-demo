## HTTPS

生成私钥文件(server.key) 和自签发的数字证书(server.crt)

openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O＝RandomCompany/OU=IT/CN=127.0.0.1/emailAddress=17610780919@163.com"
