FROM golang:1.21

ENV GOPATH=/go/src
ENV WORKSPACE=${GOPATH}/app
RUN mkdir -p ${WORKSPACE}

WORKDIR ${WORKSPACE}

COPY . ${WORKSPACE}

RUN go mod download
RUN go mod tidy

RUN ls -l

RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
ENV PATH="${GOPATH}/bin:${PATH}"