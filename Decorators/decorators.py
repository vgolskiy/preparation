def hello_decorator(func):
    def greetings_wrapper(*args, **kwargs):
        print(f"Hello, {func(*args, **kwargs)}, it was nice to meet you!")
    return greetings_wrapper


def repeat_decorator(func):
    def repeat_wrapper(*args, **kwargs):
        func(*args, **kwargs)
        func(*args, **kwargs)
    return repeat_wrapper


def time_spended(func):
    import time

    def time_wrapper(*args, **kwargs):
        start = time.time()
        res = func(*args, **kwargs)
        end = time.time()
        print(f"Time spended: {end - start}")
        return res
    return time_wrapper
