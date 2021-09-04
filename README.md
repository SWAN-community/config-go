Common storage configuration for Go. Used by projects like 
[owid-go](https://github.com/SWAN-community/owid-go) and 
[swift-go](https://github.com/SWAN-community/swift-go) to persist storage
settings.

If no values are provided in the settings file provided then the environment
variables are checked for values. See configuration.go for the values 
considered.