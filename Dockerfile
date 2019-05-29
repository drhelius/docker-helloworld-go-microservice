FROM openshift/rhel7-minimal:7.6
ADD hello /
EXPOSE 8080
CMD ["/hello"]
