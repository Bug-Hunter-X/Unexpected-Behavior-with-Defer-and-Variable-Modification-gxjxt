func main() {

        x := 10
        fmt.Println("Value of x:", x)

        xCopy := x
        defer func() {
                fmt.Println("Value of x in deferred function:", xCopy)
                xCopy++
        }()

        x = 5
        fmt.Println("Value of x after modification:", x)

}
