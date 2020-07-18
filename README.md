
# Setup

### Install s2i

Releases: https://github.com/openshift/source-to-image/releases

```shell script
curl -sSL https://github.com/openshift/source-to-image/releases/download/v1.3.0/source-to-image-v1.3.0-eed2850f-darwin-amd64.tar.gz -o /tmp/s2i.tgz
tar -xvf /tmp/s2i.tgz -C /usr/local/bin/
s2i version
```

# Basic Operation

### Build

Using example from: https://github.com/vorburger/s2i-java-example

```shell script
s2i build https://github.com/vorburger/s2i-java-example fabric8/s2i-java vorburger:s2i-java-example
```

Breakdown:

```shell script
s2i                                             
build                                           # command (in)
https://github.com/vorburger/s2i-java-example   # source (in)
fabric8/s2i-java                                # builder image (in)
vorburger:s2i-java-example                      # image name (out)
```

Result:

```
$ docker images

REPOSITORY                                TAG                                        IMAGE ID            CREATED             SIZE
vorburger                                 s2i-java-example                           6026dfa0226f        11 minutes ago      529MB
```

#### Notes

- The source is downloaded from URL using `git`. Alternatives exist to use local non-VCS sources.
- By default running the build command does not use cache meaning, for example, all dependencies are downloaded again. Option exists (`--incremental`) to enable this such as `s2i build --incremental https://github.com/vorburger/s2i-java-example fabric8/s2i-java vorburger:s2i-java-example`.
- Sources exist in runtime image by default (see [screenshot](resources/runimage-src.png)). An alternative _might_ be to provide an additional runtime image via `--runtime-image`. Attempting to use the same builder image yields in an error:
    ```
    $ s2i build --runtime-image fabric8/s2i-java https://github.com/vorburger/s2i-java-example fabric8/s2i-java vorburger:s2i-java-example
    warning: Image "fabric8/s2i-java" does not contain a value for the io.openshift.s2i.assemble-input-files label
    Build failed
    ERROR: An error occurred: no runtime artifacts to copy were specified
    ```
- Providing a seperate runtime image with `--incremental` isn't supported:
    ```
    $ s2i build --runtime-image fabric8/s2i-java --incremental https://github.com/vorburger/s2i-java-example fabric8/s2i-java vorburger:s2i-java-example
    ERROR: Incremental build with runtime image isn't supported
    ```

### Run

```shell script
docker run -p 8080:8080 vorburger:s2i-java-example
```

### Comparison

#### Cloud Native Buildpacks

```shell script
git clone https://github.com/vorburger/s2i-java-example
cd s2i-java-example
pack build -B gcr.io/paketo-buildpacks/builder:base vorburger-cnb
```

Result
```text
$ docker images

REPOSITORY                                TAG                                        IMAGE ID            CREATED             SIZE
vorburger-cnb                             latest                                     481a35d057d9        40 years ago        247MB
```

##### Notes

- Doesn't download source via `git`.
- Cache (incremental build) works by default.
- Source is not present in final runtime image only the compiled classes (see [screenshot](resources/cnb-runimage-classes.png)).