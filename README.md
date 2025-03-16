# Godor
Gondor is a tiny library that brings structured YAML configs to your Go applications.

## Motivation
I've seen lots of applications that use [flag](https://pkg.go.dev/flag) as it's main configuration tool. It's quite a flexible soultion, but not a readable one.
Structed configs are much superior. You can have multiple layers of configs applied one by one, creating unlimited number for configurations while beeing readable
and easy to edit.

## How to use
All you have to do is pass a base layer plus some number of additional layers to `gondor.Parse` and thats it. Gondor will take base layer as an origin and apply
changes from each layer one by one in the same order as they were provided. See `exmaples` folder for code examples.

## Why YAML
[YAML is a strange format](https://ruudvanasseldonk.com/2023/01/11/the-yaml-document-from-hell), nevertheless if you apply some restriction to it, it will be quite
usefull and, dare I say, enjoyable. That said, Gondor have some limitation reagarding YAML it's parsing: no anchors and no special values (inf and such). Just plain old
maps, scalars and arrays. I don't think you will ever need something else for configs, but if you do feel free to make a PR.

