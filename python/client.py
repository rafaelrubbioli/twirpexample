from twirp.context import Context
from twirp.exceptions import TwirpServerException

from proto import example_twirp, example_pb2

client = example_twirp.ExampleServerClient("http://localhost:6565")

try:
    print("Sending message: 'Hello, server!")
    response = client.MyFunction(ctx=Context(), request=example_pb2.MyFunctionRequest(client=example_pb2.PYTHON, message="Hello, server!"))
    print("Received response from server:")
    print(response.message)

    print("==============")

    print("Asking for sum (123, 456)")
    response = client.Sum(ctx=Context(), request=example_pb2.SumRequest(first=123, second=456))
    print("Received response from server:")
    print(response.result)

except TwirpServerException as e:
    print(e.code, e.message, e.meta, e.to_dict())
