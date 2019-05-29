FROM openshift/rhel7-minimal:7.6
ADD hello /
CMD ["/hello"]
