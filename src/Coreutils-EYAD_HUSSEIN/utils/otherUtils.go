package utils


func ContainsFlag(flags []string, flag string) (int, bool) {
        for index, f := range flags {
                if f == flag {
                        return index, true
                }
        }
        return -1, false
}