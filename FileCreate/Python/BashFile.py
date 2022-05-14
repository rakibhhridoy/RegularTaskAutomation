import sys

code = """#!/bin/bash
    
"""

def main():
    arg1 = sys.argv[1]
    print(arg1)

    if arg1.endswith(".sh"):
        f = open(arg1, "w")
        f.write(code)
        f.close()
    else:
        print("Bash file should have .sh extension!")

if __name__ == "__main__":
    main()

