import socket
from multiprocessing import Process

from flask import Flask, request
app = Flask(__name__)

@app.route('/v1/predict', methods = ['POST', 'GET'])
def handlerPredict():
    res = """
    {
        "a": "test data!",
        "b": "138888888"
    }
    """
    res = request.form['send1'] + request.form['send2']
    return res
if __name__ == "__main__":
    """
    服务开始运行
    """
    print('predict server start running at ::5000.')