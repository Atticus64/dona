#!/bin/bash

version=0.0.1

bin="dona"
blue="\e[0;34m"
green="\e[0;32m"
install_dir="$HOME/.dona/bin"

case $(uname -sm) in
	"Darwin x86_64") target="x86_64_news_macos"
		install_dir="/usr/local/bin" ;;
	"Darwin arm64") target="aarch64_news_macos" ;;
	*) target="x86_64_linux_dona" ;;
esac

if [ ! -d "$install_dir" ]; then
        mkdir -p "$install_dir"
fi

printf "\n\nFetching binary dona\n"

curl --fail --location --progress-bar --output dona https://github.com/Atticus64/dona/releases/download/v$version/$target

chmod +x $bin

mv $bin $install_dir

if command -v news >/dev/null; then
        printf "\e[0;33mNews\e[0m updated \e[0;32msuccesfully\e[0m\n"
        printf "Run 'dona --help' to get started\n"
else
  printf "\e[0;33mNews\e[0m installed \e[0;32msuccesfully\e[0m\n"
  case $SHELL in
        /bin/zsh)
                shell_profile="~/.zshrc"
                printf "\e[0;33mAdd\e[m manually the path of news to you $shell_profile\n\n"
                printf "export PATH=\"\$HOME/.dona/bin:\$PATH\" \n\n"
                printf "Try \e[0;32mrunning\e[m echo -n 'export PATH=\$HOME/.dona/bin:\$PATH' >> $shell_profile \n"
      printf "And reload terminal with source $shell_profile \n\n"
   ;;
        /usr/bin/fish)
                shell_profile="~/.config/fish/config.fish"
                printf "\e[0;33mAdd\e[m manually the path of news to you $shell_profile\n\n"
                printf "set PATH \$HOME/.dona/bin \$PATH \n\n"
                printf "Try \e[0;32mrunning\e[m echo \"set PATH \$HOME/.dona/bin \\\$PATH\" >> $shell_profile \n"
      printf "And reload terminal with source $shell_profile \n\n"
                ;;
        *)
                shell_profile="~/.bashrc"
                printf "\e[0;33mAdd\e[m manually the path of news to you $shell_profile\n\n"
                printf "export PATH=\"\$HOME/.dona/bin:\$PATH\" \n\n"
      printf "Try \e[0;32mrunning\e[m echo -n 'export PATH=\$HOME/.dona/bin:\$PATH' >> $shell_profile \n"
      printf "And reload terminal with source $shell_profile \n\n"
    ;;
  esac
fi
