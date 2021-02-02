HTML link parser
================

example1(html -> urls and texts):
---------------------------------
$ go run examples/ex1/main.go -h
Usage of /tmp/go-build130861059/b001/exe/main:
  -html string
        html file to parse (default "index.html")

$ go run examples/ex1/main.go -html examples/ex1/index.html
[{Href:/other-page Text:A link to another page Some span} {Href:/second-page Text:A link to a second page}]

On a bigger html:
$ go run examples/ex1/main.go -html ~/dox/git/own/sites/m4t3sz.gitlab.io/writeup/thm/overpass/index.html
[{Href://m4t3sz.gitlab.io/bsc/ Text:m4t3sz.gitlab.io/bsc} {Href://m4t3sz.gitlab.io/ Text:home} {Href://twitter.com/szilak44/ Text:twitter} {Href://github.com/matesz44/ Text:github} {Href://github.com/matesz44/dotfiles/ Text:dotfiles} {Href://github.com/matesz44/scripts/ Text:scripts} {Href:/bsc Text:about} {Href://m4t3sz.gitlab.io/bsc/showcase/ Text:showcase/} {Href://m4t3sz.gitlab.io/bsc/writeup/ Text:>writeup/} {Href://m4t3sz.gitlab.io/bsc/writeup/ctf/ Text:ctf/} {Href://m4t3sz.gitlab.io/bsc/writeup/htb/ Text:htb/} {Href://m4t3sz.gitlab.io/bsc/writeup/thm/ Text:>thm/} {Href://m4t3sz.gitlab.io/bsc/writeup/thm/overpass/ Text:>overpass/} {Href:https://tryhackme.com/room/overpass Text:room link} {Href:https://tryhackme.com/p/NinjaJc01 Text:NinjaJc01} {Href:http://10.10.248.28/downloads/src/overpass.go Text:go source} {Href:https://gchq.github.io/CyberChef/ Text:CyberChef} {Href:https://github.com/carlospolop/privilege-escalation-awesome-scripts-suite/blob/master/linPEAS/linpeas.sh Text:linpeas.sh} {Href:https://www.buymeacoffee.com/m4t35z Text:} {Href:https://github.com/karlb/smu/ Text:smu} {Href:https://git.suckless.org/sites/ Text:suckless style}]

example2(webpage -> GET -> urls):
---------------------------------
$ go run examples/ex2/main.go -h
Usage of /tmp/go-build063838438/b001/exe/main:
  -u string
        url you want to crawl (default "https://m4t3sz.gitlab.io/bsc/")

$ go run examples/ex2/main.go -u https://m4t3sz.gitlab.io/bsc/writeup/thm/overpass
https://m4t3sz.gitlab.io/bsc/
https://m4t3sz.gitlab.io/
https://twitter.com/szilak44/
https://github.com/matesz44/
https://github.com/matesz44/dotfiles/
https://github.com/matesz44/scripts/
https://m4t3sz.gitlab.io/bsc
https://m4t3sz.gitlab.io/bsc/showcase/
https://m4t3sz.gitlab.io/bsc/writeup/
https://m4t3sz.gitlab.io/bsc/writeup/ctf/
https://m4t3sz.gitlab.io/bsc/writeup/htb/
https://m4t3sz.gitlab.io/bsc/writeup/thm/
https://m4t3sz.gitlab.io/bsc/writeup/thm/overpass/
https://tryhackme.com/room/overpass
https://tryhackme.com/p/NinjaJc01
http://10.10.248.28/downloads/src/overpass.go
https://gchq.github.io/CyberChef/
https://github.com/carlospolop/privilege-escalation-awesome-scripts-suite/blob/master/linPEAS/linpeas.sh
https://www.buymeacoffee.com/m4t35z
https://github.com/karlb/smu/
https://git.suckless.org/sites/