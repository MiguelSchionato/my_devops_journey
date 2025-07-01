---
date: 2025-06-27
tags: [devops, kvm, archlinux, networking, troubleshooting, virtual-machines, libvirt, qemu, vnc]
source: # Link to the course/article/documentation
---

# Topic: KVM Networking Troubleshooting

## 1. The Problem / The Challenge

The initial objective was to set up a home server using Arch Linux with KVM for virtual machines (VMs). The core challenge quickly became establishing reliable network connectivity for these VMs, specifically getting them to communicate with the external network via a bridge (`br0`) and resolving console access issues.

---

## 2. Investigation & Solution Ideas

This journey was a real roller-coaster, full of twists and turns! Here's how I tackled each hurdle:

### 2.1 Initial Setup & Network Bridge (`br0`) Woes

1.  **Goal:** Get KVM up and running on Arch Linux, with VMs connecting to my home network via a `br0` bridge.
2.  **First Attempt (systemd-networkd):** My initial thought was to use `systemd-networkd` to create the bridge and connect my wired interface (`enp171s0`) to it. The idea was simple: VMs connect to `br0`, grab an IP from my Vodafone router, and boom, internet!
3.  **Problem:** `systemd-networkd` and Wi-Fi (`wlan0`) managed by NetworkManager just didn't play nice. It caused instability on my Wi-Fi, a classic network management conflict.
4.  **Pivot to NetworkManager:** Okay, `systemd-networkd` out, NetworkManager in! I disabled `systemd-networkd` and configured `br0` and its slave interface (`br0-slave`) directly with NetworkManager. This felt more natural since NetworkManager was already handling my Wi-Fi.
5.  **New Problem:** Even with the config, `br0` wasn't active. I had to manually kick it with `nmcli connection up br0` and `nmcli connection up br0-slave`. Finally, `br0` showed up and worked! Phew.

### 2.2 The Mysterious Empty Console

1.  **VM Creation:** Network sorted, time for the VM! I used `virt-install` to create an Ubuntu Server VM, trying to get the installer to show up in my SSH terminal using `--graphics none` and `--console pty,target_type=serial`.
2.  **Problem:** Black screen. Frozen console. The installer output just wasn't redirecting to my terminal. Super frustrating!

### 2.3 The X-Forwarding Battle

1.  **Idea:** If text console failed, maybe graphical? I tried `ssh -X homeserver` and `virt-viewer`.
2.  **Problem:** Errors galore! "Warning: untrusted X11 forwarding setup failed: xauth key data not generated" and "X11 forwarding request failed on channel 0." I checked `sshd_config` (`X11Forwarding yes`, `X11UseLocalhost no`), cleaned `xauth` files, but no dice. This pointed to a deeper issue with my local graphical environment (Hyprland with Nvidia).

### 2.4 KVM/Libvirt Under-the-Hood Errors

1.  **Discovery:** While wrestling with consoles, `libvirtd` logs revealed some nasty surprises preventing VMs from even starting.
2.  **`Cannot find 'dmidecode'`:** Libvirtd couldn't find `dmidecode`, a tool for hardware info. Easy fix: `sudo pacman -S dmidecode`.
3.  **`Unable to get XATTR trusted.libvirt.security.ref_dac`:** This was the big one. My filesystem (where `.qcow2` disks lived) didn't support or allow libvirt to write "Extended Attributes" (XATTRs), which are crucial security metadata. This literally stopped libvirt from touching my VM disks. I checked `/var/lib/libvirt/images` permissions, but the XATTR issue was deeper.

### 2.5 The Debian Netinstall & VNC Breakthrough

1.  **Strategy Shift:** With console and XATTR issues persisting, I changed tactics.
2.  **New ISO:** Switched to Debian Netinstall, known for its more robust text installer.
3.  **VNC to the Rescue:** To bypass all the console and X-forwarding headaches, I decided to expose the VM's graphical interface via VNC (`--graphics vnc,listen=0.0.0.0,port=5900` in `virt-install`).
4.  **Desktop VNC Client Challenges:** This was another mini-battle! `vncviewer` (TigerVNC) kept giving "Can't open display: :0" – confirming my local graphical environment (Hyprland with Nvidia) was the culprit. I tried `xhost` (not found), `remmina` (GLib-GIO-CRITICAL errors), and `tightvnc` (compilation errors).
5.  **VICTORY with Vinagre!** After trying what felt like every VNC client under the sun, Vinagre finally worked perfectly! I could see the VM installation on my desktop! This was a huge relief, confirming the VM was actually running on the server, and my client was the problem.

### 2.6 The Final GRUB Hurdle

1.  **Installation Progress:** Debian installed, but then...
2.  **GRUB Freeze:** The VM would freeze at "boot from hard disk." This meant GRUB (the bootloader) was installed, but the OS wasn't booting.
3.  **GRUB is Key:** Confirmed GRUB is essential for Linux boot and must be installed on the VM's main virtual disk (`/dev/vda`) during installation. No separate FAT32 partition needed for legacy BIOS VMs.
4.  **Next Step:** The plan is to re-do the VM installation carefully, focusing on the GRUB installation step and monitoring via Vinagre to ensure everything copies and configures correctly.

---

## 3. The Final Implemented Solution

The ultimate solution involved a multi-pronged approach to overcome persistent networking and console access issues. The key steps were:

1.  **NetworkManager for Bridge Configuration:** Transitioning from `systemd-networkd` to NetworkManager to reliably configure the `br0` bridge, ensuring stable network connectivity for VMs. This involved manually activating the bridge and its slave interface.
2.  **VNC for VM Console Access:** Bypassing problematic serial console and X-forwarding issues by exposing the VM's graphical interface via VNC. This allowed for visual interaction with the VM installer and subsequent troubleshooting. The `virt-install` command was modified to include `--graphics vnc,listen=0.0.0.0,port=5900`.
3.  **Addressing KVM/Libvirt Dependencies:** Installing `dmidecode` to resolve a missing dependency for `libvirtd`.
4.  **Understanding GRUB Installation:** Recognizing the critical role of GRUB in booting the VM and ensuring its correct installation on the virtual disk (`/dev/vda`) during the OS installation process.

While each step presented its own challenges, the combination of these solutions allowed for a successful Debian VM installation and access.

 >After all that struggle. I decide to install a dual-boot  on my main desktop to trouble-shoot and get rid of my Hyperland-Nvidia setup that can, sometimes, be problematic. You can check all [here](../01-Foundations/Linux-%20Debian.md) 
 

---

## 4. Key Concepts & Reference Links

*   **KVM (Kernel-based Virtual Machine)**: A virtualization infrastructure for the Linux kernel that turns it into a hypervisor.
*   **QEMU (Quick EMUlator)**: A generic and open-source machine emulator and virtualizer that works with KVM to provide hardware emulation for VMs.
*   **libvirt**: A virtualization management library that provides a common API for managing various hypervisors, including KVM/QEMU.
*   **`virt-install`**: A command-line tool for creating new virtual machines using `libvirt`.
*   **Network Bridge (`br0`)**: A virtual networking device that connects multiple network segments at the data link layer, allowing VMs to share the host's physical network interface.
*   **`systemd-networkd`**: A system service that manages network configurations.
*   **NetworkManager**: A dynamic network management daemon that handles network connections.
*   **X-Forwarding**: A feature of SSH that allows graphical applications to be run on a remote server and displayed on the local machine.
*   **`dmidecode`**: A tool for dumping a computer's DMI (Desktop Management Interface) table contents in a human-readable format.
*   **XATTRs (Extended Attributes)**: Metadata associated with filesystems, used by `libvirt` for security and other purposes.
*   **VNC (Virtual Network Computing)**: A graphical desktop sharing system that allows remote control of a computer.
*   **GRUB (GRand Unified Bootloader)**: A boot loader package from the GNU Project, used to boot a computer's operating system.

---
# Things Learned
A summary of the key takeaways from this exercise.

## New Concepts
-   **KVM/QEMU Interaction:** Gained a deeper understanding of how KVM and QEMU work together for virtualization, and how `libvirt` acts as an abstraction layer.
-   **NetworkManager for Bridges:** Learned to configure network bridges (`br0`) effectively using NetworkManager, especially when dealing with conflicts with other network services like `systemd-networkd`.
-   **VNC as a Troubleshooting Tool:** Discovered the utility of VNC for accessing VM consoles when traditional serial console or X-forwarding methods fail, particularly in complex graphical environments.
-   **XATTR Importance for Libvirt:** Understood the critical role of Extended Attributes (XATTRs) for `libvirt`'s operation and how filesystem limitations can impact VM functionality.
-   **GRUB's Role in VM Boot:** Reinforced the importance of correct GRUB installation within the VM's virtual disk for successful operating system boot.

## Reinforced Concepts
List concepts you already knew but now understand more deeply or have seen in a new practical context.
-   **Linux Networking Fundamentals:** Solidified understanding of network bridging, interface management, and the interplay between different network configuration tools in Linux.
-   **Troubleshooting Methodologies:** Applied systematic troubleshooting techniques, moving from initial hypotheses to deeper investigations and adapting solutions based on observed errors.
-   **Importance of Logs:** Reaffirmed the value of checking system and application logs (e.g., `libvirtd` logs) for diagnosing underlying issues.
-   **Hardware/Software Compatibility:** Experienced firsthand how hardware (e.g., Ethernet card) and software (e.g., kernel version, graphical environment) compatibility can impact virtualization setups.
-   **Persistence and Iteration:** The process highlighted the need for persistence and iterative problem-solving when facing complex technical challenges.
