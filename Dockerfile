FROM archlinux:latest
RUN  pacman -Syyu --noconfirm unzip
COPY ./run /root
COPY main.sh /root/
COPY buckup.sh /root/
RUN chmod 777 /root/run /root/main.sh
CMD [ "/bin/bash","-c","/root/main.sh" ]