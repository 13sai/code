def sai():
    for i in range(10):
        for j in range(10):
            if i + j > 13:
                return i, j


print(sai())


class Break(Exception):
    pass


try:
    for i in range(10):
        for j in range(10):
            if i + j > 13:
                raise Break
except Break:
    print(i, j)


def flag():
    btn = False
    for i in range(10):
        for j in range(10):
            if i + j > 13:
                btn = True
                break
        if btn:
            break
    print(i, j)


flag()


def loop():
    for i in range(10):
        for j in range(10):
            if i + j > 13:
                break
        else:
            continue
        break
    print(i, j)


loop()
