clear

cd ..

#(
#echo "building kk-schedule-ui"
#cd kk-schedule-ui && npm run generate
#)
#(
#echo "kk-schedule-server vendor"
#cd kk-schedule-server && go mod vendor
#)

#docker pull --platform linux/amd64 golang:1.24.7
#docker pull --platform linux/arm64 golang:1.24.7
#docker pull --platform linux/amd64 alpine:latest
#docker pull --platform linux/arm64 alpine:latest

export GOVERSION=1.24.7

export TAG=0.0.1


# local build test
#(
#docker build \
#--build-arg GOVERSION=${GOVERSION} \
#-t kk-schedule-local:${TAG} \
#-f ./image-build/Dockerfile .
#)

# docker login

(
docker buildx build --platform linux/amd64,linux/arm64  \
            --build-arg GOVERSION=${GOVERSION} \
            -t cruvie/kk-schedule:${TAG} \
            -f ./image-build/Dockerfile . \
            --push
)


(
docker buildx build --platform linux/amd64,linux/arm64  \
            --build-arg GOVERSION=${GOVERSION} \
            -t cruvie/kk-schedule:latest \
            -f ./image-build/Dockerfile . \
            --push
)


