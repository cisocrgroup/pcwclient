[![builds.sr.ht status](https://builds.sr.ht/~flobar/builds.sr.ht.svg)](https://builds.sr.ht/~flobar/builds.sr.ht?)
# pcwclient

Command line client for [pocoweb](http://github.com/cisocrgroup/pocoweb)

## Examples
Set pocoweb's URL: `export POCOWEB_URL=https://pocoweb.cis.lmu.de`

Authentificate: `export POCOWEB_AUTH=$(pcwclient login -F '{{.Auth}}' user email)`