module api

go 1.14

require github.com/dgrijalva/jwt-go v3.2.0+incompatible

require service v0.0.0

replace service v0.0.0 => ../service

require model v0.0.0

replace model v0.0.0 => ../model