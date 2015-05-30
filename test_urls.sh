#!/bin/bash

mkdir out && cd out
wget https://ianonavy.com/files/urls.txt
../gocraw -file=urls.txt
