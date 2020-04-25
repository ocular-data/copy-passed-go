# Copy-Passed
#### A remote clipboard application

---

[site](https://copy-passed.web.app)

[Download app](https://github.com/ocular-data/copy-passed-android/releases/latest)

---

## Features
- [x] copy from anywhere on phone
- [ ] paste from anywhere on phone
- [x] easy firebase login

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
$ echo 12 | goboard
$
$ goboard
12
```

---
 
## Install

```
go get github.com/ocular-data/copy-passed-go/cmd/goboard
```

---

inspired by [atotto/clipboard](https://github.com/atotto/clipboard)