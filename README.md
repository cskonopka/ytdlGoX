
# ytdlGoX
> This project is a small collection of ytdl examples that explore various ways of collecting video content. I decided to use go-sh instead of ytdl by rylio since I will have more flexibility if I expand the system. 

## What was used?
* go-sh [https://github.com/codeskyblue/go-sh]
* youtube-dl [https://github.com/rg3/youtube-dl]
* adapted example https://golangvedu.wordpress.com]

## Installation

* go-sh

```
go get github.com/codeskyblue/go-sh
```

* youtube-dl

To install it right away for all UNIX users (Linux, OS X, etc.), type:

```
sudo curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl
sudo chmod a+rx /usr/local/bin/youtube-dl
```

If you do not have curl, you can alternatively use a recent wget:

```
sudo wget https://yt-dl.org/downloads/latest/youtube-dl -O /usr/local/bin/youtube-dl
sudo chmod a+rx /usr/local/bin/youtube-dl
```

Windows users can download an .exe file and place it in any location on their PATH except for %SYSTEMROOT%\System32 (e.g. do not put in C:\Windows\System32).

You can also use pip:

```
sudo -H pip install --upgrade youtube-dl
```

This command will update youtube-dl if you have already installed it. See the pypi page for more information.

OS X users can install youtube-dl with Homebrew:

```
brew install youtube-dl
```

Or with MacPorts:

```
sudo port install youtube-dl
```

Alternatively, refer to the developer instructions for how to check out and work with the git repository. For further options, including PGP signatures, see the youtube-dl Download Page.




Install dependencies:

```
$ 
```

## Run development

Start dev server:

```
$ make serve
```
