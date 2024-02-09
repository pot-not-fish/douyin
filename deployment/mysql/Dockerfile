# 使用官方MySQL镜像
FROM mysql:8.0
# 设置root用户的密码
ENV MYSQL_ROOT_PASSWORD=123456
# 设置字符集
ENV LANG=C.UTF-8
# 将自定义配置文件复制到镜像中
COPY my.cnf /etc/mysql/my.cnf
# 将初始化SQL脚本复制到镜像中
COPY init.sql /docker-entrypoint-initdb.d/
# 持久化数据卷
VOLUME /var/lib/mysql