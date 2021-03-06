# Angular_Laravel_GinGonic | VICENT COLL Y ALEX ZARAZUA

<p align="center">

  <h2 align="center"><strong>KALYPSO APP</strong></h2>

</p>

## INDEX

* About this project
* Getting Started
* Preview 
* Features
* Go Packages
* Built With and technologies

         
## About This Project

KalypsoApp is an dockerized application made during the second course of DAW (Development of web applications).

In Kalypso you will be able to can buy clothes from the best brands on the market as well as save the clothes that you like the most to have them saved and be able to buy them when you most want and have them on your profile.
As you make purchases, your karma points will increase, with which you will be able to access many offers and discounts.


## Getting Started

To get the repo running locally:

 * 1.- Clone this repo
 * 2.- Install Docker Community Edition : 
   * ` docker-compose up --build `


## Preview

  * **KALYPSO HOME**


   <img src="./frontend/src/img/homePreview1.png">
       <img src="./frontend/src/img/homePreview.png">


  * **Products Shop List**

    <img src="./frontend/src/img/Preview_List.png">


  * **List Product Description**

      <img src="./frontend/src/img/preview2.png">

  * **Product Detail**

      <img src="./frontend/src/img/preview_detail.png">     

  * **Register**

     <img src="./frontend/src/img/Register_Preview.png">

  * **Login**

    <img src="./frontend/src/img/Login_Preview.png">



## Features

| Page | Features |
| - | - |
| Home | Carousel with some images,Popular Brands with GO and MySql , and Products More Visited by Go  |
| Shop | List with GO and MySql, Details with Go and MySql |
| Product | Product Info , and Favorite button and Buy Product Button |
| Settings | LogOut and  User Settings |
| Profile | User info, favourited Garment |
| PanelAdmin | Dashborad For Users Admin that they can Create and Delete products with Laravel |

<br>

| Service | Features |
| - | - |
| Register | Regular register  |
| Login | Regular login with Go if is an User client and with Laravel if is an Admin User |
| Favourites | Favourite button in each garment , favourited show up on profile  |


<br>

| Techical Feature | Where it works |
| - | - |
| Docker | Entire application is dockerized |
| Redis | In Login Service and Karma in Brands and Products for more Visited or More Popular in case of the Brands |
| Authentication | Login and Register Services with JWT |

<br>



## Go Packages

* https://github.com/kardianos/govendor
* https://github.com/pilu/fresh
* https://github.com/go-sql-driver/mysql
* https://github.com/gin-gonic/gin
* https://github.com/jinzhu/gorm
* https://github.com/go-redis/redis
* https://golang.org/x/crypto/bcrypt
* https://github.com/dgrijalva/jwt-go

## Built With

 * GinGonic
 * Laravel with PHP 7.0
 * Angular 11.0.2
 * Docker
 * Docker-Compose
 * MySQL
 * Redis


## Other Technologies

 * Conduit - starting template
 * JWT
 * CSS3
 * Toaster
 * Gorm
 * Traefik
 * Grafana
 * Prometheus


