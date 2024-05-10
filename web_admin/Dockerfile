FROM node:16.14.2 as builder

COPY ./ /web_admin

RUN cd /web_admin && \
    npm install i && \
    npm run build:prod

FROM nginx:1.26.0-perl

ENV TZ="Asia/Shanghai"
RUN echo "Asia/Shanghai" > /etc/timezone
RUN /bin/ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=builder /web_admin/dist /opt/web_admin
