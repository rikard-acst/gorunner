# GoRunner

## A simple executable runner to circumvent an annoying small problem

To run git efficiently in WSL with a repository on the C drive, you have to run the following
executable as `git`

```bash
#! /bin/sh
case "$(pwd -P)" in
    /mnt/?/*) exec /mnt/c/Program\ Files/Git/cmd/git.exe "$@" ;;
    *) exec /usr/bin/git "$@" ;;
esac
```

But when using named panes in tmux with `set pane-border-status top`, it overwrites your named pane with the
Windows executable name.  This stops it from spinning up a new shell in cmd.exe to run these commands
and allows you to keep your pane name.

The resulting git command becomes

```bash
#! /bin/sh
case "$(pwd -P)" in
    /mnt/?/*) exec gorunner /mnt/c/Program\ Files/Git/cmd/git.exe "$@" ;;
    *) exec /usr/bin/git "$@" ;;
esac
```
