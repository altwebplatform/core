FROM ubuntu:16.10

MAINTAINER millisecond@gmail.com

# To run and get shell locally:
# sh compile.sh && docker build . -t millisecond && docker run -it millisecond /bin/bash

# Install required packages and remove the apt packages cache when done.
# removed version lock of nginx, the old packages are being removed, maybe we need to host our own like the ansible scripts do
RUN apt-get update && apt-get install -y \
    build-essential=12.1ubuntu2 \
    git=1:2.9.3-1 \
    nginx=1.10.1-0ubuntu1.2 \
    wget=1.18-2ubuntu1 \
    curl=7.50.1-1ubuntu1.1 \
    net-tools=1.60+git20150829.73cef8a-2ubuntu1 \
    vim=2:7.4.1829-1ubuntu2.1 \
    libffi-dev=3.2.1-6 \
    libssl-dev=1.0.2g-1ubuntu9.1 \
    rsyslog=8.16.0-1ubuntu5 \
    rsyslog-relp=8.16.0-1ubuntu5 \
    rsyslog-gnutls=8.16.0-1ubuntu5 \
    python-setuptools=26.1.1-1 \
  && rm -rf /var/lib/apt/lists/*

RUN easy_install pip==9.0.1

RUN pip install -I --isolated \
    supervisor-wildcards==0.1.3 \
    supervisor==3.3.1

# Logging
#COPY ./conf/logging/rsyslog.conf /etc/
#COPY ./conf/logging/50-default.conf /etc/rsyslog.d/
#COPY ./conf/logging/logrotate.conf /etc/logrotate.d/rsyslog
#COPY ./conf/logging/logrotate.cron /etc/cron.hourly/logrotate
#RUN chown -h syslog:adm /var/log/webtraffic
#RUN rm -f /etc/cron.daily/logrotate

# Update certs to latest version hosted by curl
RUN curl --remote-name --time-cond /usr/local/share/ca-certificates/cacert.pem https://curl.haxx.se/ca/cacert.pem

# Supervisor
COPY ./conf/supervisor-*.conf /etc/supervisor/conf/

# Scripts
COPY ./bin/* /usr/local/bin/
RUN chmod a+x /usr/local/bin/reloader

# Configuration
#COPY ./conf/nginx.conf /etc/nginx/nginx.conf
#COPY ./conf/mime.types /etc/nginx/mime.types
COPY ./conf/supervisord.conf /etc/supervisor/

RUN mkdir -p /tmp/supervisord

#RUN chown -hR syslog:adm /tmp/rsyslogd

RUN echo "alias start='supervisord -c /etc/supervisor/supervisord.conf'" >> /root/.bash_aliases
RUN echo "alias restart='supervisorctl mrestart \"*\"'" >> /root/.bash_aliases
RUN echo "alias stop='supervisorctl mstop \"*\"'" >> /root/.bash_aliases
RUN echo "alias status='supervisorctl status'" >> /root/.bash_aliases
RUN echo "alias setenv='source /usr/local/bin/environment.sh'" >> /root/.bash_aliases

ENV TERM xterm-256color

WORKDIR /srv

CMD ["supervisord", "-n", "-c", "/etc/supervisor/supervisord.conf"]
