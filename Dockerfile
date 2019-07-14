FROM scratch

EXPOSE 8080

COPY app /

CMD ["/app"]