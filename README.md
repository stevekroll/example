## K6
#### Build Custom K6 Image
* `task k6:pull`
* `task k6:build`

#### Create New Load Test Script
* `task k6:new filename={$filename}`

#### Run Load Test Script
* `task k6:run filename=`

## Grafana
#### Configure Grafana Dashboard
* Open `http://localhost:3000`
* Select 'Create New Dashboard'
* Create a new data source and select `Prometheus`
* Use http://host.docker.internal:9090
* Import a new dashboard using JSON found in `tools/grafana/dashboards`

## Prometheus
#### Run Unit Tests
* `task prometheus:test`
