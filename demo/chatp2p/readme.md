## 模拟命令
```shell script
p login -u [email] -p [xxx]
p register -u [email] -p [xxx]
p passwordreset -u [email] -p [xxx] -pnew [xxx]
p passwordforgot -u [email] [-verify [code] -pnew [xxx]
p view -gl  # group list
p view -g [group name] [-limit [limit] -msg -member] # view group msg list, default -msg
p manage -join [group name]  # join a group
p manage -create [group name]  # create a group
p manage -del [group name]  # delete a group
p chat -g [group name] -msg [msg text] # send a msg to group
p bot [cmd] [args...]
```

## 服务流程
### client
1. login  
2. select group  
3. get history msg and subscribe group channel  
4. send msg  

### redis
1. save msg to zset  
2. publish to channel  
3. get history by limit  

## redis key
1. h:group [group name] [Group strut]   (group)  
3. h:user [user name] [user struct]     (user)  
2. s:[group name] [user name]           (goup user)  
4. z:[group name] [msg struct]          (group msg)  
5. h:version [version | url]            (version)  