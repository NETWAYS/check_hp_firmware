# Release Workflow

Specify the release version.

```
VERSION=1.0.0
```

## Issues

Check issues at https://github.com/NETWAYS/check_hp_disk_firmware

## Authors

Update the [.mailmap](.mailmap) and [AUTHORS](AUTHORS) files:

    git checkout master
    git log --use-mailmap | grep ^Author: | cut -f2- -d' ' | sort | uniq > AUTHORS

## Update metadata

Update version in `version.go`:

    sed -i 's/const Version =.*/const Version = "'"$VERSION"'"/' version.go 

## Git Tag

Commit these changes to the "master" branch:

```
git commit -v -a -m "Release version $VERSION"
git push origin master
```

And tag it with a signed tag:

```
git tag -s -m "Version $VERSION" v$VERSION
```

Push the tag.

```
git push --tags
```
