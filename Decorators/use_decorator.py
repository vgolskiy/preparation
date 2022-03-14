from decorators import repeat_decorator
from decorators import time_spended


@time_spended
@repeat_decorator
def print_string(s):
    print(s)


print_string("You are not along!")


@time_spended
def google_it():
    import requests
    response = requests.get('https://google.com')
    return response


print(google_it())
