# Persistent Volumes example

This example follows [this tutorial](https://kubernetes.io/docs/tasks/configure-pod-container/configure-persistent-volume-storage) to demonstrate persistent volumes on a minikube instance.  

After creating the `/tmp/data/index.html` on the minikube, first a `PersistentVolume` with the `HostPath` (local dir) plugin is created. Then, a `PersistentVolumeClaim` is created that will be bind to the previously created `PersistentVolume` that is "Available." Then, an Nginx container is created that mounts the `PersistentVolumeClaim` (NOT in the volume itself) to serve up the `index.html` at `localhost`.

NOTE: `storageClassName: ""` needs to be added to the PVC for the PVC to bind to the PV as outlined [here](https://github.com/kubernetes/minikube/issues/1239). Without adding this minikube would create a new PV for the PVC, leaving the intended PV to still be "Available". I'm not sure if this is necessary when in a real Kubernetes environment.  

### Thoughts
Clearly this appears to be the pattern for creating persistent storage that databases can be provisioned to use and share. In AWS you would create a new an EBS volume and use the `AWSElasticBlockStore` persistent volume type when creating a new PV. However a weakness here is that an EBS volume can only be mounted on one EC2 instance and wouldn't this restrict the provisioning of your database containers to only that one EC2 instance? There must be a better way to do this. 
