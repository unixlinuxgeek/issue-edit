### Github Issue Editor


Testing repository:
```bash
$ git clone https://github.com/unixlinuxgeek/issue-edit
$ cd issue-edit
$ go test -v  -u unixlinuxgeek  -r gopl-issues  -t <GITHUB_TOKEN> 
```
replace *unixlinuxgeek* to your account
replace *gopl-issues to* your repository (with issues)

### Run all unit tests:
### Запуск  всех модульных тестов:

```bash
$ go test -v  -u unixlinuxgeek -r gopl-issues -t <GITHUB_TOKEN> 

=== RUN   TestCreate
issue state: open, owner: unixlinuxgeek, repo: gopl-issues
    issueedit_test.go:23: TestCreate is PASSED
        
--- PASS: TestCreate (1.43s)
=== RUN   TestReadAll
------------------------------------------------------------------
repository name: gopl-issues
language: Go
open issue count: 5
------------------------------------------------------------------
issue number: 9
issue state: open
issue title: test_issue_2024-09-22_13:13
------------------------------------------------------------------
issue number: 8
issue state: open
issue title: test_issue_2024-09-22_13:09
------------------------------------------------------------------
issue number: 5
issue state: open
issue title: test_issue_2024-09-21_16:12
------------------------------------------------------------------
issue number: 4
issue state: open
issue title: test_issue_2024-09-21_12:50
------------------------------------------------------------------
issue number: 2
issue state: open
issue title: test_issue_updated_2024-09-21 21:23
------------------------------------------------------------------
    issueedit_test.go:34: TestReadAll is PASSED
        
--- PASS: TestReadAll (0.71s)
=== RUN   TestReadOne
    issueedit_test.go:45: TestReadOne is PASSED
        
--- PASS: TestReadOne (0.83s)
=== RUN   TestUpdate
issue number: 824633810816
issue id: 2535932264
issue title: 824636084928
issue url: https://api.github.com/repos/unixlinuxgeek/gopl-issues/issues/1
issue state: closed
issue labels: []
    issueedit_test.go:58: TestUpdate is PASSED
        
--- PASS: TestUpdate (0.71s)
=== RUN   TestClose
Issue state: closed, owner: unixlinuxgeek, repo: gopl-issues. issue number: 1
    issueedit_test.go:69: TestClose is PASSED
        
--- PASS: TestClose (0.51s)
PASS
ok      github.com/unixlinuxgeek/issueedit      4.19
```


### Running single unit test:
### Запуск одиночных модульных тестов:

The TestCreate creates an issue:

Модульный тест TestCreate создает issue:
```bash
$ go test -v -run TestCreate -u unixlinuxgeek -r gopl-issues  -t <SECRET>
=== RUN   TestCreate
Issue unixlinuxgeek/gopl-issues/test_issue_2024-09-21_16:12 is created
--- PASS: TestCreate (1.13s)
PASS
ok      github.com/unixlinuxgeek/issueedit      1.136s
```

The TestReadOne read one issue:

Модульный тест TestReadOne читает один issue:
```bash
$ go test -v -run TestReadOne  -u unixlinuxgeek   -r gopl-issues  -num 4  -t <GITHUB_TOKEN> 
=== RUN   TestReadOne
----------------------------------------------------------------------------
issue title: test_issue_2024-09-21_12:50
issue number: 4
issue id: 2540057339
issue state: open
issue labels: []
issue url: https://api.github.com/repos/unixlinuxgeek/gopl-issues/issues/4
----------------------------------------------------------------------------
--- PASS: TestReadOne (1.63s)
PASS
ok      github.com/unixlinuxgeek/issueedit      1.637s
```

The TestReadAll read all issues:

Модульный тест TestReadAll читает все issue:
```bash
$ go test -v -run TestReadAll -u unixlinuxgeek -r gopl-issues -t <GITHUB_TOKEN> 
=== RUN   TestReadAll
------------------------------------------------------------------
repository name: gopl-issues
language: Go
open issue count: 4
------------------------------------------------------------------
issue number: 7
issue state: open
issue title: test_issue_2024-09-22_12:09
------------------------------------------------------------------
issue number: 5
issue state: open
issue title: test_issue_2024-09-21_16:12
------------------------------------------------------------------
issue number: 4
issue state: open
issue title: test_issue_2024-09-21_12:50
------------------------------------------------------------------
issue number: 2
issue state: open
issue title: test_issue_updated_2024-09-21 21:23
------------------------------------------------------------------
--- PASS: TestReadAll (1.53s)
PASS
ok      github.com/unixlinuxgeek/issueedit      1.537s
```
The TestUpdate updates the issue:

Модульный тест TestUpdate обновляет issue:
```bash
$ go test -v -run TestUpdate -u unixlinuxgeek  -r gopl-issues -num 7  -t <GITHUB_TOKEN> 
=== RUN   TestUpdate
issue number: 824634941040
issue id: 2540722157
issue title: 824638827344
issue url: https://api.github.com/repos/unixlinuxgeek/gopl-issues/issues/7
issue state: closed
issue labels: []
--- PASS: TestUpdate (1.17s)
PASS
ok      github.com/unixlinuxgeek/issueedit      1.177s
```


The TestClose closes the issue:

Модульный тест TestClose закрывает issue:
```bash
$ go test -v -run TestClose  -u unixlinuxgeek  -r gopl-issues  -num 7  -t <GITHUB_TOKEN>
=== RUN   TestClose
Issue state: closed, owner: unixlinuxgeek, repo: gopl-issues. issue number: 7
--- PASS: TestClose (1.36s)
PASS
ok      github.com/unixlinuxgeek/issueedit      1.362s
```
