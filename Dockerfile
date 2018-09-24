FROM scratch
ADD passwd.minimal /etc/passwd
ADD hello /
USER nobody
EXPOSE 8080
CMD ["/hello"]
