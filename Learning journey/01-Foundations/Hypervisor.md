---
date: 2025-07-09
tags: [devops, gentoo, kvm, virtualization, kernel, nfs, xorg]
source: # Link to the course/article/documentation
---

# Topic: Building a Custom Type 1 Hypervisor with Gentoo

## 1. The Challenge
> My goal was to create a custom Type 1 hypervisor. Since many commercial options are proprietary or paid, I decided to build my own virtual machine environment. I don't have the expertise to write a hypervisor from scratch, but I can build a tailored environment for my VMs.
>
> Leveraging my experience with Gentoo Linux, I aimed to create a minimal, lightweight system optimized for virtualization. This setup would also serve as a foundation for future projects, including Docker containers.

### 1.2 The idea behind it.
> The idea was born from the necessity to use and learn multiple Linux Operating Systems to focus on my System Admin career path.
> I wanted not only to be able to have multiple KVMs running at the same time (which is totally possible without this customization), but I also wanted a full immersion of KVMs. The idea is to have a lightweight system with the sole objective of running KVMs, and my main OS, which will probably be Arch or Gentoo, will also be a KVM. Not only that, but I wanted the flexibility to use all the power available from my desktop, working around the 60 fps limit of the KVM by using VFIO GPU Pass-Through, for example.

---

## 2. Investigation & Solution Ideas

> The project started with a complete wipe of my existing partitions to create a clean slate for the hypervisor. The goal was to have only the hypervisor on bare metal, with all other operating systems running as VMs.

1.  **OS Installation:** I installed Gentoo, a process I was already familiar with, which made this step straightforward.

2.  **Kernel Compilation:** This was the most challenging part. The objective was to create an IoT-like kernel, including only the strictly necessary drivers for my hardware. This process was iterative and involved a lot of trial and error. I decided to get a functional system up and running first and then fine-tune the kernel from within a VM to save time on reboots.

3.  **Home Server Configuration:** To maintain a consistent environment across all VMs, I configured my home directory to be served from a central location using NFS. This was a critical step to avoid reconfiguring my applications on every new VM. I used the FNS tool with the TDP protocol to transfer my `/home/miguel` directory to the hypervisor and then to the VMs. You can find more details in the [Home Server README](https://github.com/your-username/your-repo/blob/main/Home_Server/README.md).

4.  **Networking:** I encountered significant issues with the network drivers. The recommended Realtek r8126 driver for my ethernet card didn't work. After much troubleshooting, I found that the r8169 driver worked, but only after I configured the necessary network bridge to provide internet access to the VMs.

5.  **NFS Configuration:** I also had a small battle with the NFS versions and configurations. It took several kernel recompilations to find the exact working setup. This was a valuable detective work and troubleshooting learning experience.

6.  **Graphical Interface:** With the core system functional, I chose to install the Xorg-server to have a minimal graphical environment, just enough to run `virt-manager` and manage VMs with a GUI, not just headless ones.

7.  **Nvidia Drivers:** The final hurdle was dealing with the newly released Nvidia drivers, which required some effort to get working correctly.

---

## 3. The Final Implemented Solution

> The final solution is a highly customized Gentoo Linux installation acting as a type 1 hypervisor. The system is minimal, with a custom-compiled kernel that includes only the necessary drivers.

- **OS:** Gentoo Linux
- **Virtualization:** KVM/QEMU with `virt-manager` for GUI management.
- **Networking:** A network bridge configured with the `r8169` driver to provide internet access to the VMs.
- **Shared Home Directory:** An NFS server hosts the `/home/miguel` directory, which is mounted by all VMs, ensuring a consistent user experience.
- **Graphical Environment:** A minimal Xorg server installation provides the necessary graphical interface for `virt-manager`.

This setup allows me to run multiple VMs, including my main OS, with the ability to pass through my GPU for native performance.

---

## 4. Key Concepts & Reference Links

* **[[Kernel Compilation]]**: The process of building a custom Linux kernel.
* **[[VFIO GPU Pass-Through]]**: A technique that allows a VM to have direct access to a physical GPU.
* **[[NFS (Network File System)]]**: A distributed file system protocol allowing a user on a client computer to access files over a computer network.
* **[[Linux Bridge]]**: A software switch that allows you to connect multiple network interfaces.
* **[[Xorg]]**: The most common display server for Linux.

---
# Things Learned
> A summary of the key takeaways from this exercise.

## New Concepts
> List the brand-new ideas, commands, or techniques you discovered.
>
> - The importance of a minimal kernel for a hypervisor.
> - How to configure a network bridge in Gentoo.
> - The intricacies of NFS configuration and kernel dependencies.
> - The challenges of dealing with new and sometimes unstable hardware drivers.

## Reinforced Concepts
> List concepts I already knew but now understand more deeply or have seen in a new practical context.
>
> - Troubleshooting and problem-solving skills.
> - The flexibility and power of Gentoo for creating custom systems.
> - The importance of a systematic approach to debugging complex issues.
> - The value of a shared home directory for a multi-VM environment.