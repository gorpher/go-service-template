# 指定基础镜像
FROM iron/base
# 作者与邮箱
MAINTAINER gorpher gorpher101@gmail.com

# 创建镜像,执行内部命令,创建workdir目录
RUN mkdir -p /home/root/workdir

ADD service-linux-amd64  /home/root/workdir

# 设置环境变量
ENV APP_HOME  /home/root/workdir
ENV PATH $APP_HOME:$PATH

RUN chmod -R 777  $APP_HOME

EXPOSE 8080

WORKDIR $APP_HOME

# 容器启动执行命令
ENTRYPOINT ["service-linux-amd64"]
