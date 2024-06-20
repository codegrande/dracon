FROM gcr.io/openssf/scorecard:stable

FROM alpine:3.18.3

COPY --from=0 /scorecard /scorecard

ENTRYPOINT ["/scorecard"]