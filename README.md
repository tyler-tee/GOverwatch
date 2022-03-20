# GOverwatch
GOverwatch is an early Go port of Overwatch. Still intended to leverage Masscan's speed and Nmap's versatility, but with Go's portability and minimal overhead.

Rapidly discover any open ports using Masscan, then automatically feed only those ports to Nmap for further interrogation.

## Prerequisites
- Nmap
- Masscan

## Usage

### General Use
> go run main.go

### Headless
Headless mode will allow you to run GOverwatch unattended. Set it up with a cronjob/scheduled task and you can keep an eye on your external footprint automatically.

## License
[GNUv3](https://www.gnu.org/licenses/gpl-3.0.en.html)
