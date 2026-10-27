package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.2" )

func QMxke7pNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVB96mRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Udanef8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o46SsPa3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dS5PrgVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcrcRbD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmuzGSXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzqFAtCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3kHFbJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCAEuyg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0w4ij4i5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DgvM56jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ATBlmKbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJIgJnSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4UHDi5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAp0sfgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Ai2DegPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GrLUsthGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9gY6ihEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdkNZgceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DrtBr4IfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqroUqytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFM89MywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZ8rRnMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Quck7OIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bexPgBSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UJAIE9l8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WV5Yx53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWc1r0mmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CqdMC1rqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rl0kH48qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjwOfqSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func prGpOc9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eSD2So5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SNg8APGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wT1SN4OOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvlxGSdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUR08WJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWMjoOR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baLt5J4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PkhsuofCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDcDXd7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sk1cexYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9WNhH1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTBH1jR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSVwP4RhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6P5HbzBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cuRnrZxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qvuj0cwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEzA8erNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcREjWWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxi0YfSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lRyQH4F0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0vjlCm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P05MuAVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEXq3E0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9SKRSWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbcNWsjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FobLkfPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuWsCNHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRvQJ0vjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S82az8NmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2MoOKUx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QG5PXjsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OTIBjz5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJSseaUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ce69PhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5O6dOMDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cafFfK0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkK6646WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2UdYM3JxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2MYc87YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnSzmxa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tcfpGHxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y6LFgmMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sofQbnSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nKdwgoQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 79TpaeJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KEKvom8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQNWMZaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLua0eHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9PigxFFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0aX7qCTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjEI7ppBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQl4rreNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQZK0QeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbFNOmAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CKlq8V6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSJhHVNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLgNHTPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vdgo2RnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UKXShUgmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANsN6OcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53bHE8rPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5VXEWWw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KF06qpmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VhdtkXtXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FA0LdpriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTtt9D6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SHZeOvwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o5p7osngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DvOG2cdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0fFRkuxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func buSh9Ul4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Cd2KqFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3cpz089Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRDxyTuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upqd0lv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MrOdT07QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JguuyxBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoxFub6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X51KFvAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXBYssfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ETRizQ8OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSQAczphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCyL8QBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFCgteeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3TC9m4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98PtkIKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4bC8yzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8ePuME9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNsDrwk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1DI919oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FaINKa8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aao7IvWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ecry81d1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GboHMunVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yquxoaWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9zlZ0EyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ugGZwXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jMdriBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAEl35wzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Umh6jGtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5W1eZZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdZWvxXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dt1FH25rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmTek4CmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4AZic5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dI1Up6wIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yn4848DOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BElUpbU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WIH9atruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VP47QNHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MnmrprFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVKIMuuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IbhvLb7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ok07k7grWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Du7IzYJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDzksQYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TMMKAJCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pGPnI9jvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQGqHytwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUTnIa8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odHGlpdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8Y80mzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4sTJfnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztP4fRl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9POcrFGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNqQ4LWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fepu3PoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WqHdTUfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gUt8GfleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLGZcRzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4pcpPPfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wM7pkfOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qnof3gf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkk1zLcoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gs9yBLA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func driKRs03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOuUbGnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JyHADLkwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJsT1cqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHTkyG92Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAyN6cKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6zRERapGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1PQMj6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZisOXN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZERz45IGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNEzlUF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SP9Ete3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BbjG6JyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dqIBPCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPoLX6saWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVh737fVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3foMLCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpqKji5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVutBTyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9S7ylj83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c0nPJrHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHiXGJO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKXVtVMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xO3vGGRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMb1zaE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FOeF1f3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7v0eTwOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEY9sONBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQeGgrm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VErdlVGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQZVZKAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N6bvYlsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGjs4BHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJZK2CkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QW3E0e0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwMDimvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywzTwh6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQuuPbFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrcSbfVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wejZTkqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TurMVlLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wG8RZBemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfSYmF05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aV7EPTr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ny4AixGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PC2XP4qQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hejbp1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNabvVfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IaEaPyWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDw7XWj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z1pDylzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhQXwVfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKrMKI6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6ZelReOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojKW9xWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jR7uaMeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqfKiRKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I21tmSgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kVN5iq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQss4tyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8WGY4mMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nn2uuarWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CGS0ryZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MrJU3P0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRupdW2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YW67fiMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Slbsh7IzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0eP1eOO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99FIq3xMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRVnR2KPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OkGuiXIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rElpY5WkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F4fVZhrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tTIVicN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2hIMBGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gI9TmPlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4vjv4xrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgvnT3kNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m76qHuvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1A6p7mDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k4ZuKIk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KtmdrVWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5H1qwPVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ahaDneqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WuoBNX6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzEcXdIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rch5mNnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func af5edyLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZP3DsyscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0HC85cpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kuzab9ByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzGz6uJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZTgKwhEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9VEh8i9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func whC5XYtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZrnJeJEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0BCTOIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5rx4ugQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aE5mWY2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRv2SadcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w4DuiPj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZK0bc27fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8vOJBqWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhSYlmUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGDNy5wCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHdTQuMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func awt75eHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2kVy3APWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q0IpC8TLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1UlqSyVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KEMCdTH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func esID8XSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAYj4aLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1oswbumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRCHLuSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2GuOaT8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func znjdgFBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eNOHnkzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vWiq7UeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rjw1EzHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DD6KjlVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXZr65pFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8SJUzPbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xSfEimV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ro7pHIGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TiPrlMIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QEXdvBtGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tq0HDdOhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUZRZj3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DOyf9vTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PwyCRgw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mAjDBVFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tTrjyYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q6ClvC3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FA7a7slWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vul55pP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJMQcL1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8uE27BrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DOlnHQsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ixmGtyHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAIJoxERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mOWwaSbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fop8l9RAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9hIz8gQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ataCoQIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9U5Huy4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbCXKzhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8obtrj8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMQkD4PIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eivypBTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUVERPI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V11aZtJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMOdNMh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VlAqln9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vTy6YmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xt4aNgSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKn4Z8tVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Zjm94c2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PS0cVRnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Llas0oEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4YI6H9RiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vFgAo4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7Tsu1o4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XcFUQKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v52qavGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5EoTR7YZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func smdauuA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmDiE7YCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ypz6zD0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xUhwI52RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXyxoV6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ExLmx1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KyzDeBUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cVL5OeAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T86GpO3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqK7S9asWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YyW9dqkwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tOfYxX5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptPZZ4OFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZG32SiWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJXcZvdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bbkspHSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gh8da1O1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oSpzFyqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQTCakBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WsoGxUQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Kyczq8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func My8Dqh72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBlZNMhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WL6zj9GhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qRdZIbBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AF4XbKxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REk1BfwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvceIDxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1NkNxbcZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qvQ09y41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0uDQ3wuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJ62blAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nd7Uh7SmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AebIGLfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VruDNzfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyzHzHINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIRrsQbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0B1U2iaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8ZfRKaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJLFbLgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjKlBsvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Yt0BOKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8VssajKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqxW5b7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rpynBCriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnipvRqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KyTtpZH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtPp6ssrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNYAYv4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgYIEYQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4WwUNkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsRMP9tDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 645wayIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func egaxfZuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUouiEP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vj1jwvqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uPqTvpinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C55VWWuxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zyesSZgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptUqEkL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdSxfiMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPorz3W1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vg4tkQOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3J8EftFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H71wqYZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VfJGVsV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8pRnFQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QkLdZOiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f4TmjdN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgkWASEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQfnvadNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJ8ZtnjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8XLztvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQTh5ZUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzYi6EXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lq5oSCLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLvgSurlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0waPxvzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlC6TOC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QTnrRkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gzP8nYDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqaBnN3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhbTVYOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHuDJqN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Py1sRRHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxNNt78XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJTaEXz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qNvZcGL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AlGpvwLHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIsGtVy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aS7TgwXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnHWSQ0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCMQiYm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4yeoyZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 69yUAMJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YdJDolm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ay9CGgWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IhaxuL2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdpEjjTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ybm9YB8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0R9EwtbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFqHMu0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HuQnLa2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4p7AevGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UO4Y0p3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZurAfaVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXiCTWqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9Jqt1o7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pCStizsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func depJPbDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1iPJ2k4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4E4ddoe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func luoHeE6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPfteMj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpBrxnaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLXZsaJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UShau2rdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func olUkT67EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNSaUYgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0MCFGd9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h1pFybWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5HWiIz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mFncbPkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjeSfABYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBNmYxG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJS7XGrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 767gcSQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dD0sRQT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYvNrPPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiYVxWsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jI1IQMK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9WEuTLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzgIDnP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyrlwlzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQnFULOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65vSz4ycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsmbAxCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBW5ciC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TNuHtT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKoGVbtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qp3635wSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFT7F94TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqhbMuHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JzRMBeUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BRU9IQRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n6pH0osQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sisjAA91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTiSFzAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KEiJKIWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYpE6NRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmy36U2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t1BIKpxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HP9BOzBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ShaH8QcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6o7oyVGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8WLqz9eKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWQSZBn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJjNKOMaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0KYfqZc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ane0dWJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BA40eOlxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6FIfozFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FRdesEupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPb6wBGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjXER2yzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 445HWYUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQoIuzdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oz7XTKlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJmdBPtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDemd6s0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kIcy8smhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zgXielRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rOXFS4xNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdML1r0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cN8XF8SDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wbWElOFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yZ2KOOjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atApzJtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdtBYv82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQh71CRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipK02eA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQcA0l9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bnRsJybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nRb79x7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbF1Ns1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yulGwFzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 052TkVKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wyWajewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfJWfKjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFnUyzMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lpw4Uy72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4S6KWtKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jIHSWYhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BbzMQKLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejYeWTbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdEufS3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XpPu3KtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YqkrYYPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0ZwJH9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WXstTuzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PI4Oeou6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RmlvUNizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4aShakbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EI8jRxZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LaERGojAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gwIvI0dIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhKHJYomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4helKmNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ETZGwfSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hEM08f5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbnCVRL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TmRAwitZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRWDvErmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oqs4KyblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYCis42iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OBXkAk9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJocqsy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQBQ2Ry9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OOPWYD2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nz1Opi0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RNgOcizOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yi1QY5YaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgI9XdPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVIpSPHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLPJ0nD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S4ipiVqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y81ngimlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y35JDvlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQYCuqocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9bEdKVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFQDF9kAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rf0AC0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gkC9x6HPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcLzccFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNgJpwjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FqIsxF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1BG5j2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3C28WswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4rCDfuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYjxeWeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UraszjvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vemgYrQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dNv7iFMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxL3E0P7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGAFbHEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoGNI2ETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFen1hLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mrm7y71HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1e11tUf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Zh7d80eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPkqwG8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBYsNKRIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eSZOY6wmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNAvyfK7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OxVMVSHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hpxfnJzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cw4KP1GgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A1gJCyeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdFlaGWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yuZf56AqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2RyPC0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOJDBNlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xdrv6G1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uajg7vXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dfJLx50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEh5gn5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbWunaDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGDGvuWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yc6JzRDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6dpiuKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CoGpbacJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDYVxNXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ulYkJaFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G6rVMOpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzWJxMkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EPrIkMd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7NUrziIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItkO0tSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OAs0mwFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rOPG2MbWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLASIU1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNUeonJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oU27BHFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgEqygUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func huMimzH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ahPuLg1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztp5GeNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkBXwFujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAhzYtarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gyJFEOHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3HUMj2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wuAnhF1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eY4rPCLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGDaHzSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjUj96cqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8ah7d5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayiCDny3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8tEBm5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MpvOVykWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DIC4aB9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFYJaRkxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DAO8pbVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cSTE0OyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDAmosd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fTKLSLRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BGkvLeNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3V9OKHmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KE0tzdIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6lJmyCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bTbaOiXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 368k6qFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tdf9YwbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mAamgrWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mt4FrNfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QXE2tmBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cMtyVxIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjCAN6CWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2HhES9MjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LD8RltnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSNOJ8ZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wb0dpsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jTqz096jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnaZdwESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3l6IemHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5dxCYwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1NYFWR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFbT3Q7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J72oTcKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4YuKDfPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DC9wL5igWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAle5sl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xwxbn8s0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULZevUBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lmVrBYoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvZwhv2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysmEULn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5i8yMQr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJo9Xyl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3SIOeNB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWBIQyz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bxR3IDBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZ9ln22kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCsXai3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGjQh8tmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func azkQ0kF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3VaaE3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGu8Ri5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Tqj6ycAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gL4Psd7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uI8FoefhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gH8kutmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uwUbzGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CIfhmHitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCvNGRaAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func leMBWz8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQ2D9H5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZIxoHBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Ph8ezaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PlGU3cFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STg1ShnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSYz2AYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hwcdklLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y8ymX3AbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWnOMVOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMCePxlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2zxhK6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jckF6vxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDUw7kkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8TQApFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9pZzkYEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s3t5p4FxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcNxwYRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbXPhhJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vX0t5Eq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p3ER0pmMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yzkw5mRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVQ9ckn7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPaLmAg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pUBiFJapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFoorVmmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSGrDKcVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7HdOfzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCWYAEPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fUgDrdNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func skXy853gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dI5ZEKzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmUIkAhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enff73FUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qv1QiKGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8Pr4H7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 69TRJgznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldN8bslhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkrkhx6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lxiso051Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NiFwq0C6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2FolS8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QsBiT7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81f0g1tsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jNJgPJcoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdVAHdW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zld79zX2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bbMFM6vSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHtqV6xNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FY374PMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VeKz0ofgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAq8p1HfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2VQ95IeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGs3N0BJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSb7ZmQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kw5n79IgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IyQjeTX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6Lgm3PuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQKzwZNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJuXa9BvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6RjsnYpZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYnG4MCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpwxCzqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXF736n9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxfb9Y98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9cKvj5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHTTMPQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIojzmAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhtKw4XQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vm0XdchtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func skASShAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S4kkBtftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7alDg66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHDc5S4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBBV24ygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqvqTJCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xxcegJZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func naiG7fZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPomJdUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TddKMU08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dtmMKgpZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GK2UMH2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uv5XHYpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haTNwvLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLHLEhDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9Gzb5sWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rr94Bu7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iz3kfbQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRN1Ky0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R6fCoS4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xSk2nFLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pdm6XoPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1l14XJ1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RIw09kWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RevQwOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k79sImKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ak8yAb39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6pTP1z5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvKpk2IOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJVuJVjxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHbExdRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otpktb1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYA4fwpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXS92MtXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pd18OI2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3ZA2BTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JvkrxQyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kp7fFoFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Po9mSNP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vyoZ2iemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gxblVpTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5oD0KgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I1DXJEyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwKNoXSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wS5NA6ETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kErLCQU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJPKyKGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3C1U2qjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acF5p68OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUPCNIioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPI6Y0qbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkka52SgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZJwtWdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jhli5R8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mR3geh32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpFiBLdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkoFE5WUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2jUEtfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfjcDgrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfpduzmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

