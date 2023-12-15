FROM nginx
# 在 dockerfile 的目录下运行 docker build 是最佳选择 只会构建当前目录上下文不会加载一堆无用文件
RUN echo '这是一个本地构建的nginx镜像' > /usr/share/nginx/html/index.html
