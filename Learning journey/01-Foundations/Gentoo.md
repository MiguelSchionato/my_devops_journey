# Gentoo System Implementation Report

- **Project:** Custom Gentoo Desktop Installation
- **Author:** Miguel
- **Status:** Evolving
- **Date:** 2025-07-01

---

## 1 Objective

The primary objective of this project was not merely to install a new operating system, but to undertake a deep, hands-on learning journey into the architecture and inner workings of a Linux system. The goal was to build a highly optimized, minimalist Gentoo environment from the ground up, tailored specifically to cutting-edge hardware.

This document serves as a living report of that process, with the understanding that the resulting system is not a static endpoint, but the foundation for **a constantly evolving system subject to continuous upgrades and learning.**

---

## 2 System & Prerequisites

- **Hardware:**
    - **CPU:** Ryzen 5 9600X (Zen 5)
    - **GPU:** NVIDIA RTX 5060
- **Existing Environment:**
    - Dual-boot with Windows.
    - An existing Arch Linux installation.
    - A shared `/home` partition to be used by both Arch and Gentoo.
    - An existing EFI partition.

---

## 3 Implementation & Learning Log

The installation was approached as a narrative in several acts, each presenting unique challenges and learning opportunities.

### **Act I: Architectural Foundation**

The initial phase focused on establishing the system's "DNA" via partitioning and the `make.conf` file.

- **Partitioning:** A new root partition was created for Gentoo, while the existing EFI and `/home` partitions were strategically reused to ensure interoperability.
- **`make.conf` Configuration:** This was the most critical step. A "purist" philosophy was applied:
    - **`CFLAGS="-march=znver5"`:** Optimized for the specific CPU architecture.
    - **`ACCEPT_KEYWORDS="~amd64"`:** Enabled the testing branch for necessary modern hardware support.
    - **`USE="-kde -gnome -qt5 -gtk -nls"`:** Established a minimal base by globally disabling large desktop environments and services. Every feature would have to be a conscious addition.

### **Act II: Critical Incident - The Configuration Wipe**

This phase was the most dramatic and instructive.

- **Symptoms:** A cascade of critical, seemingly unrelated errors appeared after a pause in the installation: `gcc: command not found`, `Failed to validate sane /dev`, login failures, and empty configuration files (`fstab`, `make.conf`).
- **Root Cause Analysis:** After investigating multiple dead ends, the root cause was identified: the **Stage3 tarball had been accidentally re-extracted** over the live installation, resetting the entire `/etc` directory to its default state and wiping out all prior work.
- **Lesson:** This incident provided an unforgettable lesson on the destructive power of certain commands and the absolute necessity of understanding the purpose and consequence of every step taken.

### **Act III: Meticulous System Reconstruction**

With the cause identified, the system was rebuilt with a now much deeper understanding.

- **User/Group Sync:** Recreated the user, learning the importance of synchronizing UID/GID for the shared `/home`.
- **`fstab` Hardening:** Rebuilt `fstab` using partition `LABEL`s instead of `UUID`s for improved readability.
- **Dependency Management:** Learned to use `emerge --unmerge --nodeps` to clean out "ghost" package entries left over from the configuration wipe.
- **Service Configuration:** Gained a fundamental understanding of OpenRC and the PAM authentication chain (`dbus` -> `elogind`), using `dispatch-conf` to resolve the final login issues.

### **Act IV: First Boot & GPU Troubleshooting**

The first successful boot led to the final major challenge: the graphics driver.

- **Bootloader:** Independently solved an issue where the Windows bootloader was taking priority.
- **NVIDIA Driver:** Identified the need for the `kernel-open` USE flag for the `nvidia-drivers`.
- **Performance Diagnostics:** When Hyprland felt sluggish, a crucial test using Minecraft (which ran at 500 FPS) proved that 3D acceleration was working correctly. This isolated the problem to the compositor, not the driver, demonstrating a key troubleshooting principle.

### **Act V: Refinement and Administrator Autonomy**

In the final stages, a focus was shifted from fixing problems to making informed administrative decisions about system refinement, including power management, display manager choices, and security considerations.

---

## 4 Key Configurations

The core philosophy of the system is captured in this `make.conf`:

```ini
# CPU/Compiler Optimization
COMMON_FLAGS="-O2 -pipe -march=znver5"
CFLAGS="${COMMON_FLAGS}"
CXXFLAGS="${COMMON_FLAGS}"
FCFLAGS="${COMMON_FLAGS}"
FFLAGS="${COMMON_FLAGS}"

# Build Parallelization
MAKEOPTS="-j13 -l12"

# Hardware & Drivers
VIDEO_CARDS="nvidia"
INPUT_DEVICES="libinput"

# Core System & Feature Flags (USE flags)
USE="-kde  -gnome -plasma -qt5 -qt6 -qtk -gtk3 -systemd -nls pam policykit wayland elogind dbus pipewire wireplumber alsa udev modules nvenc nptl X acl unicode readline ssl opengl opencl bzip2 zip zlib zstd fish-completion"

# Package Acceptance
ACCEPT_KEYWORDS="~amd64"
ACCEPT_LICENSE="* -@EULA @BINARY-REDISTRIBUTABLE"

# Localization
L10N="en en-US"
LINGUAS="en"
C_MESSAGES=C.utf8
```

---

## 5 Conclusion & Next Steps

This project was a success, meeting its primary goal as an intensive learning experience. The resulting system is not merely an "installation," but a system that was built, broken, understood, and repaired. The knowledge gained from debugging each unexpected issue is the most valuable outcome.

The journey does not end here. This stable, custom-built foundation is now ready for its intended purpose: **to serve as a platform for continuous learning, experimentation, and ongoing upgrades.**


---
Obsidian links - just ignore it
 [[Linux]]  [[portage]] [[Linux-Gentoo]] [[USE flags]] [[make.conf]]
