FROM gitpod/workspace-full

WORKDIR /home/gitpod
RUN go get github.com/bazelbuild/bazelisk