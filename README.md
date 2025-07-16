# 🍺 Homebrew Tap for `timestamp-converter`

This is the [Homebrew](https://brew.sh/) tap for [`timestamp-converter`](https://github.com/mikemols/timestamp-converter), a small and handy CLI tool that converts UNIX timestamps into human-readable date formats.

---

## 🚀 Install

```bash
brew tap mikemols/tap
brew install timestamp-converter
```

---

## 🛠 Usage

```bash
timestamp-converter <timestamp> [options]
```

### 🔁 Examples

Convert a UNIX timestamp:

```bash
timestamp-converter 1752574424
```

Copy the result to your clipboard:

```bash
timestamp-converter 1752574424 --copy
```

or using a short flag:

```bash
timestamp-converter 1752574424 -c
```

---

## 🧪 Sample Output

```bash
$ timestamp-converter 1752574424
Tuesday, July 15, 2025 14:53:44 UTC
```

---

## 🔄 Updating

```bash
brew update
brew upgrade timestamp-converter
```

---

## 🧠 About

This tool was built by [@mikemols](https://github.com/mikemols) for day-to-day development needs.

The tap repository is hosted here:  
👉 https://github.com/mikemols/homebrew-tap

Main CLI tool source code:  
👉 https://github.com/mikemols/timestamp-converter

---

## 📬 Feedback / Contributions

If you have ideas or bug reports, feel free to open an issue or PR on the [timestamp-converter](https://github.com/mikemols/timestamp-converter) repo.
```

---
