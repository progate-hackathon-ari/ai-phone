FROM postgres:15.4

RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
ENV TZ=Asia/Tokyo
ENV LC_COLLATE=ja_JP.UTF-8
ENV LC_CTYPE=ja_JP.UTF-8