# Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

    helm repo add prime https://udhos.github.io/prime

Update files from repo:

    helm repo update

Search prime:

    $ helm search repo prime -l --version ">=0.0.0"
    NAME       	CHART VERSION	APP VERSION	DESCRIPTION                
    prime/prime	0.1.0        	0.0.0      	A Helm chart for Kubernetes

To install the charts:

    helm install my-prime prime/prime
    #            ^        ^     ^
    #            |        |      \__________ chart
    #            |        |
    #            |         \________________ repo
    #            |
    #             \_________________________ release (chart instance installed in cluster)

To uninstall the charts:

    helm uninstall my-prime

# Source

<https://github.com/udhos/prime>
