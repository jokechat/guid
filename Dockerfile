FROM --platform=$TARGETPLATFORM alpine:3.18
ARG RELEASE_NAME
RUN echo "build using release ${RELEASE_NAME}"
COPY release/${RELEASE_NAME} /bin/guid

EXPOSE 80
EXPOSE 8080

CMD ["/bin/guid"]