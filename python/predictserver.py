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



# def handler_client(client_socket):
#     """处理客户端请求"""
#     # 获取客户端请求数据
#     request_data = client_socket.recv(1024)
#     print("request data:", request_data)

#     # 构造响应数据
#     response_start_line = "HTTP/1.1 200 OK\r\n"
#     response_headers = "Server: My server\r\n"
#     response_body = "{\"Response\": \"545080079\"}"
#     response = response_start_line + response_headers + "\r\n" + response_body
#     print("response data:", response)

#     # 向客户端返回响应数据
#     client_socket.send(bytes(response, "utf-8"))

#     # 关闭客户端连接
#     client_socket.close()
    
# if __name__ == "__main__":
#     server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
#     server_socket.bind(("", 8000))
#     server_socket.listen(120)

#     while True:
#         client_socket, client_address = server_socket.accept()
#         print("用户[%s:%s]请求本服务..." % client_address)
#         handle_client_process = Process(target=handler_client, args=(client_socket,))
#         handle_client_process.start()
#         client_socket.close()
# # 在浏览器输入 http://localhost:8000/
# # 会出现 hello itcast
