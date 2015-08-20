_This is alpha-stage software. All aspects are subject to change._

# Weaver

A standalone Markdown/RMarkdown processor.  It is used to make quick static and standalone laboratory notebook pages.

# Dependencies

`weaver` is assuming that [R][1] (>=3.1.0) is already installed on your system and accessible to the `weaver` user.

# Developer Notes

The `assets/assets.go` file embeds the HTML templates and R scripts used by `weaver`.  It allows for a standalone executable.  This is how to generate the `assets.go` file via [`go-bindata`][2]:

    go get -v -u github.com/jteeuwen/go-bindata/...
    export PATH=$GOPATH/bin:$PATH
    go generate -x ./assets
    less assets/assets.go


[1]:  https://www.r-project.org
[2]:  https://github.com/jteeuwen/go-bindata
