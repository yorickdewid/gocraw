[![license](https://img.shields.io/badge/license-BSD%20-blue.svg)](https://raw.githubusercontent.com/yorickdewid/gocraw/master/LICENSE)

## gocraw

Webcrawler written in go. Retrieves webcontent and saves it to a file.

### Examples

```
gocraw https://github.com
```

Or load the config file

```
gocraw -file gocraw.conf
```

## Configuration

List the domains per line

```
# Some comment

https://github.com/
# http://www.facebook.com # Exclude this URL
https://www.google.com
http://www.golang.org/
```
