#!/usr/bin bash

export FLASK_APP=predictserver.py
rm ./predict_execute_server
cd ../
make
./bin/predict_execute_server &
cd ./python
pwd
flask run --host=0.0.0.0 &