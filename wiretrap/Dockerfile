FROM rust:latest as builder
WORKDIR /app
COPY . /app
RUN cargo build --release

FROM gcr.io/distroless/cc
COPY --from=builder /app/target/release/wiretrap /
CMD ["./wiretrap 0.0.0.0:2408"]

