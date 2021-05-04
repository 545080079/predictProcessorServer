import json
import socket
from multiprocessing import Process

from flask import Flask, request
app = Flask(__name__)

@app.route('/v1/predict', methods = ['POST', 'GET'])
def handlerPredict():
    """
    res likes:
    {
        "Name": "n1",
        "Next": [{
            "Name": "n2",
            "Next": [None]
        }]
    }
    """
    
    #读取模型
    inputOutputMap = {}
    with open('InputOutputMap') as f:
    for line in f.readlines():
        l = line.split(' ')
        if l[1][:-1] == '':#如果没有换行，代表是最后一个
            inputOutputMap[int(l[0])] = int(l[1])
        else:  #除最后一行都会有\n,需要去掉
            inputOutputMap[int(l[0])] = int(l[1][:-1])
            
    nameCountMap = {}
    with open('NameCountMap') as f:
        for line in f.readlines():
            l = line.split(' ')
            if l[1][:-1] == '':#如果没有换行，代表是最后一个
                nameCountMap[l[0]] = int(l[1])
            else:  #除最后一行都会有\n,需要去掉
                nameCountMap[l[0]] = int(l[1][:-1])
    
    print(inputOutputMap)
    print(nameCountMap)
    
    data = json.loads(request.form['dag'])
    #根据NodeName-Output Map统计模型查找预测值
    res = {}
    while data != None:
        res[data["Name"]] = nameCountMap.get(data["Name"], -1)    #默认-1
        data = data["Next"][0]
    
    #根据Input-Output Map统计模型查找预测值
    while data != None:
        #只对NodeName-Output未给出预测的节点使用该方法
        if res[data["Name"]] == -1:
            res[data["Name"]] = inputOutputMap.get(data["Parameters"][0], -1) #默认-1
    
    
            
        
    return res


if __name__ == "__main__":
    
    """
    服务开始运行
    """
    print('predict server start running at ::5000.')