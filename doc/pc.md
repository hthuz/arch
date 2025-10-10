# Arch Linux Issue Note


## cisco dependency
```
pacman -S libsoup webkit2gtk
```

## Missing libxml2.so

```
pacman -S libxml2-legacy
```


## change to zsh

```
chsh -s $(which zsh)
```

## input method

```
pacman -S fcitx5-im fcitx5-chinese-addons
```
and add 
```
GTK_IM_MODULE=fcitx
QT_IM_MODULE=fcitx
XMODIFIERS=@im=fcitx
```
in `etc/environment`, then edit in KDE settings (virtual keyboard and input method)


## wps optional support
```

yay -S ttf-wps-fonts ttf-ms-fonts wps-office-fonts wps-office-mui-zh-cn wps-office-mime-cn
```

fix display issue for bold font for wps under version 11:
```
yay -S freetype2-wps
```


## hidpi for SDDM login 

Put `hidpi.conf` in `/etc/sddm.conf.d` with the following contents:
```

[General]
GreeterEnvironment=QT_SCREEN_SCALE_FACTORS=2,QT_FONT_DPI=192

[Wayland]
EnableHiDPI=true

[X11]
EnableHiDPI=true
ServerArguments=-nolisten tcp -dpi 192
```




