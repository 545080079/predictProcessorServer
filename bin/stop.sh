#!/usr/bin bash

#stop jupyter
resJupyter=`ps aux|grep jupyter | grep -v "grep"| awk -F' ' '{print $2}'`
echo "[success] stop jupyter lab pid:"${resJupyter}
if [ -z "${resJupyter}" ]
then
    echo "[warning]jupyter lab pid is null"
fi

kill -9 ${resJupyter}

#stop python predict server
res1=`ps aux|grep flask | grep -v "grep"| awk -F' ' '{print $2}'`
echo "[success] stop python server pid:"${res1}
if [ -z "${res1}" ]
then
    echo "[warning]python server pid is null"
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
