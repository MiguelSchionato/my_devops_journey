# My Home Server Journey: 
######  A Personal DevOps Learning Log

---
This repository serves as a personal log and configuration backup of my ongoing journey in setting up and managing a home server. My primary goal is to deepen my understanding of DevOps principles and Linux system administration through hands-on practice.

***This documentation is for personal study and reference; it is not intended as a comprehensive tutorial for others.***

---
### 1. Initial Setup: Arch Linux Installation

My journey began with installing Arch Linux on a dedicated machine. This provided a minimalist and highly customizable base for a server environment, allowing for a deeper understanding of system components.
This first step was pretty straight-forward, didn't run on any issue as I am comfortable using arch.
###### Why arch?
It was my first choice as I'm pretty comfortable using it as it's my main distro. I also have a boot pendrive with arch iso laying around, so it was an obvious choice. 

I like arch because it gives me the freedom to highly customize and know each and every package I have, which is one of the first objectives on learning the operating system.

| Operating System | Arch Linux                                           |
| ---------------- | ---------------------------------------------------- |
| Purpose          | build a robust, secure and learning-oriented server. |

---
### 2. Network Configuration: Wired Connectivity & Static IP (DHCP Reservation)

Initially, I faced challenges with Wi-Fi stability, leading to a focus on reliable network configuration.
  
  - iwctl & systemd-networkd: Utilized iwctl for Wi-Fi management and systemd-networkd for network configuration.
  - DHCP Challenge: Logs showed No IP addresses even after connecting to the Wi-Fi Access Point, pointing to a DHCP client issue or instability preventing IP acquisition.
   - Solution: DHCP Reservation via Router: The most effective and stable solution found was to configure a static DHCP lease (DHCP Reservation) on my Vodafone Wi-Fi 6 Station router. This ensures my server always receives the same IP address from the router based on its MAC address, while the server itself remains configured for automatic DHCP (simplifying client-side configuration and preventing IP conflicts).
   - Current State: Server runs reliably via wired Ethernet, providing a stable foundation.

### 3. Basic Server Security: Firewall (UFW) & SSH Configuration

Securing the server was an immediate priority, especially for remote access.

- UFW (Uncomplicated Firewall): Configured UFW to default to deny incoming connections and allow outgoing ones, which is a strong security baseline for a personal server.
- SSH Access: Configured SSH for remote access.
- Port Change: The default SSH port (22) was changed to a non-standard port (2424) for basic obscurity.
- Local Access Only: Initially, SSH access was restricted to my local network only (e.g., 192.168.1.0/24). This significantly reduces the attack surface from the public internet.

### 4. Personal Learning & Documentation with GitHub & Obsidian

To track my progress and organize learning materials, I've integrated my documentation process.

- GitHub Repository: This very repository serves as my primary storage for server configurations, scripts, and learning notes. It's an exercise in "Infrastructure as Code" and a living portfolio of my skills.
-  Obsidian Integration: I use Obsidian for detailed note-taking. To leverage Git's version control and GitHub's sharing capabilities, I've configured a specific subfolder within my Obsidian vault to be a Git repository. This allows me to selectively publish my technical learning notes (like this document, configuration snippets, and Python exercises) while keeping other personal notes private.

### 5. Future Endeavors: Diving Deeper into DevOps

My next steps will involve exploring core DevOps tools and concepts on this home server.

- Docker/Podman: Learning containerization for application isolation and deployment.
-  Ansible: Practicing infrastructure automation to manage server configurations programmatically.
- Prometheus & Grafana: Setting up monitoring solutions to gain insights into server performance and application health.
- CI/CD (e.g., GitLab CI Runner/Gitea Actions): Exploring continuous integration to automate testing and build processes.
-  Python and Shell script for Automation: Continuing to develop my Python and Shell scripting skills, specifically for automating tasks on the server and within DevOps pipelines.
---

This journey is about continuous learning and practical application. Each challenge is an opportunity to learn something new and solidify my understanding of the technologies crucial for a career in DevOps and System administration.

