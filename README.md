## Installation

> git clone https://github.com/PatrikOlin/butler-burton <br>
> cd butler-burton <br>
> make install

or download binary and run it

## Configuration

Edit config-file in `$HOME/.config/butlerburton`

Example-config:

```ỳaml
name: "Butler Burton"
color: "#46D9FF"
webhook_url: "<Teams webhook url>"
notifications: true
vab_msg: "Jag vabbar idag, försök hålla skutan flytande så är jag tillbaks imorgon"
report:
    path: "/home/burton/.butlerburton/"
    update: true
    checkin_col: "C"
    checkout_col: "D"
    lunch_col: "F"
    bl_lunch_col: "I"
    overtime_col: "R"
    vab_col: "L"
    afk_col: "G"
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

### Get report
```
docker ps (get containerId)
docker cp "containerId":/root/.butlerburton/"reportName.xlsx" .
```
