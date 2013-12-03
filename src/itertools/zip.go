package itertools

func ZipLongest(fillValue string, args ...interface{}) (out chan []interface{}) {
    out = make(chan []interface{}, GetIterBuffer())

    go func() {
        defer close(out)
        var argsSlice [][]interface{}
        maxSliceLength := 0

        for _, arg := range args {
            var argSlice []interface{}

            for p := range Iterate(arg) {

                if (p.First.(int) + 1) > maxSliceLength {
                    maxSliceLength++
                }
                argSlice = append(argSlice, p.Second)
            }
            argsSlice = append(argsSlice, argSlice)
        }

        for i := 0; i < maxSliceLength; i++ {
            var row []interface{}

            for _, a := range argsSlice {
                if len(a)-1 >= i {
                    row = append(row, a[i])
                } else {
                    row = append(row, fillValue)
                }
            }
            out <- row
        }
    }()
    return
}
