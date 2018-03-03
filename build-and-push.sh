#!/bin/sh

docker build -t wingyplus/wingymomobot . \
    && docker tag wingyplus/wingymomobot registry.heroku.com/infinite-journey-20895/web \
    && heroku container:login \
    && heroku container:push web --app infinite-journey-20895 \
    && heroku container:logout
