import sys
code = """package main
import ("fmt"
                )


func main(){
    fmt.Println("====!====")
}


"""


def main():

    arg1 = sys.argv[1]
    print(arg1)

    if arg1.endswith(".go"):
        f = open(arg1, "w")
        f.write(code)
        f.close()
    else:
        print("Golang file should have .go extension!")


if __name__ == "__main__":
    main()

