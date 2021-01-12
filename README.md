# cache-service : repo for cache microservice

## download source 
    cd $GOPATH/src
    git clone //todo
    
### for installing mockgen
    mockgen -source=<source-file-name> -destination=<destination-file-name> -package=mocks
    e.g.
    mockgen -source=service/cache_service.go -destination=./mocks/mock_cache_service.go -package=mocks
    
### for generating Swagger docs
    First install/update swaggo library using following command
    go get -u "github.com/swaggo/swag/cmd/swag"
    
    Then:  swag init -g router/router.go --parseVendor

