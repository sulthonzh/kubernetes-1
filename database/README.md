# Database

We're using a single database for the demo but ideally you should run one per service in a production setup.

## Install Steps

Launch database

```shell
for config in *.yaml; do 
	kubectl create -f $config
done
```

Create databases and grants using `platform.sql`

```shell
node=`kubectl get pods | grep pxc-node3 | awk '{print $1}'`
kubectl exec $node -i -t -- mysql -u root -p -h pxc-cluster
```
