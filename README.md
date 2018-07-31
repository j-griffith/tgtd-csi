## tgtd-csi

`tgtd-csi` is a CSI driver built to use a targetd endpoint.

targetd is a simple tool that takes any linux box and converts
it to an iscsi storage appliance using LVM2, iSCSI and a simple
REST pkg.

Based off of the Kubernetes [targetd driver](https://github.com/kubernetes-incubator/external-storage/tree/master/iscsi/targetd)

To issue requests to a [targetd](https://github.com/open-iscsi/targetd) device.
