package function

type FuncArg struct {
    Type string
    Name string
}

func FuncArgMake(typ, name string) FuncArg {
    return FuncArg{
        Type: typ,
        Name: name,
    }
}
