FROM archlinux:latest
RUN  pacman -Syyu --noconfirm go 
COPY ./run /root
CMD /root/run