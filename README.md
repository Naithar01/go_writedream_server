# Write Dream

http://ec2-18-212-96-77.compute-1.amazonaws.com:8080/

## HTTP Framwork

- <a href="https://gin-gonic.com/docs/">Gin</a>

## 1.

- Create Project
- Create Issue Table, Model, Router,
- Add Handler ( Get All Issue List )

## 2.

- Create Error Handler
- Add Handler ( Create Issue, Find Issue By Id )

## 3.

- Add Handler ( Update Issue, Delete Issue )
- Create Set Header Middleware

## 4.

- Edit Dir ( Edit Dir Handlers )
- Create Memo Table, Model, Router
- Add Handler ( Get All Memo List, Create Memo, Find Memo By Id, Delete Memo )
- Edit Handler ( Send Message )

## Update 2023-02-10 22:11

- 모든 Issue List를 가져올 때 Memo 테이블에서 Issue Id를 참조하는 Memo의 갯수를 같이 가져오도록 구현 ( OK )

## 5.

- Edit Handler ( Get All Issue List ) -> Memo 테이블에서 Issue 테이블의 id를 외래키로 받고있는데, Issue를 가져올 때 Memo 테이블에서 각 Issue에 맞는 Memo 의 행 갯수를 같이 가져오도록 추가

## Update 2023-02-10 23:38

- Id로 Issue를 찾을 때 Memo 테이블에서 해당 Issue 의 id를 외래키로 저장하고있는 행을 같이 가져오도록 구현 ( OK )

## Update 2023-02-11 16:53

- 모든 Issue List를 가져올 때 페이징 처리 기능 구현 ( Query로 같이 와야하는 Page, PageLimit이 없을 때의 상황도 고려, Memo 테이블에서 Issue 테이블의 id를 참조하는 행이 없는 Issue 테이블의 행도 같이 가져오도록 ) ( OK )

## 6.

- Fix Bug Handler ( Get All Issue List )를 가져올 때 Memo 테이블에서 Issue id를 참조하는 행만 가져오는 버그를 수정
- Create Categories Table, Model, Handler
- Add Handler ( Get All Category List, Create Category, Delete Category )
- Create Issue_Categroy Table
- Edit Dir ( IssuePaginationModel -> IssuePaginationDTO ) models -> dto
- Edit Handler ( Issue를 저장할 때 Category_Issue 테이블에 Issue 테이블의 Id와 Category 테이블의 Id를 같이 저장하도록 구현 ) ( Issue를 삭제하거나, Category를 삭제하면 Category_Issue 테이블의 행도 같이 삭제하게 구현 )

## Update 2023-02-12 24:50

- Issue List를 가져올 때 Query로 받는 Category를 기준으로 Issue를 가져오게 구현 ( OK )

## 7.

- Edit Handler ( 모든 Handler에 Rows를 Close 해주는 코드를 추가 )

## 8.

- Edit Handler ( Category 이름으로 가져오는 Issue와, 그냥 모든 Issue를 가져오는 Handler 추가 )
- Fix Bug Handler ( Issue List를 가져올 때 페이징 처리 코드 수정 )

## 9.

- Create Controllers ( Dir )
- Handler 디렉토리에서 특정 Action을 취하고 Controller에서 Router 요청을 받을 수 있게 수정

## Update 2023-02-14 18:58

- 모든 Issue List를 가져올 때 모든 Issue List의 갯수도 같이 뱉어주게 구현 ( Ok )

## Update 2023-02-19 03:31

- Issue의 상세 페이지 (ReadIssue)에서 Memo도 페이징처리

# 10.

- Dockerfile, Docker-compose, DB-migration 파일 생성
- DB 연결 코드 수정 ( Docker 전용 )
