## 알"고"리즘

---

    - problems: 정답이 제공되지 않는 문제들이 들어있습니다.
    - solutions: 각 문제 항목과 동일한 이름의 파일이 들어있습니다. 정답을 제공해요.

---

1. Num_In_list
    - 배열에 해당 숫자가 있는지 확인
```go
NumInList([]int{1,2,3,4,5}, 5) // true
NumInList([]int{3,3,3,3,3}, 5) // false
NumInList([]int{3,3,5,3,3}, 5) // true

// 빈 배열을 인자로 받은 경우
NumInList(nil, 5) // false
NumInList([]int{}, 5) // false
```

2. Sum
   - 배열의 모든 요소의 합
```go
Sum([]int{1,2,3,4,5}) // 15
Sum([]int{3,3,3,3,3}) // 15 
Sum([]int{3,3,2,5,2}) // 15
```

3. Reserve
   - 문자열 뒤집기
```go
Reserve("cat") // tac
Reserve("alphabet") // tebahpla
```