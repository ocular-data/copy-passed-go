# Copy-Passed
#### A remote clipboard application

---

[site](https://copy-passed.web.app)

---

## Features
- [x] pipe into command to copy
- [x] use command to paste
- [x] copy globally from local clipboard
- [x] save to local clipboard on global copy
- [ ] select in history what to paste
- [ ] access groups clipboard

---

## linked projects
[Firebase hosting website](https://github.com/ocular-data/copy-passed-firebase)

[Native console access](https://github.com/ocular-data/copy-passed-terminalAccess)

[Android app](https://github.com/ocular-data/copy-passed-android)

---

### explanations
This app allows uses to copy paste text cross
platforms as long as they have internet.

This app can be used to by piping into and out of 
the command on any platform, it will also add the content 
to the computers clipboard (if available). 
note that local clipboard is not available from wsl

```bash
$ goboard -h
 Usage: [PIPE IN] goboard [OPTION]
 Copy Paste across multiple interfaces
 
 Available Options:
   -h, --help                Display this help message and exit
   -c, --clipboard           Copy from the computers clipboard instead of stdin
   -r, --re-authenticate     Revoke token registered to computer and register new
                              one
   -s, --save                Save the output of the global clipboard to the local
                              one
 
 By default, goboard copys from stdin and pastes to stdout. Use the -c option to override.
```
```bash
$ echo 12 | goboard
$
$ goboard
12
```

---
 
## Install

```
go get -u github.com/ocular-data/copy-passed-go/cmd/goboard
```

---

inspired by [atotto/clipboard](https://github.com/atotto/clipboard)
