FROM nginx:latest

COPY html /usr/share/nginx/html

COPY default.conf /etc/nginx/conf.d/

ENTRYPOINT [ "/docker-entrypoint.sh" ]
CMD [ "nginx", "-g", "daemon off;" ]