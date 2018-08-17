FROM swift:4.1 as builder

COPY . /application
WORKDIR /application

RUN mkdir -p /build/lib && cp -R /usr/lib/swift/linux/*.so /build/lib
RUN swift build -c release && mv `swift build -c release --show-bin-path` /build/bin

FROM ubuntu:16.04

WORKDIR /app

RUN apt-get -qq update && apt-get install -y \
  libicu55 libxml2 libbsd0 libcurl3 libatomic1 libmysqlclient20 \
  && rm -r /var/lib/apt/lists/*

COPY Public/ ./Public/
COPY --from=builder /build/bin/hambach-admin .
COPY --from=builder /build/lib/* /usr/lib/
COPY /html/ ./html/

EXPOSE 8080
ENTRYPOINT ["./hambach-admin"]
