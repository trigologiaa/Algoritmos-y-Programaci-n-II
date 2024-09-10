package ej02

func LCS(s1, s2 string) int {
    a := len(s1)
    b := len(s2)
    ab := make([][]int, a + 1)
    for i := range ab {
        ab[i] = make([]int, b + 1)
    }
    for i := 1; i <= a; i++ {
        for j := 1; j <= b; j++ {
            if s1[i - 1] == s2[j - 1] {
                ab[i][j] = ab[i - 1][j - 1] + 1
            } else {
                ab[i][j] = maximo(ab[i - 1][j], ab[i][j - 1])
            }
        }
    }
    return ab[a][b]
}

func maximo(a, b int) int {
    if a > b {
        return a
    }
    return b
}