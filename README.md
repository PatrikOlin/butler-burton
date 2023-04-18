## Installation

> git clone https://github.com/PatrikOlin/butler-burton <br>
> cd butler-burton <br>
> make install

or download the binary and run it

<details>
<summary>Linux</summary>
You can do this.
</details>


<details>
<summary>Windows</summary>

## Install Butler-Burton on Windows

##### Download Ubuntu for Windows in Microsoft Store or use WSL in either Powershell or Commandline (We recommend you use the Ubuntu one in the store)

If you are using the latter 

``` 
wsl --install Ubuntu-{version} example 22.04 
```

If you want everything to run on WSL2 as soon as you install it, you can set it as the default version. 

``` 
wsl --set-default-version 2
```

Then run:

```
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
```

First of you should set your name and password for Ubuntu. But we are not done yet. Install both golang-go and make and update all packages 

```
sudo apt update && sudo apt install golang-go make 
```

After this step make sure you are on BL-VPN or in any of the offices 

### Installation 

`git clone https://github.com/PatrikOlin/butler-burton`

`cd butler-burton`

`make install`

If no binary file is created when running make install, download the binary file from the release page into the Downloads folder and change the name of the file to bb
[Release Page](https://github.com/PatrikOlin/butler-burton/releases)

Move the file into the bin-folder like so: 
`sudo mv /mnt/c/Users/{yourWindowsUser}/Downloads/bb /usr/bin`

then try the `bb -h` command (it should work!)

 
#### Common Errors:

##### If you already have a timesheet on sharepoint and it doesnt get the name straight or cant download it:
`bb ts g`

`bb ts s "TIRP_Firstname_Lastname_Apr-23.xlsx"`

`bb ts d`
 
##### If the DNS is acting up use this:

Clone and run the shell: 

> sudo sh ./run.sh 

Restart WSL:

> wsl --shutdown 

Then: 

`curl --resolve 'codeload.github.com:443:20.201.28.149' 'https://codeload.github.com/epomatti/wsl2-dns-fix-config/tar.gz/refs/tags/v1.0.1' -o wsl2-dnsfix.tar.gz`

`tar -xf wsl2-dnsfix.tar.gz > cd wsl2-dns-fix-config-1.0.1`

`sudo sh ./run.sh`

###### What it does
The [`run.sh`](./run.sh) script will perform these tasks: 
1. Delete the following files: `/etc/wsl.conf` and `/etc/resolv.conf`
2. Create the new ".conf" files (pre-created in the dist folder) setting Google DNS for name resolution and preventing WSL from overriding it: 
```
sh # /etc/wsl.conf [network] generateResolvConf = false 
```
```
sh # /etc/resolv.conf nameserver 8.8.8.8 
```
3. Make `/etc/resolv.conf` immutable
</details>


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
