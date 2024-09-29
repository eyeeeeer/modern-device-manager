# Modern Device Manager

#### This repository contains a source code of the modern device manager app. This app written in golang + typescript programming languages via wails + react framework stack.

#### This app is a recreation of the original Windows Device Manager app. This app target: bring a up-to-date fluent design style to this app, simplify this app for new users.

#### This app is fully open-source and free for all. Everyone welcome to contribute and fix bugs. If you like this app please star repository and if you want you can donate me via PayPal.

#### This app is still in development, so please be patient and report any bugs you find.

# How to run this app?
**Currently app still in development which is why it is unavailable for download as compiled builds. If you want to run this app on your PC you need build it first. Below you can find a guide how you can compile this app on your PC and run it.**

## Prerequisites:
- Go 1.20 or higher
- Node.js v15.0.0 or higher (recommended: v20.17.0 LTS (Latest LTS release))
- npm v7.0.0 or higher (recommended: v10.8.2 (In bundle with latest LTS node release))

1. Install Wails framework
```cmd
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
```
2. Clone this repository

    You can do this via terminal by the git clone command or via Git manager (for example Github Desktop)
```cmd
    git clone https://github.com/eyeeeeer/modern-device-manager.git
```
3. In project root folder run `wails build` command
```cmd
    cd modern-device-manager
    wails build
```

**After this actions you can find the output build file in modern-device-manager/build/bin directory.**

