HTML link parser
================

example1
--------
(html -> urls and texts)
$ go run examples/ex1/main.go -h
Usage of /tmp/go-build130861059/b001/exe/main:
  -html string
        html file to parse (default "index.html")

$ go run examples/ex1/main.go -html examples/ex1/index.html
[{Href:/other-page Text:A link to another page Some span} {Href:/second-page Text:A link to a second page}]

On a bigger html:
$ go run examples/ex1/main.go -html ~/dox/git/own/sites/m4t3sz.gitlab.io/writeup/thm/overpass/index.html
[{Href://m4t3sz.gitlab.io/bsc/ Text:m4t3sz.gitlab.io/bsc} {Href://m4t3sz.gitlab.io/ Text:home} {Href://twitter.com/szilak44/ Text:twitter} {Href://github.com/matesz44/ Text:github} {Href://github.com/matesz44/dotfiles/ Text:dotfiles} {Href://github.com/matesz44/scripts/ Text:scripts} {Href:/bsc Text:about} {Href://m4t3sz.gitlab.io/bsc/showcase/ Text:showcase/} {Href://m4t3sz.gitlab.io/bsc/writeup/ Text:>writeup/} {Href://m4t3sz.gitlab.io/bsc/writeup/ctf/ Text:ctf/} {Href://m4t3sz.gitlab.io/bsc/writeup/htb/ Text:htb/} {Href://m4t3sz.gitlab.io/bsc/writeup/thm/ Text:>thm/} {Href://m4t3sz.gitlab.io/bsc/writeup/thm/overpass/ Text:>overpass/} {Href:https://tryhackme.com/room/overpass Text:room link} {Href:https://tryhackme.com/p/NinjaJc01 Text:NinjaJc01} {Href:http://10.10.248.28/downloads/src/overpass.go Text:go source} {Href:https://gchq.github.io/CyberChef/ Text:CyberChef} {Href:https://github.com/carlospolop/privilege-escalation-awesome-scripts-suite/blob/master/linPEAS/linpeas.sh Text:linpeas.sh} {Href:https://www.buymeacoffee.com/m4t35z Text:} {Href:https://github.com/karlb/smu/ Text:smu} {Href:https://git.suckless.org/sites/ Text:suckless style}]

example2
--------
(webpage -> GET -> urls(only on provided domain) -> recursively(depth can be set))
(output can be XML (-x))
$ go run examples/ex2/main.go -h                          
Usage of /tmp/go-build579754148/b001/exe/main:
  -d int
        depth you want to follow links (default 3)
  -u string
        url you want to crawl (default "https://m4t3sz.gitlab.io/bsc/")
  -x    use this to output xml in a sitemap format

$ go run examples/ex2/main.go -u https://m4t3sz.gitlab.io/bsc/ -d 2   
https://m4t3sz.gitlab.io/bsc/showcase/
https://m4t3sz.gitlab.io/bsc/writeup/
https://m4t3sz.gitlab.io/bsc/showcase/revshellgen/
https://m4t3sz.gitlab.io/bsc/writeup/ctf/
https://m4t3sz.gitlab.io/bsc/writeup/htb/
https://m4t3sz.gitlab.io/bsc/writeup/thm/
https://m4t3sz.gitlab.io/bsc/

output xml in a sitemap format(-x):
$ go run examples/ex2/main.go -u https://m4t3sz.gitlab.io/bsc/ -d 2 -x
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="https://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://m4t3sz.gitlab.io/bsc/</loc>
  </url>
  <url>
    <loc>https://m4t3sz.gitlab.io/bsc/showcase/</loc>
  </url>
  <url>
    <loc>https://m4t3sz.gitlab.io/bsc/writeup/</loc>
  </url>
  <url>
    <loc>https://m4t3sz.gitlab.io/bsc/showcase/revshellgen/</loc>
  </url>
  <url>
    <loc>https://m4t3sz.gitlab.io/bsc/writeup/ctf/</loc>
  </url>
  <url>
    <loc>https://m4t3sz.gitlab.io/bsc/writeup/htb/</loc>
  </url>
  <url>
    <loc>https://m4t3sz.gitlab.io/bsc/writeup/thm/</loc>
  </url>
</urlset>