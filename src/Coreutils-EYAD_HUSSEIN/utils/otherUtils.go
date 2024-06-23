package utils


func ContainsFlag(flags []string, flag string) bool {
        for _, f := range flags {
                if f == flag {
                        return true
                }
        }
        return false
}