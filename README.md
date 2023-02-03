## Installation

> git clone https://github.com/PatrikOlin/butler-burton <br>
> cd butler-burton <br>
> make install

or download the binary and run it

## Configuration

Edit config-file in `$HOME/.config/butlerburton`

Example-config:

```yaml
name: "Butler Burton"
color: "#46D9FF"
webhook_url: "<Teams webhook url>"
notifications: true
vab_msg: "Jag vabbar idag, försök hålla skutan flytande så är jag tillbaks imorgon"
weekend_msg: "Trevlig helg!"
time_sheet:
    employee_id: "0000",
    path: "/home/burton/.butlerburton/"
    update: true
```

## Development

### Build docker image

```sh
docker build -t "imageName" .
```

### Start docker container

```sh
docker run -it "imageName" sh
```

### Get time sheet

```sh
docker ps (get containerId)
docker cp "containerId":/root/.butlerburton/"reportName.xlsx" .
```

### Generate man page from butlerburton.md

```sh
pandoc butlerburton.md -s -t man -o butler-burton.1
gzip butler-burton.1
```
