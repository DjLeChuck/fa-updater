# FA Updater

FA updater is a simple CLI tool that will allow you to keep your FA files up to date. It will analyze the versions you
have and check if more recent ones exist in order to download them easily.

## Installation

Go to the [latest release](https://github.com/DjLeChuck/fa-updater/releases/latest) page and download the version for
your computer:

* macOS:  `xxx_darwin_amd64.tar.gz`
* Linux: `xxx_linux_amd64.tar.gz`
* Windows: `xxx_windows_amd64.zip`

Then, extract the archive where you want. You should see a new `fa-updater` file which is the tool executable.

## Usage

Open up your favorite terminal application and go to the directory containing the `fa-updater` executable.

Then, you can execute it without argument to see if it's working:

```bash
$ ./fa-updater 
FA updater is a simple tool that will allow you to keep your FA files up to date. It will analyze the versions you have and check if more recent ones exist in order to download them easily.

First, be sure to define the directory which contains your assets (fa-updater setDirectory), then launch the update process (fa-updater updateAssets).

Usage:
  fa-updater [command]

Available Commands:
  help         Help about any command
  setDirectory Define the directory which contains your files
  updateAssets Launch the update process
  updateTokens Update all tokens

Flags:
      --config string   config file (default is $HOME/.fa-updater.yaml)
  -h, --help            help for fa-updater
  -v, --version         version for fa-updater

Use "fa-updater [command] --help" for more information about a command.
```

### Available Commands

#### help

By default, it shows the same thing as running without argument. However, if you add the name of another command, then
you will see the help of this command (eg. `./fa-updater help setDirectory`).

#### setDirectory `[type]` `[path]`

Define the directory which contains your files.

`[type]` can be one of the following:

* `dungeondraft`: will be used by the `updateAssets` command to get newer versions if exists,
* `tokens`: will be used by the `updateTokens` command to get newer versions if exists.

`[path]` should target an existing directory.

By default, you cannot override a directory already configured, an error message will be shown telling you to add the
`--force` flag (short version `-f`) to do it:

```bash
$ ./fa-updater setDirectory dungeondraft /path/to/da_assets
2023-02-04T10:23:15+01:00 | INFO  | The directory is already configured: /path/to/dungeondraft/ForgottenAdventures/
2023-02-04T10:23:15+01:00 | INFO  | Please, use the flag --force flag if you want to override the configuration
$ ./fa-updater setDirectory dungeondraft /path/to/da_assets --force
```

#### updateAssets

Launch the update process to compare the latest available packs with the ones in your assets' directory. It will also
take care of the pre-generated thumbnails.

First, you will need to get the Patreon page content, then give your Patreon session's cookie in order to be able to
download the files.

<details>
<summary>View usage details</summary>

```bash
$ ./fa-updater updateAssets
2023-02-04T10:45:45+01:00 | INFO  | Go on https://www.patreon.com/posts/56375276 with your browser. Display the source of the page (CTRL+U or ⌘+U) and copy it in the clipboard (CTRL+A and CTRL+C or ⌘+A and ⌘+C), then go back here and press ENTER.

# Go on the page, copy the source code, then press ENTER

2023-02-04T10:45:49+01:00 | INFO  | 35 packs found. Comparing to your assets directory...
2023-02-04T10:45:49+01:00 | INFO  | There are 2 packs to download.
2023-02-04T10:45:49+01:00 | INFO  | Please, look at the cookies on the Patreon page and copy the value of the one named "session_id" in the clipboard (CTRL+C or ⌘+C), then press ENTER. It should looks like a random string: LC2A4j7WAJe4cjR5Oeicycf4YmlEfQsNB_yqwYiWuh8

# Go on the page, copy the cookie value, then press ENTER

2023-02-04T10:46:09+01:00 | INFO  | Downloading FA_Assets_N_v3.02.dungeondraft_pack...
2023-02-04T10:46:09+01:00 | INFO  | 200 OK
2023-02-04T10:46:10+01:00 | INFO  | Download saved to /path/to/dungeondraft/ForgottenAdventures/FA_Assets_N_v3.02.dungeondraft_pack
2023-02-04T10:46:09+01:00 | INFO  | Downloading FA_Assets_O_v3.01.dungeondraft_pack...
2023-02-04T10:46:09+01:00 | INFO  | 200 OK
2023-02-04T10:46:10+01:00 | INFO  | Download saved to /path/to/dungeondraft/ForgottenAdventures/FA_Assets_O_v3.01.dungeondraft_pack
2023-02-04T10:46:10+01:00 | INFO  | Checking thumbnails...
2023-02-04T10:46:10+01:00 | INFO  | Thumbnails processing done.
```

</details>

##### Why such a complicated process?

Concerning the first step and sadly for us, Patreon is using Cloudflare to protect access of their website. Because of
this, it's impossible to automatically crawl the page content, it must be a human action.

As it concerns the cookie, it's necessary to be allowed to download the Patreon files. This cookie is used to
identificate yourself as a valid user which has access to them.

#### updateTokens

Launch the update process to compare the latest available tokens with the ones in your tokens assets directory.

You will need to give your Patreon session's cookie in order to be able to download the files.


<details>
<summary>View usage details</summary>

```bash
2023-02-04T11:05:29+01:00 | INFO  | Please, look at the cookies on the Patreon page and copy the value of the one named "session_id" in the clipboard (CTRL+C or ⌘+C), then press ENTER. It should looks like a random string: LC2A4j7WAJe4cjR5Oeicycf4YmlEfQsNB_yqwYiWuh8

# Go on the page, copy the cookie value, then press ENTER

2023-02-04T11:05:36+01:00 | INFO  | Processing page 1...
2023-02-04T11:05:36+01:00 | INFO  | Downloading Creature Tokens – Pack 42...
2023-02-04T11:05:37+01:00 | INFO  | 200 OK
2023-02-04T11:05:40+01:00 | INFO  | Download saved to /tmp/Creature Tokens – Pack 42
2023-02-04T11:05:40+01:00 | INFO  | Unzipping Creature Tokens – Pack 42...
2023-02-04T11:05:42+01:00 | INFO  | Processing page 2...
2023-02-04T11:05:42+01:00 | INFO  | Processing page 3...
2023-02-04T11:05:43+01:00 | INFO  | Processing page 4...
2023-02-04T11:05:44+01:00 | INFO  | Processing page 5...
2023-02-04T11:05:45+01:00 | INFO  | Processing page 6...
2023-02-04T11:05:46+01:00 | INFO  | Processing page 7...
2023-02-04T11:05:46+01:00 | INFO  | Processing page 8...
2023-02-04T11:05:47+01:00 | INFO  | Processing page 9...
```

</details>

##### Why such a complicated process?

This time, no needs to ask for the code source of the page because links are availables through a RSS feed on the
Forgotten Adventures website itself. However, the cookie is always necessary to be allowed to download the Patreon
files.

## How to contribute?

### Bug report

Please open an issue: https://github.com/DjLeChuck/fa-updater/issues

You can also provide a fix via a pull request: https://github.com/DjLeChuck/fa-updater/pulls
