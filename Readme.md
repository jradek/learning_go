# Readme

* Setup [direnv](https://direnv.net/)

    ```bash
    eval "$(direnv hook bash)"
    ```

* build example

    ```bash
    cd hello
    go build hello
    ./hello
    ```

* test example

    ```bash
    cd waiig
    go test monkey/lexer
    ```