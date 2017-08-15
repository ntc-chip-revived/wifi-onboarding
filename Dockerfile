FROM scratch

ADD http://dl-cdn.alpinelinux.org/alpine/v3.6/releases/armhf/alpine-minirootfs-3.6.2-armhf.tar.gz /
#ADD qemu-arm-static /usr/bin/qemu-arm-static

RUN cat /etc/issue && \
\
apk update &&  \
apk add \
    dnsmasq \
    hostapd \
    iptables

RUN mkdir APP
ADD ./start.sh /APP/start.sh
ADD build/linux_arm/wifi-onboarding /APP/wifi-onboarding
ADD ./static /APP/static
ADD ./view /APP/view

ADD hostapd.conf /etc/hostapd.conf
ADD dnsmasq.conf /etc/dnsmasq.conf

WORKDIR /APP
CMD ["/APP/start.sh"]
