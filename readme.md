# choco

This is a interpreter language whose host language is golang.

**IMPORTANT** This project is based on [Writing An Interpreter In Go](https://interpreterbook.com/)

## what choco can do

```bash
let thresholdForAdult = 20;
let user = {
    "name": "Tom",
    "age": thresholdForAdult + 10
}

let userrole = if (user["age"] >= thresholdForAdult) {
                 return "ADULT"
               } else {
                 return "CHILD"
               }

let showNameIfAdult = fn(user) {
    if (userrole == "ADULT") {
        puts(user["name"])
    }
}
```

you can find more example in `./examples` directory

## how to use

```
$ git clone git@github.com:vsanna/choco.git
$ cd choco

# modify according to your environment
$ echo "export PATH=\"$(pwd):\$PATH\"" >> ~/.zshrc

# to use repl
$ ichoco

# to run yourcode
$ choco your-code.choco
```

## how to build(for dev)

```bash
$ rm -f ichoco && cd app/repl && go build main.go && mv main ../../ichoco && cd ../../
$ rm -f choco && cd app/runner && go build main.go && mv main ../../choco && cd ../../
```

## TODO

- add options for cmdline
  - e.g. debug mode
- exntending hash
- make blockstatement return value

## (Off topic) why "choco"?

this lang is named after one of my dogs, and "ちょこ"っと使える(can be used a little in Japanese) sounds like "Choco".
