
# Dona 🍩

CLI to manage your dotfiles 

![](./assets/demo.gif)

## Features 🎉

* 🔍 Search dotfiles with a command and with pagination 
* 🏆 Track your dotfiles with CLI api connected to Git `dona git` 
* ✨ Save third party dotfiles in a specific folder 
* 🔐 Manage pins (key/value items to save repositories or resources) 

## Usage 🚀

### Search dotfiles
Search across GitHub repositories for dotfiles:
```bash
dona search "arch linux aesthetic"
dona search neovim --sort stars
dona search fedora --page 2
```

![dona search demo](https://i.postimg.cc/D0TK5Dmk/image.png)

### Save & manage third party dotfiles
```bash
dona save user/dotfiles    # Save dotfiles from a GitHub repo
dona list                  # List saved dotfiles
dona del user/dotfiles     # Delete a saved dotfile
dona clone user/dotfiles   # Clone a dotfile repository
```

### Pins 📌
Save repositories or resources as key/value items, like Pinterest:
```bash
dona pin add user/dotfiles -t neovim    # Pin a repo with a tag
dona pin add user/dotfiles -t fedora    # Pin with different tag
dona pin del user/dotfiles              # Delete a pin
dona pin del fedora                     # Delete all pins with tag
dona list                               # List your pins
```

### Your own dotfiles
```bash
dona init                  # Initialize your dotfiles
dona git status            # Git status of your dotfiles
dona git push              # Push your dotfiles
```

### Extras
```bash
dona please                # 🍩
dona version               # Print version
```

## Installation

### With Go 🎩

```bash
go install github.com/atticus64/dona@latest
```

### With Curl 👓

```bash
curl https://raw.githubusercontent.com/Atticus64/dona/main/install.sh | bash
```

## Build project

```
go build
```

