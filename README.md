# VD Wrapper

## What?
VD Wrapper is a small wrapper for [Virtual Desktop](https://vrdesktop.net).

It has one simple job: to launch Virtual Desktop Streamer app and automatically populate preset client usernames.

## Why?
I don't have a beefy gaming PC I could dedicate for PCVR, so I use [Shadow PC](https://shadow.tech) to play VR. With some good network setup the experience in most games is great.

However, there are caveats. Shadow PC is basically a virtual machine in a datacenter. Due to some virtualization quirks (maybe VM instance migrating between hypervisors?) the unique hardware ID of my Shadow instance often changes.

Virtual Desktop relies on this hardware ID to authenticate requests to their cloud and encrypt entered Oculus/Pico/Viveport usernames on backend. Plain text client usernames are not stored locally at all.

When hardware ID changes (almost every reboot of my Shadow instance), saved encrypted identifiers can't be used anymore, so Virtual Desktop 'forgets' the usernames I entered. This means that every time I boot up my Shadow to play some VR, I have to enter my username again and click Save. Ugh.

I asked VD devs if there is any way to store usernames locally to retain them between hardware ID changes. VD devs told me this is "never going to happen" as it would "open a backdoor". Bruh.

Hence the wrapper.

## How?
The wrapper is simple and stupid, here is what it does:

1. Kills VD if it's already running
2. Removes encrypted identifiers from the VD config (for consistency during next steps)
3. Launches VD and waits for the settings window to appear
4. Simulates keypresses to navigate through the settings window, fill out the username(s) and hit Save
5. Closes the settings window (VD stays in the system tray)

## So how do I use it?
The wrapper is configured with command line flags.

1. Place the vd-wrapper executable wherever you want
2. Create a shortcut for the executable, and add your space-separated usernames to the arguments (see example below)
3. Use this shortcut to launch VD instead of the original shortcut
4. Add this shortcut to the Windows Startup folder. Don't forget to **turn off** `Start with Windows` in VD settings

Example arguments for adding 2 Oculus usernames, 1 Pico username and 3 Viveport IDs:

```vd-wrapper.exe --oculus Username Username2 --pico Username3 --vive ID1 ID2 ID3```

At least one username must be provided. Each platform can have up to 4 usernames.

## It doesn't work for me
Hm, weird, it works on my machine.

I made this wrapper for myself, so I am not planning to support any other environments.

If you have a fix, feel free to submit a pull request. If your pull request doesn't break VD Wrapper on my machine, I might accept it.

## Your code sucks
Your mom sucks.