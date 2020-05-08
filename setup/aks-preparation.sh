# show azure sub details
az account show

# login to azure with username and password
az login

# get latest available AKS version for <region>, do not forget to change <region>, you can use "northeurope" for Europe regions
version=$(az aks get-versions -l <region> --query 'orchestrators[-1].orchestratorVersion' -o tsv)

# create resource group for AKS in the region, do not forget to change correct <region> and <resource-group> , syntax should be {yourname}-aks-rg
az group create --name <resource-group> --location <region>

# create AKS service, do not forget to change <unique-aks-cluster-name> and <region>, provide <APP_ID>, <APP_SECRET> values, version is already set above
# option 1 - basic model
az aks create --resource-group <resource-group> \
    --name <unique-aks-cluster-name> \
    --location <region> \
    --kubernetes-version $version \
    --generate-ssh-keys \
    --load-balancer-sku basic \
    --service-principal <APP_ID> \
    --client-secret <APP_SECRET>

# option 2 -- auto scaler enabled
az aks create --resource-group <resource-group> \
    --name <unique-aks-cluster-name> \
    --location <region> \
    --kubernetes-version $version \
    --generate-ssh-keys \
    --vm-set-type VirtualMachineScaleSets \
    --enable-cluster-autoscaler \
    --min-count 1 \
    --max-count 3 \
    --load-balancer-sku basic \
    --service-principal <APP_ID> \
    --client-secret <APP_SECRET>



    az aks create --resource-group yahya-aks-rg \
    --name yahya-aks \
    --location westeurope \
    --kubernetes-version 1.17.0 \
    --generate-ssh-keys \
    --load-balancer-sku basic \
    --service-principal <APP_ID> \
    --client-secret <APP_SECRET>


# to get kubeconfig
az aks get-credentials --name <unique-aks-cluster-name> --resource-group <resource-group>