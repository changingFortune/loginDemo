## test
    1.用户id 无重复整数 pseudo_encrypt   done
        sonyflake包:时间戳+机器码+自增序列
    2. RESTful和gRPC接口协议(注册登录)   done
    3.高可用,横向伸缩的注册和登录          done 
        高可用:1.使用守护进程监控进程状态,进程挂起后自动重启
              2.阈值弹性扩容
              3.故障转移
              4.ip漂移
        横向伸缩:
             1.无状态服务器,即jwt+redis实现登录的无状态,redis实现逻辑服务无状态
             2.有状态服务器,一致性哈希
    4.三段不重复用户名
    5.面积和周长问题           done
        如需要面积最长则走圆
        如需要周长最长则走方形


### 测试用例
    // 获取用户
    curl -X 'GET' \
    'HTTP://127.0.0.1:11000/v1/user' \
    -H 'accept: application/json'
    
    // 注册
    curl -X 'POST' \
    'HTTP://127.0.0.1:11000/v1/user/register' \
    -H 'accept: application/json' \
    -H 'Content-Type: application/json' \
    -d '{
    "username": "123",
    "password": "11111"
    }'
    
    // 登录
    curl -X 'POST' \
    'HTTP://127.0.0.1:11000/v1/user/login' \
    -H 'accept: application/json' \
    -H 'Content-Type: application/json' \
    -d '{
    "username": "123",
    "password": "11111"
    }'

