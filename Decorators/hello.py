def hello_decorator(func):
    def wrapper(*args, **kwargs):
        print(f"Hello, {func(*args, **kwargs)}, it was nice to meet you!")
    return wrapper


@hello_decorator
def greetings(name):
    return name


greetings("Vova")
