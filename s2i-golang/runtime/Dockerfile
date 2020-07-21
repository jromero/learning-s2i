# s2i-golang-runtime
FROM centos:7

# the maintainer
LABEL maintainer="Javier Romero <root@jromero.codes>" \
      io.openshift.s2i.scripts-url="image:///usr/libexec/s2i/bin"

# s2i scripts
COPY s2i /usr/libexec/s2i

# create user/group
RUN groupadd -g 1001 app && \
    adduser -u 1001 -g 1001 default

# default user
USER 1001