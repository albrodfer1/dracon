# Used only to build Go binaries.
FROM golang:1.23.3 AS builder

WORKDIR /wrk

# Copy only go related files.
COPY go.mod go.sum ./
COPY vendor ./vendor
COPY cmd/git-clone ./cmd/git-clone
COPY pkg ./pkg
# And builder.
COPY Makefile .

# Install make
RUN apt-get update && \
    apt-get install -y make && \
    rm -rf /var/lib/apt/lists/*

RUN make build-target

# Create a workspace to clone repos to.
RUN mkdir -p /workspace

# Used to actually run the binary in minimal image.
FROM gcr.io/distroless/base-debian12:nonroot

COPY --from=builder --chown=65534:65534 /wrk/target /bin/target
COPY --from=builder --chown=65534:65534 /workspace /workspace

# Run as UID for 'nobody' since k8s pod securityContext runAsNonRoot can't resolve the user ID:
# https://github.com/kubernetes/kubernetes/issues/40958
USER 65534

# Setting the workdir where we'll clone repositories.
WORKDIR /workspace

# Set the binary as the entry point
ENTRYPOINT ["/bin/target"]
