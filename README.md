# Project Title

Blazingly fast RESTful API starter in Golang for small to medium scale projects. 

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

For smooth execution make sure you have these installed:

* Go 1.16+
* make
* entr

### Installation

A step by step guide that will tell you how to get the development environment up and running.

```
$ make
$ make build
$ ./autoreload.sh 
$ Start Coding...😃

```

## Usage

A few examples of useful commands and/or tasks.

```
$ First example
$ Second example
$ And keep this in mind
```

## Deployment


* run ``make build`` 
* add your domain to nginx.conf in the config diretory
* once there is a binary named "api.out", copy it into the appropriate directory of your choice
* copy ``nginx.conf`` to from config directory to  ``/etc/ngnix/sites-available/{nameoftyoursite}`` on production web server
* create a symbolic link from the file you just copied and direct it  ``/etc/ngnix/sites-enabled/{nameoftyoursite}`` on your server
* the site should then be accessible by domain



### Packages

* chi
* go-Json
* godotenv
* logrus
* viper
* mysql
* gorm


## Contributions

With your help we can make this a real good starter template for starting a web service.
Contributions are welcome!

## License

All contributions will be licensed as Apache 2.0
