# pomo - Terminal Pomodoro Tool
* A sound will play when your timer is done, if you want this to work please make sure you have MPV installed
## Usage
* `pomo` - Show the current status of your timer
* `pomo new` - Start a new timer (default interval is 25 minutes)
    * `pomo new hour` - Start a new timer with the amount of time left till the top of the hour
        * Such as if it's 13:24 it will start a 36 minute timer
    * `pomo new <custom second amount>` - Will start a timer with a custom amount defined as seconds
* `pomo stop` - Stops the current running timer, even if it's not done
* `pomo pause` - Pauses the current timer
* `pomo resume` - Resumes the current timer from where it last left off
* `pomo -h` - List help menu
*  `eval "$(pomo completion zsh)"` - Put this in your zshrc to get completion for this tool
    * You can change zsh to bash or fish if you need to
## Install
`go install github.com/skykosiner/pomo@latest`
