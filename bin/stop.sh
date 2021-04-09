#!/usr/bin bash

#stop python predict server
res1=`ps aux|grep flask | grep -v "grep"| awk -F' ' '{print $2}'`
echo "[success] stop python server pid:"${res1}
if [ -z "${res1}" ]
then
    echo "[warning]python server pid is null"
    exit 1
fi

kill -9 ${res1}

#stop go server
res2=`ps aux|grep predict_execute_server | grep -v "grep"| awk -F' ' '{print $2}'`
echo "[success] stop go server pid:"${res2}
if [ -z "${res2}" ]
then 
    echo "[warning]go server pid is null"
    exit 1
fi

kill -9 ${res2}
