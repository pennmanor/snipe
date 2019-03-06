# Overview

This is a demo of interacting with the snipe api. The snipe api documentation is
available [here](https://snipe-it.readme.io/reference#api-overview) with examples
in php, python, javascript, etc.

This particular demo is in golang.

## Starting Point

Starting at line [132](https://gitlab.pennmanor.net/district/snipe-api-demo/blob/master/main.go#L131) you can see a
call a function to get a list of loaners and then loop over them. The getLoaners() function starts on line 65 is only
about 20ish lines with error checking. 

Most of the other code is boilerplate stuff to start a webserver, setup routes, render templates.

## Note on calling our instance of the snipe API

One catch with our instance of snipe is that it is behind a google identiy aware proxy. That means
these calls can't directly hit our snipe URL's from outside of the server cluster. 

To test this I have been using port forwarding to access the container directly. It is really easy to setup and I can
create service accoutns for anyone who needs access.

