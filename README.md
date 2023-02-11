# Write Dream 

## HTTP Framwork
* <a href="https://gin-gonic.com/docs/">Gin</a>

## 1. 
* Create Project
* Create Issue Table, Model, Router, 
* Add Handler ( Get All Issue List )

## 2. 
* Create Error Handler
* Add Handler ( Create Issue, Find Issue By Id )

## 3. 
* Add Handler ( Update Issue, Delete Issue )
* Create Set Header Middleware

## 4.
* Edit Dir ( Edit Dir Handlers )
* Create Memo Table, Model, Router
* Add Handler ( Get All Memo List, Create Memo, Find Memo By Id, Delete Memo )
* Edit Handler ( Send Message )

## Update 2023-02-10 22:11
* 모든 Issue List를 가져올 때 Memo 테이블에서 Issue Id를 참조하는 Memo의 갯수를 같이 가져오도록 구현 ( OK )

## 5. 
* Edit Handler ( Get All Issue List ) -> Memo 테이블에서 Issue 테이블의 id를 외래키로 받고있는데, Issue를 가져올 때 Memo 테이블에서 각 Issue에 맞는 Memo 의 행 갯수를 같이 가져오도록 추가

## Update 2023-02-10 23:38
* Id로 Issue를 찾을 때 Memo 테이블에서 해당 Issue 의 id를 외래키로 저장하고있는 행을 같이 가져오도록 구현 ( OK )

## Update 2023-02-11 16:53
* 모든 Issue List를 가져올 때 페이징 처리 기능 구현 ( Query로 같이 와야하는 Page, PageLimit이 없을 때의 상황도 고려, Memo 테이블에서 Issue 테이블의 id를 참조하는 행이 없는 Issue 테이블의 행도 같이 가져오도록)
