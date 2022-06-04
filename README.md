# Cron Parser

This this a cli application that can describe standart cron syntax.
For example if you give the following line as argument into our app.

```
*/15 0 1,15 * 1-5 /usr/bin/find
```

It will produce such result

```
minute 0 15 30 45
hour 0
day of month 1 15
month 1 2 3 4 5 6 7 8 9 10 11 12
day of week 1 2 3 4 5
command /usr/bin/find
```

# ❗️Important Notes

- This application does not know about seconds. Usualy cron jobs include seconds. But this app assumes that you will start with minutes.

- Even though this app can handle most of standart patterns. It cannot handle combinations of those patterns very well.

- Because of the problem above some of our tests do not pass. Feel free to fix our code to pass those tests

# Options on how to run this code

## Compile and run

- Go to [oficial page](https://go.dev/dl/). And install latest version of go. Then open up your terminal in the root of this project. And run the following commands.

```
go mod tidy
```

This will install any dependencies we might need. Probably you won't need any but its a good practive. And now you can compile our code with `Makefile`.

```
make
```

If you do not have make installed you could just copy the following command and run it manualy from the root.

```
go build -o bin/cron-parser cmd/cli/main.go
```

Now you should have a binary in the `bin` directory called `cron-parser`. You can just run it as follows.

```
bin/cron-parser  "*/15 0 1,15 * 1-5 /usr/bin/find
```
