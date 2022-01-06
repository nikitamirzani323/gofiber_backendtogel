FROM golang:alpine AS agenclientbuilds
WORKDIR /appbuilds
COPY . .
RUN go mod tidy
RUN go build -o binary

FROM node:lts-alpine AS agenclientsveltebuilds
WORKDIR /svelteapp
COPY [ "svelte/package.json" , "svelte/package-lock.json" , "svelte/rollup.config.js" , "./"]

FROM agenclientsveltebuilds AS agenclientsveltebuilds2
RUN npm install
RUN cp -R node_modules prod_node_modules


FROM agenclientsveltebuilds AS agenclientsveltebuilds3
COPY --from=agenclientsveltebuilds2 /svelteapp/prod_node_modules ./node_modules
COPY ./svelte .
RUN npm run build

FROM alpine:latest as agenclientrelease
WORKDIR /app
RUN apk add tzdata
COPY --from=agenclientsveltebuilds3 /svelteapp/public ./svelte/public
COPY --from=agenclientbuilds /appbuilds/binary .
COPY --from=agenclientbuilds /appbuilds/.env /app/.env
ENV PORT=7072
ENV PATH_API_BACKEND="http://128.199.241.112:7072/"
ENV TZ=Asia/Jakarta

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT [ "./binary" ]