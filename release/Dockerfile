# syntax=docker/dockerfile:1

FROM --platform=$TARGETPLATFORM alpine:3.18

WORKDIR /app

COPY ./release/entrypoint.sh .
ARG project_name
ARG TARGETPLATFORM
COPY ./dist/$TARGETPLATFORM .

ARG version
ARG tag
ENV APP_BINARY=${project_name}
ENV VERSION=${version:-develop}
ENV TAG=${tag:-master}

ENTRYPOINT ["./entrypoint.sh"]
