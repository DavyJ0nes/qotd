FROM prom/busybox
MAINTAINER DavyJ0nes <davy.jones@me.com>
ENV UPDATED_ON 11-02-2017
ENV QOTD_CACHE_FILE /srv/app/cache.txt
ENV QOTD_URL http://quotes.rest/qod.json?category=management

RUN mkdir -p /srv/app/templates/static
WORKDIR /srv/app
ADD templates/index.html /srv/app/templates/index.html
COPY templates/static/* /srv/app/templates/static/
ADD releases/qotd /srv/app
EXPOSE 8080
CMD ["./qotd"]
