
def closure_point():
    px = 0
    py = 0

    def setfn(x, y):
        nonlocal px, py
        px = x
        py = y

    def getfn():
        nonlocal px, py
        return (px, py)

    def move():
        nonlocal px, py
        px += 1
        py += 1

    return {'set': setfn, 'get': getfn, 'move': move}

def build_closure_point():
    return closure_point()


def closure():
    x = 1
    def add():
        nonlocal x
        x += 1
    def get():
        nonlocal x
        return x

    return add, get

def main():
    my_obj = closure_point()
    my_obj['set'](1, 2)
    print(f"my_obj set at: ", my_obj['get']())
    my_obj['move']()
    print(f"my_obj after move: ", my_obj['get']())

    new_obj = build_closure_point()
    print(f"new_obj created at: ", new_obj['get']())
    new_obj['move']()
    print(f"new_obj after move: ", new_obj['get']())

    print(f"my_obj still at: ", my_obj['get']())

def test():
    add1, get = closure()
    add1()
    add1()
    print(get())

#test()
main()
