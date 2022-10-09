cd /root/minecraft;
while true;
    do sleep 21600;
    if [ ! -d './buckup' ]; then
        mkdir ./buckup;
    fi;
    cp -R ./worlds ./buckup/$(date \"+"+"%s"+"\").buckup;
    cd ./buckup;
        find ./ -mtime +2 -name \"*.buckup\" -type d | xargs rm -rf;
    cd ..;
done