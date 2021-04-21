#!/usr/bin bash

export FLASK_APP=predictserver.py
export workdir=/root/workdir/predictProcessorServer
cd $workdir
make
./bin/predict_execute_server &
cd $workdir/python
pwd
flask run --host=0.0.0.0 &
