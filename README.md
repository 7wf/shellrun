### Shellrun

Run shell commands from a website inside your machine. [[demo]](https://youtu.be/I0aXh0_cJ6I)

#### Installation

##### Requirements

- **Required**: Linux, [`zsh`](https://zsh.org), `gnome-terminal`.
- **Optional**: [`Go`](https://golang.org), [`git`](https://git-scm.com/downloads).

> :warning: The current version of **Shellrun** can only run on **Linux** with `gnome-terminal` and `zsh`.
>
> - Support to other shells and terminals are planned.

##### Downloading the source code

###### with `git`

```sh
git clone https://github.com/7wf/shellrun.git
```

###### without `git`

Download the repository through GitHub by clicking `Clone or download` and `Download ZIP`.

After obtaining the `.zip` file, you can extract it to a folder.

##### Server

To run commands from the browser, you need to have setup a local server.

With the server running, the browser can communicate with your machine through a HTTP server.

###### with `go`

Run `go run server/main.go` inside the repository folder.

###### without `go`

Download the server binary from [Releases](https://github.com/7wf/shellrun/releases) and run it through terminal.

```sh
./shellrun_server-linux-amd64
```

##### Extension

###### Chromium-based

Go to `chrome://extensions`, enable the `Developer Mode`, `Load Unpacked` and select Shellrun extension folder.

#### License

[MIT](/LICENSE) &copy; Itallo Gabriel (https://github.com/7wf)
