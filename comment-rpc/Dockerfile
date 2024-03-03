FROM scratch

# 移动到工作目录：/build
WORKDIR /build

# 将配置文件复制到容器中
COPY config.yaml app /build/

# 设置时区
ENV TZ Asia/Shanghai

# 暴露端口
EXPOSE 8887

# 需要运行的命令
ENTRYPOINT ["./app"]