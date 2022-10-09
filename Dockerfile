FROM archlinux:latest
RUN  pacman -Syyu --noconfirm go 
COPY ./run /root
RUN chmod 777 /root/run
CMD /root/run