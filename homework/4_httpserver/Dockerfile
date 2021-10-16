FROM ubuntu
ENV MY_SERVICE_PORT=80
LABEL multi.label1="httpserver_homework" 
ADD bin/amd64/4_httpserver /httpserver
EXPOSE 80
ENTRYPOINT /httpserver