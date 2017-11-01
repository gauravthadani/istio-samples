FROM scratch
ADD istio-samples /

EXPOSE 8080
CMD ["/istio-samples"]

