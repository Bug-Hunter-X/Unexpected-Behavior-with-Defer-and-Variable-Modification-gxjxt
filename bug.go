func main() {


        x := 10
        fmt.Println("Value of x:", x)

        defer func() {
                fmt.Println("Value of x in deferred function:", x)
                x++
        }()

        x = 5
        fmt.Println("Value of x after modification:", x)

}
