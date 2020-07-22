# s2i-golang
FROM openshift/base-centos7

      # the maintainer
LABEL maintainer="Javier Romero <root@jromero.codes>" \
      # specify where s2i scripts are located
      io.openshift.s2i.scripts-url="image:///usr/libexec/s2i/bin"

# configuration
ARG GO_VERSION=1.14
ARG GO_INSTALL_DIR=/usr/local/

# env vars
ENV APP_ROOT=/opt/app-root
ENV BUILDER_VERSION 1.0
ENV GOPATH ${APP_ROOT}/src/go
ENV PATH=${PATH}:${GOPATH}/bin:${GO_INSTALL_DIR}/go/bin

# build dependencies
RUN curl -sSL https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz -o go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C ${GO_INSTALL_DIR} -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm -f go${GO_VERSION}.linux-amd64.tar.gz && \
    mkdir -p ${GOPATH}

# s2i scripts
COPY s2i /usr/libexec/s2i

# set permissions
RUN chown -R 1001:1001 ${APP_ROOT} && \
    chown -R 1001:1001 /usr/libexec/s2i

# default user
USER 1001

# default CMD for the image
CMD ["/usr/libexec/s2i/usage"]