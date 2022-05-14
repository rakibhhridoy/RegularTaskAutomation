import sys

code = """#!/bin/env python


def main():
    print("")








if __name__ == "__main__":
    main()

"""

def main():
    arg1 = sys.argv[1]
    print(arg1)

    if arg1.endswith(".py"):
        f = open(arg1, "w")
        f.write(code)
        f.close()
    else:
        print("Python file should have .py extension!")


if __name__ == "__main__":
    main()

