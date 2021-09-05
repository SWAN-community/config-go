# ![Secured Web Addressability Network](https://raw.githubusercontent.com/SWAN-community/swan/main/images/swan.128.pxls.100.dpi.png)
# Secured Web Addressability Network (SWAN) - Config Go

Common configuration for Go. Used by projects like 
[owid-go](https://github.com/SWAN-community/owid-go) and 
[swift-go](https://github.com/SWAN-community/swift-go) to retrieve settings at
startup.

The environment variables are checked for values to override any values 
contained in the configuration file or that are absent from the configuration
file. Field names in the configuration structure are converted from camel case
to the popular upper case underscore separated environment variable key format.
For example; the field name ServiceProvider in a structure will be read from the
environment variable SERVICE_PROVIDER if present.

Structures used with the LoadConfig function should use the mapstructure 
identifier to identify fields to be read from configuration files. For example;
the field ServiceProvider should be declared as follows.

```
ServiceProvider string `mapstructure:"serviceProvider"`
```