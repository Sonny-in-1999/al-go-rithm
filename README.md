## 알"고"리즘

---

    - problems: 정답이 제공되지 않는 문제들입니다.
    - solutions: 각 문제의 정답입니다. 문제와 동일한 파일명을 갖습니다.
    - test: solutions에 작성한 정답 파일을 검증하는 테스트 코드 입니다.
---
      다음과 같은 방법으로 작성한 로직을 검증합니다 :
      test 디렉토리에 각 solution의 테스트 코드를 작성합니다.
      터미널을 통해 test 디렉토리에서 go test . 명령어 실행
      테스트 코드를 각각 실행(파일 단위)하고 싶다면 go test <파일명>
---
1. Num_In_list
    - 배열 요소 검증
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
