# Welcome to url parser 
  
Parser will fetch html content of url based on provided and it will find tag count.Archive data will stored inside docker volumes will be used for view html page 
from locally 

# Steps to run 

  ### Step 1 :

      will build docker image parser

        docker build . -t parser

  ### Step 2 :

      will run parser with command line arguments 

         docker run parser https://www.google.com https://tutorialedge.net/golang/go-docker-tutorial/