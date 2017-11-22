# Deploy Harbor from OVA

**Prerequisites**

- You downloaded build of the OVA installer from: **location of the download url**
- Deploy the appliance to a vCenter Server instance. Deploying the appliance directly on an ESXi host is not supported.
   - vCenter 6.0.0 or higher
   - vSphere 6.0.0 or higher
- Deploy the appliance to a vCenter Server system that meets the minimum system requirements:
   - 2 vCPUs
   - 8GB RAM
   - 80GB free disk space on the datastore
- Ensure that vCenter user have the following privileges:
   - Datastore > Allocate space
   - Datastore > Low level file Operations
   - Folder > Create Folder
   - Folder > Delete Folder
   - Network > Assign network
   - Resource > Assign virtual machine to resource pool
   - Virtual machine > Configuration > Add new disk
   - Virtual Machine > Configuration > Add existing disk
   - Virtual Machine > Configuration > Add or remove device
   - Virtual Machine > Configuration > Change CPU count
   - Virtual Machine > Configuration > Change resource
   - Virtual Machine > Configuration > Memory
   - Virtual Machine > Configuration > Modify device settings
   - Virtual Machine > Configuration > Remove disk
   - Virtual Machine > Configuration > Rename
   - Virtual Machine > Configuration > Settings
   - Virtual machine > Configuration > Advanced
   - Virtual Machine > Interaction > Power off
   - Virtual Machine > Interaction > Power on
   - Virtual Machine > Inventory > Create from existing
   - Virtual Machine > Inventory > Create new
   - Virtual Machine > Inventory > Remove
   - Virtual Machine > Provisioning > Clone virtual machine
   - Virtual Machine > Provisioning > Customize
   - Virtual Machine > Provisioning > Read customization specifications
   - vApp > Import
   - Profile-driven storage -> Profile-driven storage view
- Ensure that all vCenter Server instances and ESXi hosts in the environment in which you are deploying the appliance have network time protocol (NTP) running. Running NTP prevents problems arising from clock skew between the vSphere Integrated Containers appliance, virtual container hosts, and the vSphere infrastructure.
- If your intend to use a custom certificates, need to change it in /data/harbor.cfg and reconfigure harbor.
- Use the Flex-based vSphere Web Client to deploy the appliance. You cannot deploy OVA files from the HTML5 vSphere Client or from the legacy Windows client.

**Procedure**
1. In the vSphere Web Client, right-click an object in the vCenter Server inventory, select **Deploy OVF template**
   ![Screenshot of Deploy OVF template](img/ovainstall/DeployOVFmenu.png)
2. Select template: navigate to the OVA file or input the URL of the ova file in URL field. 
   ![Screenshot of Import ova](img/ovainstall/importova.png)
3. Follow the installer prompts to perform basic configuration of the appliance and to select the vSphere resources for it to use. 
    
    - Accept or modify the appliance name
    - Select the destination datacenter or folder
    ![Screenshot of appliance name](img/ovainstall/namelocation.png)
     - Select the destination host, cluster, or resource pool
    ![Screenshot of resoure pool](img/ovainstall/resource.png)
    - Select the disk format and destination datastore
    ![Screenshot of datastore](img/ovainstall/datastore.png)
    - Select the network that the appliance connects to
    ![Screenshot of network](img/ovainstall/network.png)

4. On the **Customize template** page, under Harbor Configure, select the authentication mode and Harbor admin password, if authentication mode is set to ldap_auth, the Harbor LDAP configure is required. the Harbor Admin Password is required, it is used in step 9 to login.
    ![Screenshot of customize harbor](img/ovainstall/customizeharbor.png)
    If the Authentication Mode is set to ldap_auth, need to config Harbor LDAP settings:
    - LDAP Base DN: The base DN to lookup user
    - LDAP UID: The attribute field to match user, such as uid, cn, email or other attributes.
    - LDAP URL: The url for LDAP endpoint
    - Search DN: The user's DN who has the permission to search the LDAP Server
    - Search DN Password: The password of search DN.
    - Search Filter: The filter to search users, can be empty.
    - Search Scope: The scope to search users.
    ![Screenshot of customize ldap](img/ovainstall/customizeldap.png)

5. On the **Customize template** page, under **System**, set the root password for the appliance VM and the option for **Permit Root Login**. 

    Setting the root password for the appliance is mandatory. 

    **IMPORTANT**: You require SSH access to the vSphere Integrated Containers appliance to perform upgrades. You can also use SSH access in exceptional cases that you cannot handle through standard remote management or CLI tools. Only use SSH to access the appliance when instructed to do so in the documentation.
    ![Screenshot of customize template system](img/ovainstall/system.png)

6. Expand **Networking Properties** and optionally configure a static IP address for the appliance VM. 

    To use DHCP, leave the **Networking Properties** blank.

    **IMPORTANT**: If you set a static IP address for the appliance, use spaces to separate DNS servers. Do not use comma separation for DNS servers. 

    - Leave the text boxes blank to use auto-generated certificates.
   
7. When the deployment completes, refresh the current page and power on the appliance VM.
   ![Screenshot of power on](img/ovainstall/poweron.png)
   Press Enter to login
   ![Screenshot of login root](img/ovainstall/poweron2.png)

   Login to the console with root username and password, and get the harbor appliance address by ifconfig.

   If you deployed the appliance so that it obtains its address via DHCP, go to the **Summary** tab for the appliance VM and note the address.

8. (Optional) If you provided a static network configuration, view the network status of the appliance.

    1. In the **Summary** tab for the appliance VM, launch the VM console
    2. In the VM console, press the right arrow key. 

    The network status shows whether the network settings that you provided during the deployment match the settings with which the appliance is running. If there are mismatches, power off the appliance and select **Edit Settings** > **vApp Options** to correct the network settings.
    
9. In a browser, go to  https://<i>harbor_appliance_address</i> and when prompted, enter the username admin and the password of admin input in step 4
   ![Screenshot of login harbor](img/ovainstall/login.png)

**Result**

- You see the Harbor administration console

    ![Screenshot of after login](img/ovainstall/afterlogin.png)
