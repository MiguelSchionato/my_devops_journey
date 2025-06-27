---
date: 2025-06-26
tags:
  - "#Linux"
  - "#Debian"
source:
---

# Topic: Linux- Debian
Dual boot on my main desktop
## 1. The Problem / The Challenge

The objective is to install a dual boot Debian distro on my main machine.
 As it's not intended to be used as the main machine, I only installed on a 50G partition, using the same swap partition and, of course, the same EFI boot partition.
###### Why Debian?
	I had almost no experience with this distro, and I wanted to fill this gap in order to become a more professional SysAdm. 
	The second reason is to have a more stable but still flexible distro to debug and take the rolling-release (Arch) out of the equation.  



## 2. Investigation & Solution Ideas

### 2.1  **Installation boot:**
- It was an easy process using the Bookworm ISO of the official website and the command `sudo dd` on Arch.
### 2.2  **Booting:**
###### 2.2.1 Partitioning:
1. Manually resized and add a new partition. Used 50G for the whole system, and the same swap partition as I already had.
###### 2.2.2 Drivers compatibility: 
1. That part was a little more difficult because the ISO didn't recognize my Ethernet card, and the right driver wasn't on the list.
2. **Idea:** Manually download the right firmware from the debian.org/firmware-realtek page. 
	1. Probably because my motherboard was too recent, and the Debian ISO didn't recognize it.
	2. **Result:** It didn't go as planned; the installation was successful, but the Kernel was too old to recognize the driver (r8125 or r8169).
		 1. **Solution:** I had to manually update the Kernel on Arch, as USB Tethering wasn't recognized either.
		 2. After manually installing every Kernel package and their dependencies, it finally worked!
		 3. I had to do a few more tweaks here and there, like enabling and starting "link ip" and "systemd-networkd". But it was a more straightforward process, and it ran with ease.
### 2.3 **Fixing broken install**
1.  **Idea:** Now with a stable Ethernet connection, I could run and fix the broken manual install of everything.
	1.  Run the `apt --fix-broken install` command.
	2. Didn't go as planned, as it was a fresh install and the package manager didn't have the correct mirrors in `/etc/apt/source.list`.
		1. **Solution:** Manually copy the mirror list instructed on the official website (`deb https://deb.debian.org`).
		2. After that, it worked perfectly, and I could fix the upgrades.

### 2.4 **Installing a GUI**
1. This part should be easier, but I haven't decided on a GUI/tiling manager yet. I'll update this once I do.


---

## 3. The Final Implemented Solution

**The solution I found was manually upgrading the Kernel, Ethernet driver, and their dependencies.**
There was probably a better way to do it, but I don't know it. If you know, feel free to contact me and enlighten me.

 ```bash
dpkg -i /var/cache/apt/archives/*.deb
 ```
It was definitely more hard work than difficult. But it improved my troubleshooting skills. 

---

## 4. Key Concepts & Reference Links

*   **Dual Boot**: Installing two operating systems on a single computer, allowing the user to choose which one to load at startup.
*   **EFI Boot Partition**: A partition on a data storage device that is used by computers adhering to the UEFI (Unified Extensible Firmware Interface) specification.
*   **Kernel Update**: The process of updating the core component of an operating system, often necessary for hardware compatibility and security.
*   **`sudo dd`**: A command-line utility used for converting and copying files, often used for creating bootable USB drives from ISO images.
*   **`apt`**: Debian's package management system, used for installing, updating, and removing software packages.
*   **`/etc/apt/source.list`**: A configuration file in Debian-based systems that lists the sources from which `apt` retrieves packages.
*   **`dpkg -i`**: A command-line utility in Debian-based systems used to install, remove, and manage `.deb` packages.
*   **`systemd-networkd`**: A systemd daemon that manages network configurations.

___
# Things Learned
A summary of the key takeaways from this exercise.

## New Concepts
-   **Manual Kernel Update on Arch**: Learned the process of manually updating the Linux kernel on an Arch system to resolve driver compatibility issues for a Debian installation.
-   **Debian Ethernet Driver Specifics**: Discovered that newer hardware might require specific, manually installed Ethernet drivers (e.g., Realtek r8125/r8169) not included in older Debian ISOs.
-   **Network Configuration Commands**: Gained familiarity with enabling and starting `link ip` and `systemd-networkd` for network setup in Debian.
-   **`apt` Mirror Configuration**: Understood the importance of correctly configuring `/etc/apt/source.list` with official Debian mirrors to fix broken installations and enable updates.
-   **`dpkg -i` for Local Packages**: Learned how to use `dpkg -i` to install `.deb` packages directly from the file system.

## Reinforced Concepts
List concepts you already knew but now understand more deeply or have seen in a new practical context.
-   **Dual Boot Installation Process**: Reinforced understanding of the general steps involved in setting up a dual-boot system, including partitioning and bootloader considerations.
-   **Hardware Compatibility in Linux**: Reaffirmed the critical role of up-to-date drivers for proper hardware recognition and functionality in Linux environments.
-   **Iterative Troubleshooting**: Experienced firsthand the iterative nature of troubleshooting complex system issues, where one solution often leads to the next problem to be solved.
-   **Importance of Official Documentation**: Confirmed the value of consulting official documentation (e.g., debian.org for firmware and mirror lists) as a primary resource for resolving system-level problems.
-   **Troubleshooting Skills**: Improved general troubleshooting skills by systematically identifying and resolving issues related to network connectivity, kernel versions, and package management in a new Linux distribution.

--- 
Obsidian links (just ignore it)
[[apt]]
[[Debian]]
[[dd]]
[[dpkg]]
[[.deb]]
[[systemd-networkd]]
[[link ip]]
[[EFI]]
