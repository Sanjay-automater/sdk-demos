# Device Firmware Upgrade (DFU) Using Bluetooth Over-the-Air (OTA) Update

## Introduction

This tutorial describes how to perform a Device Firmware Upgrade (DFU) using Bluetooth Over-the-Air (OTA) update. 

Any device with OTA updates enabled in its GATT profile can be upgraded using OTA. Most example applications provided in the Bluetooth SDK include OTA support. In these examples, the DFU mode is triggered through the Silicon Labs OTA service in the application's GATT database. You can add OTA functionality by installing the OTA DFU software component in your project.
<!-- Question for engineering: Is there a specific version of the OTA DFU software component, and does it have backward compatibility? -->

> **Note:**  
> In this tutorial, you will upgrade the **Bluetooth - SoC Empty** example application to the **Bluetooth - SoC Thermometer** example application. The **Bluetooth - SoC Thermometer** application includes the Health Thermometer service, which provides visual feedback so you can verify that the user application has changed.

---

## Requirements

- Wireless Starter Kit and Bluetooth-capable radio board
- Simplicity Studio 5
- Android or iOS mobile device

---

## Before You Begin

<!-- Questions for engineering: 
1. Please provide the store links for Android and iOS installation packages to add in step 1)
2. Please provide example values for *PATH_SCMD** and **PATH_GCCARM** environment variables to add as the table in sub-step 1 of step 6.
3. Please provide the AN1086 link to add in the note in sub-step 2 of step 6.
 -->

1. Download the **Simplicity Connect** mobile app (Android or iOS). 
2. Connect the kit to your computer and select it in **Simplicity Studio**.
3. In **Simplicity Studio Launcher**, create the **Bluetooth - SoC Empty** example project.
4. Build the **Bluetooth - SoC Empty** project and flash the firmware image to the device.
5. Create the **Bluetooth - SoC Thermometer** example project.
6. Build the **Bluetooth - SoC Thermometer** project.
    1. Define the `PATH_SCMD` and `PATH_GCCARM` environment variables.

        | Environment Variable Name | Example Values |
        |---------------|------------------------|
        | `PATH_SCMD`   | Value1                 |
        | `PATH_GCCARM` | Value2                 |

    2. In the project tree, double-click the **create_bl_files.bat** (Windows) or **create_bl_files.sh** (Linux/macOS) script.

       The script creates an `output_gbl` folder in your project and generates six `.gbl` upgrade image files:
        - **application.gbl**: User application (including full Bluetooth stack)
        - **application-crc.gbl**: User application with a CRC32 checksum
        - **apploader.gbl**: AppLoader (including minimal Bluetooth stack)
        - **apploader-crc.gbl**: AppLoader with a CRC32 checksum
        - **full.gbl**: User application and AppLoader (full update) for UART DFU (not needed in this example)
        - **full-crc.gbl**: User application and AppLoader (full update) with a CRC32 checksum for UART DFU (not needed in this example)
       > **Note:**
       > A full update is required only if the AppLoader needs to be updated. For more information, see AN1086.
7. Transfer the `.gbl` files to your smartphone so the mobile app can access them.
    - You can transfer the files via USB to any folder on your phone or upload them to a cloud storage service (such as Google Drive, Dropbox, or iCloud) that is accessible from your phone.
8. Launch the **Simplicity Connect** mobile app.

---

## Perform the OTA Upgrade Using Simplicity Connect

After you launch the **Simplicity Connect** mobile app, follow these steps:

1. In the app, go to the **Bluetooth Browser**.
2. Find and connect to your kit (the default name is **Empty Example**).
3. Open the pop-up menu in the upper right corner and tap **OTA**.
4. Tap **Partial OTA** and select the `.gbl` file on your smartphone.
5. Tap **OTA** to start the upgrade.
6. When the OTA process finishes, verify that the kit is now running the **Bluetooth - SoC Thermometer** example application.  
   The kit will appear in the **Bluetooth Browser** with the new name **Thermometer Example**.

---

## Troubleshooting OTA Upload Issues

To enable Bluetooth OTA upgrades, program the target device with the **Gecko Bootloader**. The **Gecko Bootloader** is an application bootloader that requires the application to manage firmware image acquisition.

- Running a demo in **Simplicity Studio** flashes both the bootloader and the user application to the device.
- Flashing an example project flashes only the application.

If your OTA upload stops at 0% and you see the message “GATT CMD STARTED” on your Android device, the bootloader may be missing or incorrect. To fix this issue, follow these steps:

1. In **Simplicity Studio Launcher**, select **Create New Project** to create a Gecko Bootloader project for your kit.
2. Select the **Internal Storage Bootloader** example project with a configuration that matches your storage size.
3. In the `.isc` file of the bootloader project, configure the required options.
4. Build the project.
    - For Series 1 devices, find the `-combined.s37` file in the build directory named `GNU ARM - Default`. This file contains the combined image of the first stage bootloader and the main bootloader with a CRC32 checksum.
    - For Series 2 devices, find the `-crc.s37` file in the build directory. This file contains the main bootloader with a CRC32 checksum.
5. Flash the bootloader image to the device.

---

