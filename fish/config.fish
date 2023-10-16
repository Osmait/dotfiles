




set -x GOPATH $HOME/go
set -x GOBIN $GOPATH/bin
set -x GOROOT /usr/local/go

set -x PATH $PATH $GOBIN $GOROOT/bin

set -x DENO_INSTALL "/home/osmait/.deno"
set -x PATH $DENO_INSTALL/bin $PATH

set -Ux FZF_DEFAULT_OPTS "\
--reverse \
--border rounded \
--no-info \
--pointer=' ' \
--marker=' ' \
--ansi \
--color='16,bg+:-1,gutter:-1,prompt:5,pointer:5,marker:6,border:4,label:4,header:italic'"

set -Ux FZF_TMUX_OPTS "-p 55%,60%"

set -Ux FZF_CTRL_R_OPTS "--border-label=' history ' \
--prompt='  '"

# fnm
set -gx PATH "/home/osmait/.local/share/fnm" $PATH
eval (fnm env)

set fish_greeting "󰣇 Osmait 󱘗  󰛦  
  _____                       _        _____                 
 / ___ \                     (_) _    (____ \                
| |   | |  ___  ____    ____  _ | |_   _   \ \   ____  _   _ 
| |   | | /___)|    \  / _  || ||  _) | |   | | / _  )| | | |
| |___| ||___ || | | |( ( | || || |__ | |__/ / ( (/ /  \ V / 
 \_____/ (___/ |_|_|_| \_||_||_| \___)|_____/   \____)  \_/  

"

if status is-interactive

    set -U fish_key_bindings fish_vi_key_bindings
    function fish_prompt
        fish_mode_prompt
    end
    # Commands to run in interactive sessions can go here
    starship init fish | source
    fish_add_path ~/.tmux/plugins/t-smart-tmux-session-manager/bin
    zoxide init fish | source
    # ls
    alias l 'ls -lh'
    alias ll 'ls -lah'
    alias la 'ls -A'
    alias lm 'ls -m'
    alias lr 'ls -R'
    alias lg 'ls -l --group-directories-first'

    # git
    alias gcl 'git clone --depth 1'
    alias gi 'git init'
    alias ga 'git add'
    alias gc 'git commit -m'
    alias gp 'git push origin master'
    alias lvim '/home/osmait/.local/bin/lvim'
    alias poetry '~/.local/share/pypoetry/venv/bin/poetry'



    set -Ux PYENV_ROOT $HOME/.pyenv
    fish_add_path $PYENV_ROOT/bin


    # export GOPATH=$HOME/go
    # export GOBIN=$GOPATH/bin
    # export GOROOT=/usr/local/go

    # export PATH=$PATH:$GOBIN:$GOROOT/bin



    # export DENO_INSTALL="/home/osmait/.deno"
    # export PATH="$DENO_INSTALL/bin:$PATH"

    # export NVM_DIR="$HOME/.nvm"
    # [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
    # [ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion


end
